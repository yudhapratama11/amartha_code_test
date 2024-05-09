package usecases

import (
	"errors"
	"fmt"
	"time"
	"yudhapratama11/amartha_code_test/model"
	"yudhapratama11/amartha_code_test/repository"
)

type LoanUsecase struct {
	repo repository.LoanRepository
}

func NewLoanUsecase(repo repository.LoanRepository) *LoanUsecase {
	return &LoanUsecase{repo: repo}
}

func (uc *LoanUsecase) CreateLoan(loan *model.Loan) error {
	loan.ID = uc.repo.GetLoanLength() + 1
	loan.WeeksPayment = make([]float64, loan.InstallmentTerm)
	loan.WeeklyPayment = loan.Principal / float64(loan.InstallmentTerm)
	loan.Status = "open"

	// Define all to 0 for the installment term
	// Assume every payment should be handled every week
	for i := 0; i < loan.InstallmentTerm; i++ {
		loan.WeeksPayment[i] = 0
	}

	return uc.repo.CreateLoan(loan)
}

func (uc *LoanUsecase) GetLoanSchedule(id int) (map[string]float64, error) {
	fmt.Println("id", id)
	loanSchedule, err := uc.repo.GetLoanSchedule(id)
	if err != nil {
		return nil, err
	}

	dataSchedule := make(map[string]float64, 0)
	for i, data := range loanSchedule {
		dataSchedule[fmt.Sprintf("week_%d", i+1)] = data
	}

	return dataSchedule, nil
}

func (uc *LoanUsecase) GetOutstanding(id int) (float64, error) {
	loan, err := uc.repo.GetLoan(id)
	if err != nil {
		return 0, err
	}

	if loan.Status == "closed" {
		return 0, nil
	}

	var totalWeeksPayment float64
	for _, payment := range loan.WeeksPayment {
		totalWeeksPayment += payment
	}

	return (loan.Principal + (loan.Principal * 10 / 100)) - totalWeeksPayment, nil
}

func (uc *LoanUsecase) IsDelinquent(id int) (bool, error) {
	outStandingBalance, err := uc.GetOutstanding(id)
	if err != nil || outStandingBalance == 0 {
		return false, err
	}

	loan, err := uc.repo.GetLoan(id)
	if err != nil {
		return false, err
	}

	weeksElapsed := int(time.Now().Sub(loan.StartDate).Hours() / (24 * 7))

	counterNonPayment := 0
	for i := 0; i < weeksElapsed; i++ {
		if loan.WeeksPayment[i] == 0 {
			counterNonPayment++
		}
	}

	return counterNonPayment >= 2, nil
}

func (uc *LoanUsecase) MakePayment(id int, paidAmount float64) error {
	loan, err := uc.repo.GetLoan(id)
	if err != nil {
		return err
	}

	if loan.Status == "closed" {
		return errors.New("you cannot make payment since loan is closed")
	}

	weeksElapsed := int(time.Now().Sub(loan.StartDate).Hours() / (24 * 7))
	totalUnpaidWeek := loan.CheckTotalUnpaid(weeksElapsed)

	// calculating ammount should be paid
	totalWeeklyPayment := loan.WeeklyPayment * float64(totalUnpaidWeek)
	if paidAmount < totalWeeklyPayment {
		return errors.New("you cannot make payment since paid amount is less than weekly payment")
	}

	if totalUnpaidWeek < 2 {
		loan.WeeksPayment[weeksElapsed-1] = loan.WeeklyPayment
	} else {
		for i := weeksElapsed - totalUnpaidWeek; i < weeksElapsed; i++ {
			loan.WeeksPayment[i] = loan.WeeklyPayment
		}
	}

	if len(loan.WeeksPayment) == weeksElapsed {
		loan.Status = "closed"
	}

	return nil
}
