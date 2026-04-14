package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(c *customer,t transaction)error{
	trans := t.transactionType
	if trans != transactionDeposit && trans!=transactionWithdrawal {
		return  errors.New("unknown transaction type")
	}else if (trans == transactionWithdrawal) && (c.balance-t.amount>=0)  {
		c.balance-=t.amount
		return nil
	}else if trans == transactionDeposit {
		c.balance += t.amount
		return  nil
	} else {
		return errors.New("insufficient funds")
	}

}
