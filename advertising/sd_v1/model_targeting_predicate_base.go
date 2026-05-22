package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the TargetingPredicateBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetingPredicateBase{}

// TargetingPredicateBase A predicate to match against inside the TargetingPredicateNested component (only applicable to audience targeting - T00030).  * All IDs passed for category and brand-targeting predicates must be valid IDs in the Amazon Ads browse system. * Brand, price, and review predicates are optional and may only be specified if category is also specified. * Review predicates accept numbers between 0 and 5 and are inclusive. * When using either of the 'between' strings to construct a targeting expression the format of the string is 'double-double' where the first double must be smaller than the second double. Prices are not inclusive. * The 'exactProduct', 'similarProduct', 'relatedProduct', 'negative', and 'audiencesLikelyInterestedInAd' types do not utilize the value field. * The only type currently applicable to Amazon Audiences targeting is 'audienceSameAs'. * A 'relatedProduct' TargetingPredicateBase will Target an audience that has purchased a related product in the past 7,14,30,60,90,180, or 365 days. * The 'audiencesLikelyInterestedInAd' type is only supported when using landingPageType of OFF_AMAZON_LINK.
type TargetingPredicateBase struct {
	Type *string `json:"type,omitempty"`
	// The value to be targeted.
	Value *string `json:"value,omitempty"`
}

// NewTargetingPredicateBase instantiates a new TargetingPredicateBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetingPredicateBase() *TargetingPredicateBase {
	this := TargetingPredicateBase{}
	return &this
}

// NewTargetingPredicateBaseWithDefaults instantiates a new TargetingPredicateBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetingPredicateBaseWithDefaults() *TargetingPredicateBase {
	this := TargetingPredicateBase{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TargetingPredicateBase) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingPredicateBase) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TargetingPredicateBase) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TargetingPredicateBase) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TargetingPredicateBase) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingPredicateBase) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TargetingPredicateBase) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TargetingPredicateBase) SetValue(v string) {
	o.Value = &v
}

func (o TargetingPredicateBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTargetingPredicateBase struct {
	value *TargetingPredicateBase
	isSet bool
}

func (v NullableTargetingPredicateBase) Get() *TargetingPredicateBase {
	return v.value
}

func (v *NullableTargetingPredicateBase) Set(val *TargetingPredicateBase) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetingPredicateBase) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetingPredicateBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetingPredicateBase(val *TargetingPredicateBase) *NullableTargetingPredicateBase {
	return &NullableTargetingPredicateBase{value: val, isSet: true}
}

func (v NullableTargetingPredicateBase) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTargetingPredicateBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
