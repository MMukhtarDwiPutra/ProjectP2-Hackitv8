package controller

import (
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/internal/service"
	"github.com/labstack/echo/v4"            // Import Echo framework untuk pengelolaan HTTP API.
	"net/http"
	"fmt"
	"os"
	"encoding/json"
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
// @Security ApiKeyAuth
// @Security BearerAuth
// @Param Authorization header string true "Bearer Token (Example: 'Bearer <your_token>')"
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

// InvoiceWebhookHandler handles the webhook callback from Xendit
// @Summary      Handles Invoice Webhook from Xendit
// @Description  Processes the webhook sent by Xendit for invoice updates (e.g., PAID, EXPIRED).
// @Tags         Webhook
// @Accept       json
// @Produce      json
// @Param        x-callback-token  header    string  true  "Xendit Callback Token"
// @Param        body              body      entity.WebhookPayload  true  "Webhook Payload"
// @Success      204               {string}  string  "No Content"
// @Failure      400               {object}  map[string]interface{}  "Bad Request - Failed to parse JSON"
// @Failure      401               {object}  map[string]interface{}  "Unauthorized - Invalid callback token"
// @Failure      500               {object}  map[string]interface{}  "Internal Server Error"
// @Security     ApiKeyAuth
// @Security     BearerAuth
// @Param        Authorization header string true "Bearer Token (Example: 'Bearer <your_token>')"
// @Router       /invoice_webhook_url [post]
func (h *saldoController) InvoiceWebhookHandler(c echo.Context) error {
	// Get the x-callback-token from the header
	callbackToken := c.Request().Header.Get("x-callback-token")
	expectedToken := os.Getenv("XENDIT_WEBHOOK_KEY") // Set this in your environment

	// Validate the callback token
	if callbackToken != expectedToken {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid callback token")
	}

	// Parse the JSON body into the InvoiceWebhook struct
	var webhookPayload entity.WebhookPayload
	err := json.NewDecoder(c.Request().Body).Decode(&webhookPayload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse JSON")
	}

	status := h.saldoService.CallbackWebhook(webhookPayload)

	// Respond with HTTP 200 status
	return c.NoContent(status)
}
