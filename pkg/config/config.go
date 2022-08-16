package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBURL  string `mapstructure:"DB_URL"`
	Secret string `mapstructure:"SECRET"`
}

func LoadConfig() (c *Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("local")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Panicf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		log.Panicf("unable to decode into struct, %v", err)
	}

	return
}
