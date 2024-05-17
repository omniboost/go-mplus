package mplus_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetInvoices(t *testing.T) {
	req := client.NewGetInvoicesRequest()
	req.RequestBody().Request.FromFinancialDate.Day = 1
	req.RequestBody().Request.FromFinancialDate.Mon = 3
	req.RequestBody().Request.FromFinancialDate.Year = 2024
	req.RequestBody().Request.ThroughFinancialDate.Day = 1
	req.RequestBody().Request.ThroughFinancialDate.Mon = 4
	req.RequestBody().Request.ThroughFinancialDate.Year = 2024
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
