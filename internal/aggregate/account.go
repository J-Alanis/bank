package aggregate

import (
	"time"

	"github.com/google/uuid"
)

type account struct {
	id           uuid.UUID
	transactions []transaction
}

type transaction struct {
	amount    int
	tType     string
	createdAt time.Time
}

func NewAcount() account {

	return account{
		id:           uuid.New(),
		transactions: make([]transaction, 0),
	}

}
