package accounts

import "github/bank-OOP/clients"

type CheckingAccount struct {
	User           clients.Holder
	AgencyNumber   int
	AccountUser    int
	balanceAccount float64
}

func (c *CheckingAccount) ToWithdraw(withdrawAccount float64) (string, float64) {
	getCash := withdrawAccount > 0 && withdrawAccount <= c.balanceAccount
	if getCash {
		c.balanceAccount -= withdrawAccount
		return "saque realizado com sucesso:", c.balanceAccount
	} else {
		return "Saldo insificiente:", c.balanceAccount
	}
}

func (c *CheckingAccount) ToDeposite(upCash float64) (string, float64) {
	depositCash := upCash > 0
	if depositCash {
		c.balanceAccount += upCash
		return "deposito realizado:", c.balanceAccount
	} else {
		return "deposito não efetuado:", c.balanceAccount
	}
}

func (c *CheckingAccount) Transf(transfValue float64, accountDestiny *CheckingAccount) bool {
	if transfValue <= c.balanceAccount && transfValue > 0 {
		c.balanceAccount -= transfValue
		accountDestiny.ToDeposite(transfValue)
		return true
	} else {
		return false
	}
}

func (c *CheckingAccount) GetBalance() float64 {
	return c.balanceAccount
}
