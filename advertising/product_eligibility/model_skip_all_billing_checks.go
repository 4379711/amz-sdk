package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the SkipAllBillingChecks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkipAllBillingChecks{}

// SkipAllBillingChecks struct for SkipAllBillingChecks
type SkipAllBillingChecks struct {
	// Skip all billing/payments/suspension related checks
	SkipAllBillingChecks bool `json:"skipAllBillingChecks"`
}

type _SkipAllBillingChecks SkipAllBillingChecks

// NewSkipAllBillingChecks instantiates a new SkipAllBillingChecks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkipAllBillingChecks(skipAllBillingChecks bool) *SkipAllBillingChecks {
	this := SkipAllBillingChecks{}
	this.SkipAllBillingChecks = skipAllBillingChecks
	return &this
}

// NewSkipAllBillingChecksWithDefaults instantiates a new SkipAllBillingChecks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkipAllBillingChecksWithDefaults() *SkipAllBillingChecks {
	this := SkipAllBillingChecks{}
	return &this
}

// GetSkipAllBillingChecks returns the SkipAllBillingChecks field value
func (o *SkipAllBillingChecks) GetSkipAllBillingChecks() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SkipAllBillingChecks
}

// GetSkipAllBillingChecksOk returns a tuple with the SkipAllBillingChecks field value
// and a boolean to check if the value has been set.
func (o *SkipAllBillingChecks) GetSkipAllBillingChecksOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkipAllBillingChecks, true
}

// SetSkipAllBillingChecks sets field value
func (o *SkipAllBillingChecks) SetSkipAllBillingChecks(v bool) {
	o.SkipAllBillingChecks = v
}

func (o SkipAllBillingChecks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["skipAllBillingChecks"] = o.SkipAllBillingChecks
	return toSerialize, nil
}

type NullableSkipAllBillingChecks struct {
	value *SkipAllBillingChecks
	isSet bool
}

func (v NullableSkipAllBillingChecks) Get() *SkipAllBillingChecks {
	return v.value
}

func (v *NullableSkipAllBillingChecks) Set(val *SkipAllBillingChecks) {
	v.value = val
	v.isSet = true
}

func (v NullableSkipAllBillingChecks) IsSet() bool {
	return v.isSet
}

func (v *NullableSkipAllBillingChecks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkipAllBillingChecks(val *SkipAllBillingChecks) *NullableSkipAllBillingChecks {
	return &NullableSkipAllBillingChecks{value: val, isSet: true}
}

func (v NullableSkipAllBillingChecks) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSkipAllBillingChecks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
