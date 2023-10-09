package helpers

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zzzep/pismo-challenge/src/adapters/secondary/interfaces"
	"github.com/zzzep/pismo-challenge/src/application/entities"
	"net/http"
	"net/http/httptest"
)

func CreatePostContext(msg ...string) *gin.Context {
	var body *bytes.Buffer
	gin.SetMode(gin.TestMode)

	// create a request to pass to the handler - don't need engine for now
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	if len(msg) == 0 {
		body = bytes.NewBufferString("{}")
	} else {
		body = bytes.NewBufferString(msg[0])
	}
	c.Request, _ = http.NewRequest("POST", "/", body)
	c.Request.Header.Add("Content-Type", gin.MIMEPOSTForm)

	return c
}

func NewMockTransaction(hasError bool) MockTransaction {
	return MockTransaction{hasError}
}

type MockTransaction struct {
	hasError bool
}

func (m MockTransaction) Create(data entities.Transaction) bool {
	if m.hasError {
		return false
	}
	return true
}

func (m MockTransaction) GetByAccount(id int) ([]entities.Transaction, error) {
	if m.hasError {
		return []entities.Transaction{}, errors.New("mock error")
	}
	return []entities.Transaction{
		{TransactionId: 1, AccountId: 1, OperationTypeId: 4, Amount: 123.45},
	}, nil
}

func NewMockAccount(hasError bool) MockAccountRepo {
	return MockAccountRepo{hasError}
}

type MockAccountRepo struct {
	hasError bool
}

func (m MockAccountRepo) Create(data entities.Account) bool {
	if m.hasError {
		return false
	}
	return true
}

func (m MockAccountRepo) Get(id int) *entities.Account {
	if m.hasError {
		return nil
	}
	return &entities.Account{AccountId: 1, DocumentNumber: "1234"}
}

type Repos struct {
	AccRepo interfaces.IAccountsRepository
	TRepo   interfaces.ITransactionsRepository
}
type GinArgs struct {
	C *gin.Context
}
type TestCases []struct {
	Name       string
	Repos      Repos
	Args       GinArgs
	Want       interface{}
	StatusCode int
}
