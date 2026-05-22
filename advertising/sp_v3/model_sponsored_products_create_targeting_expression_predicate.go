package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateTargetingExpressionPredicate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateTargetingExpressionPredicate{}

// SponsoredProductsCreateTargetingExpressionPredicate struct for SponsoredProductsCreateTargetingExpressionPredicate
type SponsoredProductsCreateTargetingExpressionPredicate struct {
	Type SponsoredProductsCreateTargetingExpressionPredicateType `json:"type"`
	// The expression value
	Value *string `json:"value,omitempty"`
}

type _SponsoredProductsCreateTargetingExpressionPredicate SponsoredProductsCreateTargetingExpressionPredicate

// NewSponsoredProductsCreateTargetingExpressionPredicate instantiates a new SponsoredProductsCreateTargetingExpressionPredicate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateTargetingExpressionPredicate(type_ SponsoredProductsCreateTargetingExpressionPredicateType) *SponsoredProductsCreateTargetingExpressionPredicate {
	this := SponsoredProductsCreateTargetingExpressionPredicate{}
	this.Type = type_
	return &this
}

// NewSponsoredProductsCreateTargetingExpressionPredicateWithDefaults instantiates a new SponsoredProductsCreateTargetingExpressionPredicate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateTargetingExpressionPredicateWithDefaults() *SponsoredProductsCreateTargetingExpressionPredicate {
	this := SponsoredProductsCreateTargetingExpressionPredicate{}
	return &this
}

// GetType returns the Type field value
func (o *SponsoredProductsCreateTargetingExpressionPredicate) GetType() SponsoredProductsCreateTargetingExpressionPredicateType {
	if o == nil {
		var ret SponsoredProductsCreateTargetingExpressionPredicateType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetingExpressionPredicate) GetTypeOk() (*SponsoredProductsCreateTargetingExpressionPredicateType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SponsoredProductsCreateTargetingExpressionPredicate) SetType(v SponsoredProductsCreateTargetingExpressionPredicateType) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetingExpressionPredicate) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetingExpressionPredicate) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetingExpressionPredicate) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SponsoredProductsCreateTargetingExpressionPredicate) SetValue(v string) {
	o.Value = &v
}

func (o SponsoredProductsCreateTargetingExpressionPredicate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateTargetingExpressionPredicate struct {
	value *SponsoredProductsCreateTargetingExpressionPredicate
	isSet bool
}

func (v NullableSponsoredProductsCreateTargetingExpressionPredicate) Get() *SponsoredProductsCreateTargetingExpressionPredicate {
	return v.value
}

func (v *NullableSponsoredProductsCreateTargetingExpressionPredicate) Set(val *SponsoredProductsCreateTargetingExpressionPredicate) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateTargetingExpressionPredicate) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateTargetingExpressionPredicate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateTargetingExpressionPredicate(val *SponsoredProductsCreateTargetingExpressionPredicate) *NullableSponsoredProductsCreateTargetingExpressionPredicate {
	return &NullableSponsoredProductsCreateTargetingExpressionPredicate{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateTargetingExpressionPredicate) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateTargetingExpressionPredicate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
