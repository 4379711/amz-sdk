package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the Location type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Location{}

// Location The location where the person, business or institution is located.
type Location struct {
	// The state or region where the person, business or institution is located.
	StateOrRegion *string `json:"stateOrRegion,omitempty"`
	// The city where the person, business or institution is located.
	City *string `json:"city,omitempty"`
	// The two digit country code. In ISO 3166-1 alpha-2 format.
	CountryCode *string `json:"countryCode,omitempty"`
	// The postal code of that address. It contains a series of letters or digits or both, sometimes including spaces or punctuation.
	PostalCode *string `json:"postalCode,omitempty"`
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

// GetStateOrRegion returns the StateOrRegion field value if set, zero value otherwise.
func (o *Location) GetStateOrRegion() string {
	if o == nil || IsNil(o.StateOrRegion) {
		var ret string
		return ret
	}
	return *o.StateOrRegion
}

// GetStateOrRegionOk returns a tuple with the StateOrRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetStateOrRegionOk() (*string, bool) {
	if o == nil || IsNil(o.StateOrRegion) {
		return nil, false
	}
	return o.StateOrRegion, true
}

// HasStateOrRegion returns a boolean if a field has been set.
func (o *Location) HasStateOrRegion() bool {
	if o != nil && !IsNil(o.StateOrRegion) {
		return true
	}

	return false
}

// SetStateOrRegion gets a reference to the given string and assigns it to the StateOrRegion field.
func (o *Location) SetStateOrRegion(v string) {
	o.StateOrRegion = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *Location) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *Location) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *Location) SetCity(v string) {
	o.City = &v
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

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *Location) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *Location) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *Location) SetPostalCode(v string) {
	o.PostalCode = &v
}

func (o Location) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StateOrRegion) {
		toSerialize["stateOrRegion"] = o.StateOrRegion
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
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
