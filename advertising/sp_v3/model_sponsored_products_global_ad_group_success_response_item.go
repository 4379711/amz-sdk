package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalAdGroupSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalAdGroupSuccessResponseItem{}

// SponsoredProductsGlobalAdGroupSuccessResponseItem struct for SponsoredProductsGlobalAdGroupSuccessResponseItem
type SponsoredProductsGlobalAdGroupSuccessResponseItem struct {
	AdGroup *SponsoredProductsGlobalAdGroup `json:"adGroup,omitempty"`
	// the index of the adGroup in the array from the request body
	Index int32 `json:"index"`
	// the adGroup ID
	AdGroupId *string `json:"adGroupId,omitempty"`
}

type _SponsoredProductsGlobalAdGroupSuccessResponseItem SponsoredProductsGlobalAdGroupSuccessResponseItem

// NewSponsoredProductsGlobalAdGroupSuccessResponseItem instantiates a new SponsoredProductsGlobalAdGroupSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalAdGroupSuccessResponseItem(index int32) *SponsoredProductsGlobalAdGroupSuccessResponseItem {
	this := SponsoredProductsGlobalAdGroupSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsGlobalAdGroupSuccessResponseItemWithDefaults instantiates a new SponsoredProductsGlobalAdGroupSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalAdGroupSuccessResponseItemWithDefaults() *SponsoredProductsGlobalAdGroupSuccessResponseItem {
	this := SponsoredProductsGlobalAdGroupSuccessResponseItem{}
	return &this
}

// GetAdGroup returns the AdGroup field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalAdGroupSuccessResponseItem) GetAdGroup() SponsoredProductsGlobalAdGroup {
	if o == nil || IsNil(o.AdGroup) {
		var ret SponsoredProductsGlobalAdGroup
		return ret
	}
	return *o.AdGroup
}

// GetAdGroupOk returns a tuple with the AdGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalAdGroupSuccessResponseItem) GetAdGroupOk() (*SponsoredProductsGlobalAdGroup, bool) {
	if o == nil || IsNil(o.AdGroup) {
		return nil, false
	}
	return o.AdGroup, true
}

// HasAdGroup returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalAdGroupSuccessResponseItem) HasAdGroup() bool {
	if o != nil && !IsNil(o.AdGroup) {
		return true
	}

	return false
}

// SetAdGroup gets a reference to the given SponsoredProductsGlobalAdGroup and assigns it to the AdGroup field.
func (o *SponsoredProductsGlobalAdGroupSuccessResponseItem) SetAdGroup(v SponsoredProductsGlobalAdGroup) {
	o.AdGroup = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsGlobalAdGroupSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalAdGroupSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsGlobalAdGroupSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalAdGroupSuccessResponseItem) GetAdGroupId() string {
	if o == nil || IsNil(o.AdGroupId) {
		var ret string
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalAdGroupSuccessResponseItem) GetAdGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalAdGroupSuccessResponseItem) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given string and assigns it to the AdGroupId field.
func (o *SponsoredProductsGlobalAdGroupSuccessResponseItem) SetAdGroupId(v string) {
	o.AdGroupId = &v
}

func (o SponsoredProductsGlobalAdGroupSuccessResponseItem) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsGlobalAdGroupSuccessResponseItem struct {
	value *SponsoredProductsGlobalAdGroupSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsGlobalAdGroupSuccessResponseItem) Get() *SponsoredProductsGlobalAdGroupSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsGlobalAdGroupSuccessResponseItem) Set(val *SponsoredProductsGlobalAdGroupSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalAdGroupSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalAdGroupSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalAdGroupSuccessResponseItem(val *SponsoredProductsGlobalAdGroupSuccessResponseItem) *NullableSponsoredProductsGlobalAdGroupSuccessResponseItem {
	return &NullableSponsoredProductsGlobalAdGroupSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalAdGroupSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalAdGroupSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
