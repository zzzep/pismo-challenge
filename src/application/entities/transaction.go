package entities

import (
	"time"
)

type Transaction struct {
	TransactionId   int           `gorm:"type:int; primaryKey; autoIncrement:true" json:"-"`
	AccountId       int           `gorm:"type:int foreignKey:account_id" json:"account_id"`
	OperationTypeId OperationType `gorm:"type:enum('1','2','3','4')" json:"operation_type_id"`
	Amount          float64       `gorm:"type:float" json:"amount"`
	Balance         float64       `gorm:"type:float" json:"balance"`
	EventDate       time.Time     `gorm:"type:time; autoCreateTime; default:CURRENT_TIMESTAMP(3)" json:"-"`
	CreatedAt       time.Time     `gorm:"type:time; autoCreateTime; default:CURRENT_TIMESTAMP(3)" json:"-"`
	UpdatedAt       time.Time     `gorm:"type:time; autoUpdateTime; default:CURRENT_TIMESTAMP(3)" json:"-"`
}
