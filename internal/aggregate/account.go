package aggregate

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	id           uuid.UUID
	transactions []Transaction
}

type Transaction struct {
	amount    int
	tType     string
	createdAt time.Time
}

func NewAcount() Account {

	return Account{
		id:           uuid.New(),
		transactions: make([]Transaction, 0),
	}

}
