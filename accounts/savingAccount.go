package accounts

import "bank/owners"

type SavingAccount struct {
	Owner 				owners.Owner
	NumberAgency 	int
	NumberAccount int
	Operation 		int
	balance 			float64
}

func (acc *SavingAccount) Withdraw(withdrawAmount float64) string {
	ableToWithdraw := withdrawAmount <= acc.balance

	if withdrawAmount < 0 {
		return "Valor do saque negativo! Nao foi possivel realizar o saque."
	}

	if ableToWithdraw {
		acc.balance -= withdrawAmount
		return "Saque realizado com sucesso!"
	} else {
		return "Nao foi possivel realizar o saque!"	
	}
}

func (acc *SavingAccount) Deposit(depositAmount float64) (string, float64) {
	if depositAmount < 0 {
		return "Valor do deposito negativo! Nao foi possivel realizar o deposito.", 0
	}

	acc.balance += depositAmount
	return "Deposito realizado!", acc.balance
}


func (acc *SavingAccount) GetBalance() float64 {
	return acc.balance
}