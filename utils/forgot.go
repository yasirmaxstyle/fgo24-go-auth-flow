package utils

import (
	"fmt"
)

var newPassword string

func forgot() {
	fmt.Print("Enter your new password: ")
	fmt.Scanln(&newPassword)

	newPassword = hashPassword(newPassword)
	dataUser[index].password = newPassword

	redirecting("Reset password is successfull. Redirecting to login page...")

	login()
}
