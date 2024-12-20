package repository

import (
	"P2-Hacktiv8/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) CreateUser(user entity.User) (*entity.User, error) {
	args := m.Called(user)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *UserRepositoryMock) GetUserByEmail(email string) (*entity.User, error) {
	args := m.Called(email)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *UserRepositoryMock) GetUserById(id int) (*entity.User, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *UserRepositoryMock) UpdateBalance(user entity.BalanceRequest) (*entity.BalanceResponse, error) {
	args := m.Called(user)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.BalanceResponse), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *UserRepositoryMock) UpdateIsActivatedById(id int, isActivated string) (*entity.User, error) {
	args := m.Called(id, isActivated)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *UserRepositoryMock) GetUserByEmailAndToken(email, token string) (*entity.User, error) {
	args := m.Called(email, token)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *UserRepositoryMock) CreateXenditHistory(xenditWebhook entity.WebhookXenditPayment) (*entity.WebhookXenditPayment, error) {
	args := m.Called(xenditWebhook)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.WebhookXenditPayment), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *UserRepositoryMock) GetPaymentIdByInvoiceId(invoiceID string) (*entity.WebhookXenditPayment, error) {
	args := m.Called(invoiceID)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.WebhookXenditPayment), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *UserRepositoryMock) GetLastIDXendit() (*int, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).(*int), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *UserRepositoryMock) UpdateStatusWebhookXenditPayment(xenditWebhook entity.WebhookXenditPayment) (*entity.WebhookXenditPayment, error) {
	args := m.Called(xenditWebhook)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.WebhookXenditPayment), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestUserRepositoryMock(t *testing.T) {
	mockRepo := new(UserRepositoryMock)

	// Example: Mocking CreateUser
	user := entity.User{UserID: 1, Email: "test@example.com"}
	mockRepo.On("CreateUser", user).Return(&user, nil)

	createdUser, err := mockRepo.CreateUser(user)
	assert.NoError(t, err)
	assert.Equal(t, user, *createdUser)
	mockRepo.AssertExpectations(t)
}