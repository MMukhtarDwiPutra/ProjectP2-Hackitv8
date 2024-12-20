package entity

type User struct{
	UserID int `gorm:"primaryKey;autoIncrement" json:"user_id"`
	FullName string `json:"full_name"`
	Email string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Balance float32 `json:"balance"`
	JwtToken string `json:"jwt_token"`
	IsActivated string `json:"is_activated"`
}

type RegisterRequest struct{
	FullName string `json:"full_name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Balance float32 `json:"balance"`
}

type UserResponse struct{
	UserID int `json:"user_id"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Balance float32 `json:"balance"`
}

type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

type BalanceRequest struct{
	UserID int `json:"user_id" validate:"required"`
	Balance float32 `json:"top_up_balance" validate:"required"`
}

type BalanceResponse struct{
	UserID int `json:"user_id"`
	FullName string `json:"full_name"`
	Balance float32 `json:"balance"`
}

type TopUpResponse struct {
	InvoiceID    string  `json:"invoice_id"`
	Status       string  `json:"status"`
	Description  string  `json:"description"`
	Url          string  `json:"url"`
	MerchantName string  `json:"merchant_name"`
	ExternalID string `json:"external_id"`
}

type WebhookXenditPayment struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	InvoiceID string   `json:"invoice_id" gorm:"not null"`
	UserIDApp   int    `json:"user_id_app" gorm:"not null"`
	Status      string `json:"status" gorm:"not null"`
}