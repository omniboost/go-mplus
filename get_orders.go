package mplus

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/cydev/zero"
	"github.com/omniboost/go-mplus/omitempty"
	"github.com/omniboost/go-mplus/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewGetOrdersRequest() GetOrdersRequest {
	r := GetOrdersRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetOrdersRequest struct {
	client      *Client
	queryParams *GetOrdersRequestQueryParams
	pathParams  *GetOrdersRequestPathParams
	method      string
	headers     http.Header
	requestBody GetOrdersRequestBody
}

func (r GetOrdersRequest) SOAPAction() string {
	return ""
}

func (r GetOrdersRequest) NewQueryParams() *GetOrdersRequestQueryParams {
	return &GetOrdersRequestQueryParams{}
}

type GetOrdersRequestQueryParams struct {
}

func (p GetOrdersRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetOrdersRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetOrdersRequest) NewPathParams() *GetOrdersRequestPathParams {
	return &GetOrdersRequestPathParams{}
}

type GetOrdersRequestPathParams struct {
}

func (p *GetOrdersRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetOrdersRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetOrdersRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetOrdersRequest) Method() string {
	return r.method
}

func (r GetOrdersRequest) NewRequestBody() GetOrdersRequestBody {
	return GetOrdersRequestBody{}
}

type GetOrdersRequestBody struct {
	XMLName xml.Name `xml:"urn:getOrders"`

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

func (r *GetOrdersRequest) RequestBody() *GetOrdersRequestBody {
	return &r.requestBody
}

func (r *GetOrdersRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetOrdersRequest) SetRequestBody(body GetOrdersRequestBody) {
	r.requestBody = body
}

func (r *GetOrdersRequest) NewResponseBody() *GetOrdersRequestResponseBody {
	return &GetOrdersRequestResponseBody{}
}

type GetOrdersRequestResponseBody struct {
	XMLName xml.Name `xml:"GetOrdersResponse"`

	OrderList struct {
		Order []struct {
			OrderId    string `xml:"orderId"`
			ExtOrderId string `xml:"extOrderId"`
			InvoiceIds struct {
				ID string `xml:"id"`
			} `xml:"invoiceIds"`
			TransactionString  string    `xml:"transactionString"`
			SyncMarker         string    `xml:"syncMarker"`
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
		} `xml:"order"`
	} `xml:"orderList"`
}

func (r *GetOrdersRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *GetOrdersRequest) Do() (GetOrdersRequestResponseBody, error) {
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

type GetOrdersBasic struct {
}

func (c GetOrdersBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c GetOrdersBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
