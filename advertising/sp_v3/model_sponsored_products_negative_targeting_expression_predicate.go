package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsNegativeTargetingExpressionPredicate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNegativeTargetingExpressionPredicate{}

// SponsoredProductsNegativeTargetingExpressionPredicate struct for SponsoredProductsNegativeTargetingExpressionPredicate
type SponsoredProductsNegativeTargetingExpressionPredicate struct {
	Type *SponsoredProductsNegativeTargetingExpressionPredicateType `json:"type,omitempty"`
	// The expression value
	Value *string `json:"value,omitempty"`
}

// NewSponsoredProductsNegativeTargetingExpressionPredicate instantiates a new SponsoredProductsNegativeTargetingExpressionPredicate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNegativeTargetingExpressionPredicate() *SponsoredProductsNegativeTargetingExpressionPredicate {
	this := SponsoredProductsNegativeTargetingExpressionPredicate{}
	return &this
}

// NewSponsoredProductsNegativeTargetingExpressionPredicateWithDefaults instantiates a new SponsoredProductsNegativeTargetingExpressionPredicate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNegativeTargetingExpressionPredicateWithDefaults() *SponsoredProductsNegativeTargetingExpressionPredicate {
	this := SponsoredProductsNegativeTargetingExpressionPredicate{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeTargetingExpressionPredicate) GetType() SponsoredProductsNegativeTargetingExpressionPredicateType {
	if o == nil || IsNil(o.Type) {
		var ret SponsoredProductsNegativeTargetingExpressionPredicateType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetingExpressionPredicate) GetTypeOk() (*SponsoredProductsNegativeTargetingExpressionPredicateType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeTargetingExpressionPredicate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SponsoredProductsNegativeTargetingExpressionPredicateType and assigns it to the Type field.
func (o *SponsoredProductsNegativeTargetingExpressionPredicate) SetType(v SponsoredProductsNegativeTargetingExpressionPredicateType) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeTargetingExpressionPredicate) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetingExpressionPredicate) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeTargetingExpressionPredicate) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SponsoredProductsNegativeTargetingExpressionPredicate) SetValue(v string) {
	o.Value = &v
}

func (o SponsoredProductsNegativeTargetingExpressionPredicate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSponsoredProductsNegativeTargetingExpressionPredicate struct {
	value *SponsoredProductsNegativeTargetingExpressionPredicate
	isSet bool
}

func (v NullableSponsoredProductsNegativeTargetingExpressionPredicate) Get() *SponsoredProductsNegativeTargetingExpressionPredicate {
	return v.value
}

func (v *NullableSponsoredProductsNegativeTargetingExpressionPredicate) Set(val *SponsoredProductsNegativeTargetingExpressionPredicate) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNegativeTargetingExpressionPredicate) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNegativeTargetingExpressionPredicate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNegativeTargetingExpressionPredicate(val *SponsoredProductsNegativeTargetingExpressionPredicate) *NullableSponsoredProductsNegativeTargetingExpressionPredicate {
	return &NullableSponsoredProductsNegativeTargetingExpressionPredicate{value: val, isSet: true}
}

func (v NullableSponsoredProductsNegativeTargetingExpressionPredicate) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNegativeTargetingExpressionPredicate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
