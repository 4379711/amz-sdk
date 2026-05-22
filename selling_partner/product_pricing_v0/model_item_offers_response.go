package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemOffersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemOffersResponse{}

// ItemOffersResponse Schema for an individual `ItemOffersResponse`
type ItemOffersResponse struct {
	Headers *HttpResponseHeaders     `json:"headers,omitempty"`
	Status  *GetOffersHttpStatusLine `json:"status,omitempty"`
	Body    GetOffersResponse        `json:"body"`
	Request ItemOffersRequestParams  `json:"request"`
}

type _ItemOffersResponse ItemOffersResponse

// NewItemOffersResponse instantiates a new ItemOffersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemOffersResponse(body GetOffersResponse, request ItemOffersRequestParams) *ItemOffersResponse {
	this := ItemOffersResponse{}
	this.Body = body
	this.Request = request
	return &this
}

// NewItemOffersResponseWithDefaults instantiates a new ItemOffersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemOffersResponseWithDefaults() *ItemOffersResponse {
	this := ItemOffersResponse{}
	return &this
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *ItemOffersResponse) GetHeaders() HttpResponseHeaders {
	if o == nil || IsNil(o.Headers) {
		var ret HttpResponseHeaders
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOffersResponse) GetHeadersOk() (*HttpResponseHeaders, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *ItemOffersResponse) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given HttpResponseHeaders and assigns it to the Headers field.
func (o *ItemOffersResponse) SetHeaders(v HttpResponseHeaders) {
	o.Headers = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ItemOffersResponse) GetStatus() GetOffersHttpStatusLine {
	if o == nil || IsNil(o.Status) {
		var ret GetOffersHttpStatusLine
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOffersResponse) GetStatusOk() (*GetOffersHttpStatusLine, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ItemOffersResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GetOffersHttpStatusLine and assigns it to the Status field.
func (o *ItemOffersResponse) SetStatus(v GetOffersHttpStatusLine) {
	o.Status = &v
}

// GetBody returns the Body field value
func (o *ItemOffersResponse) GetBody() GetOffersResponse {
	if o == nil {
		var ret GetOffersResponse
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *ItemOffersResponse) GetBodyOk() (*GetOffersResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *ItemOffersResponse) SetBody(v GetOffersResponse) {
	o.Body = v
}

// GetRequest returns the Request field value
func (o *ItemOffersResponse) GetRequest() ItemOffersRequestParams {
	if o == nil {
		var ret ItemOffersRequestParams
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *ItemOffersResponse) GetRequestOk() (*ItemOffersRequestParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Request, true
}

// SetRequest sets field value
func (o *ItemOffersResponse) SetRequest(v ItemOffersRequestParams) {
	o.Request = v
}

func (o ItemOffersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["body"] = o.Body
	toSerialize["request"] = o.Request
	return toSerialize, nil
}

type NullableItemOffersResponse struct {
	value *ItemOffersResponse
	isSet bool
}

func (v NullableItemOffersResponse) Get() *ItemOffersResponse {
	return v.value
}

func (v *NullableItemOffersResponse) Set(val *ItemOffersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableItemOffersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableItemOffersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemOffersResponse(val *ItemOffersResponse) *NullableItemOffersResponse {
	return &NullableItemOffersResponse{value: val, isSet: true}
}

func (v NullableItemOffersResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemOffersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
