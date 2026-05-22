package vendor_direct_fulfillment_shipping_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SubmitShipmentConfirmationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitShipmentConfirmationsResponse{}

// SubmitShipmentConfirmationsResponse The response schema for the submitShipmentConfirmations operation.
type SubmitShipmentConfirmationsResponse struct {
	Payload *TransactionReference `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewSubmitShipmentConfirmationsResponse instantiates a new SubmitShipmentConfirmationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitShipmentConfirmationsResponse() *SubmitShipmentConfirmationsResponse {
	this := SubmitShipmentConfirmationsResponse{}
	return &this
}

// NewSubmitShipmentConfirmationsResponseWithDefaults instantiates a new SubmitShipmentConfirmationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitShipmentConfirmationsResponseWithDefaults() *SubmitShipmentConfirmationsResponse {
	this := SubmitShipmentConfirmationsResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *SubmitShipmentConfirmationsResponse) GetPayload() TransactionReference {
	if o == nil || IsNil(o.Payload) {
		var ret TransactionReference
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitShipmentConfirmationsResponse) GetPayloadOk() (*TransactionReference, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *SubmitShipmentConfirmationsResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given TransactionReference and assigns it to the Payload field.
func (o *SubmitShipmentConfirmationsResponse) SetPayload(v TransactionReference) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SubmitShipmentConfirmationsResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitShipmentConfirmationsResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SubmitShipmentConfirmationsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *SubmitShipmentConfirmationsResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o SubmitShipmentConfirmationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSubmitShipmentConfirmationsResponse struct {
	value *SubmitShipmentConfirmationsResponse
	isSet bool
}

func (v NullableSubmitShipmentConfirmationsResponse) Get() *SubmitShipmentConfirmationsResponse {
	return v.value
}

func (v *NullableSubmitShipmentConfirmationsResponse) Set(val *SubmitShipmentConfirmationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitShipmentConfirmationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitShipmentConfirmationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitShipmentConfirmationsResponse(val *SubmitShipmentConfirmationsResponse) *NullableSubmitShipmentConfirmationsResponse {
	return &NullableSubmitShipmentConfirmationsResponse{value: val, isSet: true}
}

func (v NullableSubmitShipmentConfirmationsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubmitShipmentConfirmationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
