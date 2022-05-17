package TossPaymentsApi

type ApprovePaymentOption struct {
	OrderId string `json:"orderId"`
	Amount  int64  `json:"amount"`
}

type CancelPaymentOption struct {
	CancelReason         string                `json:"cancelReason"`
	CancelAmount         *int64                `json:"cancelAmount"`
	RefundReceiveAccount *RefundReceiveAccount `json:"refundReceiveAccount"`
	TaxFreeAmount        *int64                `json:"taxFreeAmount"`
	RefundableAmount     *int64                `json:"refundableAmount"`
}
