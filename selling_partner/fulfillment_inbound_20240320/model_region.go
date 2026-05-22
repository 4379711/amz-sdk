package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the Region type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Region{}

// Region Representation of a location used within the inbounding experience.
type Region struct {
	// ISO 3166 standard alpha-2 country code.
	CountryCode *string `json:"countryCode,omitempty"`
	// State.
	State *string `json:"state,omitempty"`
	// An identifier for a warehouse, such as a FC, IXD, upstream storage.
	WarehouseId *string `json:"warehouseId,omitempty"`
}

// NewRegion instantiates a new Region object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegion() *Region {
	this := Region{}
	return &this
}

// NewRegionWithDefaults instantiates a new Region object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionWithDefaults() *Region {
	this := Region{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *Region) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *Region) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *Region) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Region) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Region) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Region) SetState(v string) {
	o.State = &v
}

// GetWarehouseId returns the WarehouseId field value if set, zero value otherwise.
func (o *Region) GetWarehouseId() string {
	if o == nil || IsNil(o.WarehouseId) {
		var ret string
		return ret
	}
	return *o.WarehouseId
}

// GetWarehouseIdOk returns a tuple with the WarehouseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetWarehouseIdOk() (*string, bool) {
	if o == nil || IsNil(o.WarehouseId) {
		return nil, false
	}
	return o.WarehouseId, true
}

// HasWarehouseId returns a boolean if a field has been set.
func (o *Region) HasWarehouseId() bool {
	if o != nil && !IsNil(o.WarehouseId) {
		return true
	}

	return false
}

// SetWarehouseId gets a reference to the given string and assigns it to the WarehouseId field.
func (o *Region) SetWarehouseId(v string) {
	o.WarehouseId = &v
}

func (o Region) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.WarehouseId) {
		toSerialize["warehouseId"] = o.WarehouseId
	}
	return toSerialize, nil
}

type NullableRegion struct {
	value *Region
	isSet bool
}

func (v NullableRegion) Get() *Region {
	return v.value
}

func (v *NullableRegion) Set(val *Region) {
	v.value = val
	v.isSet = true
}

func (v NullableRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegion(val *Region) *NullableRegion {
	return &NullableRegion{value: val, isSet: true}
}

func (v NullableRegion) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
