package vendor_direct_fulfillment_shipping_20211228

import (
	"github.com/bytedance/sonic"
)

// checks if the SubmitShipmentConfirmationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitShipmentConfirmationsRequest{}

// SubmitShipmentConfirmationsRequest The request schema for the submitShipmentConfirmations operation.
type SubmitShipmentConfirmationsRequest struct {
	// Array of `ShipmentConfirmation` objects. Each `ShipmentConfirmation` object represents the confirmation details for a specific shipment.
	ShipmentConfirmations []ShipmentConfirmation `json:"shipmentConfirmations,omitempty"`
}

// NewSubmitShipmentConfirmationsRequest instantiates a new SubmitShipmentConfirmationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitShipmentConfirmationsRequest() *SubmitShipmentConfirmationsRequest {
	this := SubmitShipmentConfirmationsRequest{}
	return &this
}

// NewSubmitShipmentConfirmationsRequestWithDefaults instantiates a new SubmitShipmentConfirmationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitShipmentConfirmationsRequestWithDefaults() *SubmitShipmentConfirmationsRequest {
	this := SubmitShipmentConfirmationsRequest{}
	return &this
}

// GetShipmentConfirmations returns the ShipmentConfirmations field value if set, zero value otherwise.
func (o *SubmitShipmentConfirmationsRequest) GetShipmentConfirmations() []ShipmentConfirmation {
	if o == nil || IsNil(o.ShipmentConfirmations) {
		var ret []ShipmentConfirmation
		return ret
	}
	return o.ShipmentConfirmations
}

// GetShipmentConfirmationsOk returns a tuple with the ShipmentConfirmations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitShipmentConfirmationsRequest) GetShipmentConfirmationsOk() ([]ShipmentConfirmation, bool) {
	if o == nil || IsNil(o.ShipmentConfirmations) {
		return nil, false
	}
	return o.ShipmentConfirmations, true
}

// HasShipmentConfirmations returns a boolean if a field has been set.
func (o *SubmitShipmentConfirmationsRequest) HasShipmentConfirmations() bool {
	if o != nil && !IsNil(o.ShipmentConfirmations) {
		return true
	}

	return false
}

// SetShipmentConfirmations gets a reference to the given []ShipmentConfirmation and assigns it to the ShipmentConfirmations field.
func (o *SubmitShipmentConfirmationsRequest) SetShipmentConfirmations(v []ShipmentConfirmation) {
	o.ShipmentConfirmations = v
}

func (o SubmitShipmentConfirmationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShipmentConfirmations) {
		toSerialize["shipmentConfirmations"] = o.ShipmentConfirmations
	}
	return toSerialize, nil
}

type NullableSubmitShipmentConfirmationsRequest struct {
	value *SubmitShipmentConfirmationsRequest
	isSet bool
}

func (v NullableSubmitShipmentConfirmationsRequest) Get() *SubmitShipmentConfirmationsRequest {
	return v.value
}

func (v *NullableSubmitShipmentConfirmationsRequest) Set(val *SubmitShipmentConfirmationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitShipmentConfirmationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitShipmentConfirmationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitShipmentConfirmationsRequest(val *SubmitShipmentConfirmationsRequest) *NullableSubmitShipmentConfirmationsRequest {
	return &NullableSubmitShipmentConfirmationsRequest{value: val, isSet: true}
}

func (v NullableSubmitShipmentConfirmationsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubmitShipmentConfirmationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
