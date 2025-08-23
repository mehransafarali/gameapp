package main

import (
	"GameApp/pkg/phonenumber"
	"fmt"
)

func main() {
	res, err := phonenumber.IsValid("09127642279")
	fmt.Println(res, ", ", err, "___")
}
