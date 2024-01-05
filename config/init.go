package config

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// InitApp initializes the app, this includes reading the .env file during development. If no path to is given the directory of the calling function will be used.
func InitApp(path ...string) {
	gin.SetMode(gin.DebugMode)
	log.Println("Debug mode engaged!")
	if gin.Mode() == gin.ReleaseMode {
		if len(path) == 0 {
			path = []string{".env"}
		}
		if true {
			err := godotenv.Load(path...)

			if err != nil {
				log.Fatal("Could not load the environment file, Error: ", err)
			}
		}
	}
}
