package router

import (
	"GoLoanApp/model/dtoIn"
	"GoLoanApp/services"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func ApiController() {
	var logger *log.Logger
	handler := mux.NewRouter()
	prefixPath := "/app"

	handler.HandleFunc("/v1"+prefixPath+"/createLoanAccount", func(w http.ResponseWriter, r *http.Request) {
		log.New(os.Stdout, "API Create Loan : ", log.LstdFlags)

		// Decode request JSON
		var (
			request dtoIn.LoanRequest
			err     error
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
		}

		// Calculate Installment
		request.Installment, err = services.CalculateInstallment(request)
		if err != nil {
			logger.Println(err)
			http.Error(w, "Error Calculate Installment", http.StatusInternalServerError)
			return
		}

		// Save data to database
		//accountCode, err := saveLoanAccount(db, request, tabelAngsuran)
		//if err != nil {
		//	logger.Println(err)
		//	http.Error(w, "Error save data to database", http.StatusInternalServerError)
		//	return
		//}

		jsonResponse, err := json.Marshal(request)
		if err != nil {
			logger.Println("Error marshalling JSON response")
			http.Error(w, "Error marshalling JSON response", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonResponse)
	}).Methods("POST", "OPTIONS")

	handler.Use(Middleware)
	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}
	fmt.Println("Server is running on port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
