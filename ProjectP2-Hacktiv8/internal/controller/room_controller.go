package controller

import (
	// "P2-Hacktiv8/entity"
	"P2-Hacktiv8/internal/service"
	// "fmt"
	// "github.com/go-playground/validator/v10" // Validator untuk memvalidasi request body.
	"github.com/labstack/echo/v4"            // Import Echo framework untuk pengelolaan HTTP API.
	// "net/http"
)

type roomController struct {
	roomService service.RoomService
}

// NewRoomController creates a new instance of roomController.
func NewRoomController(roomService service.RoomService) *roomController {
	return &roomController{roomService}
}

func (h *roomController) GetAllRooms(c echo.Context) error{
	status, webResponse := h.roomService.GetAllRooms()

	return c.JSON(status, webResponse)
}