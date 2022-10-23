package handler

import (
	"fmt"
	"github.com/Emmrys-Jay/altschool-go-task-3/util"
)

// userPin is default user pin
var userPin = 1234

func ChangePin() {
	util.Message("CHANGE PIN")
	fmt.Println("Enter your new PIN: ")
	var pin int
	fmt.Scan(&pin)

	userPin = pin
	fmt.Println("Your PIN has been successfully changed!")
	fmt.Println()
}

func UserIsVerified() bool {
	fmt.Println("Please insert your pin: ")
	var pin int
	fmt.Scan(&pin)

	if pin == userPin {
		return true
	}
	return false
}
