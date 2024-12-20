package controller

import (
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/internal/service"
	"github.com/labstack/echo/v4"            // Import Echo framework untuk pengelolaan HTTP API.
	"net/http"
	"fmt"
	// "time"
	// balance_and_transaction "github.com/xendit/xendit-go/v4/balance_and_transaction"
	"os"
	"encoding/json"
	"encoding/base64"
	// "io"
	// "bytes"
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

	status, webResponse := h.saldoService.TopUp(topUpRequest)

	// Return the virtual account details to the user
	return c.JSON(status, webResponse)
}

// GetAllPaymentsMethod retrieves all payment methods from Xendit API.
func (h *saldoController) GetAllPaymentsMethod(c echo.Context) error {
	xenditAPIKey := os.Getenv("3RD_PARTY_XENDIT_API")
	if xenditAPIKey == "" {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": "Xendit API key is not set in environment variables",
		})
	}

	// authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(xenditAPIKey+":"))
	url := "https://api.xendit.co/v2/payment_methods"

	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 		"status":  http.StatusInternalServerError,
	// 		"message": fmt.Sprintf("Failed to create request: %v", err),
	// 	})
	// }
	// req.Header.Add("Authorization", authHeader)

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 		"status":  http.StatusInternalServerError,
	// 		"message": fmt.Sprintf("Failed to make API request: %v", err),
	// 	})
	// }
	// defer resp.Body.Close()

	// // Debugging
	// fmt.Println("HTTP Status Code:", resp.StatusCode)
	// body, _ := io.ReadAll(resp.Body)
	// fmt.Println("Response Body:", string(body))

	// if resp.StatusCode != http.StatusOK {
	// 	return c.JSON(resp.StatusCode, map[string]interface{}{
	// 		"status":  resp.StatusCode,
	// 		"message": "Failed to fetch payment methods",
	// 	})
	// }

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Failed to get request from API: %v", err),
		})
	}

	apiKey := "Basic "+base64.StdEncoding.EncodeToString([]byte(xenditAPIKey+":"))
	request.SetBasicAuth(apiKey, "")
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
	 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Failed to get the API response: %v", err),
		})
	}

	defer response.Body.Close()

	var paymentResponse entity.PaymentMethodsResponse
	if err := json.NewDecoder(response.Body).Decode(&paymentResponse); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Failed to decode API response: %v", err),
		})
	}

	// var paymentResponse []entity.PaymentMethodsResponse
	// err = json.Unmarshal(body, &paymentResponse)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 		"status":  http.StatusInternalServerError,
	// 		"message": fmt.Sprintf("Failed to decode API response: %v", err),
	// 	})
	// }

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Successfully retrieved payment methods",
		"data":    paymentResponse,
	})
	
}