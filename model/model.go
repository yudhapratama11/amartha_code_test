package model

import "time"

type Loan struct {
	ID              int       `json:"id"`
	Principal       float64   `json:"principal"`
	WeeklyPayment   float64   `json:"weekly_payment"`
	InterestRate    float64   `json:"interest_rate"`
	WeeksPayment    []float64 `json:"weeks_payment"`
	Status          string    `json:"status"`
	InstallmentTerm int       `json:"installment_term"`
	StartDate       time.Time `json:"start_date"`
}

func (l *Loan) CheckTotalUnpaid(weeksElapsed int) int {
	ctrNonPaid := 0
	for i := 0; i < weeksElapsed; i++ {
		if l.WeeksPayment[i] == 0 {
			ctrNonPaid++
		}
	}

	return ctrNonPaid
}
