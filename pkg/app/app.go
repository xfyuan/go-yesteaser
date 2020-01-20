package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// Config is global object that holds all application level variables.
//var Config appConfig
var DB *gorm.DB
var DBErr error

// LoadConfig loads config from files
func LoadConfig(configDir string)  {
	viper.SetConfigType("yaml")

	if configDir == "" {
		workPath, err := os.Getwd()
		if err != nil {
			panic(fmt.Errorf("read work path failed: [%s]", err))
		}
		configDir = filepath.Join(workPath, "config")
	}

	viper.AddConfigPath(configDir)
	viper.SetConfigName("config")

	viper.SetEnvPrefix("yestea")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read common config failed: [%s]", err))
	}

	env := viper.GetString("ENV")
	if env == "" {
		viper.SetConfigName("dev")
	}
	viper.SetConfigName(env)
	if err := viper.MergeInConfig(); err != nil {
		panic(fmt.Errorf("load configuration failed: [%s]", err))
	}
}

func GenerateDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.dbname"),
	)
}

