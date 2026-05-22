package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDContentTargetingPredicateV31 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDContentTargetingPredicateV31{}

// SDContentTargetingPredicateV31 A predicate to match against in the content targeting expression.
type SDContentTargetingPredicateV31 struct {
	Type string `json:"type"`
	// The value to be targeted.
	Value string `json:"value"`
}

type _SDContentTargetingPredicateV31 SDContentTargetingPredicateV31

// NewSDContentTargetingPredicateV31 instantiates a new SDContentTargetingPredicateV31 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDContentTargetingPredicateV31(type_ string, value string) *SDContentTargetingPredicateV31 {
	this := SDContentTargetingPredicateV31{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewSDContentTargetingPredicateV31WithDefaults instantiates a new SDContentTargetingPredicateV31 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDContentTargetingPredicateV31WithDefaults() *SDContentTargetingPredicateV31 {
	this := SDContentTargetingPredicateV31{}
	return &this
}

// GetType returns the Type field value
func (o *SDContentTargetingPredicateV31) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SDContentTargetingPredicateV31) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SDContentTargetingPredicateV31) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *SDContentTargetingPredicateV31) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SDContentTargetingPredicateV31) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SDContentTargetingPredicateV31) SetValue(v string) {
	o.Value = v
}

func (o SDContentTargetingPredicateV31) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableSDContentTargetingPredicateV31 struct {
	value *SDContentTargetingPredicateV31
	isSet bool
}

func (v NullableSDContentTargetingPredicateV31) Get() *SDContentTargetingPredicateV31 {
	return v.value
}

func (v *NullableSDContentTargetingPredicateV31) Set(val *SDContentTargetingPredicateV31) {
	v.value = val
	v.isSet = true
}

func (v NullableSDContentTargetingPredicateV31) IsSet() bool {
	return v.isSet
}

func (v *NullableSDContentTargetingPredicateV31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDContentTargetingPredicateV31(val *SDContentTargetingPredicateV31) *NullableSDContentTargetingPredicateV31 {
	return &NullableSDContentTargetingPredicateV31{value: val, isSet: true}
}

func (v NullableSDContentTargetingPredicateV31) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDContentTargetingPredicateV31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
