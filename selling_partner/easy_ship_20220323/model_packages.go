package easy_ship_20220323

import (
	"github.com/bytedance/sonic"
)

// checks if the Packages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Packages{}

// Packages A list of packages.
type Packages struct {
	// A list of packages.
	Packages []Package `json:"packages"`
}

type _Packages Packages

// NewPackages instantiates a new Packages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackages(packages []Package) *Packages {
	this := Packages{}
	this.Packages = packages
	return &this
}

// NewPackagesWithDefaults instantiates a new Packages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackagesWithDefaults() *Packages {
	this := Packages{}
	return &this
}

// GetPackages returns the Packages field value
func (o *Packages) GetPackages() []Package {
	if o == nil {
		var ret []Package
		return ret
	}

	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value
// and a boolean to check if the value has been set.
func (o *Packages) GetPackagesOk() ([]Package, bool) {
	if o == nil {
		return nil, false
	}
	return o.Packages, true
}

// SetPackages sets field value
func (o *Packages) SetPackages(v []Package) {
	o.Packages = v
}

func (o Packages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["packages"] = o.Packages
	return toSerialize, nil
}

type NullablePackages struct {
	value *Packages
	isSet bool
}

func (v NullablePackages) Get() *Packages {
	return v.value
}

func (v *NullablePackages) Set(val *Packages) {
	v.value = val
	v.isSet = true
}

func (v NullablePackages) IsSet() bool {
	return v.isSet
}

func (v *NullablePackages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackages(val *Packages) *NullablePackages {
	return &NullablePackages{value: val, isSet: true}
}

func (v NullablePackages) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
