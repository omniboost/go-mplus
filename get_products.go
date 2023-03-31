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
		Product []struct {
			ProductNumber string `xml:"productNumber"`
			SyncMarker    string `xml:"syncMarker"`
			Description   string `xml:"description"`
			ExtraText     string `xml:"extraText"`
			ArticleList   struct {
				Article struct {
					ArticleNumber         int       `xml:"articleNumber"`
					PluNumber             int       `xml:"pluNumber"`
					SyncMarker            string    `xml:"syncMarker"`
					ChangeTimestamp       Timestamp `xml:"changeTimestamp"`
					CreateTimestamp       Timestamp `xml:"createTimestamp"`
					Active                string    `xml:"active"`
					Description           string    `xml:"description"`
					Colour                string    `xml:"colour"`
					Size                  string    `xml:"size"`
					InvoiceText           string    `xml:"invoiceText"`
					ReceiptText           string    `xml:"receiptText"`
					DisplayText           string    `xml:"displayText"`
					Barcode               string    `xml:"barcode"`
					ExtraText             string    `xml:"extraText"`
					TurnoverGroup         string    `xml:"turnoverGroup"`
					VatCode               string    `xml:"vatCode"`
					VatPercentage         float64   `xml:"vatPercentage"`
					AveragePurchasePrice  float64   `xml:"averagePurchasePrice"`
					PurchasePrice         float64   `xml:"purchasePrice"`
					PriceIncl             float64   `xml:"priceIncl"`
					PriceExcl             float64   `xml:"priceExcl"`
					SupplierArticleNumber string    `xml:"supplierArticleNumber"`
					Webshop               string    `xml:"webshop"`
					CategoryId            string    `xml:"categoryId"`
					CategoryIds           struct {
						Category string `xml:"category"`
					} `xml:"categoryIds"`
					StockArticle string `xml:"stockArticle"`
					Course       struct {
						Number              string `xml:"number"`
						Name                string `xml:"name"`
						Abbreviation        string `xml:"abbreviation"`
						SequenceNumber      string `xml:"sequenceNumber"`
						Type                string `xml:"type"`
						DoSetTable          string `xml:"doSetTable"`
						CanBeServedWithWine string `xml:"canBeServedWithWine"`
					} `xml:"course"`
					SpecialPrice                   string `xml:"specialPrice"`
					ShowPreparationMethodsDirectly string `xml:"showPreparationMethodsDirectly"`
					ShowPreparationMethodsPerGroup string `xml:"showPreparationMethodsPerGroup"`
					ImageList                      struct {
						Image struct {
							ImageId          string    `xml:"imageId"`
							ImageName        string    `xml:"imageName"`
							ImageHash        string    `xml:"imageHash"`
							ImagePath        string    `xml:"imagePath"`
							ImageUrl         string    `xml:"imageUrl"`
							ImageData        string    `xml:"imageData"`
							ImageDataResult  string    `xml:"imageDataResult"`
							ThumbHash        string    `xml:"thumbHash"`
							ThumbPath        string    `xml:"thumbPath"`
							ThumbUrl         string    `xml:"thumbUrl"`
							ThumbData        string    `xml:"thumbData"`
							ThumbDataResult  string    `xml:"thumbDataResult"`
							ImageWidth       string    `xml:"imageWidth"`
							ImageHeight      string    `xml:"imageHeight"`
							ThumbWidth       string    `xml:"thumbWidth"`
							ThumbHeight      string    `xml:"thumbHeight"`
							CreatedTimestamp Timestamp `xml:"createdTimestamp"`
							ChangedTimestamp Timestamp `xml:"changedTimestamp"`
						} `xml:"image"`
					} `xml:"imageList"`
					CustomFieldList struct {
						CustomField []struct {
							FieldName     string `xml:"fieldName"`
							DataType      string `xml:"dataType"`
							StrValue      string `xml:"strValue"`
							IntValue      string `xml:"intValue"`
							DecimalPlaces string `xml:"decimalPlaces"`
						} `xml:"customField"`
					} `xml:"customFieldList"`
					PreparationMethodList struct {
						PreparationMethod []struct {
							ArticleNumber              string  `xml:"articleNumber"`
							Description                string  `xml:"description"`
							PriceIncl                  float64 `xml:"priceIncl"`
							VatCode                    string  `xml:"vatCode"`
							VatPercentage              float64 `xml:"vatPercentage"`
							SyncMarker                 string  `xml:"syncMarker"`
							CloseAfterSelection        string  `xml:"closeAfterSelection"`
							GroupNumber                string  `xml:"groupNumber"`
							GroupName                  string  `xml:"groupName"`
							MaxOneSelection            string  `xml:"maxOneSelection"`
							SelectionRequired          string  `xml:"selectionRequired"`
							NutritionalCharacteristics struct {
								AllergenList        string `xml:"allergenList"`
								DietRestrictionList string `xml:"dietRestrictionList"`
							} `xml:"nutritionalCharacteristics"`
						} `xml:"preparationMethod"`
					} `xml:"preparationMethodList"`
					LinkedArticleList struct {
						LinkedArticle []struct {
							ArticleNumber string `xml:"articleNumber"`
							Quantity      string `xml:"quantity"`
							DecimalPlaces string `xml:"decimalPlaces"`
						} `xml:"linkedArticle"`
					} `xml:"linkedArticleList"`
					BarcodeList struct {
						Barcode struct {
							BarcodeId     string  `xml:"barcodeId"`
							Description   string  `xml:"description"`
							Quantity      string  `xml:"quantity"`
							DecimalPlaces string  `xml:"decimalPlaces"`
							PurchasePrice float64 `xml:"purchasePrice"`
							PriceIncl     float64 `xml:"priceIncl"`
							IsCurrent     string  `xml:"isCurrent"`
							Barcode       string  `xml:"barcode"`
							PriceExcl     float64 `xml:"priceExcl"`
						} `xml:"barcode"`
					} `xml:"barcodeList"`
					AskFor                     string `xml:"askFor"`
					Discontinued               string `xml:"discontinued"`
					NutritionalCharacteristics struct {
						AllergenList        string `xml:"allergenList"`
						DietRestrictionList string `xml:"dietRestrictionList"`
					} `xml:"nutritionalCharacteristics"`
					ComponentArticleList struct {
						ComponentArticle []struct {
							ArticleNumber     string  `xml:"articleNumber"`
							Quantity          string  `xml:"quantity"`
							DecimalPlaces     string  `xml:"decimalPlaces"`
							Description       string  `xml:"description"`
							InvoiceText       string  `xml:"invoiceText"`
							ReceiptText       string  `xml:"receiptText"`
							DisplayText       string  `xml:"displayText"`
							Barcode           string  `xml:"barcode"`
							TurnoverGroup     string  `xml:"turnoverGroup"`
							VatCode           string  `xml:"vatCode"`
							VatPercentage     float64 `xml:"vatPercentage"`
							PurchasePrice     float64 `xml:"purchasePrice"`
							PriceIncl         float64 `xml:"priceIncl"`
							PriceExcl         float64 `xml:"priceExcl"`
							LinkedArticleList struct {
								LinkedArticle struct {
									ArticleNumber string `xml:"articleNumber"`
									Quantity      string `xml:"quantity"`
									DecimalPlaces string `xml:"decimalPlaces"`
								} `xml:"linkedArticle"`
							} `xml:"linkedArticleList"`
						} `xml:"componentArticle"`
					} `xml:"componentArticleList"`
					PriceGroupList struct {
						PriceGroup struct {
							PriceGroupNumber   string  `xml:"priceGroupNumber"`
							Description        string  `xml:"description"`
							DiscountPercentage float64 `xml:"discountPercentage"`
						} `xml:"priceGroup"`
					} `xml:"priceGroupList"`
					OrderQuantityValue         string `xml:"orderQuantityValue"`
					OrderQuantityDecimalPlaces string `xml:"orderQuantityDecimalPlaces"`
					SupplierRelationNumber     string `xml:"supplierRelationNumber"`
					SupplierRelationName       string `xml:"supplierRelationName"`
				} `xml:"article"`
			} `xml:"articleList"`
			SortOrderGroupList struct {
				SortOrderGroup []struct {
					GroupNumber string `xml:"groupNumber"`
					SortOrder   string `xml:"sortOrder"`
				} `xml:"sortOrderGroup"`
			} `xml:"sortOrderGroupList"`
			GroupNumbers []string `xml:"groupNumbers"`
		} `xml:"product"`
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

type GetProductsBasic struct {
}

func (c GetProductsBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c GetProductsBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
