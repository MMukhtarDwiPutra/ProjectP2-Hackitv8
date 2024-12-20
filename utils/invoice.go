package utils

import (
	"bytes"
	"encoding/json"
	"P2-Hacktiv8/entity"
	"net/http"
	"os"
	"encoding/base64"
	"fmt"
	"io"
)

func CreateInvoice(user entity.User, topUpRequest entity.BalanceRequest, externalID string) (*entity.InvoiceResponse, error) {
	// Get the API key from the environment variable
	apiKey := os.Getenv("3RD_PARTY_XENDIT_KEY")
	apiUrl := os.Getenv("XENDIT_INVOICE_URL")

	fmt.Println(topUpRequest.Balance)
	// Create the invoice request
	invoiceRequest := entity.InvoiceRequest{
		ExternalID:     externalID,
		Amount:         topUpRequest.Balance * 1000,
		Description:    "Invoice Demo #123",
		InvoiceDuration: 86400,
		Currency:       "IDR",
		ReferenceID:    "test",
		CheckoutMethod: "ONE TIME PAYMENT",
	}

	// Marshal the invoice request to JSON
	reqBody, err := json.Marshal(invoiceRequest)
	if err != nil {
		return nil, err
	}

	// Create the HTTP client and request
	client := &http.Client{}
	request, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	// Base64 encode the API key and set the Authorization header
	encodedAPIKey := base64.StdEncoding.EncodeToString([]byte(apiKey + ":"))
	request.Header.Set("Authorization", "Basic "+encodedAPIKey)
	request.Header.Set("Content-Type", "application/json")

	// Make the API call
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Handle the response (decode the response JSON)
	var paymentResponse entity.InvoiceResponse
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&paymentResponse); err != nil {
		return nil, err
	}

	// Return the response
	return &paymentResponse, nil
}