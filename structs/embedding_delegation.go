package main

import "fmt"

type Account struct{
	account_number int
	account_desc string
}

type Bank struct{
	bank_name string
	routing_number int
	Account
}
//another way to embed structure
type Bank2 struct{
	bank_name string
	routing_number int
	account Account
}

func main() {

	account := Account{12345, "checking"}
	fmt.Println(account)

	bank := Bank{"Wells Fargo", 765890, account}
	fmt.Println(bank)
	fmt.Println(bank.account_number)
	fmt.Println(bank.account_desc)


	bank2 := Bank{"Wells Fargo", 765890, account}
	fmt.Println(bank2)
	fmt.Println(bank2.account_number)
	fmt.Println(bank2.account_desc)

}
