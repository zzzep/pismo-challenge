package respositories

import (
	"github.com/zzzep/pismo-challenge/src/data/entities"
)

type AccountsRepository struct {
	db gorm.Model
}

func NewAccountRepository() AccountsRepository {
	return AccountsRepository{}
}
func (a AccountsRepository) create(transaction entities.Account) bool {
	return true
}
