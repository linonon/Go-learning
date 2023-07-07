package main

import "fmt"

type User struct {
	ID int
}

func main() {
	var user = &User{}
	user.ID = 1
	fmt.Println(user.ID)

	user = nil
	fmt.Println(user.ID)

}
