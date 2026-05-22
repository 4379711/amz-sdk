package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the PackageVolume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageVolume{}

// PackageVolume Represents the volume of the package with a unit of measurement.
type PackageVolume struct {
	UnitOfMeasurement VolumeUnitOfMeasurement `json:"unitOfMeasurement"`
	// The package volume value.
	Volume float64 `json:"volume"`
}

type _PackageVolume PackageVolume

// NewPackageVolume instantiates a new PackageVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageVolume(unitOfMeasurement VolumeUnitOfMeasurement, volume float64) *PackageVolume {
	this := PackageVolume{}
	this.UnitOfMeasurement = unitOfMeasurement
	this.Volume = volume
	return &this
}

// NewPackageVolumeWithDefaults instantiates a new PackageVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageVolumeWithDefaults() *PackageVolume {
	this := PackageVolume{}
	return &this
}

// GetUnitOfMeasurement returns the UnitOfMeasurement field value
func (o *PackageVolume) GetUnitOfMeasurement() VolumeUnitOfMeasurement {
	if o == nil {
		var ret VolumeUnitOfMeasurement
		return ret
	}

	return o.UnitOfMeasurement
}

// GetUnitOfMeasurementOk returns a tuple with the UnitOfMeasurement field value
// and a boolean to check if the value has been set.
func (o *PackageVolume) GetUnitOfMeasurementOk() (*VolumeUnitOfMeasurement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitOfMeasurement, true
}

// SetUnitOfMeasurement sets field value
func (o *PackageVolume) SetUnitOfMeasurement(v VolumeUnitOfMeasurement) {
	o.UnitOfMeasurement = v
}

// GetVolume returns the Volume field value
func (o *PackageVolume) GetVolume() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value
// and a boolean to check if the value has been set.
func (o *PackageVolume) GetVolumeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Volume, true
}

// SetVolume sets field value
func (o *PackageVolume) SetVolume(v float64) {
	o.Volume = v
}

func (o PackageVolume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["unitOfMeasurement"] = o.UnitOfMeasurement
	toSerialize["volume"] = o.Volume
	return toSerialize, nil
}

type NullablePackageVolume struct {
	value *PackageVolume
	isSet bool
}

func (v NullablePackageVolume) Get() *PackageVolume {
	return v.value
}

func (v *NullablePackageVolume) Set(val *PackageVolume) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageVolume) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageVolume(val *PackageVolume) *NullablePackageVolume {
	return &NullablePackageVolume{value: val, isSet: true}
}

func (v NullablePackageVolume) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackageVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
