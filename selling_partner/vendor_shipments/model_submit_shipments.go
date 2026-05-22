package vendor_shipments

import (
	"github.com/bytedance/sonic"
)

// checks if the SubmitShipments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitShipments{}

// SubmitShipments The request schema for the SubmitShipments operation.
type SubmitShipments struct {
	// A list of one or more shipments with underlying details.
	Shipments []Shipment `json:"shipments,omitempty"`
}

// NewSubmitShipments instantiates a new SubmitShipments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitShipments() *SubmitShipments {
	this := SubmitShipments{}
	return &this
}

// NewSubmitShipmentsWithDefaults instantiates a new SubmitShipments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitShipmentsWithDefaults() *SubmitShipments {
	this := SubmitShipments{}
	return &this
}

// GetShipments returns the Shipments field value if set, zero value otherwise.
func (o *SubmitShipments) GetShipments() []Shipment {
	if o == nil || IsNil(o.Shipments) {
		var ret []Shipment
		return ret
	}
	return o.Shipments
}

// GetShipmentsOk returns a tuple with the Shipments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitShipments) GetShipmentsOk() ([]Shipment, bool) {
	if o == nil || IsNil(o.Shipments) {
		return nil, false
	}
	return o.Shipments, true
}

// HasShipments returns a boolean if a field has been set.
func (o *SubmitShipments) HasShipments() bool {
	if o != nil && !IsNil(o.Shipments) {
		return true
	}

	return false
}

// SetShipments gets a reference to the given []Shipment and assigns it to the Shipments field.
func (o *SubmitShipments) SetShipments(v []Shipment) {
	o.Shipments = v
}

func (o SubmitShipments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Shipments) {
		toSerialize["shipments"] = o.Shipments
	}
	return toSerialize, nil
}

type NullableSubmitShipments struct {
	value *SubmitShipments
	isSet bool
}

func (v NullableSubmitShipments) Get() *SubmitShipments {
	return v.value
}

func (v *NullableSubmitShipments) Set(val *SubmitShipments) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitShipments) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitShipments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitShipments(val *SubmitShipments) *NullableSubmitShipments {
	return &NullableSubmitShipments{value: val, isSet: true}
}

func (v NullableSubmitShipments) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubmitShipments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
