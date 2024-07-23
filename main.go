package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"programacao.thiagogarcia/jwt-gin/controllers/authentication"
	"programacao.thiagogarcia/jwt-gin/models"
)

func main() {
	initENV()
	models.ConnectDataBase()
	setupServer()

}

func setupServer() {
	r := gin.Default()

	public := r.Group("/api")
	public.POST("/register", authentication.Register)
	public.POST("/login", authentication.Login)

	PORT := os.Getenv("PORT")
	if err := r.Run(PORT); err != nil {
		log.Fatalln(err)
	}
}

func initENV() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
		return
	}
}
