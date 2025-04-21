package main

import "fmt"

// "fmt"
type User struct {
	name string
}

func main() {
	user, err := getUser()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(user.name)
}
func getUser() (s User, err error) {
	s = User{"john"}
	if s.name == "" {
		return s, fmt.Errorf("no user found")
	}
	return s, nil
}
