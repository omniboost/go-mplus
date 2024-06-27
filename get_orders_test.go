package mplus_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetOrders(t *testing.T) {
	req := client.NewGetOrdersRequest()
	req.RequestBody().Request.FromFinancialDate.Day = 1
	req.RequestBody().Request.FromFinancialDate.Mon = 1
	req.RequestBody().Request.FromFinancialDate.Year = 2024
	req.RequestBody().Request.ThroughFinancialDate.Day = 1
	req.RequestBody().Request.ThroughFinancialDate.Mon = 7
	req.RequestBody().Request.ThroughFinancialDate.Year = 2024
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

func TestGetOrdersAll(t *testing.T) {
	req := client.NewGetOrdersRequest()
	req.RequestBody().Request.FromFinancialDate.Day = 1
	req.RequestBody().Request.FromFinancialDate.Mon = 1
	req.RequestBody().Request.FromFinancialDate.Year = 2024
	req.RequestBody().Request.ThroughFinancialDate.Day = 1
	req.RequestBody().Request.ThroughFinancialDate.Mon = 7
	req.RequestBody().Request.ThroughFinancialDate.Year = 2024
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
