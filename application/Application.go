package application

import (
	calc "CalculatorAPI/pkg/Calculator"
	"encoding/json"
	"fmt"
	"net/http"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}
func (a *Application) RunServer(port string) error {
	http.HandleFunc("/api/v1/calculate", calculateHandler)
	return http.ListenAndServe(":"+port, nil)
}

type CalculationRequest struct {
	Expression string `json:"expression"`
}

type CalculationResponse struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusUnprocessableEntity, "Expression is not valid")
		return
	}

	result, err := calc.Calc(request.Expression)
	if err != nil {
		switch err {
		case calc.ErrEmptyExpression, calc.ErrNumberParsing, calc.ErrInvalidParentheses, calc.ErrDivisionByZero, calc.ErrInvalidExpression, calc.ErrUnrecognizedNumber, calc.ErrMissingNumber:
			respondWithError(w, http.StatusUnprocessableEntity, "Expression is not valid")
		default:
			respondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	response := CalculationResponse{Result: fmt.Sprintf("%f", result)}
	respondWithJSON(w, http.StatusOK, response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	response := CalculationResponse{Error: message}
	respondWithJSON(w, code, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		fmt.Printf("Error encoding response: %v\n", err)
	}
}
