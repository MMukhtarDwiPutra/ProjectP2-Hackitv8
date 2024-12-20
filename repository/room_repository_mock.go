package repository

import (
	"P2-Hacktiv8/entity"

	"github.com/stretchr/testify/mock"
	"fmt"
)

type RoomRepositoryMock struct {
	Mock mock.Mock
}

func (m *RoomRepositoryMock) GetAllRooms() (*[]entity.Room, error) {
	res := m.Mock.Called()

	if res.Get(0) == nil {
		return nil, res.Error(1)
	}

	rooms := res.Get(0).(*[]entity.Room)
	return rooms, res.Error(1)
}

func (m *RoomRepositoryMock) GetRoomById(id int) (*entity.Room, error) {
    res := m.Mock.Called(id)

    if res.Get(0) == nil {
        return nil, res.Error(1)
    }

    room, ok := res.Get(0).(*entity.Room)
    if !ok {
        return nil, fmt.Errorf("expected *entity.Room but got %T", res.Get(0))
    }

    return room, nil
}

func (m *RoomRepositoryMock) UpdateRoomAvailability(roomID int, avail string) (*entity.Room, error) {
	args := m.Mock.Called(roomID, avail)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Room), args.Error(1)
	}
	return nil, args.Error(1)
}

