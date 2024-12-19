package controller

import (
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/internal/service"
	"fmt"
	"github.com/labstack/echo/v4" // Import Echo framework untuk pengelolaan HTTP API.
	"net/http"
)

// bookingController handles booking-related operations.
type bookingController struct {
	bookingService service.BookingService
}

// NewBookingController creates a new instance of bookingController.
func NewBookingController(bookingService service.BookingService) *bookingController {
	return &bookingController{bookingService}
}

// BookARoom godoc
// @Summary Book a room
// @Description Books a room for a user, requiring down payment based on the room price and user's balance.
// @Tags Booking
// @Accept json
// @Produce json
// @Param request body entity.BookingRequest true "Booking Request"
// @Success 201 {object} map[string]interface{} "Booking created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request parameters"
// @Failure 404 {object} map[string]interface{} "Room or User not found"
// @Failure 402 {object} map[string]interface{} "Insufficient balance"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /booking [post]
func (h *bookingController) BookARoom(c echo.Context) error {
	var bookingRequest entity.BookingRequest

	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status" : http.StatusUnauthorized,
			"message": fmt.Sprintf("User ID is not valid!"),
		})
	}

	// Bind request body to struct.
	if err := c.Bind(&bookingRequest); err != nil {
		// Return response if request body is invalid.
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status" : http.StatusInternalServerError,
			"message": fmt.Sprintf("Booking create fail: %v", err),
		})
	}

	bookingRequest.UserID = userID
	// Validate the request data using validator.
	err := validate.Struct(bookingRequest)
	if err != nil {
		// Return response if validation fails.
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status" : http.StatusBadRequest,
			"message": fmt.Sprintf("Need: %v", err),
		})
	}

	status, webResponse := h.bookingService.BookARoom(bookingRequest)

	return c.JSON(status, webResponse)
}

// BookingReport godoc
// @Summary Get booking report for a user
// @Description Retrieves all bookings made by a specific user.
// @Tags Booking
// @Produce json
// @Success 200 {object} map[string]interface{} "Booking report retrieved successfully"
// @Failure 401 {object} map[string]interface{} "User not authorized"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /booking/report [get]
func (h *bookingController) BookingReport(c echo.Context) error {
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status" : http.StatusUnauthorized,
			"message": fmt.Sprintf("User ID is not valid!"),
		})
	}

	status, webResponse := h.bookingService.BookingReport(userID)

	return c.JSON(status, webResponse)
}
