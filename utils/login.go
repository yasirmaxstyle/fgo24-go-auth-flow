package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

var dataLogin user

var choice int

func inputChoice(message string) {
	fmt.Print(message)
	fmt.Scanln(&choice)
}

func hashPassword(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

var index int

func inputNameLogin() {
	fmt.Print("Enter your name: ")
	fmt.Scanln(&dataLogin.username)

	isRegistered := false

	for x := range dataUser {
		if dataUser[x].username == dataLogin.username {
			index = x
			isRegistered = true
		}
	}

	if !isRegistered {
		inputChoice("This name is not registered. Please go to registration or try again!\n1. Register\n2. Try again\n3. Exit\n")
		if choice == 1 {
			clear()
			register()
		} else if choice == 2 {
			clear()
			login()
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

	dataLogin.password = hashPassword(dataLogin.password)
	if dataLogin.password != dataUser[index].password {
		inputChoice("Your password is wrong. You can reset or try again!\n1. Reset\n2. Try again\n3. Exit\n")
		if choice == 1 {
			clear()
			forgot()
		} else if choice == 2 {
			fmt.Print("\033[1A\033[K")
			inputPassLogin()
		} else if choice == 3 {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}

func login() {
	fmt.Println("-----SIGN IN-----")
	inputNameLogin()

	inputPassLogin()
	clear()

	redirecting("Login is successful, redirecting to home page...")

	home()
}
