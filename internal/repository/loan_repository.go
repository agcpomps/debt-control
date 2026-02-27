package repository

import (
	"agcpomps/debt-control/internal/domain"
	"context"

	"github.com/google/uuid"
)

type LoanRepository interface {
	Create(ctx context.Context, loan *domain.Loan) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Loan, error)
}
