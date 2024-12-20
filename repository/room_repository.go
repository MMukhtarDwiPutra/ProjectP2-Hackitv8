package repository

import(
	"P2-Hacktiv8/entity"
	"gorm.io/gorm"         // ORM (Object Relational Mapping) Gorm untuk interaksi dengan database.
)

type RoomRepository interface{
	GetAllRooms() (*[]entity.Room, error)
	GetRoomById(id int) (*entity.Room, error)
	UpdateRoomAvailability(roomID int, avail string) (*entity.Room, error)
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

func (r *roomRepository) UpdateRoomAvailability(roomID int, avail string) (*entity.Room, error){
	// Increment the balance
	if err := r.db.Model(&entity.Room{}).
		Where("room_id = ?", roomID).
		Update("availability_status", gorm.Expr("?", avail)).Error; err != nil {
		return nil, err
	}

	// Retrieve the updated user
	var updatedRoom entity.Room

	if err := r.db.Where("room_id = ?", roomID).First(&updatedRoom).Error; err != nil {
		return nil, err
	}

	return &updatedRoom, nil
}