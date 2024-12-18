package service

import(
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/repository"
	"net/http"
	"gorm.io/gorm"
	"errors"
	"fmt"
)

type BookingService interface{
	BookARoom(bookingRequest entity.BookingRequest) (int, map[string]interface{})
	BookingReport(userID int) (int, map[string]interface{})
}

type bookingService struct{
	bookingRepository repository.BookingRepository
	userRepository repository.UserRepository
	roomRepository repository.RoomRepository
}

func NewBookingService(bookingRepository repository.BookingRepository, userRepository repository.UserRepository, roomRepository repository.RoomRepository) *bookingService{
	return &bookingService{bookingRepository, userRepository, roomRepository}
}

func (c *bookingService) BookARoom(bookingRequest entity.BookingRequest) (int, map[string]interface{}) {
	// Attempt to retrieve the room first
	room, err := c.roomRepository.GetRoomById(bookingRequest.RoomID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, map[string]interface{}{
				"status":  http.StatusNotFound,
				"message": "Room not found!",
			}
		}
		return http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Booking creation failed: %v", err),
		}
	}

	// Attempt to retrieve the user
	user, err := c.userRepository.GetUserById(bookingRequest.UserID)
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, map[string]interface{}{
				"status":  http.StatusNotFound,
				"message": "User not found!",
			}
		}
		return http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Booking creation failed: %v", err),
		}
	}

	// Validate user balance
	downPayment := room.Price * 40 / 100
	if user.Balance < downPayment {
		return http.StatusPaymentRequired, map[string]interface{}{
			"status":  http.StatusPaymentRequired,
			"message": "Booking creation failed: insufficient balance.",
		}
	}

	// Create the booking
	booking := entity.Booking{
		UserID: bookingRequest.UserID,
		RoomID: bookingRequest.RoomID,
	}
	bookingResp, err := c.bookingRepository.CreateBooking(booking)
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Booking creation failed: %v", err),
		}
	}
	
	// Deduct balance
	user.Balance -= downPayment
	newUserBalance := entity.BalanceRequest{
		UserID:  user.UserID,
		Balance: user.Balance,
	}
	_, err = c.userRepository.UpdateBalance(newUserBalance)
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Booking creation failed: %v", err),
		}
	}

	return http.StatusCreated, map[string]interface{}{
		"status":  http.StatusCreated,
		"message": "Booking created successfully.",
		"data":    bookingResp,
	}
}

func (c *bookingService) BookingReport(userID int) (int, map[string]interface{}){
	bookingReport, err := c.bookingRepository.GetBookingByUserId(userID)
	if err != nil{
		return http.StatusInternalServerError, map[string]interface{}{
			"status" : http.StatusInternalServerError,
			"message": fmt.Sprintf("Booking create fail: %v",err),
		}
	}

	return http.StatusOK, map[string]interface{}{
		"status" : http.StatusOK,
		"message": "Getting booking report successfully",
		"data" : bookingReport,
	}
}