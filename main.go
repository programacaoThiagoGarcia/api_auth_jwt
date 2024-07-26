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

	security := r.Group("/security")
	security.POST("/register", authentication.Register)
	security.POST("/login", authentication.Login)

	public := r.Group("/public")
	public.GET("", nil)

	private := r.Group("/private")
	private.GET("", nil)

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
