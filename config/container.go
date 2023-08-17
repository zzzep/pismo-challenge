package config

import (
	"github.com/gin-gonic/gin"
	"github.com/zzzep/pismo-challenge/src/controllers"
)

type Container struct {
	Router                *gin.Engine
	AccountController     *controllers.Account
	TransactionController *controllers.Transaction
}

func NewContainer() (c Container) {
	c.Router = gin.Default()
	c.AccountController = controllers.NewAccount()
	c.TransactionController = controllers.NewTransaction()
	return c
}
