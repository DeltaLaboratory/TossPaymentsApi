package TossPaymentsApi

// Payment
// Object containing payment information
// Definition : https://docs.tosspayments.com/reference#payment-%EA%B0%9D%EC%B2%B4
type Payment struct {
	Version        string `json:"version"`        // Version of the payment object
	PaymentKey     string `json:"paymentKey"`     // Unique key of transaction
	OrderId        string `json:"orderId"`        // Unique ID issued by the merchant of order
	OrderName      string `json:"orderName"`      // Name of the order
	Type           string `json:"type"`           // Payment type information [ Possible Values : "NORMAL", "BILLING", "CONNECTPAY" ]
	Method         string `json:"method"`         // Payment method [ Possible Values : "카드", "가상계좌", "휴대폰", "계좌이체", "상품권(문화상품권, 도서문화상품권, 게임문화상품권)" ]
	MemberId       string `json:"mid"`            // Merchant ID
	Currency       string `json:"currency"`       // Currency of the payment, MUST be "KRW"
	TotalAmount    int64  `json:"totalAmount"`    // Total payment amount
	BalanceAmount  int64  `json:"balanceAmount"`  // Cancelable Amount (current balance)
	Status         string `json:"status"`         // Payment status [ Possible Values : "READY", "IN_PROGRESS", "WAITING_FOR_DEPOSIT", "DONE", "CANCELED", "PARTIAL_CANCELLED", "ABORTED", "EXPIRED" ]
	RequestedAt    string `json:"requestedAt"`    // Requested time, [ISO 8601 format]
	ApprovedAt     string `json:"approvedAt"`     // Approved time, [ISO 8601 format]
	UseEscrow      bool   `json:"useEscrow"`      // Whether to use escrow service
	TransactionKey string `json:"transactionKey"` // Unique key of transaction
	SuppliedAmount int64  `json:"suppliedAmount"` // Supplied amount
	Vat            int64  `json:"vat"`            // VAT amount
	CultureExpense bool   `json:"cultureExpense"` // Whether to use culture expense
	TaxFreeAmount  int64  `json:"taxFreeAmount"`  // Tax-free amount
	// Payment Cancellation History, NULLABLE
	Cancels *[]struct {
		CancelAmount     int64  `json:"cancelAmount"`     // Canceled amount
		CancelReason     string `json:"cancelReason"`     // Canceled reason
		TaxFreeAmount    int64  `json:"taxFreeAmount"`    // Tax-free amount
		TaxAmount        *int64 `json:"taxAmount"`        // Tax amount, NULLABLE
		RefundableAmount int64  `json:"refundableAmount"` // Refundable amount
		CancelAt         string `json:"cancelAt"`         // Canceled time, [ISO 8601 format]
	} `json:"cancels"`
	// Card
	Card *struct {
		Company               string `json:"company"`               // Card company code
		Number                string `json:"number"`                // Card number, Partial masked
		InstallmentPlanMonths int    `json:"installmentPlanMonths"` // Installment month, 0 if lump sum payment
		IsInterestFree        bool   `json:"isInterestFree"`        // Whether interest free is available
		InterestPayer         string `json:"interestPayer"`         // Interest fee payer if interest free is available, [ Possible Values : "CUSTOMER", "MERCHANT", "CARD_COMPANY" ]
		ApproveNo             string `json:"approveNo"`             // Approval number of card company
		UseCardPoint          bool   `json:"useCardPoint"`
		CardType              string `json:"cardType"`
		OwnerType             string `json:"ownerType"`
		AcquireStatus         string `json:"acquireStatus"`
		ReceiptURL            string `json:"receiptUrl"`
	} `json:"card"`
}
