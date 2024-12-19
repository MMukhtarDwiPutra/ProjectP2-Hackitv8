package service

import(
	"P2-Hacktiv8/repository"
	"net/http"
	"fmt"
)

type RoomService interface{
	GetAllRooms() (int, map[string]interface{})
}

type roomService struct{
	roomRepository repository.RoomRepository
}

func NewRoomService(roomRepository repository.RoomRepository) *roomService{
	return &roomService{roomRepository}
}

func (s *roomService) GetAllRooms() (int, map[string]interface{}){
	rooms, err := s.roomRepository.GetAllRooms()
	if err != nil{
		return http.StatusCreated, map[string]interface{}{
			"status" : http.StatusInternalServerError,
			"message": fmt.Sprintf("Getting all rooms fail: %v",err),
		}
	}

	return http.StatusOK, map[string]interface{}{
		"status" : http.StatusOK,
		"message": "Successfully getting all rooms",
		"data": rooms,
	}
}