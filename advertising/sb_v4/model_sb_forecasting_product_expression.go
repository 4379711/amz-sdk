package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingProductExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingProductExpression{}

// SBForecastingProductExpression Expression settings for the target.
type SBForecastingProductExpression struct {
	// The expression type associated with the target. Valid value: ASIN_CATEGORY_SAME_AS, ASIN_BRAND_SAME_AS, ASIN_PRICE_LESS_THAN, ASIN_PRICE_BETWEEN, ASIN_PRICE_GREATER_THAN, ASIN_REVIEW_RATING_LESS_THAN, ASIN_REVIEW_RATING_BETWEEN, ASIN_REVIEW_RATING_GREATER_THAN, ASIN_SAME_AS.
	Type *string `json:"type,omitempty"`
	// The expression value associated with targets.
	Value *string `json:"value,omitempty"`
}

// NewSBForecastingProductExpression instantiates a new SBForecastingProductExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingProductExpression() *SBForecastingProductExpression {
	this := SBForecastingProductExpression{}
	return &this
}

// NewSBForecastingProductExpressionWithDefaults instantiates a new SBForecastingProductExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingProductExpressionWithDefaults() *SBForecastingProductExpression {
	this := SBForecastingProductExpression{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SBForecastingProductExpression) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingProductExpression) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SBForecastingProductExpression) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SBForecastingProductExpression) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SBForecastingProductExpression) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingProductExpression) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SBForecastingProductExpression) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SBForecastingProductExpression) SetValue(v string) {
	o.Value = &v
}

func (o SBForecastingProductExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSBForecastingProductExpression struct {
	value *SBForecastingProductExpression
	isSet bool
}

func (v NullableSBForecastingProductExpression) Get() *SBForecastingProductExpression {
	return v.value
}

func (v *NullableSBForecastingProductExpression) Set(val *SBForecastingProductExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingProductExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingProductExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingProductExpression(val *SBForecastingProductExpression) *NullableSBForecastingProductExpression {
	return &NullableSBForecastingProductExpression{value: val, isSet: true}
}

func (v NullableSBForecastingProductExpression) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingProductExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
