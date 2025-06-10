package utils

import (
	"fmt"
	"os"
	"time"
)

func redirecting(msg string) {
	fmt.Print(msg)
	time.Sleep(time.Second * 1)
	clear()
}

func Auth() {
	fmt.Println("-----WELCOME-----")
	fmt.Println("Go to register if you are new or login if you have registered?\n1. Register\n2. Login")
	fmt.Scanln(&choice)
	if choice == 1 {
		clear()
		register()
	} else if choice == 2 {
		clear()
		login()
	} else {
		os.Exit(1)
	}
}
