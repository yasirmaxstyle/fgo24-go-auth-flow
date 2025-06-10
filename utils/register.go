package utils

import (
	"fmt"
	"time"
)

type user struct {
	username string
	password string
}

var dataUser []user

var newReg user

func clear() {
	fmt.Print("\033[H\033[2J")
}

func confirm() {
	var confirmPassword string
	fmt.Print("Confirm your password: ")
	fmt.Scanln(&confirmPassword)
	confirmPassword = hashPassword(confirmPassword)
	if confirmPassword != newReg.password {
		fmt.Println("Passwords do not match. Try again")
    time.Sleep(time.Millisecond * 1000)
    fmt.Print("\033[1A\033[K")
		confirm()
	}
}

func register() {
	fmt.Println("-----SIGN UP-----")
	fmt.Print("Enter your name: ")
	fmt.Scanln(&newReg.username)

	fmt.Print("Enter your password: ")
	fmt.Scanln(&newReg.password)
	newReg.password = hashPassword(newReg.password)

	confirm()

	dataUser = append(dataUser, newReg)
	clear()

	redirecting("Registartion is successful, redirecting to login page...")

	login()
}
