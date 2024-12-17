package entity

type User struct{
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	FullName string `json:"full_name"`
	Email string `gorm:"unique" json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Age int `json:"age"`
}

type RegisterRequest struct{
	FullName string `json:"full_name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Age int `json:"age" validate:"required"`
}

type RegisterResponse struct{
	ID int `json:"id"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Username string `json:"username"`
	Age int `json:"age"`
}

type LoginRequest struct{
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}