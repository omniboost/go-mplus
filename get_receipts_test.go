package mplus_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetReceipts(t *testing.T) {
	req := client.NewGetReceiptsRequest()
	req.RequestBody().Request.FromFinancialDate.Day = 1
	req.RequestBody().Request.FromFinancialDate.Mon = 1
	req.RequestBody().Request.FromFinancialDate.Year = 2022
	req.RequestBody().Request.ThroughFinancialDate.Day = 1
	req.RequestBody().Request.ThroughFinancialDate.Mon = 1
	req.RequestBody().Request.ThroughFinancialDate.Year = 2023
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
