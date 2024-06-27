package mplus_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetJournals(t *testing.T) {
	req := client.NewGetJournalsRequest()
	req.RequestBody().Request.FromFinancialDate.Day = 13
	req.RequestBody().Request.FromFinancialDate.Mon = 6
	req.RequestBody().Request.FromFinancialDate.Year = 2024
	req.RequestBody().Request.ThroughFinancialDate.Day = 13
	req.RequestBody().Request.ThroughFinancialDate.Mon = 6
	req.RequestBody().Request.ThroughFinancialDate.Year = 2024

	req.RequestBody().Request.BranchNumbers = []int{1}
	req.RequestBody().Request.JournalFilterList = struct {
		JournalFilter []string "xml:\"urn:journalFilter\""
	}{
		JournalFilter: []string{
			// "JOURNAL-FILTER-RECEIPT",
			// "JOURNAL-FILTER-INVOICE",
			"JOURNAL-FILTER-ORDER",
		},
	}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
