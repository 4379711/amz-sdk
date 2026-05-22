package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the TargetingPredicate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetingPredicate{}

// TargetingPredicate A predicate to match against in the targeting expression (only applicable to contextual targeting - T00020).  * All IDs passed for category and brand-targeting predicates must be valid IDs in the Amazon Ads browse system. * Brand, price, and review predicates are optional and may only be specified if category is also specified. * Review predicates accept numbers between 0 and 5 and are inclusive. * When using either of the 'between' strings to construct a targeting expression the format of the string is 'double-double' where the first double must be smaller than the second double. Prices are not inclusive.
type TargetingPredicate struct {
	Type *string `json:"type,omitempty"`
	// The value to be targeted.
	Value *string `json:"value,omitempty"`
}

// NewTargetingPredicate instantiates a new TargetingPredicate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetingPredicate() *TargetingPredicate {
	this := TargetingPredicate{}
	return &this
}

// NewTargetingPredicateWithDefaults instantiates a new TargetingPredicate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetingPredicateWithDefaults() *TargetingPredicate {
	this := TargetingPredicate{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TargetingPredicate) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingPredicate) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TargetingPredicate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TargetingPredicate) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TargetingPredicate) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingPredicate) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TargetingPredicate) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TargetingPredicate) SetValue(v string) {
	o.Value = &v
}

func (o TargetingPredicate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTargetingPredicate struct {
	value *TargetingPredicate
	isSet bool
}

func (v NullableTargetingPredicate) Get() *TargetingPredicate {
	return v.value
}

func (v *NullableTargetingPredicate) Set(val *TargetingPredicate) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetingPredicate) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetingPredicate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetingPredicate(val *TargetingPredicate) *NullableTargetingPredicate {
	return &NullableTargetingPredicate{value: val, isSet: true}
}

func (v NullableTargetingPredicate) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTargetingPredicate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
