package utils

import "fmt"

var newpassword string

func forgot() {
	fmt.Print("Enter your new password: ")
	fmt.Scanln(&newpassword)
}
