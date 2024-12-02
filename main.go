package main

import (
	"fmt"
	"game/entity"
	"game/repository/mysql"
)


func main() {
	
}

func testUserMysqlRepo() {
	mysqlRepo := mysql.New()

	createdUser, err := mysqlRepo.Register(entity.User{
		ID: 0,
		PhoneNumber: "0914444",
		Name: "hossein",
	})
	

	if err != nil {
		fmt.Println("created user is failed", err)
	} else {
		fmt.Println("created user is successful", createdUser)
	}

	isUnique, err :=  mysqlRepo.IsPhoneNumberUnique(createdUser.PhoneNumber)

	if err != nil {
		fmt.Println("unique err", err)
	}

	fmt.Println("isUnique", isUnique)
}