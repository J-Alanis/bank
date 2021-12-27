package valueobject

import (
	"time"
)

type Transaction struct {
	amount            int
	depositOrWithdraw string
	account           uuid.UUID
	createdAt         time.Time
}
