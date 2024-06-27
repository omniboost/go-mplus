package mplus_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetReceipts(t *testing.T) {
	req := client.NewGetReceiptsRequest()
	req.RequestBody().Request.FromFinancialDate.Day = 13
	req.RequestBody().Request.FromFinancialDate.Mon = 6
	req.RequestBody().Request.FromFinancialDate.Year = 2024
	req.RequestBody().Request.ThroughFinancialDate.Day = 14
	req.RequestBody().Request.ThroughFinancialDate.Mon = 6
	req.RequestBody().Request.ThroughFinancialDate.Year = 2024
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
