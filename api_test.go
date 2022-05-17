package TossPaymentsApi

import (
	"fmt"
	"testing"
)

func TestClient_ApprovePayment(t *testing.T) {
	client := NewClient("test_sk_O6BYq7GWPVvxpOOydMa8NE5vbo1d")
	res, err := client.ApprovePayment("1", "1", 1)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(res)
		return
	}
}
