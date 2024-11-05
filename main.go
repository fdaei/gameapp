package main

import (
	"fmt"
	"gameapp/entity"
	"gameapp/repository/mysql"
)

func main() {
	fmt.Println("Hello, World!")
	testUserMysqlRepo()
}

func testUserMysqlRepo() {
	mysqlRepo := mysql.New()
	createdUser, err := mysqlRepo.Register(entity.User{
		ID:          0,
		PhoneNumber: "09390315707",
		Name:        "firoozeh",
	})

	if err != nil {
		fmt.Println("register user", err)
	} else {
		fmt.Println("created user", createdUser)
	}

	isUnique, err := mysqlRepo.IsPhoneNumberUnique(createdUser.PhoneNumber + "23")
	if err != nil {
		fmt.Println("check phone number unique", err)
	} else {
		fmt.Println("phone number unique?", isUnique)
	}

}
