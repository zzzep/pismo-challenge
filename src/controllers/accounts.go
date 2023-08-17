package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zzzep/pismo-challenge/src/data/domains"
	"github.com/zzzep/pismo-challenge/src/data/repositories"
	"strconv"
)

type Account struct {
	repo *repositories.AccountsRepository
}

func NewAccount() *Account {
	a := &Account{repo: repositories.NewAccountRepository()}
	return a
}

func (a *Account) CreateAccount(c *gin.Context) {
	b, _ := c.GetRawData()
	acc := &domains.Account{}
	_ = json.Unmarshal(b, acc)
	if a.repo.Create(*acc) {
		c.JSON(200, acc)
		return
	}
	c.JSON(500, nil)
}

func (a *Account) GetAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("accountId"))
	acc := a.repo.Get(id)
	if acc != nil {
		c.JSON(200, acc)
		return
	}
	c.JSON(404, gin.H{"message": "not found"})
}
