package repositories

import (
	"github.com/zzzep/pismo-challenge/src/application/entities"
	"github.com/zzzep/pismo-challenge/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type TransactionsRepository struct {
	Db *gorm.DB
}

func (a TransactionsRepository) GetDb() *gorm.DB {
	return a.Db
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
	m := mysql.Open(config.GetDatabaseConnection())
	db, err := gorm.Open(m, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return &TransactionsRepository{}
	}
	return &TransactionsRepository{Db: db}
}

// Create creates a new transaction in the TransactionsRepository.
//
// It takes a data parameter of type domains.Transaction.
// It returns a boolean indicating whether the creation was successful.
func (a TransactionsRepository) Create(data entities.Transaction) bool {
	r := a.Db.Create(&data)
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
func (a TransactionsRepository) GetByAccount(id int) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	r := a.Db.Where("account_id = ?", id).Find(&transactions)
	if r.Error != nil {
		log.Fatal(r.Error)
	}
	return transactions, r.Error
}

func (a TransactionsRepository) GetUnpaidBalanceByAccount(id int) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	r := a.
		Db.
		Where("account_id = ?", id).
		Where("balance <> 0").
		Order("event_date asc").
		Find(&transactions)
	if r.Error != nil {
		log.Fatal(r.Error)
	}
	return transactions, r.Error
}

func (a TransactionsRepository) Update(data entities.Transaction) error {
	r := a.Db.Save(&data)
	if r.Error != nil {
		log.Fatal(r.Error)
		return r.Error
	}
	return nil
}
