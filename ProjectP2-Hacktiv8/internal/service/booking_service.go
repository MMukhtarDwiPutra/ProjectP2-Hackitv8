package service

import(
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/repository"
	"P2-Hacktiv8/utils"
	"net/http"
	"gorm.io/gorm"
	"errors"
	"fmt"
	"time"
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
	if bookingRequest.DateIn == "" || bookingRequest.DateOut == "" {
	    return http.StatusBadRequest, map[string]interface{}{"message": "date_in and date_out cannot be empty"}
	}

	dateIn, err := time.Parse("2006-01-02", bookingRequest.DateIn)
	if err != nil {
	    return http.StatusBadRequest, map[string]interface{}{"message": "Invalid format for date_in"}
	}

	dateOut, err := time.Parse("2006-01-02", bookingRequest.DateOut)
	if err != nil {
	    return http.StatusBadRequest, map[string]interface{}{"message": "Invalid format for date_out"}
	}

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

	fmt.Println(room)
	if (room.AvailabilityStatus != "Available") {
		return http.StatusConflict, map[string]interface{}{
		    "status":  http.StatusConflict,
		    "message": "The room is not available because it has already been booked.",
		}
	}

	_, err = c.roomRepository.UpdateRoomAvailability(bookingRequest.RoomID, "Booked")
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Booking update status failed: %v", err),
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

	if user.Balance < room.Price {
		return http.StatusPaymentRequired, map[string]interface{}{
			"status":  http.StatusPaymentRequired,
			"message": "Booking creation failed: insufficient balance.",
		}
	}

	// Create the booking
	booking := entity.Booking{
		UserID: bookingRequest.UserID,
		RoomID: bookingRequest.RoomID,
		DateIn: dateIn.Format("2006-01-02"),
		DateOut: dateOut.Format("2006-01-02"),
	}
	bookingResp, err := c.bookingRepository.CreateBooking(booking)
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Booking creation failed: %v", err),
		}
	}
	
	// Deduct balance
	user.Balance -= room.Price
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

	// Send email notification
	to := user.Email
	subject := fmt.Sprintf("Booking Room %v Successfully", room.RoomID)

	// Format the content to include detailed booking information
	content := fmt.Sprintf(`
	Dear %s,

	Thank you for booking with us! Here are the details of your booking:

	- Room ID: %v
	- Price: $%.2f
	- Booking Date: %v

	We hope you enjoy your stay. If you have any questions or need further assistance, feel free to contact us.

	Best regards,
	Your Booking Team
	`, user.FullName, room.RoomID, room.Price, time.Now().Format("January 2, 2006, 3:04 PM"))

	// Send the email
	utils.SendEmailNotification(to, subject, content)

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