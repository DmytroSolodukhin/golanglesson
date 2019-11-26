package env

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	FileName string
	Api *grpcClientConnConfig
}

type grpcClientConnConfig struct {
	Host    string
	Port    int
	Timeout time.Duration
}

var Settings *Config

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	viper.SetDefault("FILE_NAME", "files/googlechrome.dmg")
	viper.SetDefault("GRPC_HOST", "localhost")
	viper.SetDefault("GRPC_PORT", 50051)
	viper.SetDefault("API_TIMEOUT", 30*time.Second)

	Settings = &Config{
		FileName: viper.GetString("FILE_NAME"),
		Api: &grpcClientConnConfig{
			Host: viper.GetString("GRPC_HOST"),
			Port: viper.GetString("GRPC_PORT"),
			Timeout: viper.GetDuration("API_TIMEOUT"),
		},
	}
}
