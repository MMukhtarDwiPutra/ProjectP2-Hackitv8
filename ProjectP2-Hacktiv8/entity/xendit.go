package entity

import "time"

type InvoiceWebhookPayload struct {
    ID                     string  `json:"id"`
    ExternalID             string  `json:"external_id"`
    UserID                 string  `json:"user_id"`
    IsHigh                 bool    `json:"is_high"`
    PaymentMethod          string  `json:"payment_method"`
    Status                 string  `json:"status"`
    MerchantName           string  `json:"merchant_name"`
    Amount                 int     `json:"amount"`
    PaidAmount             int     `json:"paid_amount"`
    BankCode               string  `json:"bank_code"`
    PaidAt                 string  `json:"paid_at"`
    PayerEmail             string  `json:"payer_email"`
    Description            string  `json:"description"`
    AdjustedReceivedAmount int     `json:"adjusted_received_amount"`
    FeesPaidAmount         int     `json:"fees_paid_amount"`
}

type PaymentResponse struct {
	BusinessID      string `json:"business_id"`      // ID of the business
	IsLivemode      bool   `json:"is_livemode"`      // Indicates if the mode is live
	ChannelCode     string `json:"channel_code"`     // Code for the payment channel
	Name            string `json:"name"`            // Name of the payment method
	Currency        string `json:"currency"`        // Currency used for the payment
	ChannelCategory string `json:"channel_category"`// Category of the payment channel
	IsEnabled       bool   `json:"is_enabled"`      // Indicates if the payment method is enabled
}

type WebhookEvent struct {
	Event      string    `json:"event"`
	Data       EventData `json:"data"`
	Created    time.Time `json:"created"`
	BusinessID string    `json:"business_id"`
	APIVersion *string   `json:"api_version"`
}

type EventData struct {
	ID                string        `json:"id"`
	Amount            int           `json:"amount"`
	Status            string        `json:"status"`
	Country           string        `json:"country"`
	Created           time.Time     `json:"created"`
	Updated           time.Time     `json:"updated"`
	Currency          string        `json:"currency"`
	Metadata          Metadata      `json:"metadata"`
	CustomerID        string        `json:"customer_id"`
	ReferenceID       string        `json:"reference_id"`
	PaymentMethod     PaymentMethod `json:"payment_method"`
	Description       *string       `json:"description"`
	FailureCode       *string       `json:"failure_code"`
	PaymentDetail     *string       `json:"payment_detail"`
	ChannelProperties *string       `json:"channel_properties"`
	PaymentRequestID  string        `json:"payment_request_id"`
}

type Metadata struct {
	SKU string `json:"sku"`
}

type PaymentMethod struct {
	ID                  string         `json:"id"`
	Card                *string        `json:"card"`
	Type                string         `json:"type"`
	Status              string         `json:"status"`
	Created             time.Time      `json:"created"`
	EWallet             *string        `json:"ewallet"`
	QRCode              *string        `json:"qr_code"`
	Updated             time.Time      `json:"updated"`
	Metadata            *string        `json:"metadata"`
	Description         *string        `json:"description"`
	Reusability         string         `json:"reusability"`
	DirectDebit         DirectDebit    `json:"direct_debit"`
	ReferenceID         string         `json:"reference_id"`
	VirtualAccount      *string        `json:"virtual_account"`
	OverTheCounter      *string        `json:"over_the_counter"`
	DirectBankTransfer  *string        `json:"direct_bank_transfer"`
}

type DirectDebit struct {
	Type           string          `json:"type"`
	DebitCard      *string         `json:"debit_card"`
	BankAccount    BankAccount     `json:"bank_account"`
	ChannelCode    string          `json:"channel_code"`
	ChannelProps   ChannelProps    `json:"channel_properties"`
}

type BankAccount struct {
	BankAccountHash           string `json:"bank_account_hash"`
	MaskedBankAccountNumber   string `json:"masked_bank_account_number"`
}

type ChannelProps struct {
	FailureReturnURL string `json:"failure_return_url"`
	SuccessReturnURL string `json:"success_return_url"`
}