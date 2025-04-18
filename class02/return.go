package main

import (
	"fmt"
)

func main() {

	user1, user2 := getUsers()

	fmt.Printf("%s %s\n", user1, user2)

}
func getUsers() (string, string) {
	return "John", "Pork"
}
