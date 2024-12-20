package controller

import (
	"P2-Hacktiv8/internal/service"
	"github.com/labstack/echo/v4" // Import Echo framework untuk pengelolaan HTTP API.
)

// roomController is the controller for room-related operations.
type roomController struct {
	roomService service.RoomService
}

// NewRoomController creates a new instance of roomController.
func NewRoomController(roomService service.RoomService) *roomController {
	return &roomController{roomService}
}

// GetAllRooms godoc
// @Summary Get all rooms
// @Description Retrieves a list of all available rooms.
// @Tags Room
// @Produce json
// @Success 200 {object} map[string]interface{} "Successfully fetched all rooms"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /rooms [get]
func (h *roomController) GetAllRooms(c echo.Context) error {
	status, webResponse := h.roomService.GetAllRooms()

	return c.JSON(status, webResponse)
}
