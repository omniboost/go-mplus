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
		BranchNumbers []int `xml:"urn:branchNumbers,omitempty"`
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
		Receipt Receipts `xml:"receipt"`
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
