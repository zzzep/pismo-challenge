package migrations

import (
	entity2 "github.com/zzzep/pismo-challenge/src/application/entities"
	"github.com/zzzep/pismo-challenge/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Run runs the function.
//
// No parameters.
// Returns an error.
func Run() (err error) {
	var db *gorm.DB
	conn := config.GetDatabaseConnection()
	m := mysql.Open(conn)
	db, err = gorm.Open(m, &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&entity2.Account{}, &entity2.Transaction{})
	if err != nil {
		return err
	}

	return nil
}
