package repository

import(
	"P2-Hacktiv8/entity"
	"gorm.io/gorm"         // ORM (Object Relational Mapping) Gorm untuk interaksi dengan database.
)

type UserRepository interface{
	CreateUser(user entity.User) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUserById(id int) (*entity.User, error)
	UpdateBalance(user entity.BalanceRequest) (*entity.BalanceResponse, error)
	UpdateIsActivatedById(id int, isActivated string) (*entity.User, error)
	GetUserByEmailAndToken(email, token string) (*entity.User, error)
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

func (r *userRepository) GetUserById(id int) (*entity.User, error){
	var user entity.User
	if err := r.db.Where("user_id = ?", id).First(&user).Error; err != nil{
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateBalance(user entity.BalanceRequest) (*entity.BalanceResponse, error){
	// Increment the balance
	if err := r.db.Model(&entity.User{}).
		Where("user_id = ?", user.UserID).
		Update("balance", gorm.Expr("?", user.Balance)).Error; err != nil {
		return nil, err
	}

	// Retrieve the updated user
	var updatedUser entity.User

	if err := r.db.Where("user_id = ?", user.UserID).First(&updatedUser).Error; err != nil {
		return nil, err
	}

	userResp := entity.BalanceResponse{
		UserID: updatedUser.UserID,
		FullName: updatedUser.FullName,
		Balance: updatedUser.Balance,
	}

	return &userResp, nil
}

func (r *userRepository) UpdateIsActivatedById(id int, isActivated string) (*entity.User, error){
	if err := r.db.Model(&entity.User{}).
		Where("user_id = ?", id).
		Update("is_activated", gorm.Expr("?", isActivated)).Error; err != nil {
		return nil, err
	}

	// Retrieve the updated user
	var updatedUser entity.User

	if err := r.db.Where("user_id = ?", id).First(&updatedUser).Error; err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (r *userRepository) GetUserByEmailAndToken(email, token string) (*entity.User, error){
	var user entity.User
	if err := r.db.Where("email = ? and token = ?", email, token).First(&user).Error; err != nil{
		return nil, err
	}

	return &user, nil
}