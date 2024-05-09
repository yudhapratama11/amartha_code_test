package main

import (
	"log"
	"net/http"
	"time"
	"yudhapratama11/amartha_code_test/handler"
	"yudhapratama11/amartha_code_test/repository"
	usecases "yudhapratama11/amartha_code_test/usecase"
)

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

func (l *Loan) checkTotalUnpaid(weeksElapsed int) int {
	ctrNonPaid := 0
	for i := 0; i < weeksElapsed; i++ {
		if l.WeeksPayment[i] == 0 {
			ctrNonPaid++
		}
	}

	return ctrNonPaid
}

func main() {
	loanRepo := repository.NewLoanRepository()
	loanUsecase := usecases.NewLoanUsecase(loanRepo)
	loanHandler := handler.NewLoanHandler(loanUsecase)

	srv := &http.Server{
		Handler:      loanHandler,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
