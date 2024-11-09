package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port    string `mapstructure:"PORT"`
	GinMode string `mapstructure:"GIN_MODE"`
}

var AppConfig *Config

func init() {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()

	failOnError(v.BindEnv("PORT"), "Failed on Bind PORT")
	failOnError(v.BindEnv("GIN_MODE"), "Failed on Bind GIN_MODE")
	err := v.ReadInConfig()
	if err != nil {
		log.Println("Load from environment variable")
	}
	err = v.Unmarshal(&AppConfig)
	if err != nil {
		failOnError(err, "Failed to read enivronment")
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
