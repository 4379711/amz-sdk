package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingPredicateNestedV31 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingPredicateNestedV31{}

// SDTargetingPredicateNestedV31 A behavioral event and list of targeting predicates that represents an audience to target (only applicable to audience targeting - T00030).  * For manual ASIN-grain targeting, the value array must contain only, 'exactProduct', 'similarProduct', 'relatedProduct' and 'lookback' TargetingPredicateBase components. The 'lookback' is mandatory and the value should be set to '7', '14', '30', '60', '90', '180' or '365'. * For manual Category-grain targeting, the value array must contain a 'lookback' and 'asinCategorySameAs' TargetingPredicateBase component, which can be further refined with optional brand, price, star-rating and shipping eligibility refinements. The 'lookback' is mandatory and the value should be set to '7', '14', '30', '60', '90', '180' or '365'. * For manual Category-grain targeting, the value array must contain a 'lookback' and 'asinCategorySameAs' TargetingPredicateBase component, which can be further refined with optional brand, price, star-rating and shipping eligibility refinements. * For Amazon Audiences targeting, the TargetingPredicateNested type should be set to 'audience' and the value array should include one TargetingPredicateBase component with type set to 'audienceSameAs'.
type SDTargetingPredicateNestedV31 struct {
	Type  string                        `json:"type"`
	Value []SDTargetingPredicateBaseV31 `json:"value"`
}

type _SDTargetingPredicateNestedV31 SDTargetingPredicateNestedV31

// NewSDTargetingPredicateNestedV31 instantiates a new SDTargetingPredicateNestedV31 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingPredicateNestedV31(type_ string, value []SDTargetingPredicateBaseV31) *SDTargetingPredicateNestedV31 {
	this := SDTargetingPredicateNestedV31{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewSDTargetingPredicateNestedV31WithDefaults instantiates a new SDTargetingPredicateNestedV31 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingPredicateNestedV31WithDefaults() *SDTargetingPredicateNestedV31 {
	this := SDTargetingPredicateNestedV31{}
	return &this
}

// GetType returns the Type field value
func (o *SDTargetingPredicateNestedV31) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SDTargetingPredicateNestedV31) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SDTargetingPredicateNestedV31) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *SDTargetingPredicateNestedV31) GetValue() []SDTargetingPredicateBaseV31 {
	if o == nil {
		var ret []SDTargetingPredicateBaseV31
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SDTargetingPredicateNestedV31) GetValueOk() ([]SDTargetingPredicateBaseV31, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *SDTargetingPredicateNestedV31) SetValue(v []SDTargetingPredicateBaseV31) {
	o.Value = v
}

func (o SDTargetingPredicateNestedV31) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableSDTargetingPredicateNestedV31 struct {
	value *SDTargetingPredicateNestedV31
	isSet bool
}

func (v NullableSDTargetingPredicateNestedV31) Get() *SDTargetingPredicateNestedV31 {
	return v.value
}

func (v *NullableSDTargetingPredicateNestedV31) Set(val *SDTargetingPredicateNestedV31) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingPredicateNestedV31) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingPredicateNestedV31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingPredicateNestedV31(val *SDTargetingPredicateNestedV31) *NullableSDTargetingPredicateNestedV31 {
	return &NullableSDTargetingPredicateNestedV31{value: val, isSet: true}
}

func (v NullableSDTargetingPredicateNestedV31) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingPredicateNestedV31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
