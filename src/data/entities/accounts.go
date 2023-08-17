package entities

type Account struct {
	AccountId      int    `json:"account_id" gorm:"primaryKey"`
	DocumentNumber string `json:"document_number"`
}
