package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignNegativeKeywordSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignNegativeKeywordSuccessResponseItem{}

// SponsoredProductsCampaignNegativeKeywordSuccessResponseItem struct for SponsoredProductsCampaignNegativeKeywordSuccessResponseItem
type SponsoredProductsCampaignNegativeKeywordSuccessResponseItem struct {
	CampaignNegativeKeyword *SponsoredProductsCampaignNegativeKeyword `json:"campaignNegativeKeyword,omitempty"`
	// the campaignNegativeKeyword ID
	CampaignNegativeKeywordId *string `json:"campaignNegativeKeywordId,omitempty"`
	// the index of the campaign in the array from the request body
	Index int32 `json:"index"`
}

type _SponsoredProductsCampaignNegativeKeywordSuccessResponseItem SponsoredProductsCampaignNegativeKeywordSuccessResponseItem

// NewSponsoredProductsCampaignNegativeKeywordSuccessResponseItem instantiates a new SponsoredProductsCampaignNegativeKeywordSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignNegativeKeywordSuccessResponseItem(index int32) *SponsoredProductsCampaignNegativeKeywordSuccessResponseItem {
	this := SponsoredProductsCampaignNegativeKeywordSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsCampaignNegativeKeywordSuccessResponseItemWithDefaults instantiates a new SponsoredProductsCampaignNegativeKeywordSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignNegativeKeywordSuccessResponseItemWithDefaults() *SponsoredProductsCampaignNegativeKeywordSuccessResponseItem {
	this := SponsoredProductsCampaignNegativeKeywordSuccessResponseItem{}
	return &this
}

// GetCampaignNegativeKeyword returns the CampaignNegativeKeyword field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignNegativeKeywordSuccessResponseItem) GetCampaignNegativeKeyword() SponsoredProductsCampaignNegativeKeyword {
	if o == nil || IsNil(o.CampaignNegativeKeyword) {
		var ret SponsoredProductsCampaignNegativeKeyword
		return ret
	}
	return *o.CampaignNegativeKeyword
}

// GetCampaignNegativeKeywordOk returns a tuple with the CampaignNegativeKeyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeywordSuccessResponseItem) GetCampaignNegativeKeywordOk() (*SponsoredProductsCampaignNegativeKeyword, bool) {
	if o == nil || IsNil(o.CampaignNegativeKeyword) {
		return nil, false
	}
	return o.CampaignNegativeKeyword, true
}

// HasCampaignNegativeKeyword returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignNegativeKeywordSuccessResponseItem) HasCampaignNegativeKeyword() bool {
	if o != nil && !IsNil(o.CampaignNegativeKeyword) {
		return true
	}

	return false
}

// SetCampaignNegativeKeyword gets a reference to the given SponsoredProductsCampaignNegativeKeyword and assigns it to the CampaignNegativeKeyword field.
func (o *SponsoredProductsCampaignNegativeKeywordSuccessResponseItem) SetCampaignNegativeKeyword(v SponsoredProductsCampaignNegativeKeyword) {
	o.CampaignNegativeKeyword = &v
}

// GetCampaignNegativeKeywordId returns the CampaignNegativeKeywordId field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignNegativeKeywordSuccessResponseItem) GetCampaignNegativeKeywordId() string {
	if o == nil || IsNil(o.CampaignNegativeKeywordId) {
		var ret string
		return ret
	}
	return *o.CampaignNegativeKeywordId
}

// GetCampaignNegativeKeywordIdOk returns a tuple with the CampaignNegativeKeywordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeywordSuccessResponseItem) GetCampaignNegativeKeywordIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignNegativeKeywordId) {
		return nil, false
	}
	return o.CampaignNegativeKeywordId, true
}

// HasCampaignNegativeKeywordId returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignNegativeKeywordSuccessResponseItem) HasCampaignNegativeKeywordId() bool {
	if o != nil && !IsNil(o.CampaignNegativeKeywordId) {
		return true
	}

	return false
}

// SetCampaignNegativeKeywordId gets a reference to the given string and assigns it to the CampaignNegativeKeywordId field.
func (o *SponsoredProductsCampaignNegativeKeywordSuccessResponseItem) SetCampaignNegativeKeywordId(v string) {
	o.CampaignNegativeKeywordId = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsCampaignNegativeKeywordSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeywordSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsCampaignNegativeKeywordSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

func (o SponsoredProductsCampaignNegativeKeywordSuccessResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignNegativeKeyword) {
		toSerialize["campaignNegativeKeyword"] = o.CampaignNegativeKeyword
	}
	if !IsNil(o.CampaignNegativeKeywordId) {
		toSerialize["campaignNegativeKeywordId"] = o.CampaignNegativeKeywordId
	}
	toSerialize["index"] = o.Index
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignNegativeKeywordSuccessResponseItem struct {
	value *SponsoredProductsCampaignNegativeKeywordSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsCampaignNegativeKeywordSuccessResponseItem) Get() *SponsoredProductsCampaignNegativeKeywordSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsCampaignNegativeKeywordSuccessResponseItem) Set(val *SponsoredProductsCampaignNegativeKeywordSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignNegativeKeywordSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignNegativeKeywordSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignNegativeKeywordSuccessResponseItem(val *SponsoredProductsCampaignNegativeKeywordSuccessResponseItem) *NullableSponsoredProductsCampaignNegativeKeywordSuccessResponseItem {
	return &NullableSponsoredProductsCampaignNegativeKeywordSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignNegativeKeywordSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignNegativeKeywordSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
