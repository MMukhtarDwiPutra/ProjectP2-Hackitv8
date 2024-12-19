package repository

import (
	"P2-Hacktiv8/entity"
	"github.com/stretchr/testify/mock"
)

type BookingRepositoryMock struct {
	Mock mock.Mock
}

func (m *BookingRepositoryMock) CreateBooking(booking entity.Booking) (*entity.Booking, error) {
	res := m.Mock.Called(booking)

	if res.Get(0) == nil {
		return nil, res.Error(1)
	}

	createdBooking := res.Get(0).(entity.Booking)
	return &createdBooking, res.Error(1)
}

func (m *BookingRepositoryMock) GetBookingByUserId(userID int) (*[]entity.Booking, error) {
	res := m.Mock.Called(userID)

	if res.Get(0) == nil {
		return nil, res.Error(1)
	}

	bookings := res.Get(0).([]entity.Booking)
	return &bookings, res.Error(1)
}
