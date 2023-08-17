package repositories

import (
	"github.com/zzzep/pismo-challenge/src/data/domains"
	"github.com/zzzep/pismo-challenge/src/enum"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type TransactionsRepository struct {
	db *gorm.DB
}

// NewTransactionRepository creates a new instance of TransactionsRepository.
//
// It establishes a connection to the MySQL database using the enum.GetDatabaseConnection() function,
// and initializes a new gorm.DB instance using the connection.
// If there is an error in establishing the connection or initializing the gorm.DB instance,
// it logs the error and returns an empty instance of TransactionsRepository.
//
// Returns a pointer to the initialized TransactionsRepository.
func NewTransactionRepository() *TransactionsRepository {
	m := mysql.Open(enum.GetDatabaseConnection())
	db, err := gorm.Open(m, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return &TransactionsRepository{}
	}
	return &TransactionsRepository{db: db}
}

// Create creates a new transaction in the TransactionsRepository.
//
// It takes a data parameter of type domains.Transaction.
// It returns a boolean indicating whether the creation was successful.
func (a TransactionsRepository) Create(data domains.Transaction) bool {
	r := a.db.Create(&data)
	if r.Error != nil {
		log.Fatal(r.Error)
		return false
	}
	return true
}

// GetByAccount retrieves transactions by account ID.
//
// It takes an integer ID as a parameter.
// It returns a slice of domains.Transaction.
func (a TransactionsRepository) GetByAccount(id int) []domains.Transaction {
	var transactions []domains.Transaction
	_ = a.db.Where("account_id = ?", id).Find(&transactions)
	return transactions
}
