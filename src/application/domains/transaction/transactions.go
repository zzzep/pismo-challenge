package transaction

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zzzep/pismo-challenge/src/adapters/secondary/interfaces"
	"github.com/zzzep/pismo-challenge/src/application/entities"
	"net/http"
	"strconv"
)

type Transaction struct {
	repo    interfaces.ITransactionsRepository
	accRepo interfaces.IAccountsRepository
}

// NewTransaction creates a new Transaction object.
//
// It takes in a TransactionsRepository object as a parameter and returns a pointer to a Transaction object.
func NewTransaction(repo interfaces.ITransactionsRepository, accRepo interfaces.IAccountsRepository) *Transaction {
	return &Transaction{repo: repo, accRepo: accRepo}
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
	transaction, err := t.fetchTransactionFromBody(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if t.accRepo.Get(transaction.AccountId) == nil {
		c.JSON(http.StatusNotFound, "account not found")
		return
	}

	t.checkValue(transaction)

	transaction.Balance = transaction.Amount
	transaction.Balance = t.calculateBalance(transaction)
	c.JSON(http.StatusOK, transaction)
}

func (t *Transaction) calculateBalance(e *entities.Transaction) float64 {
	current := e.Balance
	transactions, _ := t.repo.GetUnpaidBalanceByAccount(e.AccountId)
	for _, transaction := range transactions {
		if transaction.TransactionId == e.TransactionId {
			continue
		}
		current = t.calculateTransactionBalance(&transaction, current)
		tErr := t.repo.Update(transaction)
		if tErr != nil {
			panic(tErr)
		}
	}
	e.Balance = current
	err := t.repo.Update(*e)
	if err != nil {
		panic(err)
	}
	return current
}

func (t *Transaction) calculateTransactionBalance(transaction *entities.Transaction, balance float64) float64 {
	if transaction.Balance == 0 {
		return balance
	}

	if (balance > 0 && transaction.Balance > 0) || (balance < 0 && transaction.Balance < 0) {
		return balance
	}

	if (balance > 0 && balance >= -transaction.Balance) || (balance < 0 && balance >= transaction.Balance) {
		balance += transaction.Balance
		transaction.Balance = 0
		return balance
	}

	if balance > 0 {
		transaction.Balance = transaction.Balance + balance
	} else {
		transaction.Balance = transaction.Balance - balance
	}
	balance = 0

	return balance
}

// fetchTransactionFromBody fetches a transaction from the request body.
//
// It takes a gin.Context as a parameter and returns a pointer to entities.Transaction and an error.
func (t *Transaction) fetchTransactionFromBody(c *gin.Context) (*entities.Transaction, error) {
	data, dataErr := c.GetRawData()
	if dataErr != nil {
		return nil, dataErr
	}

	transaction := &entities.Transaction{}
	jsonErr := json.Unmarshal(data, transaction)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return transaction, nil
}

// checkValue updates the transaction amount based on the operation type.
//
// It takes a pointer to a Transaction struct as the parameter.
// There is no return value.
func (t *Transaction) checkValue(transaction *entities.Transaction) {
	if transaction.OperationTypeId == entities.Payment && transaction.Amount < 0 {
		transaction.Amount = transaction.Amount * -1
	}
	if transaction.OperationTypeId != entities.Payment && transaction.Amount > 0 {
		transaction.Amount = transaction.Amount * -1
	}
}

// GetTransactionByAccount
// @Summary List Transaction by Account
// @Accept json
// @Produce json
// @Param accountId path int true "Account ID"
// @Success 200 {object} []domains.Transaction{account_id=int,document_number=string}
// @Failure 404 {object} map[string]any{message=string}
// @Router /accounts/{accountId}/transactions [GET]
// GetTransactionByAccount retrieves transactions by account ID.
func (t *Transaction) GetTransactionByAccount(c *gin.Context) {
	accountId, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	transactions, dbErr := t.repo.GetByAccount(accountId)
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, dbErr.Error())
		return
	}

	if len(transactions) > 0 {
		c.JSON(http.StatusOK, transactions)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
}
