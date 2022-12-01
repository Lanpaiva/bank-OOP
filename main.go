package main

import (
	"fmt"
	"github/bank-OOP/accounts"
)

func main() {
	contEx := accounts.CheckingAccount{}
	contEx.ToDeposite(100)

	fmt.Println(contEx.GetBalance())
}
