package entities

type Transaction struct {
	TransactionId   int           `gorm:"primaryKey"`
	AccountId       int           `json:"account_id"`
	OperationTypeId OperationType `json:"operation_type_id"`
	Amount          float64       `json:"amount"`
	EventDate       string
}
