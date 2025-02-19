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

type PremiumAccount struct {
	Account
	Balance float64
}

func CreateAccount(id int, name string, lastName string, age uint) *Account {
	return &Account{
		Id:       id,
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

func (account *Account) UpdateName(newName string) {
	account.Name = newName
}

func CreatePremiumAccount(id int, name string, lastName string, age uint, balance float64) *PremiumAccount {
	return &PremiumAccount{
		Account: Account{
			Id:       id,
			Name:     name,
			LastName: lastName,
			Age:      age,
		},
		Balance: balance,
	}
}

func (p *PremiumAccount) Deposit(amount float64) {
	if amount > 0 {
		p.Balance += amount
		fmt.Println("Deposited: ", amount, "New Balance: ", p.Balance)
	} else {
		fmt.Println("Invalid deposit amount")
	}
}

func (p *PremiumAccount) Withdraw(amount float64) {
	if amount > 0 && amount <= p.Balance {
		p.Balance -= amount
		fmt.Println("Withdrawn: ", amount, "New Balance: ", p.Balance)
	} else {
		fmt.Println("Invalid Withdraw amount or insufficient balance")
	}
}

func (p *PremiumAccount) Info() {
	p.Account.Info()
	fmt.Println("Balance: ", p.Balance)

}
