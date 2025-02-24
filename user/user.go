package user

import (
	"fmt"
	"time"
)

var idUsers = make(map[int]string)

func clearTerminal() {
	time.Sleep(2 * time.Second)
	fmt.Print("\033[H\033[2J")
}

func SelectOption() {
	for i := 0; i < 10000; i++ {
		var selectOption int

		fmt.Print("1 - Create User \n2 - Search User \n3 - Exit\nSelect Option: ")
		fmt.Scan(&selectOption)

		if selectOption == 1 {
			createUser()
			fmt.Println("Account created!")
			clearTerminal()
		} else if selectOption == 2 {
			GetUser()
			clearTerminal()
		} else if selectOption == 3 {
			break
		} else {
			fmt.Println("Option Invalid!")
		}
	}
}

func createUser() {
	// Variables
	var nameUser string
	var idUser int

	fmt.Print("Name of your User: ")
	_, _ = fmt.Scan(&nameUser)

	// ID for User at list
	idUser = len(idUsers) + 1

	idUsers[idUser] = nameUser
}

func GetUser() {
	// Variables
	var idUser int

	fmt.Print("Id User: ")
	_, _ = fmt.Scan(&idUser)

	if name, found := idUsers[idUser]; found {
		fmt.Println(name)
	} else {
		fmt.Println("ID don't Exist!")
	}
}
