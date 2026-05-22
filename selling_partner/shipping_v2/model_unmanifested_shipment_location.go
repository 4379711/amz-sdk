package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the UnmanifestedShipmentLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnmanifestedShipmentLocation{}

// UnmanifestedShipmentLocation UnmanifestedShipmentLocation info
type UnmanifestedShipmentLocation struct {
	Address *Address `json:"address,omitempty"`
	// Its Last Manifest Date.
	LastManifestDate *string `json:"lastManifestDate,omitempty"`
}

// NewUnmanifestedShipmentLocation instantiates a new UnmanifestedShipmentLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnmanifestedShipmentLocation() *UnmanifestedShipmentLocation {
	this := UnmanifestedShipmentLocation{}
	return &this
}

// NewUnmanifestedShipmentLocationWithDefaults instantiates a new UnmanifestedShipmentLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnmanifestedShipmentLocationWithDefaults() *UnmanifestedShipmentLocation {
	this := UnmanifestedShipmentLocation{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UnmanifestedShipmentLocation) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnmanifestedShipmentLocation) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UnmanifestedShipmentLocation) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *UnmanifestedShipmentLocation) SetAddress(v Address) {
	o.Address = &v
}

// GetLastManifestDate returns the LastManifestDate field value if set, zero value otherwise.
func (o *UnmanifestedShipmentLocation) GetLastManifestDate() string {
	if o == nil || IsNil(o.LastManifestDate) {
		var ret string
		return ret
	}
	return *o.LastManifestDate
}

// GetLastManifestDateOk returns a tuple with the LastManifestDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnmanifestedShipmentLocation) GetLastManifestDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastManifestDate) {
		return nil, false
	}
	return o.LastManifestDate, true
}

// HasLastManifestDate returns a boolean if a field has been set.
func (o *UnmanifestedShipmentLocation) HasLastManifestDate() bool {
	if o != nil && !IsNil(o.LastManifestDate) {
		return true
	}

	return false
}

// SetLastManifestDate gets a reference to the given string and assigns it to the LastManifestDate field.
func (o *UnmanifestedShipmentLocation) SetLastManifestDate(v string) {
	o.LastManifestDate = &v
}

func (o UnmanifestedShipmentLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.LastManifestDate) {
		toSerialize["lastManifestDate"] = o.LastManifestDate
	}
	return toSerialize, nil
}

type NullableUnmanifestedShipmentLocation struct {
	value *UnmanifestedShipmentLocation
	isSet bool
}

func (v NullableUnmanifestedShipmentLocation) Get() *UnmanifestedShipmentLocation {
	return v.value
}

func (v *NullableUnmanifestedShipmentLocation) Set(val *UnmanifestedShipmentLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableUnmanifestedShipmentLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableUnmanifestedShipmentLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnmanifestedShipmentLocation(val *UnmanifestedShipmentLocation) *NullableUnmanifestedShipmentLocation {
	return &NullableUnmanifestedShipmentLocation{value: val, isSet: true}
}

func (v NullableUnmanifestedShipmentLocation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUnmanifestedShipmentLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
