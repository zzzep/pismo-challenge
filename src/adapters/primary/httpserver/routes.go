// Package config to provides configuration of http server
package httpserver

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zzzep/pismo-challenge/src/adapters/primary"
)

// SetRoutes function to start and serve HTTP Routes
func SetRoutes(c *primary.Container) {

	c.Router.POST("/accounts", c.AccountController.CreateAccount)
	c.Router.GET("/accounts/:accountId", c.AccountController.GetAccount)
	c.Router.POST("/transactions", c.TransactionController.CreateTransaction)
	c.Router.GET("/accounts/:accountId/transactions", c.TransactionController.GetTransactionByAccount)
	c.Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success", "status": "OK"})
	})
	c.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	c.Router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
}
