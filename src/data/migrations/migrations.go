package migrations

import (
	"github.com/zzzep/pismo-challenge/src/config"
	"github.com/zzzep/pismo-challenge/src/data/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Run() (err error) {
	var db *gorm.DB
	m := mysql.Open(config.GetDatabaseConnection())
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
