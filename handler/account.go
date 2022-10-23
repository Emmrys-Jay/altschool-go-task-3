package handler

import (
	"fmt"
	"github.com/Emmrys-Jay/altschool-go-task-3/util"
)

var balance float64

func DepositFunds() {
	util.Message("DEPOSIT FUNDS")
	fmt.Println("Input the amount you want to deposit: ")
	var amount float64
	fmt.Scan(&amount)

	balance += amount
	fmt.Print("\n", "You transaction was successful!", "\n")
	fmt.Println()
}

func WithdrawFunds() {
	util.Message("WITHDRAW FUNDS")
	fmt.Println("Input the amount you want to withdraw: ")
	var amount float64
	fmt.Scan(&amount)

	if amount > balance {
		fmt.Println()
		fmt.Println("Insufficient Funds! ")
		return
	}

	balance -= amount
	fmt.Printf("\nYou have successfully withdrawn #%.2f\n", amount)
	fmt.Println()
}

func CheckBalance() {
	util.Message("CHECK BALANCE")
	fmt.Printf("Your Account Balance is #%.2f\n", balance)
	fmt.Println()
}
