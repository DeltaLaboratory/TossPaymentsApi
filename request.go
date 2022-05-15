package TossPaymentsApi

type approvePaymentJsonRequest struct {
	OrderId string `json:"orderId"`
	Amount  int64  `json:"amount"`
}
