package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Configs struct {
	App      app
	Database database
}

type app struct {
	Name           string
	Port           int
	Debug          bool
	Version        string
	Test           string
	MigrationsPath string
	LogPath        string
	I18nPath       string
}

type database struct {
	Host            string
	Port            int
	User            string
	Password        string
	Name            string
	PrintQuery      int
	ConnMaxLifeTime int
	ConnMaxIdleTime int
	MaxIdleConns    int
	MaxOpenConns    int
}

var c Configs
var App app
var Database database

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	configPath := os.Getenv("CONFIG_PATH")
	fmt.Println(configPath)
	if configPath == "" {
		configPath = "./pkg/configs"
	}
	viper.AddConfigPath(configPath)
	viper.AddConfigPath("$HOME/.myapp")

	// Tentukan environment variables prefix
	viper.SetEnvPrefix("MYAPP")
	viper.AutomaticEnv()
	viper.BindEnv("app.name", "MYAPP_APP_NAME")
	viper.BindEnv("app.migrationsPath", "MYAPP_APP_MIGRATIONS_PATH")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return err
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		return err
	}

	App = c.App
	Database = c.Database
	// for config file change without restart app
	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	err = viper.Unmarshal(&C)
	// 	if err != nil {
	// 		applog.Info().Msg(err.Error())
	// 	}
	// })
	// viper.WatchConfig()

	return nil
}
