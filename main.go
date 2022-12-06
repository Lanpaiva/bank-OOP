package main

import (
	"fmt"
	"github/bank-OOP/accounts"
)

func PayTicket(account AccountVerify, ticketValue float64) {
	account.ToWithDraw(ticketValue)
}

type AccountVerify interface {
	ToWithDraw(value float64) string
}

func main() {
	contEx := accounts.CheckingAccount{}

	contEx.ToDeposite(100)
	PayTicket(&contEx, 60)

	fmt.Println(contEx.GetBalance())

	contaAlan := accounts.SavingAccount{}

	contaAlan.ToDeposite(600)
	PayTicket(&contaAlan, 70)

	fmt.Println(contaAlan)
}
