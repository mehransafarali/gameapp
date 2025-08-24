package main

import (
	"GameApp/entity"
	"GameApp/repository/mysql"
	"fmt"
)

func main() {

}

func TestUserMethods() {
	mysqlRepo := mysql.New()
	user, err := mysqlRepo.Register(entity.User{
		ID:          0,
		PhoneNumber: "09127644",
		Name:        "mehran safarali",
	})
	if err != nil {
		fmt.Println("can not register user: ", err)
	} else {
		fmt.Println(user)
	}

	isUnique, err := mysqlRepo.IsPhoneNumberUnique(user.PhoneNumber)
	if err != nil {
		fmt.Println("unique error: ", err)
	}
	fmt.Println("isUnique", isUnique)
}
