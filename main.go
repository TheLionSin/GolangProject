package main

import (
	"GolangProject/account"
	"fmt"
)

func main() {
	fmt.Println()
	user1 := account.CreateAccount(1, "Valodya", "Vasilevich", 25)
	user2 := account.CreateAccount(2, "Alesha", "Medvedev", 24)
	accInterface := account.AccountInterface(user1)
	fmt.Println(accInterface.GetId())
	accInterface2 := account.AccountInterface(user2)
	accInterface2.Info()
}
