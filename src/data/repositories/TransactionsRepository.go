package repositories

import (
	"github.com/zzzep/pismo-challenge/src/data/entities"
	"github.com/zzzep/pismo-challenge/src/enum"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type TransactionsRepository struct {
	db *gorm.DB
}

func NewTransactionRepository() *TransactionsRepository {
	m := mysql.Open(enum.GetDatabaseConnection())
	db, err := gorm.Open(m, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return &TransactionsRepository{}
	}
	return &TransactionsRepository{db: db}
}
func (a TransactionsRepository) create(data entities.Transaction) bool {
	r := a.db.Create(&data)
	if r.Error != nil {
		log.Fatal(r.Error)
		return false
	}
	return true
}
