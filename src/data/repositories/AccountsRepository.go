package repositories

import (
	"github.com/zzzep/pismo-challenge/src/data/domains"
	"github.com/zzzep/pismo-challenge/src/enum"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type IAccountsRepository interface {
	Create(data domains.Account) bool
	Get(id int) *domains.Account
}

type AccountsRepository struct {
	db *gorm.DB
}

// NewAccountRepository creates a new instance of AccountsRepository.
//
// It opens a MySQL database connection using the GetDatabaseConnection function from the enum package.
// It then initializes a new GORM database connection using the database connection and configuration options.
// If an error occurs during the database connection setup, it logs the error and returns an empty AccountsRepository.
// Otherwise, it returns a new instance of AccountsRepository with the initialized database connection.
func NewAccountRepository() *AccountsRepository {
	m := mysql.Open(enum.GetDatabaseConnection())
	db, err := gorm.Open(m, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return &AccountsRepository{}
	}
	return &AccountsRepository{db: db}
}

// Create creates a new account in the AccountsRepository.
//
// It takes a `domains.Account` data as a parameter and returns a boolean value.
func (a AccountsRepository) Create(data domains.Account) bool {
	r := a.db.Create(&data)
	if r.Error != nil {
		log.Fatal(r.Error)
		return false
	}
	return true
}

// Get retrieves an account from the AccountsRepository based on the provided ID.
//
// Parameters:
// - id: an integer representing the ID of the account to retrieve.
//
// Returns:
// - a pointer to a domains.Account struct representing the retrieved account.
func (a AccountsRepository) Get(id int) *domains.Account {
	acc := &domains.Account{}
	_ = a.db.First(acc, id)
	return acc
}
