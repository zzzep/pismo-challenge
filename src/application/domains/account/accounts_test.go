package account

import (
	"github.com/gin-gonic/gin"
	"github.com/zzzep/pismo-challenge/src/adapters/secondary/interfaces"
	"github.com/zzzep/pismo-challenge/src/application/entities"
	"github.com/zzzep/pismo-challenge/src/helpers"
	"reflect"
	"testing"
)

func newMockAccount(hasError bool) MockAccountRepo {
	return MockAccountRepo{hasError}
}

type MockAccountRepo struct {
	hasError bool
}

func (m MockAccountRepo) Create(data entities.AccountEntity) bool {
	if m.hasError {
		return false
	}
	return true
}

func (m MockAccountRepo) Get(id int) *entities.AccountEntity {
	if m.hasError {
		return nil
	}
	return &entities.AccountEntity{AccountId: 1, DocumentNumber: "1234"}
}

func TestAccount_CreateAccount(t *testing.T) {
	type fields struct {
		repo interfaces.IAccountsRepository
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "New AccountEntity", fields: fields{repo: newMockAccount(false)}, args: args{c: helpers.CreatePostContext()}},
		{name: "New AccountEntity", fields: fields{repo: newMockAccount(true)}, args: args{c: helpers.CreatePostContext()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				repo: tt.fields.repo,
			}
			a.CreateAccount(tt.args.c)
		})
	}
}

func TestAccount_GetAccount(t *testing.T) {
	type fields struct {
		repo interfaces.IAccountsRepository
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Get AccountEntity", fields: fields{repo: newMockAccount(false)}, args: args{c: helpers.CreatePostContext()}},
		{name: "Get AccountEntity", fields: fields{repo: newMockAccount(true)}, args: args{c: helpers.CreatePostContext()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				repo: tt.fields.repo,
			}
			a.GetAccount(tt.args.c)
		})
	}
}

func TestNewAccount(t *testing.T) {
	type args struct {
		repo interfaces.IAccountsRepository
	}
	tests := []struct {
		name string
		args args
		want *Account
	}{
		{name: "New AccountEntity", args: args{repo: newMockAccount(false)}, want: NewAccount(newMockAccount(false))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccount(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
