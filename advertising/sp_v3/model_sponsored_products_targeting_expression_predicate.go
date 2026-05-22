package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsTargetingExpressionPredicate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsTargetingExpressionPredicate{}

// SponsoredProductsTargetingExpressionPredicate struct for SponsoredProductsTargetingExpressionPredicate
type SponsoredProductsTargetingExpressionPredicate struct {
	Type *SponsoredProductsTargetingExpressionPredicateType `json:"type,omitempty"`
	// The expression value
	Value *string `json:"value,omitempty"`
}

// NewSponsoredProductsTargetingExpressionPredicate instantiates a new SponsoredProductsTargetingExpressionPredicate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsTargetingExpressionPredicate() *SponsoredProductsTargetingExpressionPredicate {
	this := SponsoredProductsTargetingExpressionPredicate{}
	return &this
}

// NewSponsoredProductsTargetingExpressionPredicateWithDefaults instantiates a new SponsoredProductsTargetingExpressionPredicate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsTargetingExpressionPredicateWithDefaults() *SponsoredProductsTargetingExpressionPredicate {
	this := SponsoredProductsTargetingExpressionPredicate{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SponsoredProductsTargetingExpressionPredicate) GetType() SponsoredProductsTargetingExpressionPredicateType {
	if o == nil || IsNil(o.Type) {
		var ret SponsoredProductsTargetingExpressionPredicateType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingExpressionPredicate) GetTypeOk() (*SponsoredProductsTargetingExpressionPredicateType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SponsoredProductsTargetingExpressionPredicate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SponsoredProductsTargetingExpressionPredicateType and assigns it to the Type field.
func (o *SponsoredProductsTargetingExpressionPredicate) SetType(v SponsoredProductsTargetingExpressionPredicateType) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SponsoredProductsTargetingExpressionPredicate) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingExpressionPredicate) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SponsoredProductsTargetingExpressionPredicate) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SponsoredProductsTargetingExpressionPredicate) SetValue(v string) {
	o.Value = &v
}

func (o SponsoredProductsTargetingExpressionPredicate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSponsoredProductsTargetingExpressionPredicate struct {
	value *SponsoredProductsTargetingExpressionPredicate
	isSet bool
}

func (v NullableSponsoredProductsTargetingExpressionPredicate) Get() *SponsoredProductsTargetingExpressionPredicate {
	return v.value
}

func (v *NullableSponsoredProductsTargetingExpressionPredicate) Set(val *SponsoredProductsTargetingExpressionPredicate) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetingExpressionPredicate) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetingExpressionPredicate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetingExpressionPredicate(val *SponsoredProductsTargetingExpressionPredicate) *NullableSponsoredProductsTargetingExpressionPredicate {
	return &NullableSponsoredProductsTargetingExpressionPredicate{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetingExpressionPredicate) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetingExpressionPredicate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
