package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the Geocode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Geocode{}

// Geocode Defines the latitude and longitude of the access point.
type Geocode struct {
	// The latitude of access point.
	Latitude *string `json:"latitude,omitempty"`
	// The longitude of access point.
	Longitude *string `json:"longitude,omitempty"`
}

// NewGeocode instantiates a new Geocode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeocode() *Geocode {
	this := Geocode{}
	return &this
}

// NewGeocodeWithDefaults instantiates a new Geocode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeocodeWithDefaults() *Geocode {
	this := Geocode{}
	return &this
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *Geocode) GetLatitude() string {
	if o == nil || IsNil(o.Latitude) {
		var ret string
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Geocode) GetLatitudeOk() (*string, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *Geocode) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given string and assigns it to the Latitude field.
func (o *Geocode) SetLatitude(v string) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *Geocode) GetLongitude() string {
	if o == nil || IsNil(o.Longitude) {
		var ret string
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Geocode) GetLongitudeOk() (*string, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *Geocode) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given string and assigns it to the Longitude field.
func (o *Geocode) SetLongitude(v string) {
	o.Longitude = &v
}

func (o Geocode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	return toSerialize, nil
}

type NullableGeocode struct {
	value *Geocode
	isSet bool
}

func (v NullableGeocode) Get() *Geocode {
	return v.value
}

func (v *NullableGeocode) Set(val *Geocode) {
	v.value = val
	v.isSet = true
}

func (v NullableGeocode) IsSet() bool {
	return v.isSet
}

func (v *NullableGeocode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeocode(val *Geocode) *NullableGeocode {
	return &NullableGeocode{value: val, isSet: true}
}

func (v NullableGeocode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGeocode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
