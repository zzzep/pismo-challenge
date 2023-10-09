package primary

import (
	"github.com/gin-gonic/gin"
	"github.com/zzzep/pismo-challenge/src/adapters/secondary/interfaces"
	"github.com/zzzep/pismo-challenge/src/adapters/secondary/repositories"
	"github.com/zzzep/pismo-challenge/src/application/domains/account"
	"github.com/zzzep/pismo-challenge/src/application/domains/transaction"
)

type Container struct {
	Router                *gin.Engine
	AccountController     *account.Account
	TransactionController *transaction.Transaction
	AccountRepo           interfaces.IAccountsRepository
	TransactionRepo       interfaces.ITransactionsRepository
}

// NewContainer initializes a new Container.
//
// Returns:
//
//	c Container: The initialized Container.
func NewContainer() *Container {
	c := &Container{}

	c.Router = gin.Default()

	// Repositories
	c.AccountRepo = repositories.NewAccountRepository()
	c.TransactionRepo = repositories.NewTransactionRepository()

	// Controllers
	c.AccountController = account.NewAccount(c.AccountRepo)
	c.TransactionController = transaction.NewTransaction(c.TransactionRepo, c.AccountRepo)

	return c
}
