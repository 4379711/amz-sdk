package replenishment20221107

import (
	"github.com/bytedance/sonic"
)

// checks if the DiscountFunding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscountFunding{}

// DiscountFunding The discount funding on the offer.
type DiscountFunding struct {
	// Filters the results to only include offers with the percentage specified.
	Percentage []float32 `json:"percentage,omitempty"`
}

// NewDiscountFunding instantiates a new DiscountFunding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscountFunding() *DiscountFunding {
	this := DiscountFunding{}
	return &this
}

// NewDiscountFundingWithDefaults instantiates a new DiscountFunding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscountFundingWithDefaults() *DiscountFunding {
	this := DiscountFunding{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *DiscountFunding) GetPercentage() []float32 {
	if o == nil || IsNil(o.Percentage) {
		var ret []float32
		return ret
	}
	return o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountFunding) GetPercentageOk() ([]float32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *DiscountFunding) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given []float32 and assigns it to the Percentage field.
func (o *DiscountFunding) SetPercentage(v []float32) {
	o.Percentage = v
}

func (o DiscountFunding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return toSerialize, nil
}

type NullableDiscountFunding struct {
	value *DiscountFunding
	isSet bool
}

func (v NullableDiscountFunding) Get() *DiscountFunding {
	return v.value
}

func (v *NullableDiscountFunding) Set(val *DiscountFunding) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscountFunding) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscountFunding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscountFunding(val *DiscountFunding) *NullableDiscountFunding {
	return &NullableDiscountFunding{value: val, isSet: true}
}

func (v NullableDiscountFunding) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDiscountFunding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
