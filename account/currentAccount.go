package account

type CheckingAccount struct {
	User           string
	AgencyNumber   int
	AccountUser    int
	BalanceAccount float64
}

func (c *CheckingAccount) toWithdraw(withdrawAccount float64) (string, float64) {
	getCash := withdrawAccount > 0 && withdrawAccount <= c.BalanceAccount
	if getCash {
		c.BalanceAccount -= withdrawAccount
		return "saque realizado com sucesso:", c.BalanceAccount
	} else {
		return "Saldo insificiente:", c.BalanceAccount
	}
}

func (c *CheckingAccount) toDeposite(upCash float64) (string, float64) {
	depositCash := upCash > 0
	if depositCash {
		c.BalanceAccount += upCash
		return "deposito realizado:", c.BalanceAccount
	} else {
		return "deposito não efetuado:", c.BalanceAccount
	}
}

func (c *CheckingAccount) Transf(transfValue float64, accountDestiny *CheckingAccount) bool {
	if transfValue <= c.BalanceAccount && transfValue > 0 {
		c.BalanceAccount -= transfValue
		accountDestiny.toDeposite(transfValue)
		return true
	} else {
		return false
	}
}
