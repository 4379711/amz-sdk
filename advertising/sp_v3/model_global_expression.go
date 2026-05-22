package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalExpression{}

// GlobalExpression struct for GlobalExpression
type GlobalExpression struct {
	// The targeting expression value for each country.
	CountryValues *map[string]string `json:"countryValues,omitempty"`
	// The targeting expression. The type property specifies the targeting option. Use CLOSE_MATCH to match your auto targeting ads closely to the specified value. Use LOOSE_MATCH to match your auto targeting ads broadly to the specified value. Use SUBSTITUTES to display your auto targeting ads along with substitutable products. Use COMPLEMENTS to display your auto targeting ads along with affiliated products. Use KEYWORD_BROAD_MATCH to broadly match your keyword targeting ads with search queries. Use KEYWORD_EXACT_MATCH to exactly match your keyword targeting ads with search queries. Use KEYWORD_PHRASE_MATCH to match your keyword targeting ads with search phrases. Use PAT_ASIN to target an ASIN that is the same as the ASIN expressed. Use PAT_CATEGORY to target the same category as the category expressed. Use PAT_CATEGORY_REFINEMENT to target all of the category refinements. These refinements allows users to specify a brand, price, rating, age range, genre and prime shipping eligibility.
	Type string `json:"type"`
}

type _GlobalExpression GlobalExpression

// NewGlobalExpression instantiates a new GlobalExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalExpression(type_ string) *GlobalExpression {
	this := GlobalExpression{}
	this.Type = type_
	return &this
}

// NewGlobalExpressionWithDefaults instantiates a new GlobalExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalExpressionWithDefaults() *GlobalExpression {
	this := GlobalExpression{}
	return &this
}

// GetCountryValues returns the CountryValues field value if set, zero value otherwise.
func (o *GlobalExpression) GetCountryValues() map[string]string {
	if o == nil || IsNil(o.CountryValues) {
		var ret map[string]string
		return ret
	}
	return *o.CountryValues
}

// GetCountryValuesOk returns a tuple with the CountryValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalExpression) GetCountryValuesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CountryValues) {
		return nil, false
	}
	return o.CountryValues, true
}

// HasCountryValues returns a boolean if a field has been set.
func (o *GlobalExpression) HasCountryValues() bool {
	if o != nil && !IsNil(o.CountryValues) {
		return true
	}

	return false
}

// SetCountryValues gets a reference to the given map[string]string and assigns it to the CountryValues field.
func (o *GlobalExpression) SetCountryValues(v map[string]string) {
	o.CountryValues = &v
}

// GetType returns the Type field value
func (o *GlobalExpression) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GlobalExpression) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GlobalExpression) SetType(v string) {
	o.Type = v
}

func (o GlobalExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryValues) {
		toSerialize["countryValues"] = o.CountryValues
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableGlobalExpression struct {
	value *GlobalExpression
	isSet bool
}

func (v NullableGlobalExpression) Get() *GlobalExpression {
	return v.value
}

func (v *NullableGlobalExpression) Set(val *GlobalExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalExpression(val *GlobalExpression) *NullableGlobalExpression {
	return &NullableGlobalExpression{value: val, isSet: true}
}

func (v NullableGlobalExpression) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
