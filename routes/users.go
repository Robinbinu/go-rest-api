package routes

import (
	"net/http"

	"example.com/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err:=context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"message":"Invalid User details"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Cannot save user"})
		return
	}
	
	context.JSON(http.StatusOK,gin.H{"message":"sign up successfull","user":user})

}