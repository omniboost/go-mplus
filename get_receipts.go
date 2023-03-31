package mplus

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/cydev/zero"
	"github.com/omniboost/go-netsuite-soap/omitempty"
	"github.com/omniboost/go-netsuite-soap/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewGetReceiptsRequest() GetReceiptsRequest {
	r := GetReceiptsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetReceiptsRequest struct {
	client      *Client
	queryParams *GetReceiptsRequestQueryParams
	pathParams  *GetReceiptsRequestPathParams
	method      string
	headers     http.Header
	requestBody GetReceiptsRequestBody
}

func (r GetReceiptsRequest) SOAPAction() string {
	return ""
}

func (r GetReceiptsRequest) NewQueryParams() *GetReceiptsRequestQueryParams {
	return &GetReceiptsRequestQueryParams{}
}

type GetReceiptsRequestQueryParams struct {
}

func (p GetReceiptsRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetReceiptsRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetReceiptsRequest) NewPathParams() *GetReceiptsRequestPathParams {
	return &GetReceiptsRequestPathParams{}
}

type GetReceiptsRequestPathParams struct {
}

func (p *GetReceiptsRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetReceiptsRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetReceiptsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetReceiptsRequest) Method() string {
	return r.method
}

func (r GetReceiptsRequest) NewRequestBody() GetReceiptsRequestBody {
	return GetReceiptsRequestBody{}
}

type GetReceiptsRequestBody struct {
	XMLName xml.Name `xml:"urn:getReceipts"`

	Request struct {
		FromFinancialDate struct {
			Day  int `xml:"urn:day"`
			Mon  int `xml:"urn:mon"`
			Year int `xml:"urn:year"`
		} `xml:"urn:fromFinancialDate"`
		ThroughFinancialDate struct {
			Day  int `xml:"urn:day"`
			Mon  int `xml:"urn:mon"`
			Year int `xml:"urn:year"`
		} `xml:"urn:throughFinancialDate"`
	} `xml:"urn:request"`
}

func (r *GetReceiptsRequest) RequestBody() *GetReceiptsRequestBody {
	return &r.requestBody
}

func (r *GetReceiptsRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetReceiptsRequest) SetRequestBody(body GetReceiptsRequestBody) {
	r.requestBody = body
}

func (r *GetReceiptsRequest) NewResponseBody() *GetReceiptsRequestResponseBody {
	return &GetReceiptsRequestResponseBody{}
}

type GetReceiptsRequestResponseBody struct {
	XMLName xml.Name `xml:"GetReceiptsResponse"`

	ReceiptList struct {
		Receipt []struct {
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
		} `xml:"receipt"`
	} `xml:"receiptList"`
}

func (r *GetReceiptsRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *GetReceiptsRequest) Do() (GetReceiptsRequestResponseBody, error) {
	var err error

	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	if err != nil {
		return *responseBody, errors.WithStack(err)
	}

	return *responseBody, nil
}

type GetReceiptsBasic struct {
}

func (c GetReceiptsBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c GetReceiptsBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
