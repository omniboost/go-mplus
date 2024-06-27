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
		SyncMarker        int `xml:"urn:syncMarker,omitempty"`
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
		Order Orders `xml:"order"`
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

func (r *GetOrdersRequest) All() (GetOrdersRequestResponseBody, error) {
	r.RequestBody().Request.SyncMarker = -1

	response := *r.NewResponseBody()
	for {
		resp, err := r.Do()
		if err != nil {
			return resp, err
		}

		// Break out of loop when no orders are found
		if len(resp.OrderList.Order) == 0 {
			break
		}

		// Get latest sync marker
		for _, p := range resp.OrderList.Order {
			if p.SyncMarker >= r.RequestBody().Request.SyncMarker {
				r.RequestBody().Request.SyncMarker = (p.SyncMarker + 1)
			}
		}

		// Add orders to list
		response.OrderList.Order = append(response.OrderList.Order, resp.OrderList.Order...)
	}

	return response, nil
}

type GetOrdersBasic struct {
}

func (c GetOrdersBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c GetOrdersBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
