package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the MultiCountryTargetingExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiCountryTargetingExpression{}

// MultiCountryTargetingExpression The targeting expression. The `type` property specifies the targeting option. Use `CLOSE_MATCH` to match your auto targeting ads closely to the specified value. Use `LOOSE_MATCH` to match your auto targeting ads broadly to the specified value. Use `SUBSTITUTES` to display your auto targeting ads along with substitutable products. Use `COMPLEMENTS` to display your auto targeting ads along with affiliated products. Use `KEYWORD_BROAD_MATCH` to broadly match your keyword targeting ads with search queries. Use `KEYWORD_EXACT_MATCH` to exactly match your keyword targeting ads with search queries. Use `KEYWORD_PHRASE_MATCH` to match your keyword targeting ads with search phrases. your keyword targeting ads with search phrases. Use `PAT_ASIN` to match your product attribute targeting ads with product ASIN. Use `PAT_CATEGORY` to match your product attribute targeting ads with product category. Use `PAT_CATEGORY_REFINEMENT` to match your product attribute targeting ads with product attribute, including brand, price, rating, prime shipping eligible, age range and genre. Use `KEYWORD_GROUP` to match your keyword targeting ads with keyword group.
type MultiCountryTargetingExpression struct {
	CountryValues *map[string]string                  `json:"countryValues,omitempty"`
	Type          MultiCountryTargetingExpressionType `json:"type"`
}

type _MultiCountryTargetingExpression MultiCountryTargetingExpression

// NewMultiCountryTargetingExpression instantiates a new MultiCountryTargetingExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiCountryTargetingExpression(type_ MultiCountryTargetingExpressionType) *MultiCountryTargetingExpression {
	this := MultiCountryTargetingExpression{}
	this.Type = type_
	return &this
}

// NewMultiCountryTargetingExpressionWithDefaults instantiates a new MultiCountryTargetingExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiCountryTargetingExpressionWithDefaults() *MultiCountryTargetingExpression {
	this := MultiCountryTargetingExpression{}
	return &this
}

// GetCountryValues returns the CountryValues field value if set, zero value otherwise.
func (o *MultiCountryTargetingExpression) GetCountryValues() map[string]string {
	if o == nil || IsNil(o.CountryValues) {
		var ret map[string]string
		return ret
	}
	return *o.CountryValues
}

// GetCountryValuesOk returns a tuple with the CountryValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiCountryTargetingExpression) GetCountryValuesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CountryValues) {
		return nil, false
	}
	return o.CountryValues, true
}

// HasCountryValues returns a boolean if a field has been set.
func (o *MultiCountryTargetingExpression) HasCountryValues() bool {
	if o != nil && !IsNil(o.CountryValues) {
		return true
	}

	return false
}

// SetCountryValues gets a reference to the given map[string]string and assigns it to the CountryValues field.
func (o *MultiCountryTargetingExpression) SetCountryValues(v map[string]string) {
	o.CountryValues = &v
}

// GetType returns the Type field value
func (o *MultiCountryTargetingExpression) GetType() MultiCountryTargetingExpressionType {
	if o == nil {
		var ret MultiCountryTargetingExpressionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MultiCountryTargetingExpression) GetTypeOk() (*MultiCountryTargetingExpressionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MultiCountryTargetingExpression) SetType(v MultiCountryTargetingExpressionType) {
	o.Type = v
}

func (o MultiCountryTargetingExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryValues) {
		toSerialize["countryValues"] = o.CountryValues
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableMultiCountryTargetingExpression struct {
	value *MultiCountryTargetingExpression
	isSet bool
}

func (v NullableMultiCountryTargetingExpression) Get() *MultiCountryTargetingExpression {
	return v.value
}

func (v *NullableMultiCountryTargetingExpression) Set(val *MultiCountryTargetingExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiCountryTargetingExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiCountryTargetingExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiCountryTargetingExpression(val *MultiCountryTargetingExpression) *NullableMultiCountryTargetingExpression {
	return &NullableMultiCountryTargetingExpression{value: val, isSet: true}
}

func (v NullableMultiCountryTargetingExpression) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMultiCountryTargetingExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
