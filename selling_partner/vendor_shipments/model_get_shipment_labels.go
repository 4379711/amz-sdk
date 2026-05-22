package vendor_shipments

import (
	"github.com/bytedance/sonic"
)

// checks if the GetShipmentLabels type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetShipmentLabels{}

// GetShipmentLabels The response schema for the GetShipmentLabels operation.
type GetShipmentLabels struct {
	Payload *TransportationLabels `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetShipmentLabels instantiates a new GetShipmentLabels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetShipmentLabels() *GetShipmentLabels {
	this := GetShipmentLabels{}
	return &this
}

// NewGetShipmentLabelsWithDefaults instantiates a new GetShipmentLabels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetShipmentLabelsWithDefaults() *GetShipmentLabels {
	this := GetShipmentLabels{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetShipmentLabels) GetPayload() TransportationLabels {
	if o == nil || IsNil(o.Payload) {
		var ret TransportationLabels
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentLabels) GetPayloadOk() (*TransportationLabels, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetShipmentLabels) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given TransportationLabels and assigns it to the Payload field.
func (o *GetShipmentLabels) SetPayload(v TransportationLabels) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetShipmentLabels) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentLabels) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetShipmentLabels) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetShipmentLabels) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetShipmentLabels) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetShipmentLabels struct {
	value *GetShipmentLabels
	isSet bool
}

func (v NullableGetShipmentLabels) Get() *GetShipmentLabels {
	return v.value
}

func (v *NullableGetShipmentLabels) Set(val *GetShipmentLabels) {
	v.value = val
	v.isSet = true
}

func (v NullableGetShipmentLabels) IsSet() bool {
	return v.isSet
}

func (v *NullableGetShipmentLabels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetShipmentLabels(val *GetShipmentLabels) *NullableGetShipmentLabels {
	return &NullableGetShipmentLabels{value: val, isSet: true}
}

func (v NullableGetShipmentLabels) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetShipmentLabels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
