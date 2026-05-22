package vendor_direct_fulfillment_shipping_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the GetPackingSlipResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPackingSlipResponse{}

// GetPackingSlipResponse Response payload with packing slip.
type GetPackingSlipResponse struct {
	Payload *PackingSlip `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetPackingSlipResponse instantiates a new GetPackingSlipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPackingSlipResponse() *GetPackingSlipResponse {
	this := GetPackingSlipResponse{}
	return &this
}

// NewGetPackingSlipResponseWithDefaults instantiates a new GetPackingSlipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPackingSlipResponseWithDefaults() *GetPackingSlipResponse {
	this := GetPackingSlipResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetPackingSlipResponse) GetPayload() PackingSlip {
	if o == nil || IsNil(o.Payload) {
		var ret PackingSlip
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPackingSlipResponse) GetPayloadOk() (*PackingSlip, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetPackingSlipResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given PackingSlip and assigns it to the Payload field.
func (o *GetPackingSlipResponse) SetPayload(v PackingSlip) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetPackingSlipResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPackingSlipResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetPackingSlipResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetPackingSlipResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetPackingSlipResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetPackingSlipResponse struct {
	value *GetPackingSlipResponse
	isSet bool
}

func (v NullableGetPackingSlipResponse) Get() *GetPackingSlipResponse {
	return v.value
}

func (v *NullableGetPackingSlipResponse) Set(val *GetPackingSlipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPackingSlipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPackingSlipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPackingSlipResponse(val *GetPackingSlipResponse) *NullableGetPackingSlipResponse {
	return &NullableGetPackingSlipResponse{value: val, isSet: true}
}

func (v NullableGetPackingSlipResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetPackingSlipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
