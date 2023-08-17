package repositories

import (
	"github.com/zzzep/pismo-challenge/src/data/entities"
	"github.com/zzzep/pismo-challenge/src/enum"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AccountsRepository struct {
	db *gorm.DB
}

func NewAccountRepository() AccountsRepository {
	m := mysql.Open(enum.GetDatabaseConnection())
	db, _ := gorm.Open(m, &gorm.Config{})
	return AccountsRepository{db: db}
}
func (a AccountsRepository) Create(transaction entities.Account) bool {
	a.db.Create(&transaction)
	return true
}
