package postgres

import (
	"agcpomps/debt-control/internal/db"
	"agcpomps/debt-control/internal/domain"
	"agcpomps/debt-control/internal/repository"
	"context"
	"strconv"

	"github.com/google/uuid"
)

type loanRepository struct {
	queries *db.Queries
}

func NewLoanRepository(q *db.Queries) repository.LoanRepository {
	return &loanRepository{queries: q}
}

func (r *loanRepository) Create(ctx context.Context, loan *domain.Loan) error {
	params := db.CreateLoanParams{
		ClientID:            loan.ClientID,
		Principal:           strconv.FormatFloat(loan.Principal, 'f', 2, 64),
		MonthlyInterestRate: strconv.FormatFloat(loan.MonthlyInterestRate, 'f', 3, 64),
		Months:              loan.Months,
		TotalAmount:         strconv.FormatFloat(loan.TotalAmount, 'f', 2, 64),
		DisbursedAt:         loan.DisbursedAt,
	}

	result, err := r.queries.CreateLoan(ctx, params)
	if err != nil {
		return err
	}

	loan.ID = result.ID
	loan.CreatedAt = result.CreatedAt

	return nil
}

func (r *loanRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Loan, error) {
	dbLoan, err := r.queries.GetLoanByID(ctx, id)
	if err != nil {
		return nil, err
	}

	principal, _ := strconv.ParseFloat(dbLoan.Principal, 64)
	rate, _ := strconv.ParseFloat(dbLoan.MonthlyInterestRate, 64)
	total, _ := strconv.ParseFloat(dbLoan.TotalAmount, 64)

	return &domain.Loan{
		ID:                  dbLoan.ID,
		ClientID:            dbLoan.ClientID,
		Principal:           principal,
		MonthlyInterestRate: rate,
		Months:              dbLoan.Months,
		TotalAmount:         total,
		DisbursedAt:         dbLoan.DisbursedAt,
		CreatedAt:           dbLoan.CreatedAt,
	}, nil
}
