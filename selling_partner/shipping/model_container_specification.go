package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the ContainerSpecification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSpecification{}

// ContainerSpecification Container specification for checking the service rate.
type ContainerSpecification struct {
	Dimensions Dimensions `json:"dimensions"`
	Weight     Weight     `json:"weight"`
}

type _ContainerSpecification ContainerSpecification

// NewContainerSpecification instantiates a new ContainerSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSpecification(dimensions Dimensions, weight Weight) *ContainerSpecification {
	this := ContainerSpecification{}
	this.Dimensions = dimensions
	this.Weight = weight
	return &this
}

// NewContainerSpecificationWithDefaults instantiates a new ContainerSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSpecificationWithDefaults() *ContainerSpecification {
	this := ContainerSpecification{}
	return &this
}

// GetDimensions returns the Dimensions field value
func (o *ContainerSpecification) GetDimensions() Dimensions {
	if o == nil {
		var ret Dimensions
		return ret
	}

	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *ContainerSpecification) GetDimensionsOk() (*Dimensions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value
func (o *ContainerSpecification) SetDimensions(v Dimensions) {
	o.Dimensions = v
}

// GetWeight returns the Weight field value
func (o *ContainerSpecification) GetWeight() Weight {
	if o == nil {
		var ret Weight
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *ContainerSpecification) GetWeightOk() (*Weight, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *ContainerSpecification) SetWeight(v Weight) {
	o.Weight = v
}

func (o ContainerSpecification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dimensions"] = o.Dimensions
	toSerialize["weight"] = o.Weight
	return toSerialize, nil
}

type NullableContainerSpecification struct {
	value *ContainerSpecification
	isSet bool
}

func (v NullableContainerSpecification) Get() *ContainerSpecification {
	return v.value
}

func (v *NullableContainerSpecification) Set(val *ContainerSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSpecification(val *ContainerSpecification) *NullableContainerSpecification {
	return &NullableContainerSpecification{value: val, isSet: true}
}

func (v NullableContainerSpecification) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableContainerSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
