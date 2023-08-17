package entities

import "time"

type Account struct {
	AccountId      int       `gorm:"type:int; primaryKey" json:"account_id"`
	DocumentNumber string    `gorm:"type:varchar(20)" json:"document_number"`
	CreatedAt      time.Time `gorm:"type:time; autoCreateTime; default:CURRENT_TIMESTAMP(3)"`
	UpdatedAt      time.Time `gorm:"type:time; autoUpdateTime; default:CURRENT_TIMESTAMP(3)"`
	DeletedAt      time.Time `gorm:"type:time; default:null"`
}
