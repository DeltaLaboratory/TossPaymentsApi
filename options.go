package TossPaymentsApi

type ApprovePaymentOption struct {
	OrderId string `json:"orderId"`
	Amount  int64  `json:"amount"`
}

type CancelPaymentOption struct {
	CancelReason         string                `json:"cancelReason"`
	CancelAmount         *float64              `json:"cancelAmount"`
	RefundReceiveAccount *RefundReceiveAccount `json:"refundReceiveAccount"`
	TaxFreeAmount        *float64              `json:"taxFreeAmount"`
	RefundableAmount     *float64              `json:"refundableAmount"`
}

type KeyinPaymentOption struct {
	Amount                 float64  `json:"amount"`
	OrderId                string   `json:"orderId"`
	CardNumber             string   `json:"cardNumber"`
	CardExpirationYear     string   `json:"cardExpirationYear"`
	CardExpirationMonth    string   `json:"cardExpirationMonth"`
	OrderName              *string  `json:"orderName"`
	CardPassword           *string  `json:"cardPassword"`
	CustomerIdentityNumber *string  `json:"customerIdentityNumber"`
	CardInstallmentPlan    *int     `json:"cardInstallmentPlan"`
	TaxFreeAmount          *float64 `json:"taxFreeAmount"`
	CustomerEmail          *string  `json:"customerEmail"`
	VBV                    *struct {
		CAVV string `json:"cavv"` // Cardholder Authentication Verification Value
		XID  string `json:"xid"`  // Transaction ID
		ECI  string `json:"eci"`  // Electronic Commerce Indicator
	} `json:"vbv"`
}
