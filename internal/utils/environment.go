package utils

import (
	"github.com/spf13/viper"
	"log"
)

// GetEnv .
func GetEnv(key string) string {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading configuration file %s", err)
		return "Could not read configuration file."
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}
