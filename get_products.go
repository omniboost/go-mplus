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

func (c *Client) NewGetProductsRequest() GetProductsRequest {
	r := GetProductsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetProductsRequest struct {
	client      *Client
	queryParams *GetProductsRequestQueryParams
	pathParams  *GetProductsRequestPathParams
	method      string
	headers     http.Header
	requestBody GetProductsRequestBody
}

func (r GetProductsRequest) SOAPAction() string {
	return ""
}

func (r GetProductsRequest) NewQueryParams() *GetProductsRequestQueryParams {
	return &GetProductsRequestQueryParams{}
}

type GetProductsRequestQueryParams struct {
}

func (p GetProductsRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetProductsRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetProductsRequest) NewPathParams() *GetProductsRequestPathParams {
	return &GetProductsRequestPathParams{}
}

type GetProductsRequestPathParams struct {
}

func (p *GetProductsRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetProductsRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetProductsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetProductsRequest) Method() string {
	return r.method
}

func (r GetProductsRequest) NewRequestBody() GetProductsRequestBody {
	return GetProductsRequestBody{}
}

type GetProductsRequestBody struct {
	XMLName xml.Name `xml:"urn:getProducts"`

	Request struct {
		SyncMarker     int   `xml:"urn:syncMarker,omitempty"`
		OnlyActive     bool  `xml:"urn:onlyActive"`
		ArticleNumbers []int `xml:"urn:articleNumbers"`
	} `xml:"urn:request"`
}

func (r *GetProductsRequest) RequestBody() *GetProductsRequestBody {
	return &r.requestBody
}

func (r *GetProductsRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetProductsRequest) SetRequestBody(body GetProductsRequestBody) {
	r.requestBody = body
}

func (r *GetProductsRequest) NewResponseBody() *GetProductsRequestResponseBody {
	return &GetProductsRequestResponseBody{}
}

type GetProductsRequestResponseBody struct {
	XMLName xml.Name `xml:"GetProductsResponse"`

	ProductList struct {
		Product Products `xml:"product"`
	} `xml:"productList"`
}

func (r *GetProductsRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *GetProductsRequest) Do() (GetProductsRequestResponseBody, error) {
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

func (r *GetProductsRequest) All() (GetProductsRequestResponseBody, error) {
	r.RequestBody().Request.SyncMarker = -1
	resp, err := r.Do()
	if err != nil {
		return resp, err
	}

	concat := *r.NewResponseBody()
	concat.ProductList.Product = resp.ProductList.Product

	for len(resp.ProductList.Product) > 0 {
		for _, p := range resp.ProductList.Product {
			if p.SyncMarker >= r.RequestBody().Request.SyncMarker {
				r.RequestBody().Request.SyncMarker = (p.SyncMarker + 1)
			}
		}

		resp, err = r.Do()
		if err != nil {
			return concat, err
		}

		concat.ProductList.Product = append(concat.ProductList.Product, resp.ProductList.Product...)
	}

	return concat, nil
}

type GetProductsBasic struct {
}

func (c GetProductsBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c GetProductsBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
