package config

import (
	"github.com/gin-gonic/gin"
	"github.com/zzzep/pismo-challenge/src/controllers"
	"github.com/zzzep/pismo-challenge/src/data/repositories"
)

type Container struct {
	Router                *gin.Engine
	AccountController     *controllers.Account
	TransactionController *controllers.Transaction
	AccountRepo           *repositories.AccountsRepository
	TransactionRepo       *repositories.TransactionsRepository
}

// NewContainer initializes a new Container.
//
// Returns:
//     c Container: The initialized Container.
func NewContainer() (c Container) {
	c.Router = gin.Default()

	// Repositories
	c.AccountRepo = repositories.NewAccountRepository()
	c.TransactionRepo = repositories.NewTransactionRepository()

	// Controllers
	c.AccountController = controllers.NewAccount(c.AccountRepo)
	c.TransactionController = controllers.NewTransaction(c.TransactionRepo)

	return c
}
