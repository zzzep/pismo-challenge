package repositories

import (
	"github.com/zzzep/pismo-challenge/src/application/entities"
	"github.com/zzzep/pismo-challenge/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type AccountsRepository struct {
	Db *gorm.DB
}

// NewAccountRepository creates a new instance of AccountsRepository.
//
// It opens a MySQL database connection using the GetDatabaseConnection function from the enum package.
// It then initializes a new GORM database connection using the database connection and configuration options.
// If an error occurs during the database connection setup, it logs the error and returns an empty AccountsRepository.
// Otherwise, it returns a new instance of AccountsRepository with the initialized database connection.
func NewAccountRepository() *AccountsRepository {
	m := mysql.Open(config.GetDatabaseConnection())
	db, err := gorm.Open(m, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return &AccountsRepository{}
	}
	return &AccountsRepository{Db: db}
}

// Create creates a new account in the AccountsRepository.
//
// It takes a `domains.Account` data as a parameter and returns a boolean value.
func (a AccountsRepository) Create(data entities.Account) bool {
	r := a.Db.Create(&data)
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
func (a AccountsRepository) Get(id int) *entities.Account {
	acc := &entities.Account{}
	_ = a.Db.First(acc, id)
	return acc
}
