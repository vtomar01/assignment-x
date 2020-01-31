package config

import (
	"github.com/spf13/viper"
)

type ClientConfig struct {
	BasePath       string
	DefaultTimeOut int
}

func NewPhoneStandardizerClientConfig() *ClientConfig {
	return &ClientConfig{
		BasePath:       viper.GetString("clients.phone-standardizer.basepath"),
		DefaultTimeOut: 1,
	}
}
