package main

import "fmt"

type checkingAccount struct {
	user           string
	agencyNumber   int
	accountUser    int
	balanceAccount float64
}

func (c *checkingAccount) toWithdraw(withdrawAccount float64) string {
	getCash := withdrawAccount > 0 && withdrawAccount <= c.balanceAccount
	if getCash {
		c.balanceAccount -= withdrawAccount
		return "saque realizado com sucesso"
	} else {
		return "Saldo insificiente"
	}
}

func (c *checkingAccount) toDeposite(upCash float64) string {
	depositCash := upCash > 1 && upCash >= c.balanceAccount
	if depositCash {
		c.balanceAccount += upCash
		return "deposito ralizado"
	} else {
		return "deposito não efetuado"
	}
}

func main() {
	silviaAccount := checkingAccount{}
	silviaAccount.user = "Silvia"
	silviaAccount.balanceAccount = 500

	fmt.Println(silviaAccount.balanceAccount)

	fmt.Println(silviaAccount.toDeposite(600))

	fmt.Println("Saldo atual:", silviaAccount.balanceAccount)
}
