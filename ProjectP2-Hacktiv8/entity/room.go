package entity

type Room struct{
	RoomID int `gorm:"primaryKey" json:"room_id"`
	Price float32 `json:"price"`
	RoomType string `json:"room_type"`
	AvailabilityStatus string `json:"availability_status"`
}