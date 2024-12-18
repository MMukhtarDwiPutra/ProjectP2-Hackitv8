package repository

import(
	"P2-Hacktiv8/entity"
	"gorm.io/gorm"         // ORM (Object Relational Mapping) Gorm untuk interaksi dengan database.
)

type UserRepository interface{
	CreateUser(user entity.User) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	UpdateBalance(user entity.TopUpRequest) (*entity.TopUpResponse, error)
}

type userRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository{
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user entity.User) (*entity.User, error){
	// Menyimpan data mahasiswa ke database menggunakan GORM.
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*entity.User, error){
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil{
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateBalance(user entity.TopUpRequest) (*entity.TopUpResponse, error){
	// Increment the balance
	if err := r.db.Model(&entity.User{}).
		Where("user_id = ?", user.UserID).
		Update("balance", gorm.Expr("balance + ?", user.Balance)).Error; err != nil {
		return nil, err
	}

	// Retrieve the updated user
	var updatedUser entity.User

	if err := r.db.Where("user_id = ?", user.UserID).First(&updatedUser).Error; err != nil {
		return nil, err
	}

	userResp := entity.TopUpResponse{
		UserID: updatedUser.UserID,
		FullName: updatedUser.FullName,
		Balance: updatedUser.Balance,
	}

	return &userResp, nil
}