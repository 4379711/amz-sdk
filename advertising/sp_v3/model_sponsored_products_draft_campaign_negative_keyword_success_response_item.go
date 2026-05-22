package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem{}

// SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem struct for SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem
type SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem struct {
	CampaignNegativeKeyword *SponsoredProductsDraftCampaignNegativeKeyword `json:"campaignNegativeKeyword,omitempty"`
	// the negativeKeyword ID
	CampaignNegativeKeywordId *string `json:"campaignNegativeKeywordId,omitempty"`
	// the index of the draft in the array from the request body
	Index int32 `json:"index"`
}

type _SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem

// NewSponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem instantiates a new SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem(index int32) *SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem {
	this := SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItemWithDefaults instantiates a new SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItemWithDefaults() *SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem {
	this := SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem{}
	return &this
}

// GetCampaignNegativeKeyword returns the CampaignNegativeKeyword field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) GetCampaignNegativeKeyword() SponsoredProductsDraftCampaignNegativeKeyword {
	if o == nil || IsNil(o.CampaignNegativeKeyword) {
		var ret SponsoredProductsDraftCampaignNegativeKeyword
		return ret
	}
	return *o.CampaignNegativeKeyword
}

// GetCampaignNegativeKeywordOk returns a tuple with the CampaignNegativeKeyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) GetCampaignNegativeKeywordOk() (*SponsoredProductsDraftCampaignNegativeKeyword, bool) {
	if o == nil || IsNil(o.CampaignNegativeKeyword) {
		return nil, false
	}
	return o.CampaignNegativeKeyword, true
}

// HasCampaignNegativeKeyword returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) HasCampaignNegativeKeyword() bool {
	if o != nil && !IsNil(o.CampaignNegativeKeyword) {
		return true
	}

	return false
}

// SetCampaignNegativeKeyword gets a reference to the given SponsoredProductsDraftCampaignNegativeKeyword and assigns it to the CampaignNegativeKeyword field.
func (o *SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) SetCampaignNegativeKeyword(v SponsoredProductsDraftCampaignNegativeKeyword) {
	o.CampaignNegativeKeyword = &v
}

// GetCampaignNegativeKeywordId returns the CampaignNegativeKeywordId field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) GetCampaignNegativeKeywordId() string {
	if o == nil || IsNil(o.CampaignNegativeKeywordId) {
		var ret string
		return ret
	}
	return *o.CampaignNegativeKeywordId
}

// GetCampaignNegativeKeywordIdOk returns a tuple with the CampaignNegativeKeywordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) GetCampaignNegativeKeywordIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignNegativeKeywordId) {
		return nil, false
	}
	return o.CampaignNegativeKeywordId, true
}

// HasCampaignNegativeKeywordId returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) HasCampaignNegativeKeywordId() bool {
	if o != nil && !IsNil(o.CampaignNegativeKeywordId) {
		return true
	}

	return false
}

// SetCampaignNegativeKeywordId gets a reference to the given string and assigns it to the CampaignNegativeKeywordId field.
func (o *SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) SetCampaignNegativeKeywordId(v string) {
	o.CampaignNegativeKeywordId = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

func (o SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem struct {
	value *SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) Get() *SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) Set(val *SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem(val *SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) *NullableSponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem {
	return &NullableSponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
