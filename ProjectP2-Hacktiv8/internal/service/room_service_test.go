package service

import (
	"P2-Hacktiv8/entity"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

var testRoomService = NewRoomService(roomRepoMock)

func TestGetAllRooms_Success(t *testing.T) {
	// Mock the response for successful room retrieval
	mockRooms := []entity.Room{
		{RoomID: 1, Price: 1000},
		{RoomID: 2, Price: 1500},
	}

	// Mock the GetAllRooms call to return the mockRooms
	roomRepoMock.Mock.On("GetAllRooms").Return(&mockRooms, nil).Once()

	// Call the service method
	status, response := testRoomService.GetAllRooms()

	// Assertions
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, "Successfully getting all rooms", response["message"])
	assert.NotNil(t, response["data"])
	assert.Equal(t, &mockRooms, response["data"])
}

func TestGetAllRooms_Error(t *testing.T) {
	// Mock an error response from GetAllRooms
	roomRepoMock.Mock.On("GetAllRooms").Return(nil, fmt.Errorf("Some error occurred")).Once()

	// Call the service method
	status, response := testRoomService.GetAllRooms()

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, "Getting all rooms fail: Some error occurred", response["message"])
	assert.Nil(t, response["data"])
}
