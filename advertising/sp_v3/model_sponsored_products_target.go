package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsTarget{}

// SponsoredProductsTarget Target promotion group's target.
type SponsoredProductsTarget struct {
	// The id of the target promotion group.
	TargetPromotionGroupId *string `json:"targetPromotionGroupId,omitempty"`
	// The id of the target.
	TargetId *string `json:"targetId,omitempty"`
	// The match type (for KEYWORDs) or the expression type (for PRODUCT). One of QUERY_BROAD_MATCHES,     QUERY_EXACT_MATCHES, QUERY_PHRASE_MATCHES, ASIN_SAME_AS, ASIN_EXPANDED_FROM
	ExpressionType *string `json:"expressionType,omitempty"`
	// The adGroupId of the manual-targeting campaign where the target belongs.
	ManualTargetingAdGroupId *string `json:"manualTargetingAdGroupId,omitempty"`
	// The keyword text or the product ASIN of the target.
	Target *string `json:"target,omitempty"`
}

// NewSponsoredProductsTarget instantiates a new SponsoredProductsTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsTarget() *SponsoredProductsTarget {
	this := SponsoredProductsTarget{}
	return &this
}

// NewSponsoredProductsTargetWithDefaults instantiates a new SponsoredProductsTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsTargetWithDefaults() *SponsoredProductsTarget {
	this := SponsoredProductsTarget{}
	return &this
}

// GetTargetPromotionGroupId returns the TargetPromotionGroupId field value if set, zero value otherwise.
func (o *SponsoredProductsTarget) GetTargetPromotionGroupId() string {
	if o == nil || IsNil(o.TargetPromotionGroupId) {
		var ret string
		return ret
	}
	return *o.TargetPromotionGroupId
}

// GetTargetPromotionGroupIdOk returns a tuple with the TargetPromotionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTarget) GetTargetPromotionGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetPromotionGroupId) {
		return nil, false
	}
	return o.TargetPromotionGroupId, true
}

// HasTargetPromotionGroupId returns a boolean if a field has been set.
func (o *SponsoredProductsTarget) HasTargetPromotionGroupId() bool {
	if o != nil && !IsNil(o.TargetPromotionGroupId) {
		return true
	}

	return false
}

// SetTargetPromotionGroupId gets a reference to the given string and assigns it to the TargetPromotionGroupId field.
func (o *SponsoredProductsTarget) SetTargetPromotionGroupId(v string) {
	o.TargetPromotionGroupId = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *SponsoredProductsTarget) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTarget) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *SponsoredProductsTarget) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *SponsoredProductsTarget) SetTargetId(v string) {
	o.TargetId = &v
}

// GetExpressionType returns the ExpressionType field value if set, zero value otherwise.
func (o *SponsoredProductsTarget) GetExpressionType() string {
	if o == nil || IsNil(o.ExpressionType) {
		var ret string
		return ret
	}
	return *o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTarget) GetExpressionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpressionType) {
		return nil, false
	}
	return o.ExpressionType, true
}

// HasExpressionType returns a boolean if a field has been set.
func (o *SponsoredProductsTarget) HasExpressionType() bool {
	if o != nil && !IsNil(o.ExpressionType) {
		return true
	}

	return false
}

// SetExpressionType gets a reference to the given string and assigns it to the ExpressionType field.
func (o *SponsoredProductsTarget) SetExpressionType(v string) {
	o.ExpressionType = &v
}

// GetManualTargetingAdGroupId returns the ManualTargetingAdGroupId field value if set, zero value otherwise.
func (o *SponsoredProductsTarget) GetManualTargetingAdGroupId() string {
	if o == nil || IsNil(o.ManualTargetingAdGroupId) {
		var ret string
		return ret
	}
	return *o.ManualTargetingAdGroupId
}

// GetManualTargetingAdGroupIdOk returns a tuple with the ManualTargetingAdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTarget) GetManualTargetingAdGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ManualTargetingAdGroupId) {
		return nil, false
	}
	return o.ManualTargetingAdGroupId, true
}

// HasManualTargetingAdGroupId returns a boolean if a field has been set.
func (o *SponsoredProductsTarget) HasManualTargetingAdGroupId() bool {
	if o != nil && !IsNil(o.ManualTargetingAdGroupId) {
		return true
	}

	return false
}

// SetManualTargetingAdGroupId gets a reference to the given string and assigns it to the ManualTargetingAdGroupId field.
func (o *SponsoredProductsTarget) SetManualTargetingAdGroupId(v string) {
	o.ManualTargetingAdGroupId = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *SponsoredProductsTarget) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTarget) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *SponsoredProductsTarget) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *SponsoredProductsTarget) SetTarget(v string) {
	o.Target = &v
}

func (o SponsoredProductsTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetPromotionGroupId) {
		toSerialize["targetPromotionGroupId"] = o.TargetPromotionGroupId
	}
	if !IsNil(o.TargetId) {
		toSerialize["targetId"] = o.TargetId
	}
	if !IsNil(o.ExpressionType) {
		toSerialize["expressionType"] = o.ExpressionType
	}
	if !IsNil(o.ManualTargetingAdGroupId) {
		toSerialize["manualTargetingAdGroupId"] = o.ManualTargetingAdGroupId
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	return toSerialize, nil
}

type NullableSponsoredProductsTarget struct {
	value *SponsoredProductsTarget
	isSet bool
}

func (v NullableSponsoredProductsTarget) Get() *SponsoredProductsTarget {
	return v.value
}

func (v *NullableSponsoredProductsTarget) Set(val *SponsoredProductsTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTarget(val *SponsoredProductsTarget) *NullableSponsoredProductsTarget {
	return &NullableSponsoredProductsTarget{value: val, isSet: true}
}

func (v NullableSponsoredProductsTarget) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
