package respositories

import "github.com/zzzep/pismo-challenge/src/data/entities"

type TransactionsRepository struct{}

func NewTransactionRepository() TransactionsRepository {
	return TransactionsRepository{}
}
func (a TransactionsRepository) create(transaction entities.Transaction) bool {
	return true
}
