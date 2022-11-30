package main

import "fmt"

type checkingAccount struct {
	user           string
	agencyNumber   int
	accountUser    int
	balanceAccount float64
}

func main() {
	alanAccount := checkingAccount{user: "Alan", agencyNumber: 589, accountUser: 123456, balanceAccount: 125.5}
	fulanoAccount := checkingAccount{user: "Fulano", agencyNumber: 555, accountUser: 111222, balanceAccount: 200.5}
	fmt.Println(alanAccount, fulanoAccount)
}
