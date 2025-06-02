package routes

import (
	"log"
	"net/http"

	"example.com/models"
	"example.com/utils"
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

func login(context *gin.Context){
	var user models.User
	err:=context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"message":"Invalid User details"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized,gin.H{"message":"User not found"})
		return
	}

	token,err:=utils.GenerateToken(user.Email,user.ID)
	if err != nil {
		log.Println("JWT generation failed",err)
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Login failed"})
		return
	}

	context.JSON(http.StatusOK,gin.H{"message":"Login successful","token":token})
}