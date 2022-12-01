package main

import (
	"fmt"
	c "github/bank-OOP/account"
)

func main() {
	silviaAccount := c.CheckingAccount{User: "Silvia", BalanceAccount: 500}
	gustavoAccount := c.CheckingAccount{User: "Gustavo", BalanceAccount: 100}

	status := silviaAccount.Transf(500, &gustavoAccount)

	fmt.Println(status)

	fmt.Println(silviaAccount, gustavoAccount)

}
