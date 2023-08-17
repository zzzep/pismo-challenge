package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zzzep/pismo-challenge/src/data/domains"
	"github.com/zzzep/pismo-challenge/src/data/repositories"
	"strconv"
)

type Transaction struct {
	repo repositories.ITransactionsRepository
}

// NewTransaction creates a new Transaction object.
//
// It takes in a TransactionsRepository object as a parameter and returns a pointer to a Transaction object.
func NewTransaction(repo repositories.ITransactionsRepository) *Transaction {
	return &Transaction{repo: repo}
}

// CreateTransaction
// @Summary Create new Transaction
// @Accept json
// @Produce json
// @Param JSON body domains.Transaction true "Transaction"
// @Success 200 {object} domains.Transaction{}
// @Failure 500 {object} nil
// @Router /transactions [POST]
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

// GetTransactionByAccount
// @Summary List Transaction by Account
// @Accept json
// @Produce json
// @Param accountId path int true "Account ID"
// @Success 200 {object} []domains.Transaction{account_id=int,document_number=string}
// @Failure 404 {object} map[string]any{message=string}
// @Router /accounts/{accountId}/transactions [GET]
func (t *Transaction) GetTransactionByAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("accountId"))
	transactions := t.repo.GetByAccount(id)
	if len(transactions) > 0 {
		c.JSON(200, transactions)
		return
	}
	c.JSON(404, gin.H{"message": "not found"})
}
