package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the SampleLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SampleLocation{}

// SampleLocation Information about a location. It uses a postal code to identify the location.
type SampleLocation struct {
	PostalCode *PostalCode `json:"postalCode,omitempty"`
}

// NewSampleLocation instantiates a new SampleLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSampleLocation() *SampleLocation {
	this := SampleLocation{}
	return &this
}

// NewSampleLocationWithDefaults instantiates a new SampleLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSampleLocationWithDefaults() *SampleLocation {
	this := SampleLocation{}
	return &this
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *SampleLocation) GetPostalCode() PostalCode {
	if o == nil || IsNil(o.PostalCode) {
		var ret PostalCode
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SampleLocation) GetPostalCodeOk() (*PostalCode, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *SampleLocation) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given PostalCode and assigns it to the PostalCode field.
func (o *SampleLocation) SetPostalCode(v PostalCode) {
	o.PostalCode = &v
}

func (o SampleLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	return toSerialize, nil
}

type NullableSampleLocation struct {
	value *SampleLocation
	isSet bool
}

func (v NullableSampleLocation) Get() *SampleLocation {
	return v.value
}

func (v *NullableSampleLocation) Set(val *SampleLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableSampleLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableSampleLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSampleLocation(val *SampleLocation) *NullableSampleLocation {
	return &NullableSampleLocation{value: val, isSet: true}
}

func (v NullableSampleLocation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSampleLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
