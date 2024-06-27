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
	FinancialDate     Date `xml:"financialDate"`
	JournalFilterList struct {
		JournalFilter []string `xml:"journalFilter"`
	} `xml:"journalFilterList"`
	BranchNumber string `xml:"branchNumber"`
	ExtBranchID  string `xml:"extBranchId"`
	VatGroupList struct {
		VatGroup []struct {
			VatCode       int     `xml:"vatCode"`
			VatPercentage float64 `xml:"vatPercentage"`
			ExclAmount    float64 `xml:"exclAmount"`
			VatAmount     float64 `xml:"vatAmount"`
		} `xml:"vatGroup"`
	} `xml:"vatGroupList"`
	TurnoverGroupList struct {
		TurnoverGroup []struct {
			TurnoverGroupType string  `xml:"turnoverGroupType"`
			TurnoverGroup     string  `xml:"turnoverGroup"`
			TurnoverGroupName string  `xml:"turnoverGroupName"`
			InclAmount        float64 `xml:"inclAmount"`
			ExclAmount        float64 `xml:"exclAmount"`
			AccountNumber     string  `xml:"accountNumber"`
		} `xml:"turnoverGroup"`
	} `xml:"turnoverGroupList"`
	PaymentList struct {
		Payment []struct {
			Method            string  `xml:"method"`
			Description       string  `xml:"description"`
			Amount            float64 `xml:"amount"`
			PaymentMethodType string  `xml:"paymentMethodType"`
			AccountNumber     string  `xml:"accountNumber"`
		} `xml:"payment"`
	} `xml:"paymentList"`
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

type Invoices []Invoice

type Invoice struct {
	InvoiceID         string `xml:"invoiceId"`
	TransactionString string `xml:"transactionString"`
	SyncMarker        string `xml:"syncMarker"`
	InvoiceNumber     struct {
		Year   string `xml:"year"`
		Number string `xml:"number"`
	} `xml:"invoiceNumber"`
	InvoiceBarcode        string    `xml:"invoiceBarcode"`
	InvoiceType           string    `xml:"invoiceType"`
	EmployeeNumber        string    `xml:"employeeNumber"`
	EmployeeName          string    `xml:"employeeName"`
	EntryTimestamp        Timestamp `xml:"entryTimestamp"`
	RelationNumber        string    `xml:"relationNumber"`
	RelationName          string    `xml:"relationName"`
	RelationCategoryID    string    `xml:"relationCategoryId"`
	FinancialDate         Date      `xml:"financialDate"`
	FinancialBranchNumber string    `xml:"financialBranchNumber"`
	WorkplaceNumber       string    `xml:"workplaceNumber"`
	EntryBranchNumber     string    `xml:"entryBranchNumber"`
	Reference             string    `xml:"reference"`
	DueDate               Date      `xml:"dueDate"`
	TotalInclAmount       float64   `xml:"totalInclAmount"`
	TotalExclAmount       float64   `xml:"totalExclAmount"`
	VatMethod             string    `xml:"vatMethod"`
	VatGroupList          struct {
		VatGroup []struct {
			VatCode       string `xml:"vatCode"`
			VatPercentage string `xml:"vatPercentage"`
			ExclAmount    string `xml:"exclAmount"`
			VatAmount     string `xml:"vatAmount"`
		} `xml:"vatGroup"`
	} `xml:"vatGroupList"`
	ChangeCounter      string    `xml:"changeCounter"`
	VersionNumber      string    `xml:"versionNumber"`
	PaidAmount         string    `xml:"paidAmount"`
	State              string    `xml:"state"`
	Finalized          string    `xml:"finalized"`
	FinalizedTimestamp Timestamp `xml:"finalizedTimestamp"`
	LineList           struct {
		Line []struct {
			LineID         string `xml:"lineId"`
			EmployeeNumber string `xml:"employeeNumber"`
			ArticleNumber  int    `xml:"articleNumber"`
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
		} `xml:"line"`
	} `xml:"lineList"`
	PaymentList struct {
		Payment []struct {
			PaymentID     string  `xml:"paymentId"`
			FinancialDate Date    `xml:"financialDate"`
			Method        string  `xml:"method"`
			Amount        float64 `xml:"amount"`
			AccountNumber string  `xml:"accountNumber"`
		} `xml:"payment"`
	} `xml:"paymentList"`
	VatChange      string `xml:"vatChange"`
	VatCountryCode string `xml:"vatCountryCode"`
	VatCountryISO3 string `xml:"vatCountryIso3"`
}

type Receipts []Receipt

