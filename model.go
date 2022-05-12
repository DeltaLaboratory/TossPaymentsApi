package TossPaymentsApi

// Payment
// Object containing payment information
// Definition : https://docs.tosspayments.com/reference#payment-%EA%B0%9D%EC%B2%B4
type Payment struct {
	Version    string `json:"version"`
	PaymentKey string `json:"paymentKey"`
	OrderId    string `json:"orderId"`
	Type       string `json:"type"` // Possible Values : "NORMAL", "BILLING", "CONNECTPAY"
}
