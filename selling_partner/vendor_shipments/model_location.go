package vendor_shipments

import (
	"github.com/bytedance/sonic"
)

// checks if the Location type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Location{}

// Location Location identifier.
type Location struct {
	// Type of location identification.
	Type *string `json:"type,omitempty"`
	// Location code.
	LocationCode *string `json:"locationCode,omitempty"`
	// The two digit country code. In ISO 3166-1 alpha-2 format.
	CountryCode *string `json:"countryCode,omitempty"`
}

// NewLocation instantiates a new Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation() *Location {
	this := Location{}
	return &this
}

// NewLocationWithDefaults instantiates a new Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationWithDefaults() *Location {
	this := Location{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Location) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Location) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Location) SetType(v string) {
	o.Type = &v
}

// GetLocationCode returns the LocationCode field value if set, zero value otherwise.
func (o *Location) GetLocationCode() string {
	if o == nil || IsNil(o.LocationCode) {
		var ret string
		return ret
	}
	return *o.LocationCode
}

// GetLocationCodeOk returns a tuple with the LocationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetLocationCodeOk() (*string, bool) {
	if o == nil || IsNil(o.LocationCode) {
		return nil, false
	}
	return o.LocationCode, true
}

// HasLocationCode returns a boolean if a field has been set.
func (o *Location) HasLocationCode() bool {
	if o != nil && !IsNil(o.LocationCode) {
		return true
	}

	return false
}

// SetLocationCode gets a reference to the given string and assigns it to the LocationCode field.
func (o *Location) SetLocationCode(v string) {
	o.LocationCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *Location) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *Location) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *Location) SetCountryCode(v string) {
	o.CountryCode = &v
}

func (o Location) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.LocationCode) {
		toSerialize["locationCode"] = o.LocationCode
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	return toSerialize, nil
}

type NullableLocation struct {
	value *Location
	isSet bool
}

func (v NullableLocation) Get() *Location {
	return v.value
}

func (v *NullableLocation) Set(val *Location) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation(val *Location) *NullableLocation {
	return &NullableLocation{value: val, isSet: true}
}

func (v NullableLocation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