type Receipt struct {
	ReceiptID string `xml:"receiptId"`
	OrderIDs  struct {
		ID string `xml:"id"`
	} `xml:"orderIds"`
	TransactionString string `xml:"transactionString"`
	SyncMarker        int    `xml:"syncMarker"`
	ReceiptNumber     struct {
		BranchNumber    string `xml:"branchNumber"`
		WorkplaceNumber string `xml:"workplaceNumber"`
		Year            string `xml:"year"`
		Number          string `xml:"number"`
	} `xml:"receiptNumber"`
	ReceiptBarcode        string    `xml:"receiptBarcode"`
	ReceiptType           string    `xml:"receiptType"`
	EmployeeNumber        string    `xml:"employeeNumber"`
	EmployeeName          string    `xml:"employeeName"`
	EntryTimestamp        Timestamp `xml:"entryTimestamp"`
	FinancialDate         Date      `xml:"financialDate"`
	FinancialBranchNumber string    `xml:"financialBranchNumber"`
	WorkplaceNumber       string    `xml:"workplaceNumber"`
	Reference             string    `xml:"reference"`
	TotalInclAmount       float64   `xml:"totalInclAmount"`
	TotalExclAmount       float64   `xml:"totalExclAmount"`
	VatMethod             string    `xml:"vatMethod"`
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
			LineID         string `xml:"lineId"`
			EmployeeNumber string `xml:"employeeNumber"`
			ArticleNumber  int    `xml:"articleNumber"`
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
					LineID         string `xml:"lineId"`
					EmployeeNumber string `xml:"employeeNumber"`
					ArticleNumber  int    `xml:"articleNumber"`
					Data           struct {
						Quantity           int     `xml:"quantity"`
						DecimalPlaces      int     `xml:"decimalPlaces"`
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
					LineType     string `xml:"lineType"`
					CourseNumber string `xml:"courseNumber"`
				} `xml:"line"`
			} `xml:"preparationList"`
		} `xml:"line"`
	} `xml:"lineList"`
	PaymentList struct {
		Payment []struct {
			PaymentID     string  `xml:"paymentId"`
			FinancialDate Date    `xml:"financialDate"`
			Method        string  `xml:"method"`
			Amount        float64 `xml:"amount"`
			AccountNumber string  `xml:"accountNumber"`
		} `xml:"payment"`
	} `xml:"paymentList"`
	VatChange string `xml:"vatChange"`
}

type TurnoverGroups []TurnoverGroup

type TurnoverGroup struct {
	TurnoverGroupType       string `xml:"turnoverGroupType"`
	TurnoverGroup           int    `xml:"turnoverGroup"`
	TurnoverGroupName       string `xml:"turnoverGroupName"`
	AllowPointsDistribution string `xml:"allowPointsDistribution"`
	AllowPointsPayment      string `xml:"allowPointsPayment"`
	AllowDiscount           string `xml:"allowDiscount"`
	BranchAccountNumberList struct {
		BranchAccountNumber []struct {
			BranchNumber  string `xml:"branchNumber"`
			AccountNumber string `xml:"accountNumber"`
		} `xml:"branchAccountNumber"`
	} `xml:"branchAccountNumberList"`
}

type Products []Product

