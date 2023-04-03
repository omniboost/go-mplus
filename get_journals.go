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

func (c *Client) NewGetJournalsRequest() GetJournalsRequest {
	r := GetJournalsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetJournalsRequest struct {
	client      *Client
	queryParams *GetJournalsRequestQueryParams
	pathParams  *GetJournalsRequestPathParams
	method      string
	headers     http.Header
	requestBody GetJournalsRequestBody
}

func (r GetJournalsRequest) SOAPAction() string {
	return ""
}

func (r GetJournalsRequest) NewQueryParams() *GetJournalsRequestQueryParams {
	return &GetJournalsRequestQueryParams{}
}

type GetJournalsRequestQueryParams struct {
}

func (p GetJournalsRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetJournalsRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetJournalsRequest) NewPathParams() *GetJournalsRequestPathParams {
	return &GetJournalsRequestPathParams{}
}

type GetJournalsRequestPathParams struct {
}

func (p *GetJournalsRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetJournalsRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetJournalsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetJournalsRequest) Method() string {
	return r.method
}

func (r GetJournalsRequest) NewRequestBody() GetJournalsRequestBody {
	return GetJournalsRequestBody{}
}

type GetJournalsRequestBody struct {
	XMLName xml.Name `xml:"urn:getJournals"`

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

func (r *GetJournalsRequest) RequestBody() *GetJournalsRequestBody {
	return &r.requestBody
}

func (r *GetJournalsRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetJournalsRequest) SetRequestBody(body GetJournalsRequestBody) {
	r.requestBody = body
}

func (r *GetJournalsRequest) NewResponseBody() *GetJournalsRequestResponseBody {
	return &GetJournalsRequestResponseBody{}
}

type GetJournalsRequestResponseBody struct {
	XMLName xml.Name `xml:"GetJournalsResponse"`

	JournalList struct {
		Journals Journals `xml:"journal"`
	} `xml:"journalList"`
}

func (r *GetJournalsRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *GetJournalsRequest) Do() (GetJournalsRequestResponseBody, error) {
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

type GetJournalsBasic struct {
}

func (c GetJournalsBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c GetJournalsBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
