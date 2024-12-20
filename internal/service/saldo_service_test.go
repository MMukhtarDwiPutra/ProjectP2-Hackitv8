package service

import (
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/utils"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"bou.ke/monkey"
)

var testSaldoService = NewSaldoService(userRepoMock)

func TestTopUp_Success(t *testing.T) {
	monkey.Patch(utils.SendEmailNotification, func(to, subject, content string) {
		// Mock behavior: Do nothing or log the invocation
	})
	defer monkey.Unpatch(utils.SendEmailNotification) // Unpatch after the test

    // Arrange
    userID := 1
    topUpAmount := 1000
    initialBalance := 5000

    topUpRequest := entity.BalanceRequest{
        UserID:  userID,
        Balance: float32(topUpAmount),
    }

    mockUser := &entity.User{
	    UserID:  userID,
	    Balance: float32(initialBalance),
	    Email:   "testing@gmail.com", // Ensure this is set correctly
	}

    // Mocking the GetLastIDXendit method to return a valid invoice ID
    mockInvoiceID := 1234
    userRepoMock.Mock.On("GetLastIDXendit").Return(&mockInvoiceID, nil).Once()

    // Create a mock behavior for CreateInvoice using monkey patch
    monkey.Patch(utils.CreateInvoice, func(mockUser entity.User, topUpRequest entity.BalanceRequest, externalID string) (*entity.InvoiceResponse, error) {
        return &entity.InvoiceResponse{
            ID:          externalID,
            Status:      "Success",
            Description: "Top Up Invoice",
            InvoiceURL:  "http://example.com/invoice/" + externalID,
            MerchantName: "Sample Merchant",
            ExternalID:  externalID,
        }, nil
    })
    // Ensure to unpatch the method after the test
    defer monkey.Unpatch(utils.CreateInvoice)

    // Set up other mocks
    userRepoMock.Mock.On("GetUserById", userID).Return(mockUser, nil).Once()

    // Correct mock setup:
	mockXenditPayment := entity.WebhookXenditPayment{
	    InvoiceID: "INV1234",
	    UserIDApp: 1,
	    Status:    "Success",
	}

	userRepoMock.Mock.On("CreateXenditHistory", mockXenditPayment).Return(&mockXenditPayment, nil).Once()

    // Create the service with the mocked dependencies
    testSaldoService := NewSaldoService(userRepoMock)

    // Act
    status, response := testSaldoService.TopUp(topUpRequest)

    // Assert
    assert.Equal(t, http.StatusOK, status)
    assert.Equal(t, "Successfully top up balance", response["message"])

    actualResponse, ok := response["data"].(entity.TopUpResponse)

    expectedResponse := entity.TopUpResponse{
    	InvoiceID: "INV1234",
    	Status: "Success",
    	Description: "Top Up Invoice",
    	Url:  "http://example.com/invoice/INV1234",
        MerchantName: "Sample Merchant",
        ExternalID:  "INV1234",

    }

    assert.True(t, ok)
	assert.Equal(t, expectedResponse.InvoiceID, actualResponse.InvoiceID)
	assert.Equal(t, expectedResponse.Status, actualResponse.Status)
	assert.Equal(t, expectedResponse.Description, actualResponse.Description)
	assert.Equal(t, expectedResponse.Url, actualResponse.Url)
	assert.Equal(t, expectedResponse.MerchantName, actualResponse.MerchantName)
	assert.Equal(t, expectedResponse.ExternalID, actualResponse.ExternalID)

    // Verify mock expectations
    userRepoMock.Mock.AssertExpectations(t)
}

func TestTopUp_Error_GetUserById(t *testing.T) {
	// Prepare the mock request
	topUpRequest := entity.BalanceRequest{
		UserID: 1,
		Balance: 1000,
	}

	// Mock an error for GetUserById
	userRepoMock.Mock.On("GetUserById", topUpRequest.UserID).Return(nil, gorm.ErrRecordNotFound).Once()

	// Call the service method
	status, response := testSaldoService.TopUp(topUpRequest)

	// Assertions
	assert.Equal(t, http.StatusNotFound, status)
	assert.Equal(t, "User not found", response["message"])
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

