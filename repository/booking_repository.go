package repository

import(
	"P2-Hacktiv8/entity"
	"gorm.io/gorm"         // ORM (Object Relational Mapping) Gorm untuk interaksi dengan database.
)

type BookingRepository interface{
	CreateBooking(booking entity.Booking) (*entity.Booking, error)
	GetBookingByUserId(userID int) (*[]entity.Booking, error)
}

type bookingRepository struct{
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *bookingRepository{
	return &bookingRepository{db}
}

func (r *bookingRepository) CreateBooking(booking entity.Booking) (*entity.Booking, error){
	// Menyimpan data order ke database menggunakan GORM.
	if err := r.db.Create(&booking).Error; err != nil {
		return nil, err
	}

	return &booking, nil
}

func (r *bookingRepository) GetBookingByUserId(userID int) (*[]entity.Booking, error){
	var bookingReport []entity.Booking

	if err := r.db.Where("user_id = ?", userID).Find(&bookingReport).Error; err != nil{
		return nil, err
	}

	return &bookingReport, nil
}