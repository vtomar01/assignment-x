package config

import (
	"github.com/gorilla/mux"
	"github.com/vtomar01/user-service/src/main/apis"
	"github.com/vtomar01/user-service/src/main/clients"
	"github.com/vtomar01/user-service/src/main/repo"
	"github.com/vtomar01/user-service/src/main/services"
	"net/http"
	"time"
)

func InitializeApplication(apiMux *mux.Router) {

	phoneStdClientConf := NewPhoneStandardizerClientConfig()
	httpClient := &http.Client{
		Timeout: time.Duration(phoneStdClientConf.DefaultTimeOut) * time.Second,
	}
	phoneStdClient := clients.NewPhoneStandardizerClient(phoneStdClientConf.BasePath, httpClient)

	userRepo := repo.NewUserRepo(GetDbConn())

	userSvc := services.NewUserService(userRepo)

	userOrch := services.NewUserOrch(userSvc, phoneStdClient)

	apis.NewUserController(apiMux, userOrch)
}
