package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalTargetingExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalTargetingExpression{}

// GlobalTargetingExpression struct for GlobalTargetingExpression
type GlobalTargetingExpression struct {
	Expression *GlobalExpression `json:"expression,omitempty"`
	// The bid value for each country.
	CountryBids map[string]float32 `json:"countryBids"`
}

type _GlobalTargetingExpression GlobalTargetingExpression

// NewGlobalTargetingExpression instantiates a new GlobalTargetingExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalTargetingExpression(countryBids map[string]float32) *GlobalTargetingExpression {
	this := GlobalTargetingExpression{}
	this.CountryBids = countryBids
	return &this
}

// NewGlobalTargetingExpressionWithDefaults instantiates a new GlobalTargetingExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalTargetingExpressionWithDefaults() *GlobalTargetingExpression {
	this := GlobalTargetingExpression{}
	return &this
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *GlobalTargetingExpression) GetExpression() GlobalExpression {
	if o == nil || IsNil(o.Expression) {
		var ret GlobalExpression
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalTargetingExpression) GetExpressionOk() (*GlobalExpression, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *GlobalTargetingExpression) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given GlobalExpression and assigns it to the Expression field.
func (o *GlobalTargetingExpression) SetExpression(v GlobalExpression) {
	o.Expression = &v
}

// GetCountryBids returns the CountryBids field value
func (o *GlobalTargetingExpression) GetCountryBids() map[string]float32 {
	if o == nil {
		var ret map[string]float32
		return ret
	}

	return o.CountryBids
}

// GetCountryBidsOk returns a tuple with the CountryBids field value
// and a boolean to check if the value has been set.
func (o *GlobalTargetingExpression) GetCountryBidsOk() (*map[string]float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryBids, true
}

// SetCountryBids sets field value
func (o *GlobalTargetingExpression) SetCountryBids(v map[string]float32) {
	o.CountryBids = v
}

func (o GlobalTargetingExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	toSerialize["countryBids"] = o.CountryBids
	return toSerialize, nil
}

type NullableGlobalTargetingExpression struct {
	value *GlobalTargetingExpression
	isSet bool
}

func (v NullableGlobalTargetingExpression) Get() *GlobalTargetingExpression {
	return v.value
}

func (v *NullableGlobalTargetingExpression) Set(val *GlobalTargetingExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalTargetingExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalTargetingExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalTargetingExpression(val *GlobalTargetingExpression) *NullableGlobalTargetingExpression {
	return &NullableGlobalTargetingExpression{value: val, isSet: true}
}

func (v NullableGlobalTargetingExpression) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalTargetingExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
