package repository

import (
	"P2-Hacktiv8/entity"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (m *UserRepositoryMock) CreateUser(user entity.User) (*entity.User, error) {
	res := m.Mock.Called(user)

	if res.Get(0) == nil {
		return nil, res.Error(1)
	}

	createdUser := res.Get(0).(*entity.User)
	return createdUser, res.Error(1)
}

func (m *UserRepositoryMock) GetUserByEmail(email string) (*entity.User, error) {
	res := m.Mock.Called(email)

	if res.Get(0) == nil {
		return nil, res.Error(1)
	}

	user := res.Get(0).(*entity.User)
	return user, res.Error(1)
}

func (m *UserRepositoryMock) GetUserById(id int) (*entity.User, error) {
	res := m.Mock.Called(id)

	if res.Get(0) == nil {
		return nil, res.Error(1)
	}

	user := res.Get(0).(*entity.User)
	return user, res.Error(1)
}

func (m *UserRepositoryMock) UpdateBalance(user entity.BalanceRequest) (*entity.BalanceResponse, error) {
	res := m.Mock.Called(user)

	if res.Get(0) == nil {
		return nil, res.Error(1)
	}

	balanceResp := res.Get(0).(*entity.BalanceResponse)
	return balanceResp, res.Error(1)
}

func (m *UserRepositoryMock) UpdateIsActivatedById(id int, isActivated string) (*entity.User, error) {
	res := m.Mock.Called(id, isActivated)

	if res.Get(0) == nil {
		return nil, res.Error(1)
	}

	user := res.Get(0).(*entity.User)
	return user, res.Error(1)
}

func (m *UserRepositoryMock) GetUserByEmailAndToken(email, token string) (*entity.User, error) {
	res := m.Mock.Called(email, token)

	if res.Get(0) == nil {
		return nil, res.Error(1)
	}

	user := res.Get(0).(*entity.User)
	return user, res.Error(1)
}
