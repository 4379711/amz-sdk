package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ListingOffersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListingOffersResponse{}

// ListingOffersResponse Schema for an individual `ListingOffersResponse`
type ListingOffersResponse struct {
	Headers *HttpResponseHeaders        `json:"headers,omitempty"`
	Status  *GetOffersHttpStatusLine    `json:"status,omitempty"`
	Body    GetOffersResponse           `json:"body"`
	Request *ListingOffersRequestParams `json:"request,omitempty"`
}

type _ListingOffersResponse ListingOffersResponse

// NewListingOffersResponse instantiates a new ListingOffersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListingOffersResponse(body GetOffersResponse) *ListingOffersResponse {
	this := ListingOffersResponse{}
	this.Body = body
	return &this
}

// NewListingOffersResponseWithDefaults instantiates a new ListingOffersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListingOffersResponseWithDefaults() *ListingOffersResponse {
	this := ListingOffersResponse{}
	return &this
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *ListingOffersResponse) GetHeaders() HttpResponseHeaders {
	if o == nil || IsNil(o.Headers) {
		var ret HttpResponseHeaders
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingOffersResponse) GetHeadersOk() (*HttpResponseHeaders, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *ListingOffersResponse) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given HttpResponseHeaders and assigns it to the Headers field.
func (o *ListingOffersResponse) SetHeaders(v HttpResponseHeaders) {
	o.Headers = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ListingOffersResponse) GetStatus() GetOffersHttpStatusLine {
	if o == nil || IsNil(o.Status) {
		var ret GetOffersHttpStatusLine
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingOffersResponse) GetStatusOk() (*GetOffersHttpStatusLine, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ListingOffersResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GetOffersHttpStatusLine and assigns it to the Status field.
func (o *ListingOffersResponse) SetStatus(v GetOffersHttpStatusLine) {
	o.Status = &v
}

// GetBody returns the Body field value
func (o *ListingOffersResponse) GetBody() GetOffersResponse {
	if o == nil {
		var ret GetOffersResponse
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *ListingOffersResponse) GetBodyOk() (*GetOffersResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *ListingOffersResponse) SetBody(v GetOffersResponse) {
	o.Body = v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *ListingOffersResponse) GetRequest() ListingOffersRequestParams {
	if o == nil || IsNil(o.Request) {
		var ret ListingOffersRequestParams
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingOffersResponse) GetRequestOk() (*ListingOffersRequestParams, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *ListingOffersResponse) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given ListingOffersRequestParams and assigns it to the Request field.
func (o *ListingOffersResponse) SetRequest(v ListingOffersRequestParams) {
	o.Request = &v
}

func (o ListingOffersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["body"] = o.Body
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	return toSerialize, nil
}

type NullableListingOffersResponse struct {
	value *ListingOffersResponse
	isSet bool
}

func (v NullableListingOffersResponse) Get() *ListingOffersResponse {
	return v.value
}

func (v *NullableListingOffersResponse) Set(val *ListingOffersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListingOffersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListingOffersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListingOffersResponse(val *ListingOffersResponse) *NullableListingOffersResponse {
	return &NullableListingOffersResponse{value: val, isSet: true}
}

func (v NullableListingOffersResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListingOffersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
