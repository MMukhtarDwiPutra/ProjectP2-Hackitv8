package service

import (
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/repository"
	"P2-Hacktiv8/utils"
	"net/http"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"bou.ke/monkey"
)

// Mock Repositories
var (
	bookingRepoMock = &repository.BookingRepositoryMock{Mock: mock.Mock{}}
	userRepoMock    = &repository.UserRepositoryMock{Mock: mock.Mock{}}
	roomRepoMock    = &repository.RoomRepositoryMock{Mock: mock.Mock{}}

	testService = NewBookingService(bookingRepoMock, userRepoMock, roomRepoMock)
)

func TestBookARoom(t *testing.T) {
	monkey.Patch(utils.SendEmailNotification, func(to, subject, content string) {
		// Mock behavior: Do nothing or log the invocation
	})
	defer monkey.Unpatch(utils.SendEmailNotification) // Unpatch after the test

	t.Run("Room Not Found", func(t *testing.T) {
		roomID, userID := 1, 1
		bookingRequest := entity.BookingRequest{RoomID: roomID, UserID: userID, DateIn: "2024-10-23", DateOut: "2024-12-23"}

		roomRepoMock.Mock.On("GetRoomById", roomID).Return(nil, gorm.ErrRecordNotFound).Once()

		status, response := testService.BookARoom(bookingRequest)

		assert.Equal(t, http.StatusNotFound, status)
		assert.Equal(t, "Room not found!", response["message"])
	})

	t.Run("User Not Found", func(t *testing.T) {
		roomID, userID := 1, 1
		bookingRequest := entity.BookingRequest{RoomID: roomID, UserID: userID, DateIn: "2024-10-23", DateOut: "2024-12-23"}

		mockRoom := &entity.Room{RoomID: roomID, Price: 1000}
		roomRepoMock.Mock.On("GetRoomById", roomID).Return(mockRoom, nil).Once()
		userRepoMock.Mock.On("GetUserById", userID).Return(nil, gorm.ErrRecordNotFound).Once()

		status, response := testService.BookARoom(bookingRequest)

		assert.Equal(t, http.StatusNotFound, status)
		assert.Equal(t, "User not found!", response["message"])
	})

	t.Run("Insufficient Balance", func(t *testing.T) {
		roomID, userID := 1, 1
		bookingRequest := entity.BookingRequest{RoomID: roomID, UserID: userID, DateIn: "2024-10-23", DateOut: "2024-12-23"}

		mockRoom := &entity.Room{RoomID: roomID, Price: 1000}
		mockUser := &entity.User{UserID: userID, Balance: 200}

		roomRepoMock.Mock.On("GetRoomById", roomID).Return(mockRoom, nil).Once()
		userRepoMock.Mock.On("GetUserById", userID).Return(mockUser, nil).Once()

		status, response := testService.BookARoom(bookingRequest)

		assert.Equal(t, http.StatusPaymentRequired, status)
		assert.Equal(t, "Booking creation failed: insufficient balance.", response["message"])
	})

	t.Run("Successful Booking", func(t *testing.T) {
		roomID, userID := 1, 1
		bookingRequest := entity.BookingRequest{RoomID: roomID, UserID: userID, DateIn: "2024-10-23", DateOut: "2024-12-23"}

		mockRoom := &entity.Room{RoomID: roomID, Price: 1000, AvailabilityStatus: "Available"}
		mockUser := &entity.User{UserID: userID, Email: "testing@gmail.com", Balance: 5000}
		mockBooking := &entity.Booking{RoomID: roomID, UserID: userID}
		mockBalanceResponse := &entity.BalanceResponse{UserID: userID, Balance: 4900}
		mockUpdatedRoom := &entity.Room{RoomID: roomID, Price: 1000, AvailabilityStatus: "Booked"}

		roomRepoMock.Mock.On("GetRoomById", roomID).Return(mockRoom, nil).Once()
		userRepoMock.Mock.On("GetUserById", userID).Return(mockUser, nil).Once()
		bookingRepoMock.Mock.On("CreateBooking", mock.Anything).Return(*mockBooking, nil).Once()
		userRepoMock.Mock.On("UpdateBalance", mock.Anything).Return(mockBalanceResponse, nil).Once()
		roomRepoMock.Mock.On("UpdateRoomAvailability", roomID, "Booked").Return(mockUpdatedRoom, nil).Once()

		status, response := testService.BookARoom(bookingRequest)

		assert.Equal(t, http.StatusCreated, status)
		assert.Equal(t, "Booking created successfully.", response["message"])
		assert.NotNil(t, response["data"])
	})
}

func TestBookingReport(t *testing.T) {
	t.Run("Successful Report", func(t *testing.T) {
		userID := 1
		mockReport := []entity.Booking{
			{RoomID: 1, UserID: userID},
			{RoomID: 2, UserID: userID},
		}

		bookingRepoMock.Mock.On("GetBookingByUserId", userID).Return(mockReport, nil).Once()

		status, response := testService.BookingReport(userID)

		assert.Equal(t, http.StatusOK, status)
		assert.Equal(t, "Getting booking report successfully", response["message"])
		assert.NotNil(t, response["data"])
		assert.Equal(t, &mockReport, response["data"])
	})

	t.Run("Error Report", func(t *testing.T) {
		userID := 1

		bookingRepoMock.Mock.On("GetBookingByUserId", userID).Return(nil, fmt.Errorf("Some error occurred")).Once()

		status, response := testService.BookingReport(userID)

		assert.Equal(t, http.StatusInternalServerError, status)
		assert.Equal(t, "Booking create fail: Some error occurred", response["message"])
		assert.Nil(t, response["data"])
	})
}
