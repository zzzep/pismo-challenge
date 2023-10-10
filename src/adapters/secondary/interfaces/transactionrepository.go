package interfaces

import (
	"github.com/zzzep/pismo-challenge/src/application/entities"
	"gorm.io/gorm"
)

type ITransactionsRepository interface {
	GetDb() *gorm.DB
	Create(data entities.Transaction) bool
	GetByAccount(id int) ([]entities.Transaction, error)
	GetUnpaidBalanceByAccount(id int) ([]entities.Transaction, error)
	Update(data entities.Transaction) error
}
