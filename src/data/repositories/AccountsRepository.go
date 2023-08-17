package repositories

import (
	"github.com/zzzep/pismo-challenge/src/data/domains"
	"github.com/zzzep/pismo-challenge/src/enum"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type AccountsRepository struct {
	db *gorm.DB
}

func NewAccountRepository() *AccountsRepository {
	m := mysql.Open(enum.GetDatabaseConnection())
	db, err := gorm.Open(m, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return &AccountsRepository{}
	}
	return &AccountsRepository{db: db}
}
func (a AccountsRepository) Create(data domains.Account) bool {
	r := a.db.Create(&data)
	if r.Error != nil {
		log.Fatal(r.Error)
		return false
	}
	return true
}

func (a AccountsRepository) Get(id int) *domains.Account {
	acc := &domains.Account{}
	_ = a.db.First(acc, id)
	return acc
}
