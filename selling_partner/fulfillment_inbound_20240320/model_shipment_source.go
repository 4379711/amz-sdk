package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ShipmentSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentSource{}

// ShipmentSource Specifies the 'ship from' address for the shipment.
type ShipmentSource struct {
	Address *Address `json:"address,omitempty"`
	// The type of source for this shipment. Possible values: `SELLER_FACILITY`.
	SourceType string `json:"sourceType"`
}

type _ShipmentSource ShipmentSource

// NewShipmentSource instantiates a new ShipmentSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentSource(sourceType string) *ShipmentSource {
	this := ShipmentSource{}
	this.SourceType = sourceType
	return &this
}

// NewShipmentSourceWithDefaults instantiates a new ShipmentSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentSourceWithDefaults() *ShipmentSource {
	this := ShipmentSource{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ShipmentSource) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentSource) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ShipmentSource) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *ShipmentSource) SetAddress(v Address) {
	o.Address = &v
}

// GetSourceType returns the SourceType field value
func (o *ShipmentSource) GetSourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *ShipmentSource) GetSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *ShipmentSource) SetSourceType(v string) {
	o.SourceType = v
}

func (o ShipmentSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	toSerialize["sourceType"] = o.SourceType
	return toSerialize, nil
}

type NullableShipmentSource struct {
	value *ShipmentSource
	isSet bool
}

func (v NullableShipmentSource) Get() *ShipmentSource {
	return v.value
}

func (v *NullableShipmentSource) Set(val *ShipmentSource) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentSource) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentSource(val *ShipmentSource) *NullableShipmentSource {
	return &NullableShipmentSource{value: val, isSet: true}
}

func (v NullableShipmentSource) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShipmentSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
