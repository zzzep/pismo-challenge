package entities

import "time"

type Account struct {
	AccountId      int       `gorm:"type:int; primaryKey" json:"account_id,omitempty"`
	DocumentNumber string    `gorm:"type:varchar(20)" json:"document_number"`
	CreatedAt      time.Time `gorm:"type:time; autoCreateTime; default:CURRENT_TIMESTAMP(3)" json:"-"`
	UpdatedAt      time.Time `gorm:"type:time; autoUpdateTime; default:CURRENT_TIMESTAMP(3)" json:"-"`
	DeletedAt      time.Time `gorm:"type:time; default:null" json:"-"`
}
