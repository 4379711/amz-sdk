package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingPredicateV31 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingPredicateV31{}

// SDTargetingPredicateV31 A predicate to match against in the Targeting Expression (only applicable to contextual targeting - T00020).  * All IDs passed for category and brand-targeting predicates must be valid IDs in the Amazon Ads browse system. * Brand, price, and review predicates are optional and may only be specified if category is also specified. * Review predicates accept numbers between 0 and 5 and are inclusive. * When using either of the 'between' strings to construct a targeting expression the format of the string is 'double-double' where the first double must be smaller than the second double. Prices are not inclusive.
type SDTargetingPredicateV31 struct {
	Type string `json:"type"`
	// The value to be targeted.
	Value *string `json:"value,omitempty"`
}

type _SDTargetingPredicateV31 SDTargetingPredicateV31

// NewSDTargetingPredicateV31 instantiates a new SDTargetingPredicateV31 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingPredicateV31(type_ string) *SDTargetingPredicateV31 {
	this := SDTargetingPredicateV31{}
	this.Type = type_
	return &this
}

// NewSDTargetingPredicateV31WithDefaults instantiates a new SDTargetingPredicateV31 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingPredicateV31WithDefaults() *SDTargetingPredicateV31 {
	this := SDTargetingPredicateV31{}
	return &this
}

// GetType returns the Type field value
func (o *SDTargetingPredicateV31) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SDTargetingPredicateV31) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SDTargetingPredicateV31) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SDTargetingPredicateV31) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingPredicateV31) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SDTargetingPredicateV31) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SDTargetingPredicateV31) SetValue(v string) {
	o.Value = &v
}

func (o SDTargetingPredicateV31) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSDTargetingPredicateV31 struct {
	value *SDTargetingPredicateV31
	isSet bool
}

func (v NullableSDTargetingPredicateV31) Get() *SDTargetingPredicateV31 {
	return v.value
}

func (v *NullableSDTargetingPredicateV31) Set(val *SDTargetingPredicateV31) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingPredicateV31) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingPredicateV31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingPredicateV31(val *SDTargetingPredicateV31) *NullableSDTargetingPredicateV31 {
	return &NullableSDTargetingPredicateV31{value: val, isSet: true}
}

func (v NullableSDTargetingPredicateV31) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingPredicateV31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
