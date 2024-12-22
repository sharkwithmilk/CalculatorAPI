package application

import (
	calc "CalculatorAPI/pkg/Calculator"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type calcMock struct{}

func (c *calcMock) Calc(expression string) (float64, error) {
	switch expression {
	case "1+1":
		return 2, nil
	case "1-1":
		return 0, nil
	case "2*2":
		return 4, nil
	case "4/2":
		return 2, nil
	case "1/0":
		return 0, calc.ErrDivisionByZero
	default:
		return 0, calc.ErrInvalidExpression
	}
}

func TestCalculateHandler_Success_Addition(t *testing.T) {
	requestPayload := CalculationRequest{Expression: "1+1"}
	requestBody, _ := json.Marshal(requestPayload)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(requestBody))
	resp := httptest.NewRecorder()

	calculateHandler(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, resp.Code)
	}

	var response CalculationResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("error decoding response: %v", err)
	}

	if response.Result != "2.000000" {
		t.Errorf("expected result 2.000000, got %s", response.Result)
	}
}

func TestCalculateHandler_Success_Subtraction(t *testing.T) {
	requestPayload := CalculationRequest{Expression: "1-1"}
	requestBody, _ := json.Marshal(requestPayload)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(requestBody))
	resp := httptest.NewRecorder()

	calculateHandler(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, resp.Code)
	}

	var response CalculationResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("error decoding response: %v", err)
	}

	if response.Result != "0.000000" {
		t.Errorf("expected result 0.000000, got %s", response.Result)
	}
}

func TestCalculateHandler_Success_Multiplication(t *testing.T) {
	requestPayload := CalculationRequest{Expression: "2*2"}
	requestBody, _ := json.Marshal(requestPayload)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(requestBody))
	resp := httptest.NewRecorder()

	calculateHandler(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, resp.Code)
	}

	var response CalculationResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("error decoding response: %v", err)
	}

	if response.Result != "4.000000" {
		t.Errorf("expected result 4.000000, got %s", response.Result)
	}
}

func TestCalculateHandler_Success_Division(t *testing.T) {
	requestPayload := CalculationRequest{Expression: "4/2"}
	requestBody, _ := json.Marshal(requestPayload)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(requestBody))
	resp := httptest.NewRecorder()

	calculateHandler(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, resp.Code)
	}

	var response CalculationResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("error decoding response: %v", err)
	}

	if response.Result != "2.000000" {
		t.Errorf("expected result 2.000000, got %s", response.Result)
	}
}

func TestCalculateHandler_InvalidExpression(t *testing.T) {
	requestPayload := CalculationRequest{Expression: "invalid"}
	requestBody, _ := json.Marshal(requestPayload)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(requestBody))
	resp := httptest.NewRecorder()

	calculateHandler(resp, req)

	if resp.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected status %d, got %d", http.StatusUnprocessableEntity, resp.Code)
	}

	var response CalculationResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("error decoding response: %v", err)
	}

	if response.Error != "Expression is not valid" {
		t.Errorf("expected error message 'Expression is not valid', got '%s'", response.Error)
	}
}

func TestCalculateHandler_DivisionByZero(t *testing.T) {
	requestPayload := CalculationRequest{Expression: "1/0"}
	requestBody, _ := json.Marshal(requestPayload)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(requestBody))
	resp := httptest.NewRecorder()

	calculateHandler(resp, req)

	if resp.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected status %d, got %d", http.StatusUnprocessableEntity, resp.Code)
	}

	var response CalculationResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("error decoding response: %v", err)
	}

	if response.Error != "Expression is not valid" {
		t.Errorf("expected error message 'Expression is not valid', got '%s'", response.Error)
	}
}

func TestCalculateHandler_InvalidMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/calculate", nil)
	resp := httptest.NewRecorder()

	calculateHandler(resp, req)

	if resp.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status %d, got %d", http.StatusMethodNotAllowed, resp.Code)
	}

	body := resp.Body.String()
	if body != "Method not allowed\n" {
		t.Errorf("expected body 'Method not allowed', got '%s'", body)
	}
}
