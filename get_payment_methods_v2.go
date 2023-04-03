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

func (c *Client) NewGetPaymentMethodsV2Request() GetPaymentMethodsV2Request {
	r := GetPaymentMethodsV2Request{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetPaymentMethodsV2Request struct {
	client      *Client
	queryParams *GetPaymentMethodsV2RequestQueryParams
	pathParams  *GetPaymentMethodsV2RequestPathParams
	method      string
	headers     http.Header
	requestBody GetPaymentMethodsV2RequestBody
}

func (r GetPaymentMethodsV2Request) SOAPAction() string {
	return ""
}

func (r GetPaymentMethodsV2Request) NewQueryParams() *GetPaymentMethodsV2RequestQueryParams {
	return &GetPaymentMethodsV2RequestQueryParams{}
}

type GetPaymentMethodsV2RequestQueryParams struct {
}

func (p GetPaymentMethodsV2RequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetPaymentMethodsV2Request) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetPaymentMethodsV2Request) NewPathParams() *GetPaymentMethodsV2RequestPathParams {
	return &GetPaymentMethodsV2RequestPathParams{}
}

type GetPaymentMethodsV2RequestPathParams struct {
}

func (p *GetPaymentMethodsV2RequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetPaymentMethodsV2Request) PathParams() PathParams {
	return r.pathParams
}

func (r *GetPaymentMethodsV2Request) SetMethod(method string) {
	r.method = method
}

func (r *GetPaymentMethodsV2Request) Method() string {
	return r.method
}

func (r GetPaymentMethodsV2Request) NewRequestBody() GetPaymentMethodsV2RequestBody {
	return GetPaymentMethodsV2RequestBody{}
}

type GetPaymentMethodsV2RequestBody struct {
	XMLName xml.Name `xml:"urn:getPaymentMethodsV2"`

	Request struct{} `xml:"urn:request"`
}

func (r *GetPaymentMethodsV2Request) RequestBody() *GetPaymentMethodsV2RequestBody {
	return &r.requestBody
}

func (r *GetPaymentMethodsV2Request) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetPaymentMethodsV2Request) SetRequestBody(body GetPaymentMethodsV2RequestBody) {
	r.requestBody = body
}

func (r *GetPaymentMethodsV2Request) NewResponseBody() *GetPaymentMethodsV2RequestResponseBody {
	return &GetPaymentMethodsV2RequestResponseBody{}
}

type GetPaymentMethodsV2RequestResponseBody struct {
	XMLName xml.Name `xml:"GetPaymentMethodsResponse"`

	PaymentMethods PaymentMethods `xml:"paymentMethodList>paymentMethod"`
}

func (r *GetPaymentMethodsV2Request) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *GetPaymentMethodsV2Request) Do() (GetPaymentMethodsV2RequestResponseBody, error) {
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

type GetPaymentMethodsV2Basic struct {
}

func (c GetPaymentMethodsV2Basic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c GetPaymentMethodsV2Basic) IsEmpty() bool {
	return zero.IsZero(c)
}
