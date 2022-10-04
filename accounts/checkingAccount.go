package accounts

import "bank/owners"

type CheckingAccount struct {
	Owner 				owners.Owner
	NumberAgency 	int
	NumberAccount int
	balance 			float64	
}

func (acc *CheckingAccount) Withdraw(withdrawAmount float64) string {
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

func (acc *CheckingAccount) Deposit(depositAmount float64) (string, float64) {
	if depositAmount < 0 {
		return "Valor do deposito negativo! Nao foi possivel realizar o deposito.", 0
	}

	acc.balance += depositAmount
	return "Deposito realizado!", acc.balance
}

func (acc *CheckingAccount) TransferTo(targetAccount *CheckingAccount, transferAmount float64) bool {
	ableToTransfer := acc.balance >= transferAmount && transferAmount > 0
	
	if ableToTransfer {
		acc.Withdraw(transferAmount)
		targetAccount.Deposit(transferAmount)
		return true
	} else {
		return false
	}
}

func (acc *CheckingAccount) GetBalance() float64 {
	return acc.balance
}