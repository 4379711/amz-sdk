package vendor_shipments

import (
	"github.com/bytedance/sonic"
)

// checks if the TransportShipmentMeasurements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransportShipmentMeasurements{}

// TransportShipmentMeasurements Shipment measurement details.
type TransportShipmentMeasurements struct {
	// Total number of cartons present in the shipment. Provide the cartonCount only for non-palletized shipments.
	TotalCartonCount *int32 `json:"totalCartonCount,omitempty"`
	// Total number of Stackable Pallets present in the shipment.
	TotalPalletStackable *int32 `json:"totalPalletStackable,omitempty"`
	// Total number of Non Stackable Pallets present in the shipment.
	TotalPalletNonStackable *int32  `json:"totalPalletNonStackable,omitempty"`
	ShipmentWeight          *Weight `json:"shipmentWeight,omitempty"`
	ShipmentVolume          *Volume `json:"shipmentVolume,omitempty"`
}

// NewTransportShipmentMeasurements instantiates a new TransportShipmentMeasurements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportShipmentMeasurements() *TransportShipmentMeasurements {
	this := TransportShipmentMeasurements{}
	return &this
}

// NewTransportShipmentMeasurementsWithDefaults instantiates a new TransportShipmentMeasurements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportShipmentMeasurementsWithDefaults() *TransportShipmentMeasurements {
	this := TransportShipmentMeasurements{}
	return &this
}

// GetTotalCartonCount returns the TotalCartonCount field value if set, zero value otherwise.
func (o *TransportShipmentMeasurements) GetTotalCartonCount() int32 {
	if o == nil || IsNil(o.TotalCartonCount) {
		var ret int32
		return ret
	}
	return *o.TotalCartonCount
}

// GetTotalCartonCountOk returns a tuple with the TotalCartonCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportShipmentMeasurements) GetTotalCartonCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCartonCount) {
		return nil, false
	}
	return o.TotalCartonCount, true
}

// HasTotalCartonCount returns a boolean if a field has been set.
func (o *TransportShipmentMeasurements) HasTotalCartonCount() bool {
	if o != nil && !IsNil(o.TotalCartonCount) {
		return true
	}

	return false
}

// SetTotalCartonCount gets a reference to the given int32 and assigns it to the TotalCartonCount field.
func (o *TransportShipmentMeasurements) SetTotalCartonCount(v int32) {
	o.TotalCartonCount = &v
}

// GetTotalPalletStackable returns the TotalPalletStackable field value if set, zero value otherwise.
func (o *TransportShipmentMeasurements) GetTotalPalletStackable() int32 {
	if o == nil || IsNil(o.TotalPalletStackable) {
		var ret int32
		return ret
	}
	return *o.TotalPalletStackable
}

// GetTotalPalletStackableOk returns a tuple with the TotalPalletStackable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportShipmentMeasurements) GetTotalPalletStackableOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPalletStackable) {
		return nil, false
	}
	return o.TotalPalletStackable, true
}

// HasTotalPalletStackable returns a boolean if a field has been set.
func (o *TransportShipmentMeasurements) HasTotalPalletStackable() bool {
	if o != nil && !IsNil(o.TotalPalletStackable) {
		return true
	}

	return false
}

// SetTotalPalletStackable gets a reference to the given int32 and assigns it to the TotalPalletStackable field.
func (o *TransportShipmentMeasurements) SetTotalPalletStackable(v int32) {
	o.TotalPalletStackable = &v
}

// GetTotalPalletNonStackable returns the TotalPalletNonStackable field value if set, zero value otherwise.
func (o *TransportShipmentMeasurements) GetTotalPalletNonStackable() int32 {
	if o == nil || IsNil(o.TotalPalletNonStackable) {
		var ret int32
		return ret
	}
	return *o.TotalPalletNonStackable
}

// GetTotalPalletNonStackableOk returns a tuple with the TotalPalletNonStackable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportShipmentMeasurements) GetTotalPalletNonStackableOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPalletNonStackable) {
		return nil, false
	}
	return o.TotalPalletNonStackable, true
}

