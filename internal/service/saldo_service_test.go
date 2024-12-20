package service

import (
	"P2-Hacktiv8/entity"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

var testSaldoService = NewSaldoService(userRepoMock)

// func TestTopUp_Success(t *testing.T) {
//     topUpRequest := entity.BalanceRequest{
//         UserID:  1,
//         Balance: 1000,  // Top-up amount
//     }

//     mockUser := &entity.User{
//         UserID:  1,
//         Balance: 5000,  // Initial balance
//     }

//     // The balance after top-up should be 6000
//     expectedBalanceResponse := &entity.BalanceResponse{
//         UserID:  1,
//         Balance: 6000,  // Updated balance
//     }

//     // Setup mocks
//     userRepoMock.Mock.On("GetUserById", topUpRequest.UserID).Return(mockUser, nil).Once()

//     // Ensure UpdateBalance is called with topUpRequest, but we expect a BalanceRequest object here
//     userRepoMock.Mock.On("UpdateBalance", topUpRequest).Return(expectedBalanceResponse, nil).Once()

//     // Call the method under test
//     status, response := testSaldoService.TopUp(topUpRequest)

//     // Check that the response status and message are as expected
//     assert.Equal(t, http.StatusOK, status)
//     assert.Equal(t, "Successfully top up balance", response["message"])

//     // Check that the returned data matches the expected balance response
//     actualBalanceResponse := response["data"].(*entity.BalanceResponse)
//     assert.Equal(t, expectedBalanceResponse.UserID, actualBalanceResponse.UserID)
//     assert.Equal(t, expectedBalanceResponse.Balance, actualBalanceResponse.Balance)

//     // Optionally, print the response for debugging
//     fmt.Printf("Returned balance response: %+v\n", actualBalanceResponse)

//     // Verify that the mocks were called as expected
//     userRepoMock.Mock.AssertExpectations(t)
// }

func TestTopUp_Error_GetUserById(t *testing.T) {
	// Prepare the mock request
	topUpRequest := entity.BalanceRequest{
		UserID: 1,
		Balance: 1000,
	}

	// Mock an error for GetUserById
	userRepoMock.Mock.On("GetUserById", topUpRequest.UserID).Return(nil, fmt.Errorf("User not found")).Once()

	// Call the service method
	status, response := testSaldoService.TopUp(topUpRequest)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, "Top up fail in database: User not found", response["message"])
}

// func TestTopUp_Error_UpdateBalance(t *testing.T) {
// 	// Prepare the mock request
// 	topUpRequest := entity.BalanceRequest{
// 		UserID: 1,
// 		Balance: 1000,
// 	}

// 	// Mock the response for GetUserById
// 	mockUser := &entity.User{
// 		UserID: 1,
// 		Balance: 5000, // Initial balance should be 5000
// 	}

// 	// Mock the GetUserById method to return the mockUser
// 	userRepoMock.Mock.On("GetUserById", topUpRequest.UserID).Return(mockUser, nil).Once()

// 	// Mock an error for UpdateBalance (simulate the failure scenario)
// 	userRepoMock.Mock.On("UpdateBalance", topUpRequest).Return(nil, fmt.Errorf("Error updating balance")).Once()

// 	// Call the service method
// 	status, response := testSaldoService.TopUp(topUpRequest)

// 	// Assertions
// 	assert.Equal(t, http.StatusInternalServerError, status)
// 	assert.Equal(t, "Top up fail in database: Error updating balance", response["message"])

// 	// Ensure balance is not modified by the failed top-up
// 	actualBalanceResponse := response["data"]
// 	assert.Nil(t, actualBalanceResponse)
// }

