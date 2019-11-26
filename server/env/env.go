package env

import (
	"github.com/spf13/viper"
)

type Config struct {
	MngoDB *mongoConnConfig
	Api *grpcServerConnConfig
}

type grpcServerConnConfig struct {
	Port    int
}

type mongoConnConfig struct {
	Host 	string
	Port 	int
	DB 		string
}

var Settings *Config

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	viper.SetDefault("GRPC_PORT", "50051")
	viper.SetDefault("MONGO_HOST", "localhost")
	viper.SetDefault("MONGO_PORT", 27017)
	viper.SetDefault("MONGO_DB", "test")

	Settings = &Config{
		Api: &grpcServerConnConfig{
			Port: viper.getString("GRPC_PORT"),
		},
		MngoDB: &mongoConnConfig{
			Host: viper.GetString("MONGO_HOST"),
			Port: viper.GetInt("MONGO_PORT"),
			DB: viper.getString("MONGO_DB"),
		},
	}
}
