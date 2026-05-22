package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the FeaturedOfferExpectedPriceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeaturedOfferExpectedPriceResponse{}

// FeaturedOfferExpectedPriceResponse Schema for an individual FOEP response.
type FeaturedOfferExpectedPriceResponse struct {
	// A mapping of additional HTTP headers to send or receive for an individual request within a batch.
	Headers map[string]string                       `json:"headers"`
	Status  HttpStatusLine                          `json:"status"`
	Request FeaturedOfferExpectedPriceRequestParams `json:"request"`
	Body    *FeaturedOfferExpectedPriceResponseBody `json:"body,omitempty"`
}

type _FeaturedOfferExpectedPriceResponse FeaturedOfferExpectedPriceResponse

// NewFeaturedOfferExpectedPriceResponse instantiates a new FeaturedOfferExpectedPriceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeaturedOfferExpectedPriceResponse(headers map[string]string, status HttpStatusLine, request FeaturedOfferExpectedPriceRequestParams) *FeaturedOfferExpectedPriceResponse {
	this := FeaturedOfferExpectedPriceResponse{}
	this.Headers = headers
	this.Status = status
	this.Request = request
	return &this
}

// NewFeaturedOfferExpectedPriceResponseWithDefaults instantiates a new FeaturedOfferExpectedPriceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeaturedOfferExpectedPriceResponseWithDefaults() *FeaturedOfferExpectedPriceResponse {
	this := FeaturedOfferExpectedPriceResponse{}
	return &this
}

// GetHeaders returns the Headers field value
func (o *FeaturedOfferExpectedPriceResponse) GetHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *FeaturedOfferExpectedPriceResponse) GetHeadersOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headers, true
}

// SetHeaders sets field value
func (o *FeaturedOfferExpectedPriceResponse) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetStatus returns the Status field value
func (o *FeaturedOfferExpectedPriceResponse) GetStatus() HttpStatusLine {
	if o == nil {
		var ret HttpStatusLine
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FeaturedOfferExpectedPriceResponse) GetStatusOk() (*HttpStatusLine, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *FeaturedOfferExpectedPriceResponse) SetStatus(v HttpStatusLine) {
	o.Status = v
}

// GetRequest returns the Request field value
func (o *FeaturedOfferExpectedPriceResponse) GetRequest() FeaturedOfferExpectedPriceRequestParams {
	if o == nil {
		var ret FeaturedOfferExpectedPriceRequestParams
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *FeaturedOfferExpectedPriceResponse) GetRequestOk() (*FeaturedOfferExpectedPriceRequestParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Request, true
}

// SetRequest sets field value
func (o *FeaturedOfferExpectedPriceResponse) SetRequest(v FeaturedOfferExpectedPriceRequestParams) {
	o.Request = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *FeaturedOfferExpectedPriceResponse) GetBody() FeaturedOfferExpectedPriceResponseBody {
	if o == nil || IsNil(o.Body) {
		var ret FeaturedOfferExpectedPriceResponseBody
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeaturedOfferExpectedPriceResponse) GetBodyOk() (*FeaturedOfferExpectedPriceResponseBody, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *FeaturedOfferExpectedPriceResponse) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given FeaturedOfferExpectedPriceResponseBody and assigns it to the Body field.
func (o *FeaturedOfferExpectedPriceResponse) SetBody(v FeaturedOfferExpectedPriceResponseBody) {
	o.Body = &v
}

func (o FeaturedOfferExpectedPriceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["headers"] = o.Headers
	toSerialize["status"] = o.Status
	toSerialize["request"] = o.Request
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	return toSerialize, nil
}

type NullableFeaturedOfferExpectedPriceResponse struct {
	value *FeaturedOfferExpectedPriceResponse
	isSet bool
}

func (v NullableFeaturedOfferExpectedPriceResponse) Get() *FeaturedOfferExpectedPriceResponse {
	return v.value
}

func (v *NullableFeaturedOfferExpectedPriceResponse) Set(val *FeaturedOfferExpectedPriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFeaturedOfferExpectedPriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFeaturedOfferExpectedPriceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeaturedOfferExpectedPriceResponse(val *FeaturedOfferExpectedPriceResponse) *NullableFeaturedOfferExpectedPriceResponse {
	return &NullableFeaturedOfferExpectedPriceResponse{value: val, isSet: true}
}

func (v NullableFeaturedOfferExpectedPriceResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeaturedOfferExpectedPriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
