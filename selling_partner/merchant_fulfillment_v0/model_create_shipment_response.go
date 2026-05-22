package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateShipmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateShipmentResponse{}

// CreateShipmentResponse Response schema.
type CreateShipmentResponse struct {
	Payload *Shipment `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCreateShipmentResponse instantiates a new CreateShipmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShipmentResponse() *CreateShipmentResponse {
	this := CreateShipmentResponse{}
	return &this
}

// NewCreateShipmentResponseWithDefaults instantiates a new CreateShipmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShipmentResponseWithDefaults() *CreateShipmentResponse {
	this := CreateShipmentResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *CreateShipmentResponse) GetPayload() Shipment {
	if o == nil || IsNil(o.Payload) {
		var ret Shipment
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShipmentResponse) GetPayloadOk() (*Shipment, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *CreateShipmentResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given Shipment and assigns it to the Payload field.
func (o *CreateShipmentResponse) SetPayload(v Shipment) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateShipmentResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShipmentResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateShipmentResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CreateShipmentResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CreateShipmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCreateShipmentResponse struct {
	value *CreateShipmentResponse
	isSet bool
}

func (v NullableCreateShipmentResponse) Get() *CreateShipmentResponse {
	return v.value
}

func (v *NullableCreateShipmentResponse) Set(val *CreateShipmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShipmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShipmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShipmentResponse(val *CreateShipmentResponse) *NullableCreateShipmentResponse {
	return &NullableCreateShipmentResponse{value: val, isSet: true}
}

func (v NullableCreateShipmentResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateShipmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
