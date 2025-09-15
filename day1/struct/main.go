package main

import (
	"encoding/json"
	"fmt"
)

// TODO: Add methods
type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}

func main() {
	u := User{
		Email:    "alice@example.com",
		Username: "alice123",
		Name:     "Alice Doe",
		Age:      25,
	}

	// userA := User{
	// 	Email:    "a@a.local",
	// 	Username: "a",
	// 	Name:     "a",
	// 	Age:      1,
	// }
	// fmt.Println(userA)

	jsonData, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Println("Struct to JSON:")
	fmt.Println(string(jsonData))

	input := `{"email":"bob@example.com","username":"bob99","name":"Bob Smith","age":30}`
	var u2 User
	err = json.Unmarshal([]byte(input), &u2)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nJSON to Struct:")
	fmt.Printf("%+v\n", u2)
}
