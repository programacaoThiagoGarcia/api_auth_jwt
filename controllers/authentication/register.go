package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"programacao.thiagogarcia/jwt-gin/models"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(ctx *gin.Context) {
	var input RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}
	u.Username = input.Username
	u.Password = input.Password

	_, err := u.SaveUser()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created with success"})

}
