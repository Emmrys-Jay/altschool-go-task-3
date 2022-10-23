package util

import (
	"fmt"
	"os"
	"strings"
)

func NewLine(n int) {
	fmt.Print(strings.Repeat("\n", n))
}

func Cancel() {
	fmt.Println()
	fmt.Println("Thank you for banking with us!")
	fmt.Println("See you next time!")
	fmt.Println()
	os.Exit(0)
}

func AnotherOperation() bool {
	NewLine(1)
	var input string
	fmt.Println("Do you want to perform another operation? Y/N")
	fmt.Scan(&input)
	switch input {
	case "y", "Y", "yes", "Yes", "YES":
		return true
	default:
		return false
	}
}

func Message(msg string) {
	fmt.Print(strings.Repeat("~", 15), msg, strings.Repeat("~", 15), "\n")
	fmt.Println(strings.Repeat("~", 30+len(msg)))
}
