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

func (c *Client) NewGetTurnoverGroupsRequest() GetTurnoverGroupsRequest {
	r := GetTurnoverGroupsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetTurnoverGroupsRequest struct {
	client      *Client
	queryParams *GetTurnoverGroupsRequestQueryParams
	pathParams  *GetTurnoverGroupsRequestPathParams
	method      string
	headers     http.Header
	requestBody GetTurnoverGroupsRequestBody
}

func (r GetTurnoverGroupsRequest) SOAPAction() string {
	return ""
}

func (r GetTurnoverGroupsRequest) NewQueryParams() *GetTurnoverGroupsRequestQueryParams {
	return &GetTurnoverGroupsRequestQueryParams{}
}

type GetTurnoverGroupsRequestQueryParams struct {
}

func (p GetTurnoverGroupsRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetTurnoverGroupsRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetTurnoverGroupsRequest) NewPathParams() *GetTurnoverGroupsRequestPathParams {
	return &GetTurnoverGroupsRequestPathParams{}
}

type GetTurnoverGroupsRequestPathParams struct {
}

func (p *GetTurnoverGroupsRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTurnoverGroupsRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetTurnoverGroupsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTurnoverGroupsRequest) Method() string {
	return r.method
}

func (r GetTurnoverGroupsRequest) NewRequestBody() GetTurnoverGroupsRequestBody {
	return GetTurnoverGroupsRequestBody{}
}

type GetTurnoverGroupsRequestBody struct {
	XMLName xml.Name `xml:"urn:getTurnoverGroups"`
}

func (r *GetTurnoverGroupsRequest) RequestBody() *GetTurnoverGroupsRequestBody {
	return &r.requestBody
}

func (r *GetTurnoverGroupsRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetTurnoverGroupsRequest) SetRequestBody(body GetTurnoverGroupsRequestBody) {
	r.requestBody = body
}

func (r *GetTurnoverGroupsRequest) NewResponseBody() *GetTurnoverGroupsRequestResponseBody {
	return &GetTurnoverGroupsRequestResponseBody{}
}

type GetTurnoverGroupsRequestResponseBody struct {
	XMLName xml.Name `xml:"GetTurnoverGroupsResponse"`

	TurnoverGroupList struct {
		Text          string `xml:",chardata"`
		TurnoverGroup []struct {
			Text                    string `xml:",chardata"`
			TurnoverGroupType       string `xml:"turnoverGroupType"`
			TurnoverGroup           string `xml:"turnoverGroup"`
			TurnoverGroupName       string `xml:"turnoverGroupName"`
			AllowPointsDistribution string `xml:"allowPointsDistribution"`
			AllowPointsPayment      string `xml:"allowPointsPayment"`
			AllowDiscount           string `xml:"allowDiscount"`
			BranchAccountNumberList struct {
				Text                string `xml:",chardata"`
				BranchAccountNumber []struct {
					Text          string `xml:",chardata"`
					BranchNumber  string `xml:"branchNumber"`
					AccountNumber string `xml:"accountNumber"`
				} `xml:"branchAccountNumber"`
			} `xml:"branchAccountNumberList"`
		} `xml:"turnoverGroup"`
	} `xml:"turnoverGroupList"`
}

func (r *GetTurnoverGroupsRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *GetTurnoverGroupsRequest) Do() (GetTurnoverGroupsRequestResponseBody, error) {
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

type GetTurnoverGroupsBasic struct {
}

func (c GetTurnoverGroupsBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c GetTurnoverGroupsBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
