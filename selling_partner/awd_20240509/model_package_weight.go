package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the PackageWeight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageWeight{}

// PackageWeight Represents the weight of the package with a unit of measurement.
type PackageWeight struct {
	UnitOfMeasurement WeightUnitOfMeasurement `json:"unitOfMeasurement"`
	// The package weight value.
	Weight float64 `json:"weight"`
}

type _PackageWeight PackageWeight

// NewPackageWeight instantiates a new PackageWeight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageWeight(unitOfMeasurement WeightUnitOfMeasurement, weight float64) *PackageWeight {
	this := PackageWeight{}
	this.UnitOfMeasurement = unitOfMeasurement
	this.Weight = weight
	return &this
}

// NewPackageWeightWithDefaults instantiates a new PackageWeight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageWeightWithDefaults() *PackageWeight {
	this := PackageWeight{}
	return &this
}

// GetUnitOfMeasurement returns the UnitOfMeasurement field value
func (o *PackageWeight) GetUnitOfMeasurement() WeightUnitOfMeasurement {
	if o == nil {
		var ret WeightUnitOfMeasurement
		return ret
	}

	return o.UnitOfMeasurement
}

// GetUnitOfMeasurementOk returns a tuple with the UnitOfMeasurement field value
// and a boolean to check if the value has been set.
func (o *PackageWeight) GetUnitOfMeasurementOk() (*WeightUnitOfMeasurement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitOfMeasurement, true
}

// SetUnitOfMeasurement sets field value
func (o *PackageWeight) SetUnitOfMeasurement(v WeightUnitOfMeasurement) {
	o.UnitOfMeasurement = v
}

// GetWeight returns the Weight field value
func (o *PackageWeight) GetWeight() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *PackageWeight) GetWeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *PackageWeight) SetWeight(v float64) {
	o.Weight = v
}

func (o PackageWeight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["unitOfMeasurement"] = o.UnitOfMeasurement
	toSerialize["weight"] = o.Weight
	return toSerialize, nil
}

type NullablePackageWeight struct {
	value *PackageWeight
	isSet bool
}

func (v NullablePackageWeight) Get() *PackageWeight {
	return v.value
}

func (v *NullablePackageWeight) Set(val *PackageWeight) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageWeight) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageWeight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageWeight(val *PackageWeight) *NullablePackageWeight {
	return &NullablePackageWeight{value: val, isSet: true}
}

func (v NullablePackageWeight) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackageWeight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
