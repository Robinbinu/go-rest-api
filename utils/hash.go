package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string)(string,error){
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(password),14)
	return string(hashedPassword),err
}

func CheckPasswordHash(hashedPassword,password string)bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
	return err == nil
}