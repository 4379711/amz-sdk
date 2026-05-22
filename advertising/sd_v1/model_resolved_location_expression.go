package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the ResolvedLocationExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResolvedLocationExpression{}

// ResolvedLocationExpression struct for ResolvedLocationExpression
type ResolvedLocationExpression struct {
	Type *LocationPredicate `json:"type,omitempty"`
	// The human-readable location name.
	Value *string `json:"value,omitempty"`
}

// NewResolvedLocationExpression instantiates a new ResolvedLocationExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResolvedLocationExpression() *ResolvedLocationExpression {
	this := ResolvedLocationExpression{}
	return &this
}

// NewResolvedLocationExpressionWithDefaults instantiates a new ResolvedLocationExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResolvedLocationExpressionWithDefaults() *ResolvedLocationExpression {
	this := ResolvedLocationExpression{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ResolvedLocationExpression) GetType() LocationPredicate {
	if o == nil || IsNil(o.Type) {
		var ret LocationPredicate
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolvedLocationExpression) GetTypeOk() (*LocationPredicate, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ResolvedLocationExpression) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given LocationPredicate and assigns it to the Type field.
func (o *ResolvedLocationExpression) SetType(v LocationPredicate) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ResolvedLocationExpression) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolvedLocationExpression) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ResolvedLocationExpression) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ResolvedLocationExpression) SetValue(v string) {
	o.Value = &v
}

func (o ResolvedLocationExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableResolvedLocationExpression struct {
	value *ResolvedLocationExpression
	isSet bool
}

func (v NullableResolvedLocationExpression) Get() *ResolvedLocationExpression {
	return v.value
}

func (v *NullableResolvedLocationExpression) Set(val *ResolvedLocationExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableResolvedLocationExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableResolvedLocationExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResolvedLocationExpression(val *ResolvedLocationExpression) *NullableResolvedLocationExpression {
	return &NullableResolvedLocationExpression{value: val, isSet: true}
}

func (v NullableResolvedLocationExpression) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableResolvedLocationExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
