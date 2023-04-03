package mplus

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/omniboost/go-mplus/omitempty"
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

type Receipts []Receipt

type Receipt struct {
	ReceiptID string `xml:"receiptId"`
	OrderIDs  struct {
		ID string `xml:"id"`
	} `xml:"orderIds"`
	TransactionString string `xml:"transactionString"`
	SyncMarker        string `xml:"syncMarker"`
	ReceiptNumber     struct {
		BranchNumber    string `xml:"branchNumber"`
		WorkplaceNumber string `xml:"workplaceNumber"`
		Year            string `xml:"year"`
		Number          string `xml:"number"`
	} `xml:"receiptNumber"`
	ReceiptBarcode string `xml:"receiptBarcode"`
	ReceiptType    string `xml:"receiptType"`
	EmployeeNumber string `xml:"employeeNumber"`
	EmployeeName   string `xml:"employeeName"`
	EntryTimestamp struct {
		Sec      string `xml:"sec"`
		Min      string `xml:"min"`
		Hour     string `xml:"hour"`
		Day      string `xml:"day"`
		Mon      string `xml:"mon"`
		Year     string `xml:"year"`
		Isdst    string `xml:"isdst"`
		Timezone string `xml:"timezone"`
	} `xml:"entryTimestamp"`
	FinancialDate struct {
		Day  string `xml:"day"`
		Mon  string `xml:"mon"`
		Year string `xml:"year"`
	} `xml:"financialDate"`
	FinancialBranchNumber string `xml:"financialBranchNumber"`
	WorkplaceNumber       string `xml:"workplaceNumber"`
	Reference             string `xml:"reference"`
	TotalInclAmount       string `xml:"totalInclAmount"`
	TotalExclAmount       string `xml:"totalExclAmount"`
	VatMethod             string `xml:"vatMethod"`
	VatGroupList          struct {
		VatGroup []struct {
			VatCode       string `xml:"vatCode"`
			VatPercentage string `xml:"vatPercentage"`
			ExclAmount    string `xml:"exclAmount"`
			VatAmount     string `xml:"vatAmount"`
		} `xml:"vatGroup"`
	} `xml:"vatGroupList"`
	ChangeCounter string `xml:"changeCounter"`
	VersionNumber string `xml:"versionNumber"`
	PaidAmount    string `xml:"paidAmount"`
	State         string `xml:"state"`
	LineList      struct {
		Line []struct {
			Chardata       string `xml:",chardata"`
			LineID         string `xml:"lineId"`
			EmployeeNumber string `xml:"employeeNumber"`
			ArticleNumber  string `xml:"articleNumber"`
			LineType       string `xml:"lineType"`
			Data           struct {
				Quantity           int     `xml:"quantity"`
				DecimalPlaces      string  `xml:"decimalPlaces"`
				Price              float64 `xml:"price"`
				PriceExcl          float64 `xml:"priceExcl"`
				OriginalPrice      float64 `xml:"originalPrice"`
				OriginalPriceExcl  float64 `xml:"originalPriceExcl"`
				PurchasePrice      float64 `xml:"purchasePrice"`
				TurnoverGroup      int     `xml:"turnoverGroup"`
				TurnoverGroupType  string  `xml:"turnoverGroupType"`
				VatCode            int     `xml:"vatCode"`
				VatPercentage      float64 `xml:"vatPercentage"`
				DiscountType       string  `xml:"discountType"`
				DiscountPercentage float64 `xml:"discountPercentage"`
				DiscountAmount     float64 `xml:"discountAmount"`
				TotalInclAmount    float64 `xml:"totalInclAmount"`
				TotalExclAmount    float64 `xml:"totalExclAmount"`
				PriceType          string  `xml:"priceType"`
			} `xml:"data"`
			CourseNumber    string `xml:"courseNumber"`
			PreparationList struct {
				Line []struct {
					Chardata       string `xml:",chardata"`
					LineID         string `xml:"lineId"`
					EmployeeNumber string `xml:"employeeNumber"`
					ArticleNumber  string `xml:"articleNumber"`
					Data           struct {
						Quantity           int     `xml:"quantity"`
						DecimalPlaces      int     `xml:"decimalPlaces"`
						Price              float64 `xml:"price"`
						PriceExcl          float64 `xml:"priceExcl"`
						OriginalPrice      float64 `xml:"originalPrice"`
						OriginalPriceExcl  float64 `xml:"originalPriceExcl"`
						PurchasePrice      float64 `xml:"purchasePrice"`
						TurnoverGroup      string  `xml:"turnoverGroup"`
						TurnoverGroupType  string  `xml:"turnoverGroupType"`
						VatCode            int     `xml:"vatCode"`
						VatPercentage      float64 `xml:"vatPercentage"`
						DiscountType       string  `xml:"discountType"`
						DiscountPercentage float64 `xml:"discountPercentage"`
						DiscountAmount     float64 `xml:"discountAmount"`
						TotalInclAmount    float64 `xml:"totalInclAmount"`
						TotalExclAmount    float64 `xml:"totalExclAmount"`
						PriceType          string  `xml:"priceType"`
					} `xml:"data"`
					LineType     string `xml:"lineType"`
					CourseNumber string `xml:"courseNumber"`
				} `xml:"line"`
			} `xml:"preparationList"`
		} `xml:"line"`
	} `xml:"lineList"`
	PaymentList struct {
		Payment struct {
			PaymentID     string `xml:"paymentId"`
			FinancialDate struct {
				Day  string `xml:"day"`
				Mon  string `xml:"mon"`
				Year string `xml:"year"`
			} `xml:"financialDate"`
			Method        string  `xml:"method"`
			Amount        float64 `xml:"amount"`
			AccountNumber string  `xml:"accountNumber"`
		} `xml:"payment"`
	} `xml:"paymentList"`
	VatChange string `xml:"vatChange"`
}
