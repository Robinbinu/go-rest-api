package models

import (
	"errors"

	"example.com/db"
	"example.com/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error{
	query := `INSERT INTO users(email,password) VALUES(?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	//auto close after func
	defer stmt.Close()
	hashedPassword,err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email,hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id
	return nil
}

func (u User) ValidateCredentials()error{
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query,u.Email)
	var retrievedPassword string
	err:=row.Scan(&retrievedPassword)
	if err != nil {
		return errors.New("Invalid user")
	}

	isValidPassword := utils.CheckPasswordHash(retrievedPassword,u.Password)

	if !isValidPassword{
		return errors.New("Invalid user")
	}

	return nil

}