// HasTotalPalletNonStackable returns a boolean if a field has been set.
func (o *TransportShipmentMeasurements) HasTotalPalletNonStackable() bool {
	if o != nil && !IsNil(o.TotalPalletNonStackable) {
		return true
	}

	return false
}

// SetTotalPalletNonStackable gets a reference to the given int32 and assigns it to the TotalPalletNonStackable field.
func (o *TransportShipmentMeasurements) SetTotalPalletNonStackable(v int32) {
	o.TotalPalletNonStackable = &v
}

// GetShipmentWeight returns the ShipmentWeight field value if set, zero value otherwise.
func (o *TransportShipmentMeasurements) GetShipmentWeight() Weight {
	if o == nil || IsNil(o.ShipmentWeight) {
		var ret Weight
		return ret
	}
	return *o.ShipmentWeight
}

// GetShipmentWeightOk returns a tuple with the ShipmentWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportShipmentMeasurements) GetShipmentWeightOk() (*Weight, bool) {
	if o == nil || IsNil(o.ShipmentWeight) {
		return nil, false
	}
	return o.ShipmentWeight, true
}

// HasShipmentWeight returns a boolean if a field has been set.
func (o *TransportShipmentMeasurements) HasShipmentWeight() bool {
	if o != nil && !IsNil(o.ShipmentWeight) {
		return true
	}

	return false
}

// SetShipmentWeight gets a reference to the given Weight and assigns it to the ShipmentWeight field.
func (o *TransportShipmentMeasurements) SetShipmentWeight(v Weight) {
	o.ShipmentWeight = &v
}

// GetShipmentVolume returns the ShipmentVolume field value if set, zero value otherwise.
func (o *TransportShipmentMeasurements) GetShipmentVolume() Volume {
	if o == nil || IsNil(o.ShipmentVolume) {
		var ret Volume
		return ret
	}
	return *o.ShipmentVolume
}

// GetShipmentVolumeOk returns a tuple with the ShipmentVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportShipmentMeasurements) GetShipmentVolumeOk() (*Volume, bool) {
	if o == nil || IsNil(o.ShipmentVolume) {
		return nil, false
	}
	return o.ShipmentVolume, true
}

// HasShipmentVolume returns a boolean if a field has been set.
func (o *TransportShipmentMeasurements) HasShipmentVolume() bool {
	if o != nil && !IsNil(o.ShipmentVolume) {
		return true
	}

	return false
}

// SetShipmentVolume gets a reference to the given Volume and assigns it to the ShipmentVolume field.
func (o *TransportShipmentMeasurements) SetShipmentVolume(v Volume) {
	o.ShipmentVolume = &v
}

func (o TransportShipmentMeasurements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCartonCount) {
		toSerialize["totalCartonCount"] = o.TotalCartonCount
	}
	if !IsNil(o.TotalPalletStackable) {
		toSerialize["totalPalletStackable"] = o.TotalPalletStackable
	}
	if !IsNil(o.TotalPalletNonStackable) {
		toSerialize["totalPalletNonStackable"] = o.TotalPalletNonStackable
	}
	if !IsNil(o.ShipmentWeight) {
		toSerialize["shipmentWeight"] = o.ShipmentWeight
	}
	if !IsNil(o.ShipmentVolume) {
		toSerialize["shipmentVolume"] = o.ShipmentVolume
	}
	return toSerialize, nil
}

type NullableTransportShipmentMeasurements struct {
	value *TransportShipmentMeasurements
	isSet bool
}

func (v NullableTransportShipmentMeasurements) Get() *TransportShipmentMeasurements {
	return v.value
}

func (v *NullableTransportShipmentMeasurements) Set(val *TransportShipmentMeasurements) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportShipmentMeasurements) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportShipmentMeasurements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportShipmentMeasurements(val *TransportShipmentMeasurements) *NullableTransportShipmentMeasurements {
	return &NullableTransportShipmentMeasurements{value: val, isSet: true}
}

func (v NullableTransportShipmentMeasurements) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTransportShipmentMeasurements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
