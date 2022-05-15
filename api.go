package TossPaymentsApi

import (
	"encoding/base64"
	"errors"
	"github.com/imroc/req/v3"
)

func NewClient(apiKey string) *Client {
	client := &Client{}
	client.httpClient = req.C().
		SetBaseURL("https://api.tosspayments.com").
		SetCommonHeader("Authorization", "Basic "+base64.URLEncoding.EncodeToString([]byte(apiKey+":"))).
		SetCommonHeader("Content-Type", "application/json")
	return client
}

type Client struct {
	httpClient *req.Client
}

func (c *Client) ApprovePayment(paymentKey string, orderId string, amount int64) (*Payment, error) {
	payment := &Payment{}
	errorMessage := &Error{}
	resp, err := c.httpClient.R().
		SetBody(approvePaymentJsonRequest{
			OrderId: orderId,
			Amount:  amount,
		}).
		SetResult(payment).
		SetError(errorMessage).
		Post("/v1/payments/" + paymentKey)
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() != true {
		return nil, errors.New(errorMessage.Code + ":" + errorMessage.Message)
	}
	return payment, nil
}
