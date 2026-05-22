package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the GoalProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoalProduct{}

// GoalProduct A product an advertisers wants to advertise. Recommendations will be made for specified goal products.
type GoalProduct struct {
	// Amazon Standard Identification Number
	Asin string `json:"asin" validate:"regexp=[a-zA-Z0-9]{10}"`
}

type _GoalProduct GoalProduct

// NewGoalProduct instantiates a new GoalProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoalProduct(asin string) *GoalProduct {
	this := GoalProduct{}
	this.Asin = asin
	return &this
}

// NewGoalProductWithDefaults instantiates a new GoalProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoalProductWithDefaults() *GoalProduct {
	this := GoalProduct{}
	return &this
}

// GetAsin returns the Asin field value
func (o *GoalProduct) GetAsin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asin
}

// GetAsinOk returns a tuple with the Asin field value
// and a boolean to check if the value has been set.
func (o *GoalProduct) GetAsinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asin, true
}

// SetAsin sets field value
func (o *GoalProduct) SetAsin(v string) {
	o.Asin = v
}

func (o GoalProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asin"] = o.Asin
	return toSerialize, nil
}

type NullableGoalProduct struct {
	value *GoalProduct
	isSet bool
}

func (v NullableGoalProduct) Get() *GoalProduct {
	return v.value
}

func (v *NullableGoalProduct) Set(val *GoalProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableGoalProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableGoalProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoalProduct(val *GoalProduct) *NullableGoalProduct {
	return &NullableGoalProduct{value: val, isSet: true}
}

func (v NullableGoalProduct) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGoalProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
