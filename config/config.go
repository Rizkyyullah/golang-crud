package config

import (
  "log"
  "github.com/spf13/viper"
)

type Config struct {
  APP_PORT    uint16
  DB_DRIVER   string
  DB_HOST     string
  DB_PORT     uint16
  DB_USER     string
  DB_PASSWORD string
  DB_NAME     string
}

var ENV *Config

func LoadConfig() {
  viper.AddConfigPath(".")
  viper.SetConfigName(".env")
  viper.SetConfigType("env")
  
  if err := viper.ReadInConfig(); err != nil {
    panic(err)
  }
  
  log.Println("Load Configuration")
  if err := viper.Unmarshal(&ENV); err != nil {
    panic(err)
  }
}
