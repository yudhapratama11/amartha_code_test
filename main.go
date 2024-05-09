package main

import (
	"log"
	"net/http"
	"time"
	"yudhapratama11/amartha_code_test/handler"
	"yudhapratama11/amartha_code_test/repository"
	usecases "yudhapratama11/amartha_code_test/usecase"
)

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
