package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the TargetingExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetingExpression{}

// TargetingExpression The targeting expression. The `type` property specifies the targeting option. Use `CLOSE_MATCH` to match your auto targeting ads closely to the specified value. Use `LOOSE_MATCH` to match your auto targeting ads broadly to the specified value. Use `SUBSTITUTES` to display your auto targeting ads along with substitutable products. Use `COMPLEMENTS` to display your auto targeting ads along with affiliated products. Use `KEYWORD_BROAD_MATCH` to broadly match your keyword targeting ads with search queries. Use `KEYWORD_EXACT_MATCH` to exactly match your keyword targeting ads with search queries. Use `KEYWORD_PHRASE_MATCH` to match your keyword targeting ads with search phrases. your keyword targeting ads with search phrases.
type TargetingExpression struct {
	Type string `json:"type"`
	// The targeting expression value.
	Value *string `json:"value,omitempty"`
}

type _TargetingExpression TargetingExpression

// NewTargetingExpression instantiates a new TargetingExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetingExpression(type_ string) *TargetingExpression {
	this := TargetingExpression{}
	this.Type = type_
	return &this
}

// NewTargetingExpressionWithDefaults instantiates a new TargetingExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetingExpressionWithDefaults() *TargetingExpression {
	this := TargetingExpression{}
	return &this
}

// GetType returns the Type field value
func (o *TargetingExpression) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TargetingExpression) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TargetingExpression) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TargetingExpression) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingExpression) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TargetingExpression) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TargetingExpression) SetValue(v string) {
	o.Value = &v
}

func (o TargetingExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTargetingExpression struct {
	value *TargetingExpression
	isSet bool
}

func (v NullableTargetingExpression) Get() *TargetingExpression {
	return v.value
}

func (v *NullableTargetingExpression) Set(val *TargetingExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetingExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetingExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetingExpression(val *TargetingExpression) *NullableTargetingExpression {
	return &NullableTargetingExpression{value: val, isSet: true}
}

func (v NullableTargetingExpression) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTargetingExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
