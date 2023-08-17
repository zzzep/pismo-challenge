package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zzzep/pismo-challenge/src/data/respositories"
)

type Account struct {
	repo respositories.AccountsRepository
}

func NewAccount() *Account {
	a := &Account{repo: respositories.NewAccountRepository()}
	return a
}

func (a *Account) CreateAccount(c *gin.Context) {

}

func (a *Account) GetAccount(c *gin.Context) {

}