type Product struct {
	ProductNumber int    `xml:"productNumber"`
	SyncMarker    int    `xml:"syncMarker"`
	Description   string `xml:"description"`
	ExtraText     string `xml:"extraText"`
	ArticleList   struct {
		Article []struct {
			ArticleNumber         int       `xml:"articleNumber"`
			ExternalArticleID     string    `xml:"extArticleId"`
			PluNumber             string    `xml:"pluNumber"`
			SyncMarker            int       `xml:"syncMarker"`
			ChangeTimestamp       Timestamp `xml:"changeTimestamp"`
			CreateTimestamp       Timestamp `xml:"createTimestamp"`
			Active                string    `xml:"active"`
			Description           string    `xml:"description"`
			Colour                string    `xml:"colour"`
			Size                  string    `xml:"size"`
			InvoiceText           string    `xml:"invoiceText"`
			ReceiptText           string    `xml:"receiptText"`
			DisplayText           string    `xml:"displayText"`
			Barcode               string    `xml:"barcode"`
			ExtraText             string    `xml:"extraText"`
			TurnoverGroup         int       `xml:"turnoverGroup"`
			VatCode               string    `xml:"vatCode"`
			VatPercentage         float64   `xml:"vatPercentage"`
			AveragePurchasePrice  float64   `xml:"averagePurchasePrice"`
			PurchasePrice         float64   `xml:"purchasePrice"`
			PriceIncl             float64   `xml:"priceIncl"`
			PriceExcl             float64   `xml:"priceExcl"`
			SupplierArticleNumber string    `xml:"supplierArticleNumber"`
			Webshop               string    `xml:"webshop"`
			CategoryId            string    `xml:"categoryId"`
			CategoryIds           struct {
				Category string `xml:"category"`
			} `xml:"categoryIds"`
			StockArticle string `xml:"stockArticle"`
			Course       struct {
				Number              string `xml:"number"`
				Name                string `xml:"name"`
				Abbreviation        string `xml:"abbreviation"`
				SequenceNumber      string `xml:"sequenceNumber"`
				Type                string `xml:"type"`
				DoSetTable          string `xml:"doSetTable"`
				CanBeServedWithWine string `xml:"canBeServedWithWine"`
			} `xml:"course"`
			SpecialPrice                   string `xml:"specialPrice"`
			ShowPreparationMethodsDirectly string `xml:"showPreparationMethodsDirectly"`
			ShowPreparationMethodsPerGroup string `xml:"showPreparationMethodsPerGroup"`
			ImageList                      struct {
				Image struct {
					ImageId          string    `xml:"imageId"`
					ImageName        string    `xml:"imageName"`
					ImageHash        string    `xml:"imageHash"`
					ImagePath        string    `xml:"imagePath"`
					ImageUrl         string    `xml:"imageUrl"`
					ImageData        string    `xml:"imageData"`
					ImageDataResult  string    `xml:"imageDataResult"`
					ThumbHash        string    `xml:"thumbHash"`
					ThumbPath        string    `xml:"thumbPath"`
					ThumbUrl         string    `xml:"thumbUrl"`
					ThumbData        string    `xml:"thumbData"`
					ThumbDataResult  string    `xml:"thumbDataResult"`
					ImageWidth       string    `xml:"imageWidth"`
					ImageHeight      string    `xml:"imageHeight"`
					ThumbWidth       string    `xml:"thumbWidth"`
					ThumbHeight      string    `xml:"thumbHeight"`
					CreatedTimestamp Timestamp `xml:"createdTimestamp"`
					ChangedTimestamp Timestamp `xml:"changedTimestamp"`
				} `xml:"image"`
			} `xml:"imageList"`
			CustomFieldList struct {
				CustomField []struct {
					FieldName     string `xml:"fieldName"`
					DataType      string `xml:"dataType"`
					StrValue      string `xml:"strValue"`
					IntValue      string `xml:"intValue"`
					DecimalPlaces string `xml:"decimalPlaces"`
				} `xml:"customField"`
			} `xml:"customFieldList"`
			PreparationMethodList struct {
				PreparationMethod []struct {
					ArticleNumber              int     `xml:"articleNumber"`
					Description                string  `xml:"description"`
					PriceIncl                  float64 `xml:"priceIncl"`
					VatCode                    string  `xml:"vatCode"`
					VatPercentage              float64 `xml:"vatPercentage"`
					SyncMarker                 int     `xml:"syncMarker"`
					CloseAfterSelection        string  `xml:"closeAfterSelection"`
					GroupNumber                string  `xml:"groupNumber"`
					GroupName                  string  `xml:"groupName"`
					MaxOneSelection            string  `xml:"maxOneSelection"`
					SelectionRequired          string  `xml:"selectionRequired"`
					NutritionalCharacteristics struct {
						AllergenList        string `xml:"allergenList"`
						DietRestrictionList string `xml:"dietRestrictionList"`
					} `xml:"nutritionalCharacteristics"`
				} `xml:"preparationMethod"`
			} `xml:"preparationMethodList"`
			LinkedArticleList struct {
				LinkedArticle []struct {
					ArticleNumber int    `xml:"articleNumber"`
					Quantity      string `xml:"quantity"`
					DecimalPlaces string `xml:"decimalPlaces"`
				} `xml:"linkedArticle"`
			} `xml:"linkedArticleList"`
			BarcodeList struct {
				Barcode struct {
					BarcodeId     string  `xml:"barcodeId"`
					Description   string  `xml:"description"`
					Quantity      string  `xml:"quantity"`
					DecimalPlaces string  `xml:"decimalPlaces"`
					PurchasePrice float64 `xml:"purchasePrice"`
					PriceIncl     float64 `xml:"priceIncl"`
					IsCurrent     string  `xml:"isCurrent"`
					Barcode       string  `xml:"barcode"`
					PriceExcl     float64 `xml:"priceExcl"`
				} `xml:"barcode"`
			} `xml:"barcodeList"`
			AskFor                     string `xml:"askFor"`
			Discontinued               string `xml:"discontinued"`
			NutritionalCharacteristics struct {
				AllergenList        string `xml:"allergenList"`
				DietRestrictionList string `xml:"dietRestrictionList"`
			} `xml:"nutritionalCharacteristics"`
			ComponentArticleList struct {
				ComponentArticle []struct {
					ArticleNumber     int     `xml:"articleNumber"`
					Quantity          string  `xml:"quantity"`
					DecimalPlaces     string  `xml:"decimalPlaces"`
					Description       string  `xml:"description"`
					InvoiceText       string  `xml:"invoiceText"`
					ReceiptText       string  `xml:"receiptText"`
					DisplayText       string  `xml:"displayText"`
					Barcode           string  `xml:"barcode"`
					TurnoverGroup     int     `xml:"turnoverGroup"`
					VatCode           string  `xml:"vatCode"`
					VatPercentage     float64 `xml:"vatPercentage"`
					PurchasePrice     float64 `xml:"purchasePrice"`
					PriceIncl         float64 `xml:"priceIncl"`
					PriceExcl         float64 `xml:"priceExcl"`
					LinkedArticleList struct {
						LinkedArticle struct {
							ArticleNumber int    `xml:"articleNumber"`
							Quantity      string `xml:"quantity"`
							DecimalPlaces string `xml:"decimalPlaces"`
						} `xml:"linkedArticle"`
					} `xml:"linkedArticleList"`
				} `xml:"componentArticle"`
			} `xml:"componentArticleList"`
			PriceGroupList struct {
				PriceGroup struct {
					PriceGroupNumber   string  `xml:"priceGroupNumber"`
					Description        string  `xml:"description"`
					DiscountPercentage float64 `xml:"discountPercentage"`
				} `xml:"priceGroup"`
			} `xml:"priceGroupList"`
			OrderQuantityValue         string `xml:"orderQuantityValue"`
			OrderQuantityDecimalPlaces string `xml:"orderQuantityDecimalPlaces"`
			SupplierRelationNumber     string `xml:"supplierRelationNumber"`
			SupplierRelationName       string `xml:"supplierRelationName"`
		} `xml:"article"`
	} `xml:"articleList"`
	SortOrderGroupList struct {
		SortOrderGroup []struct {
			GroupNumber string `xml:"groupNumber"`
			SortOrder   string `xml:"sortOrder"`
		} `xml:"sortOrderGroup"`
	} `xml:"sortOrderGroupList"`
	GroupNumbers []string `xml:"groupNumbers"`
}

