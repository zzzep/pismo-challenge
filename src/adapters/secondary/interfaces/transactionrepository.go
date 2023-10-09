package interfaces

import (
	"github.com/zzzep/pismo-challenge/src/application/entities"
)

type ITransactionsRepository interface {
	Create(data entities.TransactionEntity) bool
	GetByAccount(id int) []entities.TransactionEntity
}
