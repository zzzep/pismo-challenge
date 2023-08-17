package migrations

import (
	"github.com/zzzep/pismo-challenge/src/data/entities"
	"github.com/zzzep/pismo-challenge/src/enum"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Run() (err error) {
	var db *gorm.DB
	m := mysql.Open(enum.GetDatabaseConnection())
	db, err = gorm.Open(m, &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&entities.Account{}, &entities.Transaction{})
	if err != nil {
		return err
	}

	return nil
}
