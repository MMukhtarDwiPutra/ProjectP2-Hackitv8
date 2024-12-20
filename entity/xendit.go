package entity

import "time"

// type XenditPaymentResponse struct {
// 	ID             string                   `json:"id"`
// 	Currency       string                   `json:"currency"`
// 	Amount         int                      `json:"amount"`
// 	CustomerID     string                   `json:"customer_id"`
// 	BusinessID     string                   `json:"business_id"`
// 	Status         string                   `json:"status"`
// 	PaymentMethod  PaymentMethod            `json:"payment_method"`
// 	ChannelProperties ChannelProperties     `json:"channel_properties"`
// 	Actions        []Action                 `json:"actions"`
// 	Created        string                   `json:"created"`
// 	Updated        string                   `json:"updated"`
// 	Metadata       map[string]interface{}   `json:"metadata"`
// }

// type EWallet struct {
// 	ChannelCode       string               `json:"channel_code"`
// 	ChannelProperties EWalletChannelProperties `json:"channel_properties"`
// 	Account           EWalletAccount       `json:"account"`
// }

// type EWalletChannelProperties struct {
// 	SuccessReturnURL string               `json:"success_return_url"`
// }

// type EWalletAccount struct {
// 	AccountDetails interface{}            `json:"account_details"`
// 	Name           interface{}            `json:"name"`
// 	Balance        interface{}            `json:"balance"`
// 	PointBalance   interface{}            `json:"point_balance"`
// }

// type Action struct {
// 	Action   string                      `json:"action"`
// 	URLType  string                      `json:"url_type"`
// 	URL      string                      `json:"url"`
// 	Method   string                      `json:"method"`
// }

// type PaymentMethodsResponse struct {
// 	Data    []PaymentMethod `json:"data"`
// 	HasMore bool            `json:"has_more"`
// }

// type PaymentMethod struct {
// 	ID                  string              `json:"id"`
// 	Card                interface{}         `json:"card"`
// 	Type                string              `json:"type"`
// 	Status              string              `json:"status"`
// 	Actions             []interface{}       `json:"actions"`
// 	Country             string              `json:"country"`
// 	Created             string              `json:"created"`
// 	EWallet             interface{}         `json:"ewallet"`
// 	QRCode              interface{}         `json:"qr_code"`
// 	Updated             string              `json:"updated"`
// 	Metadata            interface{}         `json:"metadata"`
// 	CustomerID          string              `json:"customer_id"`
// 	Description         interface{}         `json:"description"`
// 	Reusability         string              `json:"reusability"`
// 	DirectDebit         *DirectDebit        `json:"direct_debit"`
// 	FailureCode         interface{}         `json:"failure_code"`
// 	ReferenceID         string              `json:"reference_id"`
// 	VirtualAccount      interface{}         `json:"virtual_account"`
// 	OverTheCounter      interface{}         `json:"over_the_counter"`
// 	BillingInformation  interface{}         `json:"billing_information"`
// 	DirectBankTransfer  interface{}         `json:"direct_bank_transfer"`
// 	BusinessID          string              `json:"business_id"`
// }

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

// type Metadata struct {
// 	BranchArea string `json:"branch_area"`
// 	BranchCity string `json:"branch_city"`
// 	BranchCode string `json:"branch_code"`
// 	StoreBranch string `json:"store_branch"`
// }

// type Actions struct {
// 	DesktopWebCheckoutURL  *string `json:"desktop_web_checkout_url"`
// 	MobileWebCheckoutURL   *string `json:"mobile_web_checkout_url"`
// 	MobileDeeplinkCheckoutURL string `json:"mobile_deeplink_checkout_url"`
// 	QrCheckoutString       string `json:"qr_checkout_string"`
// }

