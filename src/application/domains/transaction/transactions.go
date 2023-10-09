package transaction

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zzzep/pismo-challenge/src/adapters/secondary/interfaces"
	entity2 "github.com/zzzep/pismo-challenge/src/application/entities"
	"strconv"
)

type Transaction struct {
	repo interfaces.ITransactionsRepository
}

// NewTransaction creates a new TransactionEntity object.
//
// It takes in a TransactionsRepository object as a parameter and returns a pointer to a TransactionEntity object.
func NewTransaction(repo interfaces.ITransactionsRepository) *Transaction {
	return &Transaction{repo: repo}
}

// CreateTransaction
// @Summary Create new TransactionEntity
// @Accept json
// @Produce json
// @Param JSON body domains.TransactionEntity true "TransactionEntity"
// @Success 200 {object} domains.TransactionEntity{}
// @Failure 500 {object} nil
// @Router /transactions [POST]
func (t *Transaction) CreateTransaction(c *gin.Context) {
	b, _ := c.GetRawData()
	acc := &entity2.TransactionEntity{}
	_ = json.Unmarshal(b, acc)
	if acc.OperationTypeId == entity2.Payment && acc.Amount < 0 {
		acc.Amount = acc.Amount * -1
	}
	if acc.OperationTypeId != entity2.Payment && acc.Amount > 0 {
		acc.Amount = acc.Amount * -1
	}
	if t.repo.Create(*acc) {
		c.JSON(200, acc)
		return
	}
	c.JSON(500, nil)
}

// GetTransactionByAccount
// @Summary List TransactionEntity by Account
// @Accept json
// @Produce json
// @Param accountId path int true "Account ID"
// @Success 200 {object} []domains.TransactionEntity{account_id=int,document_number=string}
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
