package user

import "fmt"

func GetUser() {
	var idUser int
	idUsers := map[int]string{
		1:  "Noah",
		2:  "Lanh",
		3:  "James",
		4:  "Carl",
		5:  "John",
		6:  "James",
		7:  "Olivia",
		8:  "Sophia",
		9:  "Emma",
		10: "Jack",
	}

	fmt.Print("Id User: ")
	_, _ = fmt.Scan(&idUser)

	if name, found := idUsers[idUser]; found {
		fmt.Println(name)
	} else {
		fmt.Println("ID don't Exist!")
	}
}
