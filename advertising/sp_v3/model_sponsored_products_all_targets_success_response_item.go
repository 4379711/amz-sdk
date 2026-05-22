package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsAllTargetsSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsAllTargetsSuccessResponseItem{}

// SponsoredProductsAllTargetsSuccessResponseItem struct for SponsoredProductsAllTargetsSuccessResponseItem
type SponsoredProductsAllTargetsSuccessResponseItem struct {
	// the index of the targetingClause in the array from the request body
	Index  int32                        `json:"index"`
	Target *SponsoredProductsAllTargets `json:"target,omitempty"`
}

type _SponsoredProductsAllTargetsSuccessResponseItem SponsoredProductsAllTargetsSuccessResponseItem

// NewSponsoredProductsAllTargetsSuccessResponseItem instantiates a new SponsoredProductsAllTargetsSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsAllTargetsSuccessResponseItem(index int32) *SponsoredProductsAllTargetsSuccessResponseItem {
	this := SponsoredProductsAllTargetsSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsAllTargetsSuccessResponseItemWithDefaults instantiates a new SponsoredProductsAllTargetsSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsAllTargetsSuccessResponseItemWithDefaults() *SponsoredProductsAllTargetsSuccessResponseItem {
	this := SponsoredProductsAllTargetsSuccessResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsAllTargetsSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAllTargetsSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsAllTargetsSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *SponsoredProductsAllTargetsSuccessResponseItem) GetTarget() SponsoredProductsAllTargets {
	if o == nil || IsNil(o.Target) {
		var ret SponsoredProductsAllTargets
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAllTargetsSuccessResponseItem) GetTargetOk() (*SponsoredProductsAllTargets, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *SponsoredProductsAllTargetsSuccessResponseItem) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given SponsoredProductsAllTargets and assigns it to the Target field.
func (o *SponsoredProductsAllTargetsSuccessResponseItem) SetTarget(v SponsoredProductsAllTargets) {
	o.Target = &v
}

func (o SponsoredProductsAllTargetsSuccessResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	return toSerialize, nil
}

type NullableSponsoredProductsAllTargetsSuccessResponseItem struct {
	value *SponsoredProductsAllTargetsSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsAllTargetsSuccessResponseItem) Get() *SponsoredProductsAllTargetsSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsAllTargetsSuccessResponseItem) Set(val *SponsoredProductsAllTargetsSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAllTargetsSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAllTargetsSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAllTargetsSuccessResponseItem(val *SponsoredProductsAllTargetsSuccessResponseItem) *NullableSponsoredProductsAllTargetsSuccessResponseItem {
	return &NullableSponsoredProductsAllTargetsSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsAllTargetsSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAllTargetsSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
