package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingNegativeProductExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingNegativeProductExpression{}

// SBForecastingNegativeProductExpression Negative expression settings for the target.
type SBForecastingNegativeProductExpression struct {
	// The negative expression type associated with the target. Valid value: ASIN_BRAND_SAME_AS, ASIN_SAME_AS.
	Type *string `json:"type,omitempty"`
	// The expression value associated with targets.
	Value *string `json:"value,omitempty"`
}

// NewSBForecastingNegativeProductExpression instantiates a new SBForecastingNegativeProductExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingNegativeProductExpression() *SBForecastingNegativeProductExpression {
	this := SBForecastingNegativeProductExpression{}
	return &this
}

// NewSBForecastingNegativeProductExpressionWithDefaults instantiates a new SBForecastingNegativeProductExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingNegativeProductExpressionWithDefaults() *SBForecastingNegativeProductExpression {
	this := SBForecastingNegativeProductExpression{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SBForecastingNegativeProductExpression) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingNegativeProductExpression) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SBForecastingNegativeProductExpression) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SBForecastingNegativeProductExpression) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SBForecastingNegativeProductExpression) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingNegativeProductExpression) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SBForecastingNegativeProductExpression) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SBForecastingNegativeProductExpression) SetValue(v string) {
	o.Value = &v
}

func (o SBForecastingNegativeProductExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSBForecastingNegativeProductExpression struct {
	value *SBForecastingNegativeProductExpression
	isSet bool
}

func (v NullableSBForecastingNegativeProductExpression) Get() *SBForecastingNegativeProductExpression {
	return v.value
}

func (v *NullableSBForecastingNegativeProductExpression) Set(val *SBForecastingNegativeProductExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingNegativeProductExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingNegativeProductExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingNegativeProductExpression(val *SBForecastingNegativeProductExpression) *NullableSBForecastingNegativeProductExpression {
	return &NullableSBForecastingNegativeProductExpression{value: val, isSet: true}
}

func (v NullableSBForecastingNegativeProductExpression) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingNegativeProductExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
