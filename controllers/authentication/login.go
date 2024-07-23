package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"

	jwtapi "programacao.thiagogarcia/jwt-gin/controllers/jwtAPI"
	"programacao.thiagogarcia/jwt-gin/models"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	var input LoginInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}
	u.Username = input.Username
	u.Password = input.Password

	err := u.CheckLogin()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := jwtapi.CreateToken(u.Username)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