type FinancialGroups []FinancialGroup

type FinancialGroup struct {
	FinancialGroupType    string  `xml:"financialGroupType"`
	FinancialGroupSource  string  `xml:"financialGroupSource"`
	FinancialGroupNumber  string  `xml:"financialGroupNumber"`
	FinancialGroupName    string  `xml:"financialGroupName"`
	AccountNumber         int     `xml:"accountNumber"`
	BranchNumber          int     `xml:"branchNumber"`
	ExtBranchId           int     `xml:"extBranchId"`
	FromFinancialDate     Date    `xml:"fromFinancialDate"`
	ThroughFinancialDate  Date    `xml:"throughFinancialDate"`
	FinancialPeriodClosed string  `xml:"financialPeriodClosed"`
	Quantity              int     `xml:"quantity"`
	DecimalPlaces         int     `xml:"decimalPlaces"`
	InclAmount            float64 `xml:"inclAmount"`
	ExclAmount            float64 `xml:"exclAmount"`
	VatGroupList          struct {
		VatGroup []struct {
			VatCode       int     `xml:"vatCode"`
			VatPercentage float64 `xml:"vatPercentage"`
			ExclAmount    float64 `xml:"exclAmount"`
			VatAmount     float64 `xml:"vatAmount"`
		} `xml:"vatGroup"`
	} `xml:"vatGroupList"`
}

type Date struct {
	Day  string `xml:"day"`
	Mon  string `xml:"mon"`
	Year string `xml:"year"`
}

type Orders []Order

