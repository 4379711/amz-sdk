package vendor_direct_fulfillment_shipping_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SubmitShipmentStatusUpdatesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitShipmentStatusUpdatesResponse{}

// SubmitShipmentStatusUpdatesResponse The response schema for the submitShipmentStatusUpdates operation.
type SubmitShipmentStatusUpdatesResponse struct {
	Payload *TransactionReference `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewSubmitShipmentStatusUpdatesResponse instantiates a new SubmitShipmentStatusUpdatesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitShipmentStatusUpdatesResponse() *SubmitShipmentStatusUpdatesResponse {
	this := SubmitShipmentStatusUpdatesResponse{}
	return &this
}

// NewSubmitShipmentStatusUpdatesResponseWithDefaults instantiates a new SubmitShipmentStatusUpdatesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitShipmentStatusUpdatesResponseWithDefaults() *SubmitShipmentStatusUpdatesResponse {
	this := SubmitShipmentStatusUpdatesResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *SubmitShipmentStatusUpdatesResponse) GetPayload() TransactionReference {
	if o == nil || IsNil(o.Payload) {
		var ret TransactionReference
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitShipmentStatusUpdatesResponse) GetPayloadOk() (*TransactionReference, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *SubmitShipmentStatusUpdatesResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given TransactionReference and assigns it to the Payload field.
func (o *SubmitShipmentStatusUpdatesResponse) SetPayload(v TransactionReference) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SubmitShipmentStatusUpdatesResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitShipmentStatusUpdatesResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SubmitShipmentStatusUpdatesResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *SubmitShipmentStatusUpdatesResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o SubmitShipmentStatusUpdatesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSubmitShipmentStatusUpdatesResponse struct {
	value *SubmitShipmentStatusUpdatesResponse
	isSet bool
}

func (v NullableSubmitShipmentStatusUpdatesResponse) Get() *SubmitShipmentStatusUpdatesResponse {
	return v.value
}

func (v *NullableSubmitShipmentStatusUpdatesResponse) Set(val *SubmitShipmentStatusUpdatesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitShipmentStatusUpdatesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitShipmentStatusUpdatesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitShipmentStatusUpdatesResponse(val *SubmitShipmentStatusUpdatesResponse) *NullableSubmitShipmentStatusUpdatesResponse {
	return &NullableSubmitShipmentStatusUpdatesResponse{value: val, isSet: true}
}

func (v NullableSubmitShipmentStatusUpdatesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubmitShipmentStatusUpdatesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
