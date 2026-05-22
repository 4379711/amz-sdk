package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the BoxRequirements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BoxRequirements{}

// BoxRequirements The requirements for a box in the packing option.
type BoxRequirements struct {
	Weight WeightRange `json:"weight"`
}

type _BoxRequirements BoxRequirements

// NewBoxRequirements instantiates a new BoxRequirements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoxRequirements(weight WeightRange) *BoxRequirements {
	this := BoxRequirements{}
	this.Weight = weight
	return &this
}

// NewBoxRequirementsWithDefaults instantiates a new BoxRequirements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoxRequirementsWithDefaults() *BoxRequirements {
	this := BoxRequirements{}
	return &this
}

// GetWeight returns the Weight field value
func (o *BoxRequirements) GetWeight() WeightRange {
	if o == nil {
		var ret WeightRange
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *BoxRequirements) GetWeightOk() (*WeightRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *BoxRequirements) SetWeight(v WeightRange) {
	o.Weight = v
}

func (o BoxRequirements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["weight"] = o.Weight
	return toSerialize, nil
}

type NullableBoxRequirements struct {
	value *BoxRequirements
	isSet bool
}

func (v NullableBoxRequirements) Get() *BoxRequirements {
	return v.value
}

func (v *NullableBoxRequirements) Set(val *BoxRequirements) {
	v.value = val
	v.isSet = true
}

func (v NullableBoxRequirements) IsSet() bool {
	return v.isSet
}

func (v *NullableBoxRequirements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoxRequirements(val *BoxRequirements) *NullableBoxRequirements {
	return &NullableBoxRequirements{value: val, isSet: true}
}

func (v NullableBoxRequirements) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBoxRequirements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