type Order struct {
	OrderId    string `xml:"orderId"`
	ExtOrderId string `xml:"extOrderId"`
	InvoiceIds struct {
		ID string `xml:"id"`
	} `xml:"invoiceIds"`
	TransactionString  string    `xml:"transactionString"`
	SyncMarker         int       `xml:"syncMarker"`
	OrderType          string    `xml:"orderType"`
	EntryBranchNumber  string    `xml:"entryBranchNumber"`
	EmployeeNumber     string    `xml:"employeeNumber"`
	EmployeeName       string    `xml:"employeeName"`
	EntryTimestamp     Timestamp `xml:"entryTimestamp"`
	RelationNumber     string    `xml:"relationNumber"`
	RelationName       string    `xml:"relationName"`
	RelationCategoryId string    `xml:"relationCategoryId"`
	DeliveryAddress    struct {
		AddressId string `xml:"addressId"`
		Contact   string `xml:"contact"`
		Address   string `xml:"address"`
		Zipcode   string `xml:"zipcode"`
		City      string `xml:"city"`
		Country   string `xml:"country"`
	} `xml:"deliveryAddress"`
	FinancialDate struct {
		Day  string `xml:"day"`
		Mon  string `xml:"mon"`
		Year string `xml:"year"`
	} `xml:"financialDate"`
	FinancialBranchNumber string  `xml:"financialBranchNumber"`
	WorkplaceNumber       string  `xml:"workplaceNumber"`
	Reference             string  `xml:"reference"`
	TotalInclAmount       float64 `xml:"totalInclAmount"`
	TotalExclAmount       float64 `xml:"totalExclAmount"`
	VatMethod             string  `xml:"vatMethod"`
	VatGroupList          struct {
		VatGroup struct {
			VatCode       string  `xml:"vatCode"`
			VatPercentage float64 `xml:"vatPercentage"`
			ExclAmount    float64 `xml:"exclAmount"`
			VatAmount     float64 `xml:"vatAmount"`
		} `xml:"vatGroup"`
	} `xml:"vatGroupList"`
	DeliveryDate struct {
		Day  string `xml:"day"`
		Mon  string `xml:"mon"`
		Year string `xml:"year"`
	} `xml:"deliveryDate"`
	DeliveryPeriodBegin struct {
		Sec      string `xml:"sec"`
		Min      string `xml:"min"`
		Hour     string `xml:"hour"`
		Day      string `xml:"day"`
		Mon      string `xml:"mon"`
		Year     string `xml:"year"`
		Isdst    string `xml:"isdst"`
		Timezone string `xml:"timezone"`
	} `xml:"deliveryPeriodBegin"`
	DeliveryPeriodEnd struct {
		Sec      string `xml:"sec"`
		Min      string `xml:"min"`
		Hour     string `xml:"hour"`
		Day      string `xml:"day"`
		Mon      string `xml:"mon"`
		Year     string `xml:"year"`
		Isdst    string `xml:"isdst"`
		Timezone string `xml:"timezone"`
	} `xml:"deliveryPeriodEnd"`
	ChangeCounter string  `xml:"changeCounter"`
	VersionNumber string  `xml:"versionNumber"`
	PrepaidAmount float64 `xml:"prepaidAmount"`
	FullyPaid     string  `xml:"fullyPaid"`
	DeliveryState string  `xml:"deliveryState"`
	CancelState   string  `xml:"cancelState"`
	CompleteState string  `xml:"completeState"`
	OrderNumber   struct {
		Year   string `xml:"year"`
		Number string `xml:"number"`
	} `xml:"orderNumber"`
	OrderBarcode string `xml:"orderBarcode"`
	LineList     struct {
		Line struct {
			Chardata       string `xml:",chardata"`
			LineId         string `xml:"lineId"`
			EmployeeNumber string `xml:"employeeNumber"`
			ArticleNumber  string `xml:"articleNumber"`
			Data           struct {
				Quantity           string  `xml:"quantity"`
				DecimalPlaces      string  `xml:"decimalPlaces"`
				Price              float64 `xml:"price"`
				PriceExcl          float64 `xml:"priceExcl"`
				OriginalPrice      float64 `xml:"originalPrice"`
				OriginalPriceExcl  float64 `xml:"originalPriceExcl"`
				PurchasePrice      float64 `xml:"purchasePrice"`
				TurnoverGroup      string  `xml:"turnoverGroup"`
				TurnoverGroupType  string  `xml:"turnoverGroupType"`
				VatCode            string  `xml:"vatCode"`
				VatPercentage      float64 `xml:"vatPercentage"`
				DiscountType       string  `xml:"discountType"`
				DiscountPercentage float64 `xml:"discountPercentage"`
				DiscountAmount     float64 `xml:"discountAmount"`
				TotalInclAmount    float64 `xml:"totalInclAmount"`
				TotalExclAmount    float64 `xml:"totalExclAmount"`
				PriceType          string  `xml:"priceType"`
			} `xml:"data"`
			LineType string `xml:"lineType"`
		} `xml:"line"`
	} `xml:"lineList"`
	PaymentList    string `xml:"paymentList"`
	VatChange      string `xml:"vatChange"`
	VatCountryCode string `xml:"vatCountryCode"`
	VatCountryIso3 string `xml:"vatCountryIso3"`
}
