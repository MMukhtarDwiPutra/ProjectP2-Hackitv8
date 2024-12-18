package entity

type Booking struct{
	BookingID int `gorm:"primaryKey;" json:"booking_id"`
	UserID int `json:"user_id"`
	RoomID int `json:"room_id"`
}

type BookingRequest struct{
	UserID int `json:"user_id"`
	RoomID int `json:"room_id" validate:"required""`
}