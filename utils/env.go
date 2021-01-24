package utils

import (
	"os"

	"github.com/spf13/viper"
)

var Viper *viper.Viper

func init() {
	Viper = viper.New()
	Viper.SetConfigFile(".env")
	err := Viper.ReadInConfig()
	if err != nil {
		print(err)
		os.Exit(1)
	}
}
