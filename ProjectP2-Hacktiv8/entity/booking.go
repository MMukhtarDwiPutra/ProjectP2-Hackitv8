package entity

type Booking struct{
	BookingID int `gorm:"primaryKey;" json:"booking_id"`
	UserID int `json:"user_id"`
	RoomID int `json:"room_id"`
	DateIn string `json:"date_in"`
	DateOut string `json:"date_out"`
}

type BookingRequest struct{
	UserID int `json:"user_id"`
	RoomID int `json:"room_id" validate:"required"`
	DateIn string `json:"date_in" validate:"required"`
	DateOut string `json:"date_out" validate:"required"`
}

type Invoice struct {
	ID         string `json:"id"`
	InvoiceUrl string `json:"invoice_url"`
}