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
	ciclanoAccount := checkingAccount{"Ciclano", 777, 654321, 521.5}
	fmt.Println(alanAccount, ciclanoAccount)
}
