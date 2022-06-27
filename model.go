package TossPaymentsApi

// Payment
// Object containing payment information
// Definition : https://docs.tosspayments.com/reference#payment-%EA%B0%9D%EC%B2%B4
// Version : 2022-06-08
type Payment struct {
	Version        string  `json:"version"`        // Version of the payment object
	PaymentKey     string  `json:"paymentKey"`     // Unique key of transaction
	OrderId        string  `json:"orderId"`        // Unique ID issued by the merchant of order
	OrderName      string  `json:"orderName"`      // Name of the order
	Type           string  `json:"type"`           // Payment type information [ Possible Values : https://docs.tosspayments.com/reference/enum-codes#%EA%B2%B0%EC%A0%9C-%EC%88%98%EB%8B%A8 ]
	Method         string  `json:"method"`         // Payment method [ Possible Values : https://docs.tosspayments.com/reference/enum-codes#%EA%B2%B0%EC%A0%9C-%EC%88%98%EB%8B%A8 ]
	MemberId       string  `json:"mid"`            // Merchant ID
	Currency       string  `json:"currency"`       // Currency of the payment, MUST be "KRW"
	TotalAmount    float64 `json:"totalAmount"`    // Total payment amount
	BalanceAmount  float64 `json:"balanceAmount"`  // Cancelable Amount (current balance)
	Status         string  `json:"status"`         // Payment status [ Possible Values : "READY", "IN_PROGRESS", "WAITING_FOR_DEPOSIT", "DONE", "CANCELED", "PARTIAL_CANCELLED", "ABORTED", "EXPIRED" ]
	RequestedAt    string  `json:"requestedAt"`    // Requested time, [ISO 8601 format]
	ApprovedAt     string  `json:"approvedAt"`     // Approved time, [ISO 8601 format]
	UseEscrow      bool    `json:"useEscrow"`      // Whether to use escrow service
	TransactionKey string  `json:"transactionKey"` // Unique key of transaction
	SuppliedAmount float64 `json:"suppliedAmount"` // Supplied amount
	Vat            float64 `json:"vat"`            // VAT amount
	CultureExpense bool    `json:"cultureExpense"` // Whether to use culture expense
	TaxFreeAmount  float64 `json:"taxFreeAmount"`  // Tax-free amount
	// Payment Cancellation History, NULLABLE
	Cancels *[]struct {
		CancelAmount     float64  `json:"cancelAmount"`     // Canceled amount
		CancelReason     string   `json:"cancelReason"`     // Canceled reason
		TaxFreeAmount    float64  `json:"taxFreeAmount"`    // Tax-free amount
		TaxAmount        *float64 `json:"taxAmount"`        // Tax amount, NULLABLE
		RefundableAmount float64  `json:"refundableAmount"` // Refundable amount
		CancelAt         string   `json:"cancelAt"`         // Canceled time, [ISO 8601 format]
	} `json:"cancels"`
	IsPartialCancelable bool `json:"isPartialCancelable"` // Whether partial cancellation is available
	// Card information, NULLABLE
	Card *struct {
		Company               string `json:"company"`               // Card company code
		Number                string `json:"number"`                // Card number, Partial masked
		InstallmentPlanMonths int    `json:"installmentPlanMonths"` // Installment month, 0 if lump sum payment
		IsInterestFree        bool   `json:"isInterestFree"`        // Whether interest free is available
		InterestPayer         string `json:"interestPayer"`         // Interest fee payer if interest free is available, [ Possible Values : "CUSTOMER", "MERCHANT", "CARD_COMPANY" ]
		ApproveNo             string `json:"approveNo"`             // Approval number of card company
		UseCardPoint          bool   `json:"useCardPoint"`          // Whether to use card point
		CardType              string `json:"cardType"`              // Card type, [ Possible Values : "신용", "체크", "기프트" ]
		OwnerType             string `json:"ownerType"`             // Card owner type, [ Possible Values : "개인", "법인" ]
		AcquireStatus         string `json:"acquireStatus"`         // Acquire status, [ Possible Values : "READY", "REQUESTED", "COMPLETED", "CANCEL_REQUESTED", "CANCELED" ]
		ReceiptURL            string `json:"receiptUrl"`            // Card sales slip URL
	} `json:"card"`
	VirtualAccount struct {
		AccountNumber    string `json:"accountNumber"`    // Virtual account number
		AccountType      string `json:"accountType"`      // Virtual account type, [ Possible Values : "일반", "고정" ]
		Bank             string `json:"bank"`             // Virtual account bank
		CustomerName     string `json:"customerName"`     // Owner name who issued virtual account
		DueDate          string `json:"dueDate"`          // Due date of virtual account
		Expired          bool   `json:"expired"`          // Whether virtual account is expired
		SettlementStatus string `json:"settlementStatus"` // Settlement status of virtual account, [ Possible Values : "INCOMPLETE", "COMPLETE" ]
		RefundStatus     string `json:"refundStatus"`     // Refund status of virtual account, [ Possible Values : "NONE", "FAILED", "PENDING", "PARTIAL_FAILED", "COMPLETED" ]
	} `json:"virtualAccount"`
	Secret      *string `json:"secret"` // Secret which is using for validate virtual account payment callback, NULLABLE
	MobilePhone *struct {
		Carrier             string `json:"carrier"`             // Deprecated: Deleted in Api Version 2020-06-08, SHOULD NOT BE USED
		CustomerMobilePhone string `json:"customerMobilePhone"` // Mobile phone number which used for mobile purchase
		SettlementStatus    string `json:"settlementStatus"`    // Settlement status of mobile phone payment, [ Possible Values : "INCOMPLETE", "COMPLETE" ]
		ReceiptUrl          string `json:"receiptUrl"`          // Mobile phone payment receipt URL
	} `json:"mobilePhone"` // Mobile phone payment information, NULLABLE
	GiftCertificate *struct {
		ApproveNo        string `json:"approveNo"`        // Approval number of gift certificate company
		SettlementStatus string `json:"settlementStatus"` // Settlement status of gift certificate payment, [ Possible Values : "INCOMPLETE", "COMPLETE" ]
	} `json:"giftCertificate"` // Gift certificate payment information, NULLABLE
	Transfer *struct {
		Bank             string `json:"bank"`             // Bank name of transfer payment
		SettlementStatus string `json:"settlementStatus"` // Settlement status of transfer payment, [ Possible Values : "INCOMPLETE", "COMPLETE" ]
	} `json:"transfer"` // Transfer payment information, NULLABLE
	EasyPay *struct {
		Amount         float64 `json:"amount"`         // Payment amount
		Provider       string  `json:"provider"`       // EasyPay provider information [ Possible Values : https://docs.tosspayments.com/reference/enum-codes#%EA%B0%84%ED%8E%B8%EA%B2%B0%EC%A0%9C-%EA%B2%B0%EC%A0%9C-%EC%88%98%EB%8B%A8 ]
		DiscountAmount float64 `json:"discountAmount"` // Discounted amount
	} `json:"easyPay"` // EasyPay information, NULLABLE
	Failure *struct {
		Code    string `json:"code"`    // Error code
		Message string `json:"message"` // Reason of error
	} `json:"failure"` // Error information, NULLABLE
	Country     string `json:"country"` // ISO 3166-1 alpha-2 country code
	CashReceipt *struct {
		Type          string  `json:"type"`          // Cash receipt type, [ Possible Values : "소득공제", "지출증빙" ]
		Amount        float64 `json:"amount"`        // Cash receipt amount
		TaxFreeAmount int64   `json:"taxFreeAmount"` // Tax-free
		IssueNumber   string  `json:"issueNumber"`   // Cash receipt issue number
		ReceiptUrl    string  `json:"receiptUrl"`    // Cash receipt receipt URL
	} `json:"cashReceipt"` // Cash receipt information, NULLABLE
	Discount *struct {
		Amount float64 `json:"amount"` // Discount amount
	} `json:"discount"` // Discount information, NULLABLE
}

type RefundReceiveAccount struct {
	Bank          string `json:"bank"`          // Bank code
	AccountNumber string `json:"accountNumber"` // Account number
	HolderName    string `json:"holderName"`    // Account holder name
}

type Error struct {
	Code    string `json:"code"`    // Error code
	Message string `json:"message"` // An error message which describe error
}
