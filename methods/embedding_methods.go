package main

import "fmt"

type Account struct{
	name string
	description string
	account_number int
}

type Bank struct{
	name string
	description string
	routing_number int
	Account
}

func (account Account) getDiscount() string{
	return "Account level discount applied"
}

func (bank Bank) getDiscount() string{
	return "Bank Level discount applied"
}

func main() {
	a := Account{"checking", "checking account desc", 123345}
	b := Bank{"Wells Fargo", "Bank for all", 564789, a}

	fmt.Println(b.getDiscount()) // there is no method overriding in GO
	fmt.Println(b.Account.getDiscount())


}
