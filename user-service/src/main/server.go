package main

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/vtomar01/user-service/src/main/config"
	"github.com/vtomar01/user-service/src/main/logging"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logging.Init()
	config.SetUpEnvironment()
	config.SetUpDatabase()
	s := NewServer()
	s.SetupComponents(config.InitializeApplication)
	s.ServeHTTP()

}

type Server struct {
	*mux.Router
	Address string
}

func NewServer() *Server {
	r := mux.NewRouter()
	addr := viper.GetString("server-ip") + ":" + viper.GetString("server-port")
	s := Server{r, addr}
	return &s
}

func (s Server) SetupComponents(applicationInitializer func(router *mux.Router)) {
	apiMux := s.PathPrefix("/api").Subrouter()
	applicationInitializer(apiMux)
}

func (s Server) ServeHTTP() {
	loggedRouter := handlers.LoggingHandler(os.Stdout, s.Router)
	srv := &http.Server{
		Handler:      loggedRouter,
		Addr:         s.Address,
		WriteTimeout: time.Minute,
		ReadTimeout:  time.Minute,
	}

	logging.Log.Info("Server starting at addr: ", s.Address)

	go func() {
		logging.Log.Info("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			logging.Log.Fatal(err)
		}
	}()

	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	os.Exit(0)
}
