package transaction

import (
	"github.com/stretchr/testify/assert"
	"github.com/zzzep/pismo-challenge/src/application/entities"
	"github.com/zzzep/pismo-challenge/src/helpers"
	"net/http"
	"reflect"
	"testing"
)

func TestNewTransaction(t *testing.T) {
	accRepo := helpers.NewMockAccount(false)
	testCases := helpers.TestCases{
		{
			Name:  "New Transaction with no error on database",
			Repos: helpers.Repos{TRepo: helpers.NewMockTransaction(false)},
			Want:  NewTransaction(helpers.NewMockTransaction(false), accRepo),
		},
		{
			Name:  "New Transaction with database error",
			Repos: helpers.Repos{TRepo: helpers.NewMockTransaction(true)},
			Want:  NewTransaction(helpers.NewMockTransaction(true), accRepo),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			got := NewTransaction(tc.Repos.TRepo, accRepo)
			if !reflect.DeepEqual(got, tc.Want) {
				t.Errorf("NewTransaction() = %v, Want %v", got, tc.Want)
			}
		})
	}
}

func TestTransaction_CreateTransaction(t *testing.T) {
	testCases := helpers.TestCases{
		{
			Name:       "Create Transaction with no error on database",
			StatusCode: http.StatusOK,
			Repos:      helpers.Repos{TRepo: helpers.NewMockTransaction(false)},
			Args:       helpers.GinArgs{C: helpers.CreatePostContext()},
		},
		{
			Name:       "Create Transaction with no error and positive amount",
			StatusCode: http.StatusOK,
			Repos:      helpers.Repos{TRepo: helpers.NewMockTransaction(false)},
			Args:       helpers.GinArgs{C: helpers.CreatePostContext("{\"operation_type_id\":1,\"amount\":10}")},
		},
		{
			Name:       "Create Transaction with no error and negative amount",
			StatusCode: http.StatusOK,
			Repos:      helpers.Repos{TRepo: helpers.NewMockTransaction(false)},
			Args:       helpers.GinArgs{C: helpers.CreatePostContext("{\"operation_type_id\":4,\"amount\":-10}")},
		},
		{
			Name:       "Create Transaction with no error and positive amount and inverted operation",
			StatusCode: http.StatusOK,
			Repos:      helpers.Repos{TRepo: helpers.NewMockTransaction(false)},
			Args:       helpers.GinArgs{C: helpers.CreatePostContext("{\"operation_type_id\":4,\"amount\":10}")},
		},
		{
			Name:       "Create Transaction with no error and negative amount and inverted operation",
			StatusCode: http.StatusOK,
			Repos:      helpers.Repos{TRepo: helpers.NewMockTransaction(false)},
			Args:       helpers.GinArgs{C: helpers.CreatePostContext("{\"operation_type_id\":1,\"amount\":-10}")},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			transaction := &Transaction{
				repo:    testCase.Repos.TRepo,
				accRepo: helpers.NewMockAccount(false),
			}
			transaction.CreateTransaction(testCase.Args.C)

			assert.Equal(t, testCase.StatusCode, testCase.Args.C.Writer.Status(), testCase.Name)
		})
	}
}

func TestTransaction_GetTransactionByAccount(t *testing.T) {
	validData := helpers.CreatePostContext(`{"accountId":1}`)
	validData.AddParam("accountId", "1")
	serverErrorData := helpers.CreatePostContext(`{"accountId":1}`)
	serverErrorData.AddParam("accountId", "1")

	testCases := helpers.TestCases{
		{
			Name:       "Get Transaction By Account with valid data",
			StatusCode: http.StatusOK,
			Repos:      helpers.Repos{TRepo: helpers.NewMockTransaction(false)},
			Args:       helpers.GinArgs{C: validData},
		},
		{
			Name:       "Get Transaction By Account with database error",
			StatusCode: http.StatusInternalServerError,
			Repos:      helpers.Repos{TRepo: helpers.NewMockTransaction(true)},
			Args:       helpers.GinArgs{C: serverErrorData},
		},
		{
			Name:       "Get Transaction By Account with empty data",
			StatusCode: http.StatusBadRequest,
			Repos:      helpers.Repos{TRepo: helpers.NewMockTransaction(false)},
			Args:       helpers.GinArgs{C: helpers.CreatePostContext()},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			transaction := NewTransaction(testCase.Repos.TRepo, helpers.NewMockAccount(false))
			transaction.GetTransactionByAccount(testCase.Args.C)

			assert.Equal(t, testCase.StatusCode, testCase.Args.C.Writer.Status(), testCase.Name)
		})
	}
}

func TestTransaction_updateTransaction(t1 *testing.T) {
	t1.Run("Calculate balance - example 1", func(t1 *testing.T) {
		t := NewTransaction(helpers.NewMockTransaction(false), helpers.NewMockAccount(false))
		transactions := []entities.Transaction{
			{Amount: -50, Balance: -50},
			{Amount: -23.5, Balance: -23.5},
		}
		current := -18.7
		want := -92.2

		for i, transaction := range transactions {
			current = t.calculateTransactionBalance(&transaction, current)
			transactions[i] = transaction
		}

		balance := current
		for _, transaction := range transactions {
			balance += transaction.Balance
		}

		if balance != want {
			t1.Errorf("balance = %v, want %v", balance, want)
		}
	})
	t1.Run("Calculate balance - example 2", func(t1 *testing.T) {
		t := NewTransaction(helpers.NewMockTransaction(false), helpers.NewMockAccount(false))
		transactions := []entities.Transaction{
			{Amount: -50, Balance: -50},
			{Amount: -23.5, Balance: -23.5},
			{Amount: -18.7, Balance: -18.7},
		}
		current := 60.0
		want := -32.2

		for i, transaction := range transactions {
			current = t.calculateTransactionBalance(&transaction, current)
			transactions[i] = transaction
		}

		balance := 0.0
		for _, transaction := range transactions {
			balance += transaction.Balance
		}

		if balance != want {
			t1.Errorf("current = %v, want %v", current, want)
		}
	})
	t1.Run("Calculate balance - example 3", func(t1 *testing.T) {
		t := NewTransaction(helpers.NewMockTransaction(false), helpers.NewMockAccount(false))
		transactions := []entities.Transaction{
			{Amount: -50, Balance: 0},
			{Amount: -23.5, Balance: -13.5},
			{Amount: -18.7, Balance: -18.7},
			{Amount: 60, Balance: 0},
		}
		current := 100.0
		want := 67.8

		for i, transaction := range transactions {
			current = t.calculateTransactionBalance(&transaction, current)
			transactions[i] = transaction
		}

		balance := current
		for _, transaction := range transactions {
			balance += transaction.Balance
		}

		if balance != want {
			t1.Errorf("balance = %v, want %v", balance, want)
		}
	})
}
