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
		SetCommonHeader("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(apiKey+":"))).
		SetCommonHeader("Content-Type", "application/json")
	return client
}

type Client struct {
	httpClient *req.Client
}

func (c *Client) ApprovePayment(paymentKey string, option *ApprovePaymentOption) (*Payment, error) {
	payment := &Payment{}
	errorMessage := &Error{}
	resp, err := c.httpClient.R().
		SetBody(*option).
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

func (c *Client) InquiryPaymentByPaymentKey(paymentKey string) (*Payment, error) {
	payment := &Payment{}
	errorMessage := &Error{}
	resp, err := c.httpClient.R().
		SetResult(payment).
		SetError(errorMessage).
		Get("/v1/payments/" + paymentKey)
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() != true {
		return nil, errors.New(errorMessage.Code + ":" + errorMessage.Message)
	}
	return payment, nil
}

func (c *Client) InquiryPaymentByOrderId(orderId string) (*Payment, error) {
	payment := &Payment{}
	errorMessage := &Error{}
	resp, err := c.httpClient.R().
		SetResult(payment).
		SetError(errorMessage).
		Get("/v1/payments/orders/" + orderId)
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() != true {
		return nil, errors.New(errorMessage.Code + ":" + errorMessage.Message)
	}
	return payment, nil
}

func (c *Client) CancelPayment(paymentKey string, option *CancelPaymentOption) (*Payment, error) {
	payment := &Payment{}
	errorMessage := &Error{}
	resp, err := c.httpClient.R().
		SetBody(*option).
		SetResult(payment).
		SetError(errorMessage).
		Post("/v1/payments/" + paymentKey + "/cancel")
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() != true {
		return nil, errors.New(errorMessage.Code + ":" + errorMessage.Message)
	}
	return payment, nil
}

func (c *Client) KeyinPaymentRequest(option *KeyinPaymentOption) (*Payment, error) {
	payment := &Payment{}
	errorMessage := &Error{}
	resp, err := c.httpClient.R().
		SetBody(*option).
		SetResult(payment).
		SetError(errorMessage).
		Post("/v1/payments/key-in")
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() != true {
		return nil, errors.New(errorMessage.Code + ":" + errorMessage.Message)
	}
	return payment, nil
}
