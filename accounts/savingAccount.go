package accounts

import "github/bank-OOP/clients"

type SavingAccount struct {
	User           clients.Holder
	AgencyNumber   int
	AccountUser    int
	Operation      int
	balanceAccount float64
}

func (c *SavingAccount) ToWithDraw(withdrawAccount float64) string {
	getCash := withdrawAccount > 0 && withdrawAccount <= c.balanceAccount
	if getCash {
		c.balanceAccount -= withdrawAccount
		return "saque realizado com sucesso:"
	} else {
		return "Saldo insificiente:"
	}
}

func (c *SavingAccount) ToDeposite(upCash float64) (string, float64) {
	depositCash := upCash > 0
	if depositCash {
		c.balanceAccount += upCash
		return "deposito realizado:", c.balanceAccount
	} else {
		return "deposito não efetuado:", c.balanceAccount
	}
}

func (c *SavingAccount) GetBalance() float64 {
	return c.balanceAccount
}
