package domain

import (
	"time"

	"github.com/google/uuid"
)

type Loan struct {
	ID                  uuid.UUID
	ClientID            uuid.UUID
	Principal           float64
	MonthlyInterestRate float64
	Months              int32
	TotalAmount         float64
	DisbursedAt         time.Time
	CreatedAt           time.Time
}
