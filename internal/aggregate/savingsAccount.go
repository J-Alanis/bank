package aggregate

import (
	"github.com/J-Alanis/bank/internal/entity"
	"github.com/J-Alanis/bank/internal/valueobject"
	"github.com/google/uuid"
)

type SavingsAccount struct {
	account      *entity.Account
	transactions []valueobject.Transaction
}

func NewSavingsAcount() SavingsAccount {

	account := &entity.Account{
		ID:      uuid.New(),
		Balance: 0,
	}

	return SavingsAccount{
		account:      account,
		transactions: make([]valueobject.Transaction, 0),
	}

}
