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

func NewAccount(repo *repositories.AccountsRepository) *Account {
	return &Account{repo: repo}
}

// CreateAccount
// @Summary Create new Account
// @Accept json
// @Produce json
// @Param JSON body domains.Account true "Document Number"
// @Success 200 {object} domains.Account{document_number=string}
// @Failure 500 {object} nil
// @Router /accounts [POST]
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

// GetAccount
// @Summary Create new Account
// @Accept json
// @Produce json
// @Param accountId path int true "Account ID"
// @Success 200 {object} domains.Account{account_id=int,document_number=string}
// @Failure 404 {object} map[string]any{message=string}
// @Router /accounts/{accountId} [GET]
func (a *Account) GetAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("accountId"))
	acc := a.repo.Get(id)
	if acc != nil {
		c.JSON(200, acc)
		return
	}
	c.JSON(404, gin.H{"message": "not found"})
}
