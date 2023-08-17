package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zzzep/pismo-challenge/src/data/domains"
	"github.com/zzzep/pismo-challenge/src/data/repositories"
	"strconv"
)

type Transaction struct {
	repo *repositories.TransactionsRepository
}

func NewTransaction(repo *repositories.TransactionsRepository) *Transaction {
	return &Transaction{repo: repo}
}

func (t *Transaction) CreateTransaction(c *gin.Context) {
	b, _ := c.GetRawData()
	acc := &domains.Transaction{}
	_ = json.Unmarshal(b, acc)
	if acc.OperationTypeId == domains.Payment && acc.Amount < 0 {
		acc.Amount = acc.Amount * -1
	}
	if acc.OperationTypeId != domains.Payment && acc.Amount > 0 {
		acc.Amount = acc.Amount * -1
	}
	if t.repo.Create(*acc) {
		c.JSON(200, acc)
		return
	}
	c.JSON(500, nil)
}

func (t *Transaction) GetTransactionByAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("accountId"))
	transactions := t.repo.GetByAccount(id)
	if len(transactions) > 0 {
		c.JSON(200, transactions)
		return
	}
	c.JSON(404, gin.H{"message": "not found"})
}
