// TODO:
// validate config

package config

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	DocumentStore string
	MongoDB       mongoDB
}

type mongoDB struct {
	Name   string
	URI    string
	DBName string
}

var legal = &struct {
	documentStore []string
}{
	documentStore: []string{"mongodb"},
}

func New() (*config, error) {
	conf := new(config)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Fail: find config file")
			return nil, err
		} else {
			log.Printf("Fail: read config file\n%s", err)
			return nil, err
		}
	}

	if err := viper.Unmarshal(conf); err != nil {
		log.Printf("Fail: decode config\n%s", err)
		return nil, err
	}

	return conf, nil
}
