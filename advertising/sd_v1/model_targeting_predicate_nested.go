package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the TargetingPredicateNested type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetingPredicateNested{}

// TargetingPredicateNested A behavioral event and list of targeting predicates that represents an audience to target (only applicable to audience targeting - T00030).  * For manual ASIN-grain targeting, the value array must contain only, 'exactProduct', 'similarProduct', 'relatedProduct' and 'lookback' TargetingPredicateBase components. The 'lookback' is mandatory and the value should be set to '7', '14', '30', '60', '90', '180' or '365'. * For manual Category-grain targeting, the value array must contain a 'lookback' and 'asinCategorySameAs' TargetingPredicateBase component, which can be further refined with optional brand, price, star-rating and shipping eligibility refinements. The 'lookback' is mandatory and the value should be set to '7', '14', '30', '60', '90', '180' or '365'. * For Amazon Audiences targeting, the TargetingPredicateNested type should be set to 'audience' and the value array should include one TargetingPredicateBase component with type set to 'audienceSameAs'.
type TargetingPredicateNested struct {
	Type  *string                  `json:"type,omitempty"`
	Value []TargetingPredicateBase `json:"value,omitempty"`
}

// NewTargetingPredicateNested instantiates a new TargetingPredicateNested object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetingPredicateNested() *TargetingPredicateNested {
	this := TargetingPredicateNested{}
	return &this
}

// NewTargetingPredicateNestedWithDefaults instantiates a new TargetingPredicateNested object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetingPredicateNestedWithDefaults() *TargetingPredicateNested {
	this := TargetingPredicateNested{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TargetingPredicateNested) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingPredicateNested) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TargetingPredicateNested) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TargetingPredicateNested) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TargetingPredicateNested) GetValue() []TargetingPredicateBase {
	if o == nil || IsNil(o.Value) {
		var ret []TargetingPredicateBase
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingPredicateNested) GetValueOk() ([]TargetingPredicateBase, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TargetingPredicateNested) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []TargetingPredicateBase and assigns it to the Value field.
func (o *TargetingPredicateNested) SetValue(v []TargetingPredicateBase) {
	o.Value = v
}

func (o TargetingPredicateNested) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTargetingPredicateNested struct {
	value *TargetingPredicateNested
	isSet bool
}

func (v NullableTargetingPredicateNested) Get() *TargetingPredicateNested {
	return v.value
}

func (v *NullableTargetingPredicateNested) Set(val *TargetingPredicateNested) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetingPredicateNested) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetingPredicateNested) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetingPredicateNested(val *TargetingPredicateNested) *NullableTargetingPredicateNested {
	return &NullableTargetingPredicateNested{value: val, isSet: true}
}

func (v NullableTargetingPredicateNested) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTargetingPredicateNested) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
