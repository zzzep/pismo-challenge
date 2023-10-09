package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/zzzep/pismo-challenge/src/adapters/secondary/interfaces"
	"github.com/zzzep/pismo-challenge/src/application/entities"
	"github.com/zzzep/pismo-challenge/src/helpers"
	"reflect"
	"testing"
)

func newMockTransaction(hasError bool) MockTransaction {
	return MockTransaction{hasError}
}

type MockTransaction struct {
	hasError bool
}

func (m MockTransaction) Create(data entities.TransactionEntity) bool {
	if m.hasError {
		return false
	}
	return true
}

func (m MockTransaction) GetByAccount(id int) []entities.TransactionEntity {
	if m.hasError {
		return []entities.TransactionEntity{}
	}
	return []entities.TransactionEntity{
		{TransactionId: 1, AccountId: 1, OperationTypeId: 4, Amount: 123.45},
	}
}

func TestNewTransaction(t *testing.T) {
	type args struct {
		repo interfaces.ITransactionsRepository
	}
	tests := []struct {
		name string
		args args
		want *Transaction
	}{
		{name: "New TransactionEntity", args: args{repo: newMockTransaction(false)}, want: NewTransaction(newMockTransaction(false))},
		{name: "New TransactionEntity", args: args{repo: newMockTransaction(true)}, want: NewTransaction(newMockTransaction(true))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTransaction(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_CreateTransaction(t1 *testing.T) {
	type fields struct {
		repo interfaces.ITransactionsRepository
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Create TransactionEntity", fields: fields{repo: newMockTransaction(false)}, args: args{c: helpers.CreatePostContext()}},
		{name: "Create TransactionEntity", fields: fields{repo: newMockTransaction(true)}, args: args{c: helpers.CreatePostContext()}},
		{name: "Create TransactionEntity", fields: fields{repo: newMockTransaction(false)}, args: args{c: helpers.CreatePostContext("{\"operation_type_id\":1,\"amount\":10}")}},
		{name: "Create TransactionEntity", fields: fields{repo: newMockTransaction(false)}, args: args{c: helpers.CreatePostContext("{\"operation_type_id\":4,\"amount\":-10}")}},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transaction{
				repo: tt.fields.repo,
			}
			t.CreateTransaction(tt.args.c)
		})
	}
}

func TestTransaction_GetTransactionByAccount(t1 *testing.T) {
	type fields struct {
		repo interfaces.ITransactionsRepository
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Get TransactionEntity By AccountEntity", fields: fields{repo: newMockTransaction(false)}, args: args{c: helpers.CreatePostContext()}},
		{name: "Get TransactionEntity By AccountEntity", fields: fields{repo: newMockTransaction(true)}, args: args{c: helpers.CreatePostContext()}},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transaction{
				repo: tt.fields.repo,
			}
			t.GetTransactionByAccount(tt.args.c)
		})
	}
}
