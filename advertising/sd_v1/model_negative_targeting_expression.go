package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the NegativeTargetingExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NegativeTargetingExpression{}

// NegativeTargetingExpression struct for NegativeTargetingExpression
type NegativeTargetingExpression struct {
	// The intent type. See the [targeting topic](https://advertising.amazon.com/help#GQCBASRVERXSARL3) in the Amazon Ads support center for more information.
	Type *string `json:"type,omitempty"`
	// The value to be negatively targeted. Used only in manual expressions.
	Value *string `json:"value,omitempty"`
}

// NewNegativeTargetingExpression instantiates a new NegativeTargetingExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNegativeTargetingExpression() *NegativeTargetingExpression {
	this := NegativeTargetingExpression{}
	return &this
}

// NewNegativeTargetingExpressionWithDefaults instantiates a new NegativeTargetingExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNegativeTargetingExpressionWithDefaults() *NegativeTargetingExpression {
	this := NegativeTargetingExpression{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NegativeTargetingExpression) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeTargetingExpression) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NegativeTargetingExpression) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NegativeTargetingExpression) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *NegativeTargetingExpression) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeTargetingExpression) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *NegativeTargetingExpression) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *NegativeTargetingExpression) SetValue(v string) {
	o.Value = &v
}

func (o NegativeTargetingExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableNegativeTargetingExpression struct {
	value *NegativeTargetingExpression
	isSet bool
}

func (v NullableNegativeTargetingExpression) Get() *NegativeTargetingExpression {
	return v.value
}

func (v *NullableNegativeTargetingExpression) Set(val *NegativeTargetingExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableNegativeTargetingExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableNegativeTargetingExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNegativeTargetingExpression(val *NegativeTargetingExpression) *NullableNegativeTargetingExpression {
	return &NullableNegativeTargetingExpression{value: val, isSet: true}
}

func (v NullableNegativeTargetingExpression) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableNegativeTargetingExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
