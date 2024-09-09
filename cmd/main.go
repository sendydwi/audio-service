package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sendydwi/audio-service/database"
	"github.com/sendydwi/audio-service/service/audioservice"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	database := database.Init()
	r := gin.Default()

	registerRestHandler(r, database)

	r.Run(":8182")
}

func registerRestHandler(r *gin.Engine, database *gorm.DB) {
	// register service handler
	base := r.Group("")

	audioRestHandler := audioservice.NewRestHandler(database)
	audioRestHandler.RegisterHandlerRoutes(base)
}
