package controllers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/zzzep/pismo-challenge/src/data/domains"
	"github.com/zzzep/pismo-challenge/src/data/repositories"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func newMockAccount(hasError bool) MockAccountRepo {
	return MockAccountRepo{hasError}
}

type MockAccountRepo struct {
	hasError bool
}

func (m MockAccountRepo) Create(data domains.Account) bool {
	if m.hasError {
		return false
	}
	return true
}

func (m MockAccountRepo) Get(id int) *domains.Account {
	if m.hasError {
		return nil
	}
	return &domains.Account{AccountId: 1, DocumentNumber: "1234"}
}

func TestAccount_CreateAccount(t *testing.T) {
	type fields struct {
		repo repositories.IAccountsRepository
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "New Account", fields: fields{repo: newMockAccount(false)}, args: args{c: context()}},
		{name: "New Account", fields: fields{repo: newMockAccount(true)}, args: args{c: context()}},
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
		repo repositories.IAccountsRepository
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Get Account", fields: fields{repo: newMockAccount(false)}, args: args{c: context()}},
		{name: "Get Account", fields: fields{repo: newMockAccount(true)}, args: args{c: context()}},
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
		repo repositories.IAccountsRepository
	}
	tests := []struct {
		name string
		args args
		want *Account
	}{
		{name: "New Account", args: args{repo: newMockAccount(false)}, want: NewAccount(newMockAccount(false))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccount(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func context(msg ...string) *gin.Context {
	var body *bytes.Buffer
	gin.SetMode(gin.TestMode)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	if len(msg) == 0 {
		body = bytes.NewBufferString("Fetch binary post data")
	} else {
		body = bytes.NewBufferString(msg[0])
	}
	c.Request, _ = http.NewRequest("POST", "/", body)
	c.Request.Header.Add("Content-Type", gin.MIMEPOSTForm)

	return c
}
