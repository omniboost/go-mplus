package mplus

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/omniboost/go-netsuite-soap/omitempty"
)

type Customer struct {
}

func (c Customer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c Customer) IsEmpty() bool {
	return zero.IsZero(c)
}

type Journals []Journal

type Journal struct {
	FinancialDate struct {
		Day  int `xml:"day"`
		Mon  int `xml:"mon"`
		Year int `xml:"year"`
	} `xml:"financialDate"`
	JournalFilters []string `xml:"journalFilterList>journalFilter"`
	BranchNumber   int      `xml:"branchNumber"`
}

type PaymentMethods []PaymentMethod

type PaymentMethod struct {
	Method                    string `xml:"method"`
	Description               string `xml:"description"`
	ButtonText                string `xml:"buttonText"`
	AllowNegativeAmount       bool   `xml:"allowNegativeAmount"`
	Active                    bool   `xml:"active"`
	Type                      string `xml:"type"`
	HasExternalPaymentWebhook bool   `xml:"hasExternalPaymentWebhook"`
}

type Timestamp struct {
	Sec      string `xml:"sec"`
	Min      string `xml:"min"`
	Hour     string `xml:"hour"`
	Day      string `xml:"day"`
	Mon      string `xml:"mon"`
	Year     string `xml:"year"`
	Isdst    string `xml:"isdst"`
	Timezone string `xml:"timezone"`
}
