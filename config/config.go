package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Server struct {
	RunMode      string        `mapstructure:"run-mode"`
	HttpPort     int           `mapstructure:"http-port"`
	ReadTimeout  time.Duration `mapstructure:"read-timeout-ms"`
	WriteTimeout time.Duration `mapstructure:"write-timeout-ms"`
}

type Database struct {
	Type        string `mapstructure:"type"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Host        string `mapstructure:"host"`
	Name        string `mapstructure:"name"`
	TablePrefix string `mapstructure:"table-prefix"`
}

type Config struct {
	ServerConfiguration   Server   `mapstructure:"server"`
	DataBaseConfiguration Database `mapstructure:"database"`
}

var Configuration Config

func Setup() {
	viper.SetConfigFile("config.yml")
	viper.AddConfigPath("/etc/mail-sender-pool/")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to load config file: %s", err))
	}
	if err := viper.Unmarshal(&Configuration); err != nil {
		panic(fmt.Errorf("failed to load config file: %s", err))
	}
}
