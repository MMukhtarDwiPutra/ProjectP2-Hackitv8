package controller

import (
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/internal/service"
	"github.com/labstack/echo/v4"            // Import Echo framework untuk pengelolaan HTTP API.
	"net/http"
	"fmt"
	// "time"
	xendit "github.com/xendit/xendit-go/v4"
	// balance_and_transaction "github.com/xendit/xendit-go/v4/balance_and_transaction"
	"os"
	"context"
	"encoding/json"
	"encoding/base64"
)

// saldoController is the controller for saldo-related operations.
type saldoController struct {
	saldoService service.SaldoService
}

// NewSaldoController creates a new instance of saldoController.
func NewSaldoController(saldoService service.SaldoService) *saldoController {
	return &saldoController{saldoService}
}

// TopUp godoc
// @Summary Top up user balance
// @Description Top up the balance of a user by providing a `BalanceRequest` object which includes the amount to be added to the balance.
// @Tags Saldo
// @Accept json
// @Produce json
// @Param request body entity.BalanceRequest true "Top Up Request"
// @Success 200 {object} map[string]interface{} "Balance successfully topped up"
// @Failure 400 {object} map[string]interface{} "Invalid request parameters"
// @Failure 401 {object} map[string]interface{} "Unauthorized - Invalid user ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /top-up [post]
func (h *saldoController) TopUp(c echo.Context) error {
	var topUpRequest entity.BalanceRequest

	// Extract the user ID from the context
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": "User ID is not valid!",
		})
	}

	// Bind the request body to the BalanceRequest struct
	if err := c.Bind(&topUpRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Binding failed: %v", err),
		})
	}

	// Set the UserID in the topUpRequest
	topUpRequest.UserID = userID

	// Validate the request body
	err := validate.Struct(topUpRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": fmt.Sprintf("Validation error: %v", err),
		})
	}

	// The selected balance type
    // (default to "CASH")
    accountType := "CASH" // [OPTIONAL] | string

    // Currency for filter for customers with multi currency accounts
    currency := "IDR" // [OPTIONAL] | string

    // The sub-account user-id that you want to make this transaction for. This header
    // is only used if you have access to xenPlatform. See xenPlatform for more
    // information
    forUserId := "5dbf20d7c8eb0c0896f811b6" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.BalanceApi.GetBalance(context.Background()).
        AccountType(accountType).
        Currency(currency).
        ForUserId(forUserId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalanceApi.GetBalance``: %v\n", err.Error())

        b, _ := json.Marshal(err)
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBalance`: Balance
    fmt.Fprintf(os.Stdout, "Response from `BalanceApi.GetBalance`: %v\n", resp)

	// Save the top-up transaction in the database (use a service layer)
	// transaction := entity.Transaction{
	// 	UserID:      userID,
	// 	Amount:      topUpRequest.Balance,
	// 	ExternalID:  params.ExternalID,
	// 	VAAccount:   virtualAccount.AccountNumber,
	// 	Status:      "PENDING",
	// 	CreatedAt:   time.Now(),
	// }
	// if err := h.transactionService.SaveTransaction(transaction); err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 		"status":  http.StatusInternalServerError,
	// 		"message": fmt.Sprintf("Failed to save transaction: %v", err),
	// 	})
	// }

	// Return the virtual account details to the user
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  http.StatusCreated,
		"message": "Top-up initiated successfully.",
	})
}

// GetAllPaymentsMethod retrieves all payment methods from Xendit API.
func (h *saldoController) GetAllPaymentsMethod(c echo.Context) error {
	// Define the API key and encode it for Basic Auth
	xenditAPIKey := os.Getenv("3RD_PARTY_XENDIT_API")
	
	if xenditAPIKey == "" {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": "Xendit API key is not set in environment variables",
		})
	}

	// Encode the API key to Base64 for Basic Auth
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(xenditAPIKey+":"))
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(xenditAPIKey+":")))

	// Xendit API endpoint to retrieve payment methods
	url := "https://api.xendit.co/payment_methods" // Replace with the actual endpoint if needed

	// Create an HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Failed to create request: %v", err),
		})
	}

	// Add the Authorization header
	req.Header.Add("Authorization", authHeader)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Failed to make API request: %v", err),
		})
	}
	defer resp.Body.Close()

	// Handle non-200 HTTP status codes
	if resp.StatusCode != http.StatusOK {
		return c.JSON(resp.StatusCode, map[string]interface{}{
			"status":  resp.StatusCode,
			"message": "Failed to fetch payment methods",
		})
	}

	// Parse the response body
	var paymentResponse []entity.PaymentMethod // Assuming this is a list of payment methods
	err = json.NewDecoder(resp.Body).Decode(&paymentResponse)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Failed to decode API response: %v", err),
		})
	}

	// Return the payment methods as JSON
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Successfully retrieved payment methods",
		"data":    paymentResponse,
	})
}

// func (h *saldoController) TopUpCallback(c echo.Context) error {
// 	var callbackRequest xendit.CallbackRequest

// 	// Bind the incoming callback data
// 	if err := c.Bind(&callbackRequest); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"status":  http.StatusBadRequest,
// 			"message": fmt.Sprintf("Invalid callback data: %v", err),
// 		})
// 	}

// 	// Verify the callback signature (optional, based on Xendit requirements)
// 	if !h.verifyXenditCallback(c) {
// 		return c.JSON(http.StatusForbidden, map[string]interface{}{
// 			"status":  http.StatusForbidden,
// 			"message": "Invalid callback signature",
// 		})
// 	}

// 	// Update the transaction in the database based on ExternalID
// 	transaction, err := h.transactionService.GetTransactionByExternalID(callbackRequest.ExternalID)
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, map[string]interface{}{
// 			"status":  http.StatusNotFound,
// 			"message": "Transaction not found",
// 		})
// 	}

// 	if callbackRequest.Status == "SUCCESS" {
// 		// Update transaction status
// 		transaction.Status = "COMPLETED"
// 		h.transactionService.UpdateTransaction(transaction)

// 		// Update user's balance
// 		user, err := h.userService.GetUserByID(transaction.UserID)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 				"status":  http.StatusInternalServerError,
// 				"message": fmt.Sprintf("Failed to fetch user: %v", err),
// 			})
// 		}
// 		user.Balance += transaction.Amount
// 		h.userService.UpdateUserBalance(user)
// 	}

// 	// Return success response
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"status":  http.StatusOK,
// 		"message": "Callback processed successfully",
// 	})
// }