package service

import(
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/repository"
	"net/http"
	// "gorm.io/gorm"
	"fmt"
)

type BookingService interface{
	BookARoom(bookingRequest entity.BookingRequest) (int, map[string]interface{})
}

type bookingService struct{
	bookingRepository repository.BookingRepository
}

func NewBookingService(bookingRepository repository.BookingRepository) *bookingService{
	return &bookingService{bookingRepository}
}

func (c *bookingService) BookARoom(bookingRequest entity.BookingRequest) (int, map[string]interface{}){
	booking := entity.Booking{
		UserID: bookingRequest.UserID,
		RoomID: bookingRequest.RoomID,
	}

	bookingResp, err := c.bookingRepository.CreateBooking(booking)
	if err != nil{
		return http.StatusInternalServerError, map[string]interface{}{
			"status" : http.StatusInternalServerError,
			"message": fmt.Sprintf("Booking create fail: %v",err),
		}
	}

	return http.StatusCreated, map[string]interface{}{
		"status" : http.StatusCreated,
		"message": "Booking created successfully",
		"data" : bookingResp,
	}
}