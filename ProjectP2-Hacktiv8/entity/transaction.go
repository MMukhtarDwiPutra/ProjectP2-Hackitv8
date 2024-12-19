package entity

import(
	"time"
)

type Transaction struct {
    UserID     int       `json:"user_id"`
    Amount     float32   `json:"amount"`       // Assuming Amount is a float64 type
    ExternalID string    `json:"external_id"`  // ExternalID as string
    VAAccount  string    `json:"va_account"`   // VAAccount as string
    Status     string    `json:"status"`
    CreatedAt  time.Time `json:"created_at"`
}