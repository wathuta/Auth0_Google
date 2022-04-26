package main

import (
	"Oauth2/google/configs"
	"Oauth2/google/logger"
	"Oauth2/google/service"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	// Initialize Viper across the application
	configs.InitializeViper()

	// Initialize Logger across the application
	logger.InitializeCustomLogger()

	// Initialize Oauth2 Services
	service.InitializeOauthGoogle()

	// Routes for the application
	http.HandleFunc("/", service.HandleMain)
	http.HandleFunc("/login-gl", service.HandleGoogleLogin)
	http.HandleFunc("/callback-gl", service.CallBackFromGoogle)

	//logger.Log.Info("Started running on http://localhost:" + viper.GetString("port"))
	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), nil))
}
