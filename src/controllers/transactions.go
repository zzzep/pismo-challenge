package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zzzep/pismo-challenge/src/data/respositories"
)

type Transaction struct {
	repo respositories.TransactionsRepository
}

func NewTransaction() *Transaction {
	t := &Transaction{repo: respositories.NewTransactionRepository()}
	return t
}

func (a *Transaction) CreateTransaction(c *gin.Context) {
	
}
