package interfaces

import (
	"github.com/zzzep/pismo-challenge/src/application/entities"
)

type IAccountsRepository interface {
	Create(data entities.Account) bool
	Get(id int) *entities.Account
}
