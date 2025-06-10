package utils

import (
	"fmt"
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

func register() {
	fmt.Println("-----SIGN UP-----")
	fmt.Print("Enter your name: ")
	fmt.Scanln(&newReg.username)

	fmt.Print("Enter your password: ")
	fmt.Scanln(&newReg.password)
	newReg.password = hashPassword(newReg.password)

	dataUser = append(dataUser, newReg)
	clear()

	redirecting("Registartion is successful, redirecting to login page...")

	login()
}
