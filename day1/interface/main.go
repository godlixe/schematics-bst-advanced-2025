package main

import "fmt"

type Database interface {
	GetData(int) (string, error)
	CreateData(string) (string, error)
	UpdateData(int, string) (string, error)
	DeleteData(int) error
}

type Controller struct {
	db Database
}

func NewController(db Database) Controller {
	return Controller{
		db: db,
	}
}

func main() {
	mongodb := NewMongoAdapter()
	controller := NewController(mongodb)

	res, err := controller.db.CreateData("test")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	res, err = controller.db.GetData(1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

}
