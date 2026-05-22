package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem{}

// SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem struct for SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem
type SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem struct {
	CampaignNegativeKeyword *SponsoredProductsGlobalCampaignNegativeKeyword `json:"campaignNegativeKeyword,omitempty"`
	// the campaignNegativeKeyword ID
	CampaignNegativeKeywordId *string `json:"campaignNegativeKeywordId,omitempty"`
	// the index of the campaign in the array from the request body
	Index int32 `json:"index"`
}

type _SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem

// NewSponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem instantiates a new SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem(index int32) *SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem {
	this := SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItemWithDefaults instantiates a new SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItemWithDefaults() *SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem {
	this := SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem{}
	return &this
}

// GetCampaignNegativeKeyword returns the CampaignNegativeKeyword field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) GetCampaignNegativeKeyword() SponsoredProductsGlobalCampaignNegativeKeyword {
	if o == nil || IsNil(o.CampaignNegativeKeyword) {
		var ret SponsoredProductsGlobalCampaignNegativeKeyword
		return ret
	}
	return *o.CampaignNegativeKeyword
}

// GetCampaignNegativeKeywordOk returns a tuple with the CampaignNegativeKeyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) GetCampaignNegativeKeywordOk() (*SponsoredProductsGlobalCampaignNegativeKeyword, bool) {
	if o == nil || IsNil(o.CampaignNegativeKeyword) {
		return nil, false
	}
	return o.CampaignNegativeKeyword, true
}

// HasCampaignNegativeKeyword returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) HasCampaignNegativeKeyword() bool {
	if o != nil && !IsNil(o.CampaignNegativeKeyword) {
		return true
	}

	return false
}

// SetCampaignNegativeKeyword gets a reference to the given SponsoredProductsGlobalCampaignNegativeKeyword and assigns it to the CampaignNegativeKeyword field.
func (o *SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) SetCampaignNegativeKeyword(v SponsoredProductsGlobalCampaignNegativeKeyword) {
	o.CampaignNegativeKeyword = &v
}

// GetCampaignNegativeKeywordId returns the CampaignNegativeKeywordId field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) GetCampaignNegativeKeywordId() string {
	if o == nil || IsNil(o.CampaignNegativeKeywordId) {
		var ret string
		return ret
	}
	return *o.CampaignNegativeKeywordId
}

// GetCampaignNegativeKeywordIdOk returns a tuple with the CampaignNegativeKeywordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) GetCampaignNegativeKeywordIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignNegativeKeywordId) {
		return nil, false
	}
	return o.CampaignNegativeKeywordId, true
}

// HasCampaignNegativeKeywordId returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) HasCampaignNegativeKeywordId() bool {
	if o != nil && !IsNil(o.CampaignNegativeKeywordId) {
		return true
	}

	return false
}

// SetCampaignNegativeKeywordId gets a reference to the given string and assigns it to the CampaignNegativeKeywordId field.
func (o *SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) SetCampaignNegativeKeywordId(v string) {
	o.CampaignNegativeKeywordId = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

func (o SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem struct {
	value *SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) Get() *SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) Set(val *SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem(val *SponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) *NullableSponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem {
	return &NullableSponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalCampaignNegativeKeywordSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
