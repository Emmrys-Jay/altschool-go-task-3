package main

import (
	"fmt"
	"github.com/Emmrys-Jay/altschool-go-task-3/handler"
	"github.com/Emmrys-Jay/altschool-go-task-3/util"
)

func main() {
	util.Message("WELCOME TO THE ATM CLI APP")
	fmt.Println("Your default PIN is 1234")

	for {
		if handler.UserIsVerified() {
			menu()

			if util.AnotherOperation() {
				continue
			} else {
				util.Cancel()
			}
		} else {
			fmt.Println("Invalid PIN!")
			continue
		}

	}

}

func menu() {
	util.NewLine(1)
	fmt.Println("What operation do you want to perform? ")
	fmt.Println("Select:\n1 to Change PIN\t\t 3 to Check Account Balance")
	fmt.Println("2 to Deposit Funds\t 4 to Withdraw Funds")
	fmt.Println("\t   5 to Cancel/Exit")

	var input int
	fmt.Scan(&input)

	switch input {
	case 1:
		handler.ChangePin()
	case 2:
		handler.DepositFunds()
	case 3:
		handler.CheckBalance()
	case 4:
		handler.WithdrawFunds()
	case 5:
		util.Cancel()
	default:
		fmt.Println("Invalid Input")
	}
}
