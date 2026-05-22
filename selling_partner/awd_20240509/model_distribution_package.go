package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the DistributionPackage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DistributionPackage{}

// DistributionPackage Represents an AWD distribution package.
type DistributionPackage struct {
	Contents     DistributionPackageContents `json:"contents"`
	Measurements MeasurementData             `json:"measurements"`
	Type         DistributionPackageType     `json:"type"`
}

type _DistributionPackage DistributionPackage

// NewDistributionPackage instantiates a new DistributionPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistributionPackage(contents DistributionPackageContents, measurements MeasurementData, type_ DistributionPackageType) *DistributionPackage {
	this := DistributionPackage{}
	this.Contents = contents
	this.Measurements = measurements
	this.Type = type_
	return &this
}

// NewDistributionPackageWithDefaults instantiates a new DistributionPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributionPackageWithDefaults() *DistributionPackage {
	this := DistributionPackage{}
	return &this
}

// GetContents returns the Contents field value
func (o *DistributionPackage) GetContents() DistributionPackageContents {
	if o == nil {
		var ret DistributionPackageContents
		return ret
	}

	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value
// and a boolean to check if the value has been set.
func (o *DistributionPackage) GetContentsOk() (*DistributionPackageContents, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contents, true
}

// SetContents sets field value
func (o *DistributionPackage) SetContents(v DistributionPackageContents) {
	o.Contents = v
}

// GetMeasurements returns the Measurements field value
func (o *DistributionPackage) GetMeasurements() MeasurementData {
	if o == nil {
		var ret MeasurementData
		return ret
	}

	return o.Measurements
}

// GetMeasurementsOk returns a tuple with the Measurements field value
// and a boolean to check if the value has been set.
func (o *DistributionPackage) GetMeasurementsOk() (*MeasurementData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Measurements, true
}

// SetMeasurements sets field value
func (o *DistributionPackage) SetMeasurements(v MeasurementData) {
	o.Measurements = v
}

// GetType returns the Type field value
func (o *DistributionPackage) GetType() DistributionPackageType {
	if o == nil {
		var ret DistributionPackageType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DistributionPackage) GetTypeOk() (*DistributionPackageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DistributionPackage) SetType(v DistributionPackageType) {
	o.Type = v
}

func (o DistributionPackage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contents"] = o.Contents
	toSerialize["measurements"] = o.Measurements
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableDistributionPackage struct {
	value *DistributionPackage
	isSet bool
}

func (v NullableDistributionPackage) Get() *DistributionPackage {
	return v.value
}

func (v *NullableDistributionPackage) Set(val *DistributionPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionPackage(val *DistributionPackage) *NullableDistributionPackage {
	return &NullableDistributionPackage{value: val, isSet: true}
}

func (v NullableDistributionPackage) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDistributionPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
