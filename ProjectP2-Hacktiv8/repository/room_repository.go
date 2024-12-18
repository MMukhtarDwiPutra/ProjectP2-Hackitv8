package repository

import(
	"P2-Hacktiv8/entity"
	"gorm.io/gorm"         // ORM (Object Relational Mapping) Gorm untuk interaksi dengan database.
)

type RoomRepository interface{
	GetAllRooms() (*[]entity.Room, error)
	GetRoomById(id int) (*entity.Room, error)
}

type roomRepository struct{
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *roomRepository{
	return &roomRepository{db}
}

func (r *roomRepository) GetAllRooms() (*[]entity.Room, error){
	var rooms []entity.Room
	if err := r.db.Find(&rooms).Error; err != nil{
		return nil, err
	}

	return &rooms, nil
}

func (r *roomRepository) GetRoomById(id int) (*entity.Room, error){
	var room entity.Room
	if err := r.db.Where("room_id = ?", id).Find(&room).Error; err != nil{
		return nil, err
	}

	return &room, nil
}