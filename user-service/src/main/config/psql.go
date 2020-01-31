package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"github.com/vtomar01/user-service/src/main/logging"
)

// DB connection
var dbConn *gorm.DB

func SetUpDatabase() {

	databaseConfig := getDatabaseConfig()

	db, err := gorm.Open("postgres", databaseConfig)
	if err != nil {
		logging.Log.Panic("failed to connect database with error: ", err.Error())
	}
	maxIdleConn := viper.GetInt("database.maxIdleConn")
	maxOpenConn := viper.GetInt("database.maxOpenConn")
	logMode := viper.GetBool("database.debugEnabled")
	db.LogMode(logMode)
	dbConn = db
	db.DB().SetMaxIdleConns(maxIdleConn)
	db.DB().SetMaxOpenConns(maxOpenConn)
}

func GetDbConn() *gorm.DB {

	return dbConn
}

func getDatabaseConfig() interface{} {
	user := viper.GetString("database.username")
	password := viper.GetString("database.password")
	database := viper.GetString("database.name")
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	psqlConf := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	logging.Log.Info("Connecting to : ", psqlConf)

	return psqlConf
}
