package utils

import (
	"fmt"
	"os"
)

var dataLogin user

var choice int

func inputChoice(message string) {
	fmt.Print(message)
	fmt.Scanln(&choice)
}

func inputNameLogin() {
	fmt.Print("Enter your name: ")
	fmt.Scanln(&dataLogin.username)
	if dataLogin.username != dataUser.username {
		inputChoice("This name is not registered. Please go to registration or try again!\n1. Register\n2. Try again\n3. Exit\n")
		if choice == 1 {
			Auth()
		} else if choice == 2 {
			inputNameLogin()
		} else if choice == 3 {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}

func inputPassLogin() {
	fmt.Print("Enter your password: ")
	fmt.Scanln(&dataLogin.password)
	if dataLogin.password != dataUser.password {
		inputChoice("Your password is wrong. You can reset or try again!\n1. Reset\n2. Try again\n3. Exit\n")
		if choice == 1 {
			forgot()
		} else if choice == 2 {
			inputPassLogin()
		} else if choice == 3 {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}

func login() {
	inputNameLogin()
	inputPassLogin()
}
