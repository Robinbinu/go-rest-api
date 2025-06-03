package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const key ="ilockikey"

func GenerateToken(email string,userId int64) (string,error){
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":email,
		"userId":userId,
		"exp":time.Now().Add(time.Hour*2).Unix(),
	})

	return token.SignedString([]byte(key))
}

func VerifyToken(token string) (int64,error){

	parsedToken,err:=jwt.Parse(token,func(t *jwt.Token) (interface{}, error) {
		//check signing method for token
		_,ok:=t.Method.(*jwt.SigningMethodHMAC)

		if !ok{
			return nil,errors.New("invalid signing method")
		}

		return []byte(key),nil
	})

	if err!=nil{
		return 0,err
	}

	isValidToken := parsedToken.Valid

	if !isValidToken{
		return 0,errors.New("token verification Failed")
	}

	claims,ok:=parsedToken.Claims.(jwt.MapClaims)
	if !ok{
		return 0,errors.New("invalid claims type in parsed token")
	}
	userId := int64(claims["userId"].(float64))


	return userId,nil
}