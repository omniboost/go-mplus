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

func (c *Client) NewGetFinancialJournalRequest() GetFinancialJournalRequest {
	r := GetFinancialJournalRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetFinancialJournalRequest struct {
	client      *Client
	queryParams *GetFinancialJournalRequestQueryParams
	pathParams  *GetFinancialJournalRequestPathParams
	method      string
	headers     http.Header
	requestBody GetFinancialJournalRequestBody
}

func (r GetFinancialJournalRequest) SOAPAction() string {
	return ""
}

func (r GetFinancialJournalRequest) NewQueryParams() *GetFinancialJournalRequestQueryParams {
	return &GetFinancialJournalRequestQueryParams{}
}

type GetFinancialJournalRequestQueryParams struct {
}

func (p GetFinancialJournalRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetFinancialJournalRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetFinancialJournalRequest) NewPathParams() *GetFinancialJournalRequestPathParams {
	return &GetFinancialJournalRequestPathParams{}
}

type GetFinancialJournalRequestPathParams struct {
}

func (p *GetFinancialJournalRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetFinancialJournalRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetFinancialJournalRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetFinancialJournalRequest) Method() string {
	return r.method
}

func (r GetFinancialJournalRequest) NewRequestBody() GetFinancialJournalRequestBody {
	return GetFinancialJournalRequestBody{}
}

type GetFinancialJournalRequestBody struct {
	XMLName xml.Name `xml:"urn:getFinancialJournal"`

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

func (r *GetFinancialJournalRequest) RequestBody() *GetFinancialJournalRequestBody {
	return &r.requestBody
}

func (r *GetFinancialJournalRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetFinancialJournalRequest) SetRequestBody(body GetFinancialJournalRequestBody) {
	r.requestBody = body
}

func (r *GetFinancialJournalRequest) NewResponseBody() *GetFinancialJournalRequestResponseBody {
	return &GetFinancialJournalRequestResponseBody{}
}

type GetFinancialJournalRequestResponseBody struct {
	XMLName xml.Name `xml:"searchResponse"`
}

func (r *GetFinancialJournalRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *GetFinancialJournalRequest) Do() (GetFinancialJournalRequestResponseBody, error) {
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

type GetFinancialJournalBasic struct {
}

func (c GetFinancialJournalBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c GetFinancialJournalBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
