package repository

import (
	"errors"
	"time"
	"yudhapratama11/amartha_code_test/model"
)

type LoanRepository struct {
	loans map[int]*model.Loan
}

// Assumption if i have valid data
func NewLoanRepository() LoanRepository {
	return LoanRepository{
		loans: map[int]*model.Loan{
			1: {
				ID:              1,
				Principal:       5000000,
				WeeklyPayment:   1100000,
				InterestRate:    10,
				WeeksPayment:    []float64{1100000, 1100000, 0, 0, 0},
				Status:          "open",
				InstallmentTerm: 5,
				StartDate:       time.Date(2024, time.April, 9, 0, 0, 0, 0, time.UTC),
			},
			2: {
				ID:              2,
				Principal:       8000000,
				WeeklyPayment:   440000,
				InterestRate:    10,
				WeeksPayment:    []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				Status:          "open",
				InstallmentTerm: 20,
				StartDate:       time.Date(2024, time.February, 9, 0, 0, 0, 0, time.UTC),
			},
		},
	}
}

func (r *LoanRepository) CreateLoan(loan *model.Loan) error {
	if loan == nil {
		return errors.New("invalid params")
	}

	r.loans[loan.ID] = loan
	return nil
}

func (r *LoanRepository) GetLoanLength() int {
	return len(r.loans)
}

func (r *LoanRepository) GetLoan(id int) (*model.Loan, error) {
	loan, ok := r.loans[id]
	if !ok {
		return nil, errors.New("not found")
	}

	return loan, nil
}

func (r *LoanRepository) GetLoanSchedule(id int) ([]float64, error) {
	loan, ok := r.loans[id]
	if !ok {
		return nil, errors.New("not found")
	}

	return loan.WeeksPayment, nil
}

func (r *LoanRepository) GetLoanStatus(id int) (string, error) {
	loan, ok := r.loans[id]
	if !ok {
		return "", errors.New("not found")
	}

	return loan.Status, nil
}
