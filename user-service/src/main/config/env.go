package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/vtomar01/user-service/src/main/logging"
)

func SetUpEnvironment() {
	err := viper.BindEnv("BUILD_PATH")
	if err != nil {
		logging.Log.Fatal("BUILD_PATH load failed")
	}

	viper.SetConfigName("local")

	viper.AddConfigPath(viper.GetString("BUILD_PATH") + "/user-service/src/main/resources/")
	viper.AddConfigPath("/app/")

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {
		panic(err)
	}
	viper.SetConfigType("json")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logging.Log.Info("Config file changed:", e.Name)
	})
}
