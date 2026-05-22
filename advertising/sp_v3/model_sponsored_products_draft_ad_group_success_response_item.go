package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftAdGroupSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftAdGroupSuccessResponseItem{}

// SponsoredProductsDraftAdGroupSuccessResponseItem struct for SponsoredProductsDraftAdGroupSuccessResponseItem
type SponsoredProductsDraftAdGroupSuccessResponseItem struct {
	AdGroup *SponsoredProductsDraftAdGroup `json:"adGroup,omitempty"`
	// the index of the draftAdGroup in the array from the request body
	Index int32 `json:"index"`
	// the adGroup ID
	AdGroupId *string `json:"adGroupId,omitempty"`
}

type _SponsoredProductsDraftAdGroupSuccessResponseItem SponsoredProductsDraftAdGroupSuccessResponseItem

// NewSponsoredProductsDraftAdGroupSuccessResponseItem instantiates a new SponsoredProductsDraftAdGroupSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftAdGroupSuccessResponseItem(index int32) *SponsoredProductsDraftAdGroupSuccessResponseItem {
	this := SponsoredProductsDraftAdGroupSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsDraftAdGroupSuccessResponseItemWithDefaults instantiates a new SponsoredProductsDraftAdGroupSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftAdGroupSuccessResponseItemWithDefaults() *SponsoredProductsDraftAdGroupSuccessResponseItem {
	this := SponsoredProductsDraftAdGroupSuccessResponseItem{}
	return &this
}

// GetAdGroup returns the AdGroup field value if set, zero value otherwise.
func (o *SponsoredProductsDraftAdGroupSuccessResponseItem) GetAdGroup() SponsoredProductsDraftAdGroup {
	if o == nil || IsNil(o.AdGroup) {
		var ret SponsoredProductsDraftAdGroup
		return ret
	}
	return *o.AdGroup
}

// GetAdGroupOk returns a tuple with the AdGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroupSuccessResponseItem) GetAdGroupOk() (*SponsoredProductsDraftAdGroup, bool) {
	if o == nil || IsNil(o.AdGroup) {
		return nil, false
	}
	return o.AdGroup, true
}

// HasAdGroup returns a boolean if a field has been set.
func (o *SponsoredProductsDraftAdGroupSuccessResponseItem) HasAdGroup() bool {
	if o != nil && !IsNil(o.AdGroup) {
		return true
	}

	return false
}

// SetAdGroup gets a reference to the given SponsoredProductsDraftAdGroup and assigns it to the AdGroup field.
func (o *SponsoredProductsDraftAdGroupSuccessResponseItem) SetAdGroup(v SponsoredProductsDraftAdGroup) {
	o.AdGroup = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsDraftAdGroupSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroupSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsDraftAdGroupSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *SponsoredProductsDraftAdGroupSuccessResponseItem) GetAdGroupId() string {
	if o == nil || IsNil(o.AdGroupId) {
		var ret string
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroupSuccessResponseItem) GetAdGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *SponsoredProductsDraftAdGroupSuccessResponseItem) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given string and assigns it to the AdGroupId field.
func (o *SponsoredProductsDraftAdGroupSuccessResponseItem) SetAdGroupId(v string) {
	o.AdGroupId = &v
}

func (o SponsoredProductsDraftAdGroupSuccessResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdGroup) {
		toSerialize["adGroup"] = o.AdGroup
	}
	toSerialize["index"] = o.Index
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftAdGroupSuccessResponseItem struct {
	value *SponsoredProductsDraftAdGroupSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsDraftAdGroupSuccessResponseItem) Get() *SponsoredProductsDraftAdGroupSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsDraftAdGroupSuccessResponseItem) Set(val *SponsoredProductsDraftAdGroupSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftAdGroupSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftAdGroupSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftAdGroupSuccessResponseItem(val *SponsoredProductsDraftAdGroupSuccessResponseItem) *NullableSponsoredProductsDraftAdGroupSuccessResponseItem {
	return &NullableSponsoredProductsDraftAdGroupSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftAdGroupSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftAdGroupSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
