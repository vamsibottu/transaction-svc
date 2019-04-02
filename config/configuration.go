package config

import "github.com/spf13/viper"

// all configuration keys
const (
	ListeningPort = "api.handlers.listeningport"
)

func init() {
	viper.SetDefault(ListeningPort, ":8000")
}
