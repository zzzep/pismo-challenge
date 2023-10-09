package account

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zzzep/pismo-challenge/src/adapters/secondary/interfaces"
	"github.com/zzzep/pismo-challenge/src/application/entities"
	"net/http"
	"strconv"
)

type Account struct {
	repo interfaces.IAccountsRepository
}

// NewAccount creates a new Account instance.
//
// It takes a pointer to an AccountsRepository as a parameter.
// It returns a pointer to an Account.
func NewAccount(repo interfaces.IAccountsRepository) *Account {
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
	data, dataErr := c.GetRawData()
	if dataErr != nil {
		c.JSON(http.StatusBadRequest, dataErr.Error())
		return
	}
	acc := &entities.Account{}
	jsonErr := json.Unmarshal(data, acc)
	if jsonErr != nil {
		c.JSON(http.StatusBadRequest, jsonErr.Error())
		return
	}
	if a.repo.Create(*acc) {
		c.JSON(http.StatusOK, acc)
		return
	}
	c.JSON(http.StatusInternalServerError, nil)
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
	id, idErr := strconv.Atoi(c.Param("accountId"))
	if idErr != nil {
		c.JSON(http.StatusBadRequest, idErr.Error())
		return
	}
	acc := a.repo.Get(id)
	if acc != nil {
		c.JSON(http.StatusOK, acc)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "id not found"})
}
