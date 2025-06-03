package middlewares

import (
	"log"
	"net/http"

	"example.com/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context){
	//jwt token verfication
	token := context.Request.Header.Get("Authorization")
	if token ==""{
		log.Println("token is null")
		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message":"Not authorized"})
		return
	}

	userId,err:=utils.VerifyToken(token)
	
	if err != nil {
		log.Println(err)
		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message":"Not authorized"})
		return
	}

	//push userId data to context
	context.Set("userID",userId)
	log.Println(userId)
	//execute the next request handler is exists
	context.Next()

}