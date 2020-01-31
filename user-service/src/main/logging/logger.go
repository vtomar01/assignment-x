package logging

import (
	"go.uber.org/zap"
	"log"
)

var Log *zap.SugaredLogger

func Init() {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	Log = l.Sugar()
}
