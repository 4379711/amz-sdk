package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalBrands type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalBrands{}

// GlobalBrands List of Brands for the countries provided in the request.
type GlobalBrands struct {
	// Map containing Brands for countries passed in the request.
	CountryBrands *map[string][]Brand `json:"countryBrands,omitempty"`
	// List of errors encountered while processing the response.
	Errors []GlobalProductAttributeTargetingErrorModel `json:"errors,omitempty"`
}

// NewGlobalBrands instantiates a new GlobalBrands object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalBrands() *GlobalBrands {
	this := GlobalBrands{}
	return &this
}

// NewGlobalBrandsWithDefaults instantiates a new GlobalBrands object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalBrandsWithDefaults() *GlobalBrands {
	this := GlobalBrands{}
	return &this
}

// GetCountryBrands returns the CountryBrands field value if set, zero value otherwise.
func (o *GlobalBrands) GetCountryBrands() map[string][]Brand {
	if o == nil || IsNil(o.CountryBrands) {
		var ret map[string][]Brand
		return ret
	}
	return *o.CountryBrands
}

// GetCountryBrandsOk returns a tuple with the CountryBrands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalBrands) GetCountryBrandsOk() (*map[string][]Brand, bool) {
	if o == nil || IsNil(o.CountryBrands) {
		return nil, false
	}
	return o.CountryBrands, true
}

// HasCountryBrands returns a boolean if a field has been set.
func (o *GlobalBrands) HasCountryBrands() bool {
	if o != nil && !IsNil(o.CountryBrands) {
		return true
	}

	return false
}

// SetCountryBrands gets a reference to the given map[string][]Brand and assigns it to the CountryBrands field.
func (o *GlobalBrands) SetCountryBrands(v map[string][]Brand) {
	o.CountryBrands = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GlobalBrands) GetErrors() []GlobalProductAttributeTargetingErrorModel {
	if o == nil || IsNil(o.Errors) {
		var ret []GlobalProductAttributeTargetingErrorModel
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalBrands) GetErrorsOk() ([]GlobalProductAttributeTargetingErrorModel, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GlobalBrands) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []GlobalProductAttributeTargetingErrorModel and assigns it to the Errors field.
func (o *GlobalBrands) SetErrors(v []GlobalProductAttributeTargetingErrorModel) {
	o.Errors = v
}

func (o GlobalBrands) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryBrands) {
		toSerialize["countryBrands"] = o.CountryBrands
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGlobalBrands struct {
	value *GlobalBrands
	isSet bool
}

func (v NullableGlobalBrands) Get() *GlobalBrands {
	return v.value
}

func (v *NullableGlobalBrands) Set(val *GlobalBrands) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalBrands) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalBrands) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalBrands(val *GlobalBrands) *NullableGlobalBrands {
	return &NullableGlobalBrands{value: val, isSet: true}
}

func (v NullableGlobalBrands) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalBrands) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
