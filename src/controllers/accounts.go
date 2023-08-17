package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zzzep/pismo-challenge/src/data/entities"
	"github.com/zzzep/pismo-challenge/src/data/repositories"
)

type Account struct {
	repo repositories.AccountsRepository
}

func NewAccount() *Account {
	a := &Account{repo: repositories.NewAccountRepository()}
	return a
}

func (a *Account) CreateAccount(c *gin.Context) {
	b, _ := c.GetRawData()
	acc := &entities.Account{}
	_ = json.Unmarshal(b, acc)
	a.repo.Create(*acc)
	c.JSON(200, gin.H{"message": "success", "status": "OK"})
}

func (a *Account) GetAccount(c *gin.Context) {

}
