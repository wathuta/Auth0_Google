package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitializeViper() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath("./")

	//Enable viper to read env vars
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}
