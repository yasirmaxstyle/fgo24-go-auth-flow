package utils

import "fmt"

type user struct {
	username string
	password string
}

var dataUser user

func register() {
	fmt.Print("Enter your name: ")
	fmt.Scanln(&dataUser.username)

  fmt.Print("Enter your password: ")
	fmt.Scanln(&dataUser.password)
}
