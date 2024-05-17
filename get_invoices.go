package mplus

import (
	"context"
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/cydev/zero"
	"github.com/omniboost/go-mplus/omitempty"
	"github.com/omniboost/go-mplus/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewGetInvoicesRequest() GetInvoicesRequest {
	r := GetInvoicesRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetInvoicesRequest struct {
	client      *Client
	queryParams *GetInvoicesRequestQueryParams
	pathParams  *GetInvoicesRequestPathParams
	method      string
	headers     http.Header
	requestBody GetInvoicesRequestBody
}

func (r GetInvoicesRequest) SOAPAction() string {
	return ""
}

func (r GetInvoicesRequest) NewQueryParams() *GetInvoicesRequestQueryParams {
	return &GetInvoicesRequestQueryParams{}
}

type GetInvoicesRequestQueryParams struct {
}

func (p GetInvoicesRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetInvoicesRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetInvoicesRequest) NewPathParams() *GetInvoicesRequestPathParams {
	return &GetInvoicesRequestPathParams{}
}

type GetInvoicesRequestPathParams struct {
}

func (p *GetInvoicesRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetInvoicesRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetInvoicesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetInvoicesRequest) Method() string {
	return r.method
}

func (r GetInvoicesRequest) NewRequestBody() GetInvoicesRequestBody {
	return GetInvoicesRequestBody{}
}

type GetInvoicesRequestBody struct {
	XMLName xml.Name `xml:"urn:getInvoices"`

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

func (r *GetInvoicesRequest) RequestBody() *GetInvoicesRequestBody {
	return &r.requestBody
}

func (r *GetInvoicesRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetInvoicesRequest) SetRequestBody(body GetInvoicesRequestBody) {
	r.requestBody = body
}

func (r *GetInvoicesRequest) NewResponseBody() *GetInvoicesRequestResponseBody {
	return &GetInvoicesRequestResponseBody{}
}

type GetInvoicesRequestResponseBody struct {
	XMLName xml.Name `xml:"GetInvoicesResponse"`

	InvoiceList struct {
		Invoice Invoices `xml:"invoice"`
	} `xml:"invoiceList"`
}

func (r *GetInvoicesRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *GetInvoicesRequest) Do() (GetInvoicesRequestResponseBody, error) {
	var err error

	// Create http request
	req, err := r.client.NewRequest(context.TODO(), r)
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

type GetInvoicesBasic struct {
}

func (c GetInvoicesBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c GetInvoicesBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
