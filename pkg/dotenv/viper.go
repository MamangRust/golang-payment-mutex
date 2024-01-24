package dotenv

import (
	"fmt"

	"github.com/spf13/viper"
)

func Viper() error {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	return err
}
