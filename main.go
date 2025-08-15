package main

import (
	"GameApp/service/userservice"
	"fmt"
)

func main() {
	res, err := userservice.IsPhoneValid("091")
	fmt.Println(res, ", ", err, "___")
}
