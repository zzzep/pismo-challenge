package migrations

import (
	"github.com/zzzep/pismo-challenge/config"
	"github.com/zzzep/pismo-challenge/src/data/entities"
)

func Run() (err error) {
	var db *gorm.DB
	db, err = gorm.Open(mysql.Open(config.GetDatabaseConnection()), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&entities.Account{}, &entities.Transaction{})
	if err != nil {
		return err
	}

	return nil
}
