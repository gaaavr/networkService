package data

import (
	"io/ioutil"
	"math"
	"os"
)

type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

func GetResultsBilling() {
	defer wg.Done()
	var (
		billing BillingData
		sum     uint8
		counter float64
	)
	file, err := os.Open("..\\simulator\\billing.data")
	if err != nil {
		Result.Lock()
		Result.Billing = billing
		Result.Unlock()
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		Result.Lock()
		Result.Billing = billing
		Result.Unlock()
		return
	}
	for i := len(data) - 1; i >= 0; i-- {
		if string(data[i]) == "1" {
			sum += uint8(math.Pow(2, counter))
		}
		counter++
	}
	billing.CreateCustomer = sum&(1<<5) != 0
	billing.Purchase = sum&(1<<4) != 0
	billing.Payout = sum&(1<<3) != 0
	billing.Recurring = sum&(1<<2) != 0
	billing.FraudControl = sum&(1<<1) != 0
	billing.CheckoutPage = sum&1 != 0
	Result.Lock()
	Result.Billing = billing
	Result.Unlock()
	return
}
