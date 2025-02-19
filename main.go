package main

import (
	"GolangProject/account"
	"fmt"
)

func main() {
	fmt.Println()
	user1 := account.CreateAccount(1, "Valodya", "Vasilevich", 25)
	user2 := account.CreateAccount(2, "Alesha", "Medvedev", 24)
	user1.UpdateName("Vladimir")
	accInterface := account.AccountInterface(user1)
	fmt.Println(accInterface.GetId())
	accInterface2 := account.AccountInterface(user2)
	accInterface2.Info()

	premAcc := account.CreatePremiumAccount(3, "Stas", "Blagov", 29, 25000.00)
	premInterface := account.AccountInterface(premAcc)

	fmt.Println(premInterface.GetId())
	premInterface.Info()
	premAcc.Deposit(5000)
	premAcc.Withdraw(1000)
	premInterface.Info()
}
