package controller

import (
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/internal/service"
	"fmt"
	// "github.com/go-playground/validator/v10" // Validator untuk memvalidasi request body.
	"github.com/labstack/echo/v4"            // Import Echo framework untuk pengelolaan HTTP API.
	"net/http"
)

type bookingController struct {
	bookingService service.BookingService
}

// NewBookingController creates a new instance of bookingController.
func NewBookingController(bookingService service.BookingService) *bookingController {
	return &bookingController{bookingService}
}

func (h *bookingController) BookARoom(c echo.Context) error{
	var bookingRequest entity.BookingRequest

	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status" : http.StatusUnauthorized,
			"message": fmt.Sprintf("User ID is not valid!"),
		})
	}

	// Melakukan bind request body ke struct.
	if err := c.Bind(&bookingRequest); err != nil {
		// Mengembalikan respons jika request body tidak valid.
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status" : http.StatusInternalServerError,
			"message": fmt.Sprintf("Booking create fail: %v",err),
		})
	}

	bookingRequest.UserID = userID
	// Memvalidasi data request menggunakan validator.
	err := validate.Struct(bookingRequest)
	if err != nil {
		// Mengembalikan respons jika validasi gagal.
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status" : http.StatusBadRequest,
			"message": fmt.Sprintf("Need: %v",err),
		})
	}

	fmt.Println(bookingRequest)

	status, webResponse := h.bookingService.BookARoom(bookingRequest)
	// status := 200
	// webResponse := map[string]string{
	// 	"messsage": "testing",
	// }

	return c.JSON(status, webResponse)
}

func (h *bookingController) BookingReport(c echo.Context) error{
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