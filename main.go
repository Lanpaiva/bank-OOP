package main

import "fmt"

type checkingAccount struct {
	user           string
	agencyNumber   int
	accountUser    int
	balanceAccount float64
}

func (c *checkingAccount) toWithdraw(withdrawAccount float64) (string, float64) {
	getCash := withdrawAccount > 0 && withdrawAccount <= c.balanceAccount
	if getCash {
		c.balanceAccount -= withdrawAccount
		return "saque realizado com sucesso:", c.balanceAccount
	} else {
		return "Saldo insificiente:", c.balanceAccount
	}
}

func (c *checkingAccount) toDeposite(upCash float64) (string, float64) {
	depositCash := upCash > 0
	if depositCash {
		c.balanceAccount += upCash
		return "deposito realizado:", c.balanceAccount
	} else {
		return "deposito não efetuado:", c.balanceAccount
	}
}

func (c *checkingAccount) transf(transfValue float64, accountDestiny *checkingAccount) bool {
	if transfValue <= c.balanceAccount && transfValue > 0 {
		c.balanceAccount -= transfValue
		accountDestiny.toDeposite(transfValue)
		return true
	} else {
		return false
	}
}

func main() {
	silviaAccount := checkingAccount{user: "Silvia", balanceAccount: 500}
	gustavoAccount := checkingAccount{user: "Gustavo", balanceAccount: 100}

	status := silviaAccount.transf(500, &gustavoAccount)

	fmt.Println(status)

	fmt.Println(silviaAccount, gustavoAccount)

}
