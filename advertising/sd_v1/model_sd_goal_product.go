package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDGoalProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDGoalProduct{}

// SDGoalProduct A product an advertisers wants to advertise. Recommendations will be made for specified goal products.
type SDGoalProduct struct {
	// Amazon Standard Identification Number
	Asin string `json:"asin" validate:"regexp=[a-zA-Z0-9]{10}"`
}

type _SDGoalProduct SDGoalProduct

// NewSDGoalProduct instantiates a new SDGoalProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDGoalProduct(asin string) *SDGoalProduct {
	this := SDGoalProduct{}
	this.Asin = asin
	return &this
}

// NewSDGoalProductWithDefaults instantiates a new SDGoalProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDGoalProductWithDefaults() *SDGoalProduct {
	this := SDGoalProduct{}
	return &this
}

// GetAsin returns the Asin field value
func (o *SDGoalProduct) GetAsin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asin
}

// GetAsinOk returns a tuple with the Asin field value
// and a boolean to check if the value has been set.
func (o *SDGoalProduct) GetAsinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asin, true
}

// SetAsin sets field value
func (o *SDGoalProduct) SetAsin(v string) {
	o.Asin = v
}

func (o SDGoalProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asin"] = o.Asin
	return toSerialize, nil
}

type NullableSDGoalProduct struct {
	value *SDGoalProduct
	isSet bool
}

func (v NullableSDGoalProduct) Get() *SDGoalProduct {
	return v.value
}

func (v *NullableSDGoalProduct) Set(val *SDGoalProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableSDGoalProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableSDGoalProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDGoalProduct(val *SDGoalProduct) *NullableSDGoalProduct {
	return &NullableSDGoalProduct{value: val, isSet: true}
}

func (v NullableSDGoalProduct) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDGoalProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
