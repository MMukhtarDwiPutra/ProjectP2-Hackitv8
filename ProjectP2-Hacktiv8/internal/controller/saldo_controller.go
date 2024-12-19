package controller

import (
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/internal/service"
	"github.com/labstack/echo/v4"            // Import Echo framework untuk pengelolaan HTTP API.
	"net/http"
	"fmt"
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
			"message": fmt.Sprintf("User ID is not valid!"),
		})
	}

	// Bind the request body to the BalanceRequest struct
	if err := c.Bind(&topUpRequest); err != nil {
		// Return response if binding fails
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Booking create fail: %v", err),
		})
	}

	// Set the UserID in the topUpRequest
	topUpRequest.UserID = userID

	// Validate the request body
	err := validate.Struct(topUpRequest)
	if err != nil {
		// Return response if validation fails
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": fmt.Sprintf("Need: %v", err),
		})
	}

	// Call the service to perform the top-up operation
	status, webResponse := h.saldoService.TopUp(topUpRequest)

	// Return the response
	return c.JSON(status, webResponse)
}
