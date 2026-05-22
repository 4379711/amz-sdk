package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDProductTargetingThemeExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDProductTargetingThemeExpression{}

// SDProductTargetingThemeExpression The expression used to define the contextual targeting theme.
type SDProductTargetingThemeExpression struct {
	// The contextual targeting grammar used to define the targeting theme. Note asinAsBestSeller is currently not supported.
	Type string `json:"type"`
}

type _SDProductTargetingThemeExpression SDProductTargetingThemeExpression

// NewSDProductTargetingThemeExpression instantiates a new SDProductTargetingThemeExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDProductTargetingThemeExpression(type_ string) *SDProductTargetingThemeExpression {
	this := SDProductTargetingThemeExpression{}
	this.Type = type_
	return &this
}

// NewSDProductTargetingThemeExpressionWithDefaults instantiates a new SDProductTargetingThemeExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDProductTargetingThemeExpressionWithDefaults() *SDProductTargetingThemeExpression {
	this := SDProductTargetingThemeExpression{}
	return &this
}

// GetType returns the Type field value
func (o *SDProductTargetingThemeExpression) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SDProductTargetingThemeExpression) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SDProductTargetingThemeExpression) SetType(v string) {
	o.Type = v
}

func (o SDProductTargetingThemeExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableSDProductTargetingThemeExpression struct {
	value *SDProductTargetingThemeExpression
	isSet bool
}

func (v NullableSDProductTargetingThemeExpression) Get() *SDProductTargetingThemeExpression {
	return v.value
}

func (v *NullableSDProductTargetingThemeExpression) Set(val *SDProductTargetingThemeExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableSDProductTargetingThemeExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableSDProductTargetingThemeExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDProductTargetingThemeExpression(val *SDProductTargetingThemeExpression) *NullableSDProductTargetingThemeExpression {
	return &NullableSDProductTargetingThemeExpression{value: val, isSet: true}
}

func (v NullableSDProductTargetingThemeExpression) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDProductTargetingThemeExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
