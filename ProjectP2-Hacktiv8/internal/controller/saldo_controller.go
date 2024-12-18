package controller

import (
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/internal/service"
	"github.com/labstack/echo/v4"            // Import Echo framework untuk pengelolaan HTTP API.
	"net/http"
	"fmt"
)

type saldoController struct {
	saldoService service.SaldoService
}

// NewSaldoController creates a new instance of saldoController.
func NewSaldoController(saldoService service.SaldoService) *saldoController {
	return &saldoController{saldoService}
}

func (h *saldoController) TopUp(c echo.Context) error{
	var topUpRequest entity.TopUpRequest

	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status" : http.StatusUnauthorized,
			"message": fmt.Sprintf("User ID is not valid!"),
		})
	}

	// Melakukan bind request body ke struct.
	if err := c.Bind(&topUpRequest); err != nil {
		// Mengembalikan respons jika request body tidak valid.
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status" : http.StatusInternalServerError,
			"message": fmt.Sprintf("Booking create fail: %v",err),
		})
	}

	topUpRequest.UserID = userID
	err := validate.Struct(topUpRequest)
	if err != nil {
		// Mengembalikan respons jika validasi gagal.
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status" : http.StatusBadRequest,
			"message": fmt.Sprintf("Need: %v",err),
		})
	}
	
	status, webResponse := h.saldoService.TopUp(topUpRequest)

	return c.JSON(status, webResponse)
}