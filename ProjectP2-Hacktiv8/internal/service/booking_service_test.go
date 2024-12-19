package service

import (
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/repository"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"fmt"
)

// Mock Repositories
var bookingRepoMock = &repository.BookingRepositoryMock{Mock: mock.Mock{}}
var userRepoMock = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var roomRepoMock = &repository.RoomRepositoryMock{Mock: mock.Mock{}}

// Service under test
var testBookingService = NewBookingService(
	bookingRepoMock,
	userRepoMock,
	roomRepoMock,
)

// Test: Room not found
func TestBookARoom_RoomNotFound(t *testing.T) {
	roomID := 1
	userID := 1
	bookingRequest := entity.BookingRequest{RoomID: roomID, UserID: userID}

	roomRepoMock.Mock.On("GetRoomById", roomID).Return(nil, gorm.ErrRecordNotFound).Once()

	status, response := testBookingService.BookARoom(bookingRequest)

	assert.Equal(t, http.StatusNotFound, status)
	assert.Equal(t, "Room not found!", response["message"])
}

func TestBookARoom_UserNotFound(t *testing.T) {
    roomID := 1
    userID := 1
    bookingRequest := entity.BookingRequest{RoomID: roomID, UserID: userID}

    // Mock a valid room response
    mockRoom := &entity.Room{RoomID: roomID, Price: 1000}
    roomRepoMock.Mock.On("GetRoomById", roomID).Return(mockRoom, nil).Once()

    // Mock a user not found error
    userRepoMock.Mock.On("GetUserById", userID).Return(nil, gorm.ErrRecordNotFound).Once()

    // Call the service method
    status, response := testBookingService.BookARoom(bookingRequest)

    // Assertions
    assert.Equal(t, http.StatusNotFound, status)
    assert.Equal(t, "User not found!", response["message"])
}

// Test: Insufficient balance
func TestBookARoom_InsufficientBalance(t *testing.T) {
	roomID := 1
	userID := 1
	bookingRequest := entity.BookingRequest{RoomID: roomID, UserID: userID}

	mockRoom := &entity.Room{RoomID: roomID, Price: 1000}
	mockUser := &entity.User{UserID: userID, Balance: 200}

	roomRepoMock.Mock.On("GetRoomById", roomID).Return(mockRoom, nil).Once()
	userRepoMock.Mock.On("GetUserById", userID).Return(mockUser, nil).Once()

	status, response := testBookingService.BookARoom(bookingRequest)

	assert.Equal(t, http.StatusPaymentRequired, status)
	assert.Equal(t, "Booking creation failed: insufficient balance.", response["message"])
}

// Test: Successful booking
func TestBookARoom_Success(t *testing.T) {
	roomID := 1
	userID := 1
	bookingRequest := entity.BookingRequest{RoomID: roomID, UserID: userID}

	mockRoom := &entity.Room{RoomID: roomID, Price: 1000}
	mockUser := &entity.User{UserID: userID, Balance: 5000}
	mockBooking := &entity.Booking{RoomID: roomID, UserID: userID}
	mockBalanceResponse := &entity.BalanceResponse{
	    UserID:  userID,
	    Balance: 4900,  // Example updated balance after deduction
	}

	roomRepoMock.Mock.On("GetRoomById", roomID).Return(mockRoom, nil).Once()
	userRepoMock.Mock.On("GetUserById", userID).Return(mockUser, nil).Once()
	bookingRepoMock.Mock.On("CreateBooking", mock.Anything).Return(*mockBooking, nil).Once()
	userRepoMock.Mock.On("UpdateBalance", mock.Anything).Return(mockBalanceResponse, nil).Once()

	status, response := testBookingService.BookARoom(bookingRequest)

	assert.Equal(t, http.StatusCreated, status)
	assert.Equal(t, "Booking created successfully.", response["message"])
	assert.NotNil(t, response["data"])
}

func TestBookingReport_Success(t *testing.T) {
	userID := 1

	// Mock a successful booking report response
	mockBookingReport := []entity.Booking{
		{RoomID: 1, UserID: userID},
		{RoomID: 2, UserID: userID},
	}

	// Mock the GetBookingByUserId call to return the mocked booking report
	bookingRepoMock.Mock.On("GetBookingByUserId", userID).Return(mockBookingReport, nil).Once()

	// Call the service method
	status, response := testBookingService.BookingReport(userID)

	// Assertions
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, "Getting booking report successfully", response["message"])
	assert.NotNil(t, response["data"])
	assert.Equal(t, &mockBookingReport, response["data"])
}

func TestBookingReport_Error(t *testing.T) {
	userID := 1

	// Mock an error response from GetBookingByUserId
	bookingRepoMock.Mock.On("GetBookingByUserId", userID).Return(nil, fmt.Errorf("Some error occurred")).Once()

	// Call the service method
	status, response := testBookingService.BookingReport(userID)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, "Booking create fail: Some error occurred", response["message"])
	assert.Nil(t, response["data"])
}
