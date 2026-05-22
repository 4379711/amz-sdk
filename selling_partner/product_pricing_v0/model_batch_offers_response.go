package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the BatchOffersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchOffersResponse{}

// BatchOffersResponse Common schema that present in `ItemOffersResponse` and `ListingOffersResponse`
type BatchOffersResponse struct {
	Headers *HttpResponseHeaders     `json:"headers,omitempty"`
	Status  *GetOffersHttpStatusLine `json:"status,omitempty"`
	Body    GetOffersResponse        `json:"body"`
}

type _BatchOffersResponse BatchOffersResponse

// NewBatchOffersResponse instantiates a new BatchOffersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchOffersResponse(body GetOffersResponse) *BatchOffersResponse {
	this := BatchOffersResponse{}
	this.Body = body
	return &this
}

// NewBatchOffersResponseWithDefaults instantiates a new BatchOffersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchOffersResponseWithDefaults() *BatchOffersResponse {
	this := BatchOffersResponse{}
	return &this
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *BatchOffersResponse) GetHeaders() HttpResponseHeaders {
	if o == nil || IsNil(o.Headers) {
		var ret HttpResponseHeaders
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchOffersResponse) GetHeadersOk() (*HttpResponseHeaders, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *BatchOffersResponse) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given HttpResponseHeaders and assigns it to the Headers field.
func (o *BatchOffersResponse) SetHeaders(v HttpResponseHeaders) {
	o.Headers = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BatchOffersResponse) GetStatus() GetOffersHttpStatusLine {
	if o == nil || IsNil(o.Status) {
		var ret GetOffersHttpStatusLine
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchOffersResponse) GetStatusOk() (*GetOffersHttpStatusLine, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BatchOffersResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GetOffersHttpStatusLine and assigns it to the Status field.
func (o *BatchOffersResponse) SetStatus(v GetOffersHttpStatusLine) {
	o.Status = &v
}

// GetBody returns the Body field value
func (o *BatchOffersResponse) GetBody() GetOffersResponse {
	if o == nil {
		var ret GetOffersResponse
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *BatchOffersResponse) GetBodyOk() (*GetOffersResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *BatchOffersResponse) SetBody(v GetOffersResponse) {
	o.Body = v
}

func (o BatchOffersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["body"] = o.Body
	return toSerialize, nil
}

type NullableBatchOffersResponse struct {
	value *BatchOffersResponse
	isSet bool
}

func (v NullableBatchOffersResponse) Get() *BatchOffersResponse {
	return v.value
}

func (v *NullableBatchOffersResponse) Set(val *BatchOffersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchOffersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchOffersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchOffersResponse(val *BatchOffersResponse) *NullableBatchOffersResponse {
	return &NullableBatchOffersResponse{value: val, isSet: true}
}

func (v NullableBatchOffersResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBatchOffersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
