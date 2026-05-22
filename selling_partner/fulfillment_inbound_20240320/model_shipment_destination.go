package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ShipmentDestination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentDestination{}

// ShipmentDestination The Amazon fulfillment center address and warehouse ID.
type ShipmentDestination struct {
	Address *Address `json:"address,omitempty"`
	// The type of destination for this shipment. Possible values: `AMAZON_OPTIMIZED`, `AMAZON_WAREHOUSE`.
	DestinationType string `json:"destinationType"`
	// The warehouse that the shipment should be sent to. Empty if the destination type is `AMAZON_OPTIMIZED`.
	WarehouseId *string `json:"warehouseId,omitempty"`
}

type _ShipmentDestination ShipmentDestination

// NewShipmentDestination instantiates a new ShipmentDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentDestination(destinationType string) *ShipmentDestination {
	this := ShipmentDestination{}
	this.DestinationType = destinationType
	return &this
}

// NewShipmentDestinationWithDefaults instantiates a new ShipmentDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentDestinationWithDefaults() *ShipmentDestination {
	this := ShipmentDestination{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ShipmentDestination) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDestination) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ShipmentDestination) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *ShipmentDestination) SetAddress(v Address) {
	o.Address = &v
}

// GetDestinationType returns the DestinationType field value
func (o *ShipmentDestination) GetDestinationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationType
}

// GetDestinationTypeOk returns a tuple with the DestinationType field value
// and a boolean to check if the value has been set.
func (o *ShipmentDestination) GetDestinationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationType, true
}

// SetDestinationType sets field value
func (o *ShipmentDestination) SetDestinationType(v string) {
	o.DestinationType = v
}

// GetWarehouseId returns the WarehouseId field value if set, zero value otherwise.
func (o *ShipmentDestination) GetWarehouseId() string {
	if o == nil || IsNil(o.WarehouseId) {
		var ret string
		return ret
	}
	return *o.WarehouseId
}

// GetWarehouseIdOk returns a tuple with the WarehouseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDestination) GetWarehouseIdOk() (*string, bool) {
	if o == nil || IsNil(o.WarehouseId) {
		return nil, false
	}
	return o.WarehouseId, true
}

// HasWarehouseId returns a boolean if a field has been set.
func (o *ShipmentDestination) HasWarehouseId() bool {
	if o != nil && !IsNil(o.WarehouseId) {
		return true
	}

	return false
}

// SetWarehouseId gets a reference to the given string and assigns it to the WarehouseId field.
func (o *ShipmentDestination) SetWarehouseId(v string) {
	o.WarehouseId = &v
}

func (o ShipmentDestination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	toSerialize["destinationType"] = o.DestinationType
	if !IsNil(o.WarehouseId) {
		toSerialize["warehouseId"] = o.WarehouseId
	}
	return toSerialize, nil
}

type NullableShipmentDestination struct {
	value *ShipmentDestination
	isSet bool
}

func (v NullableShipmentDestination) Get() *ShipmentDestination {
	return v.value
}

func (v *NullableShipmentDestination) Set(val *ShipmentDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentDestination(val *ShipmentDestination) *NullableShipmentDestination {
	return &NullableShipmentDestination{value: val, isSet: true}
}

func (v NullableShipmentDestination) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShipmentDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
