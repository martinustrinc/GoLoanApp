package router

import (
	"GoLoanApp/model/dtoIn"
	"GoLoanApp/model/dtoOut"
	"GoLoanApp/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func ApiController() {
	logger := log.New(os.Stdout, "API ", log.LstdFlags)
	handler := mux.NewRouter()
	prefixPath := "/app"

	handler.HandleFunc("/v1"+prefixPath+"/createLoanAccount", func(w http.ResponseWriter, r *http.Request) {
		logger.Println("Received request for createLoanAccount")

		// Decode request JSON
		var (
			request  dtoIn.LoanRequest
			response dtoOut.LoanResponse
			err      error
		)

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&request); err != nil {
			logger.Println("Invalid request body")
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validation input
		errs := services.ValidateRequestCreateLoanAccount(request)
		if errs != "" {
			logger.Println(errs)
			http.Error(w, errs, http.StatusBadRequest)
			return
		}

		// Calculate Installment
		response.AccountLoan, err = services.CalculateInstallment(request)
		if err != nil {
			logger.Println(err)
			http.Error(w, "Error Calculate Installment", http.StatusInternalServerError)
			return
		}

		// Save data to database
		//accountCode, err := saveLoanAccount(db, request, AccountLoan)
		//if err != nil {
		//	logger.Println(err)
		//	http.Error(w, "Error save data to database", http.StatusInternalServerError)
		//	return
		//}

		response.AccountCode = "123456"
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			logger.Println("Error marshalling JSON response")
			http.Error(w, "Error marshalling JSON response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonResponse)
		if err != nil {
			logger.Println("Error writing JSON response")
			http.Error(w, "Error writing JSON response", http.StatusInternalServerError)
			return
		}
	}).Methods("POST", "OPTIONS")

	handler.Use(Middleware)
	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8091"
	}
	logger.Println("Server is running on port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
