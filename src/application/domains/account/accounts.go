package account

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zzzep/pismo-challenge/src/adapters/secondary/interfaces"
	"github.com/zzzep/pismo-challenge/src/application/entities"
	"strconv"
)

type Account struct {
	repo interfaces.IAccountsRepository
}

// NewAccount creates a new AccountEntity instance.
//
// It takes a pointer to an AccountsRepository as a parameter.
// It returns a pointer to an AccountEntity.
func NewAccount(repo interfaces.IAccountsRepository) *Account {
	return &Account{repo: repo}
}

// CreateAccount
// @Summary Create new AccountEntity
// @Accept json
// @Produce json
// @Param JSON body domains.AccountEntity true "Document Number"
// @Success 200 {object} domains.AccountEntity{document_number=string}
// @Failure 500 {object} nil
// @Router /accounts [POST]
func (a *Account) CreateAccount(c *gin.Context) {
	b, _ := c.GetRawData()
	acc := &entities.AccountEntity{}
	_ = json.Unmarshal(b, acc)
	if a.repo.Create(*acc) {
		c.JSON(200, acc)
		return
	}
	c.JSON(500, nil)
}

// GetAccount
// @Summary Create new AccountEntity
// @Accept json
// @Produce json
// @Param accountId path int true "AccountEntity ID"
// @Success 200 {object} domains.AccountEntity{account_id=int,document_number=string}
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
