package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Nome  string `json:"Nome"`
	Valor int    `json:"Valor"`
	Conta *Conta
}
type Conta struct {
	Dono  *User
	Value int
}

var users = []User{
	{Nome: "mario", Valor: 20},
	{Nome: "jonas", Valor: 100},
	{Nome: "joao", Valor: 50},
}

func main() {

	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", postUsers)
	router.GET("/users/:valor", getUserByValor)
	router.Run("localhost:8080")

}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
func getUserByValor(c *gin.Context) {
	valor := c.Param("valor")

	for _, user := range users {
		if strconv.Itoa(user.Valor) == valor {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
func postUsers(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)

}
