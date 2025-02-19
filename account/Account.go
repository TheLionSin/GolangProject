package account

import "fmt"

type AccountInterface interface {
	Info()
	GetId() int
}
type Account struct {
	Id       int
	Name     string
	LastName string
	Age      uint
}

func CreateAccount(id int, name string, lastName string, age uint) *Account {
	return &Account{
		Name:     name,
		LastName: lastName,
		Age:      age,
	}
}

func (account *Account) Info() {
	fmt.Println("Account Info")
	fmt.Println(account.Name, account.LastName, account.Age)
}

func (account *Account) GetId() int {
	fmt.Print("Account Id: ")
	return account.Id
}
