package entity

import "time"

type DirectDebit struct {
	Type            string            `json:"type"`
	DebitCard       *DebitCard        `json:"debit_card"`
	BankAccount     interface{}       `json:"bank_account"`
	ChannelCode     string            `json:"channel_code"`
	ChannelProperties ChannelProperties `json:"channel_properties"`
	Cards Cards `json:"cards"`
}

type DebitCard struct {
	MobileNumber  string `json:"mobile_number"`
	CardLastFour  string `json:"card_last_four"`
	CardExpiry    string `json:"card_expiry"`
	Email         string `json:"email"`
}

type ChannelProperties struct {
	RedeemPoints string                  `json:"redeem_points"`
	MobileNumber  string `json:"mobile_number"`
	CardLastFour  string `json:"card_last_four"`
	CardExpiry    string `json:"card_expiry"`
	Email         string `json:"email"`
	SuccessRedirectURL string `json:"success_redirect_url"`
	Cards Cards `json:"cards"`
}

// InvoiceRequest represents the request body for creating an invoice
type InvoiceRequest struct {
	ExternalID      string `json:"external_id"`
	Amount          float32    `json:"amount"`
	Description     string `json:"description"`
	InvoiceDuration int    `json:"invoice_duration"`
	Currency        string `json:"currency"`
	ReferenceID     string `json:"reference_id"`
	CheckoutMethod  string `json:"checkout_method"`
}


type AllowedTerms struct {
	Issuer string `json:"issuer"`
	Terms  []int  `json:"terms"`
}

type InstallmentConfiguration struct {
	AllowInstallment  bool          `json:"allow_installment"`
	AllowFullPayment  bool          `json:"allow_full_payment"`
	AllowedTerms      []AllowedTerms `json:"allowed_terms"`
}

type Cards struct {
	AllowedBins            []string                `json:"allowed_bins"`
	InstallmentConfiguration InstallmentConfiguration `json:"installment_configuration"`
}

// Root structure representing the invoice response
type InvoiceResponse struct {
	ID                        string                      `json:"id"`
	ExternalID                string                      `json:"external_id"`
	UserID                    string                      `json:"user_id"`
	Status                    string                      `json:"status"`
	MerchantName              string                      `json:"merchant_name"`
	MerchantProfilePictureURL string                      `json:"merchant_profile_picture_url"`
	Amount                    int                         `json:"amount"`
	Description               string                      `json:"description"`
	ExpiryDate                time.Time                   `json:"expiry_date"`
	InvoiceURL                string                      `json:"invoice_url"`
	AvailableBanks            []Bank                      `json:"available_banks"`
	AvailableRetailOutlets    []RetailOutlet              `json:"available_retail_outlets"`
	AvailableEwallets         []Ewallet                   `json:"available_ewallets"`
	AvailableQRCodes          []QRCode                    `json:"available_qr_codes"`
	AvailablePaylaters        []Paylater                  `json:"available_paylaters"`
	ShouldExcludeCreditCard   bool                        `json:"should_exclude_credit_card"`
	ShouldSendEmail           bool                        `json:"should_send_email"`
	SuccessRedirectURL        string                      `json:"success_redirect_url"`
	FailureRedirectURL        string                      `json:"failure_redirect_url"`
	Created                   time.Time                   `json:"created"`
	Updated                   time.Time                   `json:"updated"`
	Currency                  string                      `json:"currency"`
}

// Bank represents a bank in the available banks list
type Bank struct {
	BankCode          string `json:"bank_code"`
	CollectionType   string `json:"collection_type"`
	TransferAmount   int    `json:"transfer_amount"`
	BankBranch       string `json:"bank_branch"`
	AccountHolderName string `json:"account_holder_name"`
	IdentityAmount   int    `json:"identity_amount"`
}

// RetailOutlet represents a retail outlet in the available retail outlets list
type RetailOutlet struct {
	RetailOutletName string `json:"retail_outlet_name"`
}

// Ewallet represents an e-wallet in the available ewallets list
type Ewallet struct {
	EwalletType string `json:"ewallet_type"`
}

// QRCode represents a QR code type in the available QR codes list
type QRCode struct {
	QRCodeType string `json:"qr_code_type"`
}

// Paylater represents a paylater type in the available paylaters list
type Paylater struct {
	PaylaterType string `json:"paylater_type"`
}

// CustomerNotificationPref represents customer notification preferences
type CustomerNotificationPref struct {
	InvoiceCreated    []string `json:"invoice_created"`
	InvoiceReminder   []string `json:"invoice_reminder"`
	InvoiceExpired    []string `json:"invoice_expired"`
	InvoicePaid       []string `json:"invoice_paid"`
}

type WebhookPayload struct {
	ID                     string    `json:"id"`
	ExternalID             string    `json:"external_id"`
	UserID                 string    `json:"user_id"`
	IsHigh                 bool      `json:"is_high"`
	PaymentMethod          string    `json:"payment_method"`
	Status                 string    `json:"status"`
	MerchantName           string    `json:"merchant_name"`
	Amount                 float64   `json:"amount"`
	PaidAmount             float64   `json:"paid_amount"`
	BankCode               string    `json:"bank_code"`
	PaidAt                 time.Time `json:"paid_at"`
	PayerEmail             string    `json:"payer_email"`
	Description           string    `json:"description"`
	AdjustedReceivedAmount float64   `json:"adjusted_received_amount"`
	FeesPaidAmount         float64   `json:"fees_paid_amount"`
	Updated                time.Time `json:"updated"`
	Created                time.Time `json:"created"`
	Currency               string    `json:"currency"`
	PaymentChannel         string    `json:"payment_channel"`
	PaymentDestination     string    `json:"payment_destination"`
}