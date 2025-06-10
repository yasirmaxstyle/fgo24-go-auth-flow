package utils

import (
	"fmt"
	"os"
)

func showList() {
	for x := range dataUser {
		fmt.Printf("-----USER %d------\n", x)
		fmt.Printf("username: %s\n", dataUser[x].username)
		fmt.Printf("password: %s\n", dataUser[x].password)
		fmt.Println("-----------------")

	}
	fmt.Println("1. Back to home\n\n0. Exit")
	fmt.Scanln(&choice)
	if choice == 1 {
		clear()
		home()
	} else if choice == 0 {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func home() {
	fmt.Println("-----HOME-----")
	fmt.Println("1. List User\n2. Logout\n\n0. Exit")
	fmt.Scanln(&choice)
	if choice == 1 {
		showList()
	} else if choice == 2 {
		redirecting("Logging out...")
		Auth()
	} else if choice == 0 {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
