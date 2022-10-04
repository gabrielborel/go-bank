package main

import (
	"bank/accounts"
	"bank/owners"
	"fmt"
)

func PayBill(account VerifyACcount, billValue float64) {
	account.Withdraw(billValue)
	fmt.Println("Boleto pago com sucesso!")
}

type VerifyACcount interface {
	Withdraw(value float64) string 
}

func main() {
	gabriel := owners.Owner{Name: "Gabriel Borel", CPF: "123123123", Role: "Developer"}
	gabrielSavingAccount := accounts.SavingAccount{
		Owner: gabriel,
		Operation: 1,
		NumberAgency: 200,
		NumberAccount: 300,
	}

	gabrielCheckingAccount := accounts.CheckingAccount{
		Owner: gabriel,
		NumberAgency: 1000,
		NumberAccount: 3000,
	}

	gabrielCheckingAccount.Deposit(500)
	gabrielSavingAccount.Deposit(500)

	PayBill(&gabrielCheckingAccount, 200)
	PayBill(&gabrielSavingAccount, 100)

	fmt.Println(gabrielSavingAccount.GetBalance())
	fmt.Println(gabrielCheckingAccount.GetBalance())
}
