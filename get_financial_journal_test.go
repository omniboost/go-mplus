package mplus_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetFinancialJournal(t *testing.T) {
	req := client.NewGetFinancialJournalRequest()
	req.RequestBody().Request.FromFinancialDate.Day = 4
	req.RequestBody().Request.FromFinancialDate.Mon = 5
	req.RequestBody().Request.FromFinancialDate.Year = 2021
	req.RequestBody().Request.ThroughFinancialDate.Day = 5
	req.RequestBody().Request.ThroughFinancialDate.Mon = 5
	req.RequestBody().Request.ThroughFinancialDate.Year = 2021
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