// type ChargeResponse struct {
// 	ID                  string            `json:"id"`
// 	BusinessID          string            `json:"business_id"`
// 	ReferenceID         string            `json:"reference_id"`
// 	Status              string            `json:"status"`
// 	Currency            string            `json:"currency"`
// 	ChargeAmount        int               `json:"charge_amount"`
// 	CaptureAmount       int               `json:"capture_amount"`
// 	RefundedAmount      *int              `json:"refunded_amount"`
// 	CheckoutMethod      string            `json:"checkout_method"`
// 	ChannelCode         string            `json:"channel_code"`
// 	ChannelProperties   ChannelProperties `json:"channel_properties"`
// 	Actions             Actions           `json:"actions"`
// 	IsRedirectRequired  bool              `json:"is_redirect_required"`
// 	CallbackURL         string            `json:"callback_url"`
// 	Created             time.Time         `json:"created"`
// 	Updated             time.Time         `json:"updated"`
// 	VoidStatus          *string           `json:"void_status"`
// 	VoidedAt            *time.Time        `json:"voided_at"`
// 	CaptureNow          bool              `json:"capture_now"`
// 	CustomerID          *string           `json:"customer_id"`
// 	PaymentMethodID     *string           `json:"payment_method_id"`
// 	FailureCode         *string           `json:"failure_code"`
// 	Basket              *string           `json:"basket"`
// 	Metadata            Metadata          `json:"metadata"`
// }

// type Address struct {
// 	City        string `json:"city"`
// 	Country     string `json:"country"`
// 	PostalCode  string `json:"postal_code"`
// 	State       string `json:"state"`
// 	StreetLine1 string `json:"street_line1"`
// 	StreetLine2 string `json:"street_line2"`
// }

// type Customer struct {
// 	GivenNames   string    `json:"given_names"`
// 	Surname      string    `json:"surname"`
// 	Email        string    `json:"email"`
// 	MobileNumber string    `json:"mobile_number"`
// 	Addresses    []Address `json:"addresses"`
// }

// type NotificationPreference struct {
// 	InvoiceCreated []string `json:"invoice_created"`
// 	InvoiceReminder []string `json:"invoice_reminder"`
// 	InvoicePaid     []string `json:"invoice_paid"`
// }

// type Item struct {
// 	Name     string `json:"name"`
// 	Quantity int    `json:"quantity"`
// 	Price    int    `json:"price"`
// 	Category string `json:"category"`
// 	URL      string `json:"url"`
// }

// type Fee struct {
// 	Type  string `json:"type"`
// 	Value int    `json:"value"`
// }

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

// type PaymentRequest struct {
// 	ExternalID                    string                 `json:"external_id"`
// 	Amount                        int                    `json:"amount"`
// 	Description                   string                 `json:"description"`
// 	InvoiceDuration               int                    `json:"invoice_duration"`
// 	Customer                      Customer               `json:"customer"`
// 	CustomerNotificationPreference NotificationPreference `json:"customer_notification_preference"`
// 	SuccessRedirectURL            string                 `json:"success_redirect_url"`
// 	FailureRedirectURL            string                 `json:"failure_redirect_url"`
// 	Currency                      string                 `json:"currency"`
// 	Items                         []Item                 `json:"items"`
// 	Fees                          []Fee                  `json:"fees"`
// 	PaymentMethods                []string               `json:"payment_methods"`
// 	ChannelProperties             ChannelProperties      `json:"channel_properties"`
// 	Metadata                      Metadata               `json:"metadata"`
// }

// type CustomerNotificationPreference struct {
// 	InvoiceCreated  []string `json:"invoice_created"`
// 	InvoiceReminder []string `json:"invoice_reminder"`
// 	InvoicePaid     []string `json:"invoice_paid"`
// }

// type AvailableBank struct {
// 	BankCode          string `json:"bank_code"`
// 	CollectionType    string `json:"collection_type"`
// 	TransferAmount    int    `json:"transfer_amount"`
// 	BankBranch        string `json:"bank_branch"`
// 	AccountHolderName string `json:"account_holder_name"`
// 	IdentityAmount    int    `json:"identity_amount"`
// }

// type AvailableRetailOutlet struct {
// 	RetailOutletName string `json:"retail_outlet_name"`
// }

// type AvailableEwallet struct {
// 	EwalletType string `json:"ewallet_type"`
// }

// type AvailableQRCode struct {
// 	QRCodeType string `json:"qr_code_type"`
// }

// type AvailableDirectDebit struct {
// 	DirectDebitType string `json:"direct_debit_type"`
// }

// type AvailablePaylater struct {
// 	PaylaterType string `json:"paylater_type"`
// }

// type AllowedTerm struct {
// 	Issuer string `json:"issuer"`
// 	Terms  []int  `json:"terms"`
// }

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