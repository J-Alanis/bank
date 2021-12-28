package valueobject

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	amount            int
	depositOrWithdraw string
	account           uuid.UUID
	createdAt         time.Time
}
