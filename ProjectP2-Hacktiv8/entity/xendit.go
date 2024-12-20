package entity

import "time"

type XenditPaymentResponse struct {
	ID             string                   `json:"id"`
	Currency       string                   `json:"currency"`
	Amount         int                      `json:"amount"`
	CustomerID     string                   `json:"customer_id"`
	BusinessID     string                   `json:"business_id"`
	Status         string                   `json:"status"`
	PaymentMethod  PaymentMethod            `json:"payment_method"`
	ChannelProperties ChannelProperties     `json:"channel_properties"`
	Actions        []Action                 `json:"actions"`
	Created        string                   `json:"created"`
	Updated        string                   `json:"updated"`
	Metadata       map[string]interface{}   `json:"metadata"`
}

// type PaymentMethod struct {
// 	ID             string                   `json:"id"`
// 	Type           string                   `json:"type"`
// 	Reusability    string                   `json:"reusability"`
// 	Status         string                   `json:"status"`
// 	EWallet        EWallet                  `json:"ewallet"`
// 	DirectDebit    interface{}              `json:"direct_debit"`
// }

type EWallet struct {
	ChannelCode       string               `json:"channel_code"`
	ChannelProperties EWalletChannelProperties `json:"channel_properties"`
	Account           EWalletAccount       `json:"account"`
}

type EWalletChannelProperties struct {
	SuccessReturnURL string               `json:"success_return_url"`
}

type EWalletAccount struct {
	AccountDetails interface{}            `json:"account_details"`
	Name           interface{}            `json:"name"`
	Balance        interface{}            `json:"balance"`
	PointBalance   interface{}            `json:"point_balance"`
}

// type ChannelProperties struct {
// 	RedeemPoints string                  `json:"redeem_points"`
// }

type Action struct {
	Action   string                      `json:"action"`
	URLType  string                      `json:"url_type"`
	URL      string                      `json:"url"`
	Method   string                      `json:"method"`
}

type PaymentMethodsResponse struct {
	Data    []PaymentMethod `json:"data"`
	HasMore bool            `json:"has_more"`
}

type PaymentMethod struct {
	ID                  string              `json:"id"`
	Card                interface{}         `json:"card"`
	Type                string              `json:"type"`
	Status              string              `json:"status"`
	Actions             []interface{}       `json:"actions"`
	Country             string              `json:"country"`
	Created             string              `json:"created"`
	EWallet             interface{}         `json:"ewallet"`
	QRCode              interface{}         `json:"qr_code"`
	Updated             string              `json:"updated"`
	Metadata            interface{}         `json:"metadata"`
	CustomerID          string              `json:"customer_id"`
	Description         interface{}         `json:"description"`
	Reusability         string              `json:"reusability"`
	DirectDebit         *DirectDebit        `json:"direct_debit"`
	FailureCode         interface{}         `json:"failure_code"`
	ReferenceID         string              `json:"reference_id"`
	VirtualAccount      interface{}         `json:"virtual_account"`
	OverTheCounter      interface{}         `json:"over_the_counter"`
	BillingInformation  interface{}         `json:"billing_information"`
	DirectBankTransfer  interface{}         `json:"direct_bank_transfer"`
	BusinessID          string              `json:"business_id"`
}

type DirectDebit struct {
	Type            string            `json:"type"`
	DebitCard       *DebitCard        `json:"debit_card"`
	BankAccount     interface{}       `json:"bank_account"`
	ChannelCode     string            `json:"channel_code"`
	ChannelProperties ChannelProperties `json:"channel_properties"`
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
}

type PaymentRequest struct {
	ReferenceID     string            `json:"reference_id"`
	Currency        string            `json:"currency"`
	Amount          int               `json:"amount"`
	CheckoutMethod  string            `json:"checkout_method"`
	ChannelCode     string            `json:"channel_code"`
	ChannelProperties ChannelProperties `json:"channel_properties"`
	Metadata        Metadata          `json:"metadata"`
}

type Metadata struct {
	BranchArea string `json:"branch_area"`
	BranchCity string `json:"branch_city"`
	BranchCode string `json:"branch_code"`
}

type Actions struct {
	DesktopWebCheckoutURL  *string `json:"desktop_web_checkout_url"`
	MobileWebCheckoutURL   *string `json:"mobile_web_checkout_url"`
	MobileDeeplinkCheckoutURL string `json:"mobile_deeplink_checkout_url"`
	QrCheckoutString       string `json:"qr_checkout_string"`
}

type ChargeResponse struct {
	ID                  string            `json:"id"`
	BusinessID          string            `json:"business_id"`
	ReferenceID         string            `json:"reference_id"`
	Status              string            `json:"status"`
	Currency            string            `json:"currency"`
	ChargeAmount        int               `json:"charge_amount"`
	CaptureAmount       int               `json:"capture_amount"`
	RefundedAmount      *int              `json:"refunded_amount"`
	CheckoutMethod      string            `json:"checkout_method"`
	ChannelCode         string            `json:"channel_code"`
	ChannelProperties   ChannelProperties `json:"channel_properties"`
	Actions             Actions           `json:"actions"`
	IsRedirectRequired  bool              `json:"is_redirect_required"`
	CallbackURL         string            `json:"callback_url"`
	Created             time.Time         `json:"created"`
	Updated             time.Time         `json:"updated"`
	VoidStatus          *string           `json:"void_status"`
	VoidedAt            *time.Time        `json:"voided_at"`
	CaptureNow          bool              `json:"capture_now"`
	CustomerID          *string           `json:"customer_id"`
	PaymentMethodID     *string           `json:"payment_method_id"`
	FailureCode         *string           `json:"failure_code"`
	Basket              *string           `json:"basket"`
	Metadata            Metadata          `json:"metadata"`
}

func main() {
	// Example usage of the ChargeResponseRequest struct
}
