package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsAdGroupSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsAdGroupSuccessResponseItem{}

// SponsoredProductsAdGroupSuccessResponseItem struct for SponsoredProductsAdGroupSuccessResponseItem
type SponsoredProductsAdGroupSuccessResponseItem struct {
	AdGroup *SponsoredProductsAdGroup `json:"adGroup,omitempty"`
	// the index of the adGroup in the array from the request body
	Index int32 `json:"index"`
	// the adGroup ID
	AdGroupId *string `json:"adGroupId,omitempty"`
}

type _SponsoredProductsAdGroupSuccessResponseItem SponsoredProductsAdGroupSuccessResponseItem

// NewSponsoredProductsAdGroupSuccessResponseItem instantiates a new SponsoredProductsAdGroupSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsAdGroupSuccessResponseItem(index int32) *SponsoredProductsAdGroupSuccessResponseItem {
	this := SponsoredProductsAdGroupSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsAdGroupSuccessResponseItemWithDefaults instantiates a new SponsoredProductsAdGroupSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsAdGroupSuccessResponseItemWithDefaults() *SponsoredProductsAdGroupSuccessResponseItem {
	this := SponsoredProductsAdGroupSuccessResponseItem{}
	return &this
}

// GetAdGroup returns the AdGroup field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupSuccessResponseItem) GetAdGroup() SponsoredProductsAdGroup {
	if o == nil || IsNil(o.AdGroup) {
		var ret SponsoredProductsAdGroup
		return ret
	}
	return *o.AdGroup
}

// GetAdGroupOk returns a tuple with the AdGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupSuccessResponseItem) GetAdGroupOk() (*SponsoredProductsAdGroup, bool) {
	if o == nil || IsNil(o.AdGroup) {
		return nil, false
	}
	return o.AdGroup, true
}

// HasAdGroup returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupSuccessResponseItem) HasAdGroup() bool {
	if o != nil && !IsNil(o.AdGroup) {
		return true
	}

	return false
}

// SetAdGroup gets a reference to the given SponsoredProductsAdGroup and assigns it to the AdGroup field.
func (o *SponsoredProductsAdGroupSuccessResponseItem) SetAdGroup(v SponsoredProductsAdGroup) {
	o.AdGroup = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsAdGroupSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsAdGroupSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupSuccessResponseItem) GetAdGroupId() string {
	if o == nil || IsNil(o.AdGroupId) {
		var ret string
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupSuccessResponseItem) GetAdGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupSuccessResponseItem) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given string and assigns it to the AdGroupId field.
func (o *SponsoredProductsAdGroupSuccessResponseItem) SetAdGroupId(v string) {
	o.AdGroupId = &v
}

func (o SponsoredProductsAdGroupSuccessResponseItem) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsAdGroupSuccessResponseItem struct {
	value *SponsoredProductsAdGroupSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsAdGroupSuccessResponseItem) Get() *SponsoredProductsAdGroupSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsAdGroupSuccessResponseItem) Set(val *SponsoredProductsAdGroupSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAdGroupSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAdGroupSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAdGroupSuccessResponseItem(val *SponsoredProductsAdGroupSuccessResponseItem) *NullableSponsoredProductsAdGroupSuccessResponseItem {
	return &NullableSponsoredProductsAdGroupSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsAdGroupSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAdGroupSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
