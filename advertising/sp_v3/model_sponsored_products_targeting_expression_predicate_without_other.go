package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsTargetingExpressionPredicateWithoutOther type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsTargetingExpressionPredicateWithoutOther{}

// SponsoredProductsTargetingExpressionPredicateWithoutOther struct for SponsoredProductsTargetingExpressionPredicateWithoutOther
type SponsoredProductsTargetingExpressionPredicateWithoutOther struct {
	Type SponsoredProductsTargetingExpressionPredicateTypeWithoutOther `json:"type"`
	// The expression value
	Value *string `json:"value,omitempty"`
}

type _SponsoredProductsTargetingExpressionPredicateWithoutOther SponsoredProductsTargetingExpressionPredicateWithoutOther

// NewSponsoredProductsTargetingExpressionPredicateWithoutOther instantiates a new SponsoredProductsTargetingExpressionPredicateWithoutOther object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsTargetingExpressionPredicateWithoutOther(type_ SponsoredProductsTargetingExpressionPredicateTypeWithoutOther) *SponsoredProductsTargetingExpressionPredicateWithoutOther {
	this := SponsoredProductsTargetingExpressionPredicateWithoutOther{}
	this.Type = type_
	return &this
}

// NewSponsoredProductsTargetingExpressionPredicateWithoutOtherWithDefaults instantiates a new SponsoredProductsTargetingExpressionPredicateWithoutOther object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsTargetingExpressionPredicateWithoutOtherWithDefaults() *SponsoredProductsTargetingExpressionPredicateWithoutOther {
	this := SponsoredProductsTargetingExpressionPredicateWithoutOther{}
	return &this
}

// GetType returns the Type field value
func (o *SponsoredProductsTargetingExpressionPredicateWithoutOther) GetType() SponsoredProductsTargetingExpressionPredicateTypeWithoutOther {
	if o == nil {
		var ret SponsoredProductsTargetingExpressionPredicateTypeWithoutOther
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingExpressionPredicateWithoutOther) GetTypeOk() (*SponsoredProductsTargetingExpressionPredicateTypeWithoutOther, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SponsoredProductsTargetingExpressionPredicateWithoutOther) SetType(v SponsoredProductsTargetingExpressionPredicateTypeWithoutOther) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SponsoredProductsTargetingExpressionPredicateWithoutOther) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingExpressionPredicateWithoutOther) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SponsoredProductsTargetingExpressionPredicateWithoutOther) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SponsoredProductsTargetingExpressionPredicateWithoutOther) SetValue(v string) {
	o.Value = &v
}

func (o SponsoredProductsTargetingExpressionPredicateWithoutOther) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSponsoredProductsTargetingExpressionPredicateWithoutOther struct {
	value *SponsoredProductsTargetingExpressionPredicateWithoutOther
	isSet bool
}

func (v NullableSponsoredProductsTargetingExpressionPredicateWithoutOther) Get() *SponsoredProductsTargetingExpressionPredicateWithoutOther {
	return v.value
}

func (v *NullableSponsoredProductsTargetingExpressionPredicateWithoutOther) Set(val *SponsoredProductsTargetingExpressionPredicateWithoutOther) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetingExpressionPredicateWithoutOther) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetingExpressionPredicateWithoutOther) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetingExpressionPredicateWithoutOther(val *SponsoredProductsTargetingExpressionPredicateWithoutOther) *NullableSponsoredProductsTargetingExpressionPredicateWithoutOther {
	return &NullableSponsoredProductsTargetingExpressionPredicateWithoutOther{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetingExpressionPredicateWithoutOther) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetingExpressionPredicateWithoutOther) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
