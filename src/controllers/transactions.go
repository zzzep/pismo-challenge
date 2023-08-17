package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zzzep/pismo-challenge/src/data/repositories"
)

type Transaction struct {
	repo *repositories.TransactionsRepository
}

func NewTransaction() *Transaction {
	t := &Transaction{repo: repositories.NewTransactionRepository()}
	return t
}

func (a *Transaction) CreateTransaction(c *gin.Context) {

}
