package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the GetShipmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetShipmentResponse{}

// GetShipmentResponse The response schema for the getShipment operation.
type GetShipmentResponse struct {
	Payload *Shipment `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetShipmentResponse instantiates a new GetShipmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetShipmentResponse() *GetShipmentResponse {
	this := GetShipmentResponse{}
	return &this
}

// NewGetShipmentResponseWithDefaults instantiates a new GetShipmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetShipmentResponseWithDefaults() *GetShipmentResponse {
	this := GetShipmentResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetShipmentResponse) GetPayload() Shipment {
	if o == nil || IsNil(o.Payload) {
		var ret Shipment
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentResponse) GetPayloadOk() (*Shipment, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetShipmentResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given Shipment and assigns it to the Payload field.
func (o *GetShipmentResponse) SetPayload(v Shipment) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetShipmentResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetShipmentResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetShipmentResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetShipmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetShipmentResponse struct {
	value *GetShipmentResponse
	isSet bool
}

func (v NullableGetShipmentResponse) Get() *GetShipmentResponse {
	return v.value
}

func (v *NullableGetShipmentResponse) Set(val *GetShipmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetShipmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetShipmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetShipmentResponse(val *GetShipmentResponse) *NullableGetShipmentResponse {
	return &NullableGetShipmentResponse{value: val, isSet: true}
}

func (v NullableGetShipmentResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetShipmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
