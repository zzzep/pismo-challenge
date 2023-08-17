// Package config to provides configuration of http server
package config

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetRoutes function to start and serve HTTP Routes
func (c *Container) SetRoutes() {

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
