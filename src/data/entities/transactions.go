package entities

import "time"

type Transaction struct {
	TransactionId   int           `gorm:"type:int; primaryKey"`
	Account         Account       `gorm:"type:int; references:AccountId" json:"account_id"`
	OperationTypeId OperationType `gorm:"type:enum('1','2','3','4')" json:"operation_type_id"`
	Amount          float64       `gorm:"type:float" json:"amount"`
	EventDate       time.Time     `gorm:"type:time; autoCreateTime; default:CURRENT_TIMESTAMP(3)"`
	CreatedAt       time.Time     `gorm:"type:time; autoCreateTime; default:CURRENT_TIMESTAMP(3)"`
	UpdatedAt       time.Time     `gorm:"type:time; autoUpdateTime; default:CURRENT_TIMESTAMP(3)"`
	DeletedAt       time.Time     `gorm:"type:time; default:null"`
}
