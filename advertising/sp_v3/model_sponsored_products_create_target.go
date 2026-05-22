package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateTarget{}

// SponsoredProductsCreateTarget Target created in the target promotion group.
type SponsoredProductsCreateTarget struct {
	// The id of the target promotion group.
	TargetPromotionGroupId *string `json:"targetPromotionGroupId,omitempty"`
	// The id of the target that got created.
	TargetId *string `json:"targetId,omitempty"`
	// The adGroupId of the manual-targeting campaign where the target belongs.
	ManualTargetingAdGroupId *string `json:"manualTargetingAdGroupId,omitempty"`
}

// NewSponsoredProductsCreateTarget instantiates a new SponsoredProductsCreateTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateTarget() *SponsoredProductsCreateTarget {
	this := SponsoredProductsCreateTarget{}
	return &this
}

// NewSponsoredProductsCreateTargetWithDefaults instantiates a new SponsoredProductsCreateTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateTargetWithDefaults() *SponsoredProductsCreateTarget {
	this := SponsoredProductsCreateTarget{}
	return &this
}

// GetTargetPromotionGroupId returns the TargetPromotionGroupId field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTarget) GetTargetPromotionGroupId() string {
	if o == nil || IsNil(o.TargetPromotionGroupId) {
		var ret string
		return ret
	}
	return *o.TargetPromotionGroupId
}

// GetTargetPromotionGroupIdOk returns a tuple with the TargetPromotionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTarget) GetTargetPromotionGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetPromotionGroupId) {
		return nil, false
	}
	return o.TargetPromotionGroupId, true
}

// HasTargetPromotionGroupId returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTarget) HasTargetPromotionGroupId() bool {
	if o != nil && !IsNil(o.TargetPromotionGroupId) {
		return true
	}

	return false
}

// SetTargetPromotionGroupId gets a reference to the given string and assigns it to the TargetPromotionGroupId field.
func (o *SponsoredProductsCreateTarget) SetTargetPromotionGroupId(v string) {
	o.TargetPromotionGroupId = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTarget) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTarget) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTarget) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *SponsoredProductsCreateTarget) SetTargetId(v string) {
	o.TargetId = &v
}

// GetManualTargetingAdGroupId returns the ManualTargetingAdGroupId field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTarget) GetManualTargetingAdGroupId() string {
	if o == nil || IsNil(o.ManualTargetingAdGroupId) {
		var ret string
		return ret
	}
	return *o.ManualTargetingAdGroupId
}

// GetManualTargetingAdGroupIdOk returns a tuple with the ManualTargetingAdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTarget) GetManualTargetingAdGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ManualTargetingAdGroupId) {
		return nil, false
	}
	return o.ManualTargetingAdGroupId, true
}

// HasManualTargetingAdGroupId returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTarget) HasManualTargetingAdGroupId() bool {
	if o != nil && !IsNil(o.ManualTargetingAdGroupId) {
		return true
	}

	return false
}

// SetManualTargetingAdGroupId gets a reference to the given string and assigns it to the ManualTargetingAdGroupId field.
func (o *SponsoredProductsCreateTarget) SetManualTargetingAdGroupId(v string) {
	o.ManualTargetingAdGroupId = &v
}

func (o SponsoredProductsCreateTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetPromotionGroupId) {
		toSerialize["targetPromotionGroupId"] = o.TargetPromotionGroupId
	}
	if !IsNil(o.TargetId) {
		toSerialize["targetId"] = o.TargetId
	}
	if !IsNil(o.ManualTargetingAdGroupId) {
		toSerialize["manualTargetingAdGroupId"] = o.ManualTargetingAdGroupId
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateTarget struct {
	value *SponsoredProductsCreateTarget
	isSet bool
}

func (v NullableSponsoredProductsCreateTarget) Get() *SponsoredProductsCreateTarget {
	return v.value
}

func (v *NullableSponsoredProductsCreateTarget) Set(val *SponsoredProductsCreateTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateTarget(val *SponsoredProductsCreateTarget) *NullableSponsoredProductsCreateTarget {
	return &NullableSponsoredProductsCreateTarget{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateTarget) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
