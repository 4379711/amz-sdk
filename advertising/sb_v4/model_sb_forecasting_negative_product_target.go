package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingNegativeProductTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingNegativeProductTarget{}

// SBForecastingNegativeProductTarget The negative target associated with the ad group.
type SBForecastingNegativeProductTarget struct {
	Expressions []SBForecastingNegativeProductExpression `json:"expressions,omitempty"`
}

// NewSBForecastingNegativeProductTarget instantiates a new SBForecastingNegativeProductTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingNegativeProductTarget() *SBForecastingNegativeProductTarget {
	this := SBForecastingNegativeProductTarget{}
	return &this
}

// NewSBForecastingNegativeProductTargetWithDefaults instantiates a new SBForecastingNegativeProductTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingNegativeProductTargetWithDefaults() *SBForecastingNegativeProductTarget {
	this := SBForecastingNegativeProductTarget{}
	return &this
}

// GetExpressions returns the Expressions field value if set, zero value otherwise.
func (o *SBForecastingNegativeProductTarget) GetExpressions() []SBForecastingNegativeProductExpression {
	if o == nil || IsNil(o.Expressions) {
		var ret []SBForecastingNegativeProductExpression
		return ret
	}
	return o.Expressions
}

// GetExpressionsOk returns a tuple with the Expressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingNegativeProductTarget) GetExpressionsOk() ([]SBForecastingNegativeProductExpression, bool) {
	if o == nil || IsNil(o.Expressions) {
		return nil, false
	}
	return o.Expressions, true
}

// HasExpressions returns a boolean if a field has been set.
func (o *SBForecastingNegativeProductTarget) HasExpressions() bool {
	if o != nil && !IsNil(o.Expressions) {
		return true
	}

	return false
}

// SetExpressions gets a reference to the given []SBForecastingNegativeProductExpression and assigns it to the Expressions field.
func (o *SBForecastingNegativeProductTarget) SetExpressions(v []SBForecastingNegativeProductExpression) {
	o.Expressions = v
}

func (o SBForecastingNegativeProductTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expressions) {
		toSerialize["expressions"] = o.Expressions
	}
	return toSerialize, nil
}

type NullableSBForecastingNegativeProductTarget struct {
	value *SBForecastingNegativeProductTarget
	isSet bool
}

func (v NullableSBForecastingNegativeProductTarget) Get() *SBForecastingNegativeProductTarget {
	return v.value
}

func (v *NullableSBForecastingNegativeProductTarget) Set(val *SBForecastingNegativeProductTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingNegativeProductTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingNegativeProductTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingNegativeProductTarget(val *SBForecastingNegativeProductTarget) *NullableSBForecastingNegativeProductTarget {
	return &NullableSBForecastingNegativeProductTarget{value: val, isSet: true}
}

func (v NullableSBForecastingNegativeProductTarget) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingNegativeProductTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
