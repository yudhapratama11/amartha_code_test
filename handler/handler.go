package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yudhapratama11/amartha_code_test/model"
	usecase "yudhapratama11/amartha_code_test/usecase"
	usecases "yudhapratama11/amartha_code_test/usecase"
)

func NewLoanHandler(usecase *usecase.LoanUsecase) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/loans", createLoanHandler(usecase))
	mux.HandleFunc("/loans/{id}/schedule", getScheduleHandler(usecase))
	mux.HandleFunc("/loans/{id}/outstandingbalance", getBalanceHandler(usecase))
	mux.HandleFunc("/loans/{id}/status", getStatusHandler(usecase))
	mux.HandleFunc("/loans/{id}/repay", repayLoanHandler(usecase))

	return mux
}

func createLoanHandler(usecase *usecases.LoanUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var loan model.Loan
		err := json.NewDecoder(r.Body).Decode(&loan)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)

		err = usecase.CreateLoan(&loan)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		json.NewEncoder(w).Encode(loan)
	}
}

func getScheduleHandler(usecase *usecases.LoanUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid loan ID", http.StatusBadRequest)
			return
		}

		schedule, err := usecase.GetLoanSchedule(id)
		resp := map[string]interface{}{
			"id": id,
		}

		if err != nil {
			resp["err_msg"] = err.Error()
		} else {
			resp["schedule"] = schedule
		}

		json.NewEncoder(w).Encode(resp)
	}
}

func getBalanceHandler(usecase *usecases.LoanUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid loan ID", http.StatusBadRequest)
			return
		}

		balance, err := usecase.GetOutstanding(id)
		resp := map[string]interface{}{
			"id": id,
		}

		if err != nil {
			resp["err_msg"] = err.Error()
		} else {
			resp["outstanding_balance"] = balance
		}

		json.NewEncoder(w).Encode(resp)
	}
}

func getStatusHandler(usecase *usecases.LoanUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid loan ID", http.StatusBadRequest)
			return
		}

		status, err := usecase.IsDelinquent(id)
		resp := map[string]interface{}{
			"id": id,
		}

		if err != nil {
			resp["err_msg"] = err.Error()
		} else {
			resp["is_deliquent"] = status
		}

		json.NewEncoder(w).Encode(resp)
	}

}

func repayLoanHandler(usecase *usecases.LoanUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid loan ID", http.StatusBadRequest)
			return
		}

		var repayment struct {
			PaidAmount float64 `json:"paid_amount"`
		}

		err = json.NewDecoder(r.Body).Decode(&repayment)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		resp := map[string]interface{}{
			"id": id,
		}

		err = usecase.MakePayment(id, repayment.PaidAmount)
		if err != nil {
			resp["err_msg"] = err.Error()
		} else {
			resp["msg"] = "Payment success"
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}
