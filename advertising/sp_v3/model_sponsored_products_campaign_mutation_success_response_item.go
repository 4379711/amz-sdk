package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignMutationSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignMutationSuccessResponseItem{}

// SponsoredProductsCampaignMutationSuccessResponseItem struct for SponsoredProductsCampaignMutationSuccessResponseItem
type SponsoredProductsCampaignMutationSuccessResponseItem struct {
	// the campaign ID
	CampaignId *string `json:"campaignId,omitempty"`
	// the index of the campaign in the array from the request body
	Index    int32                      `json:"index"`
	Campaign *SponsoredProductsCampaign `json:"campaign,omitempty"`
}

type _SponsoredProductsCampaignMutationSuccessResponseItem SponsoredProductsCampaignMutationSuccessResponseItem

// NewSponsoredProductsCampaignMutationSuccessResponseItem instantiates a new SponsoredProductsCampaignMutationSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignMutationSuccessResponseItem(index int32) *SponsoredProductsCampaignMutationSuccessResponseItem {
	this := SponsoredProductsCampaignMutationSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsCampaignMutationSuccessResponseItemWithDefaults instantiates a new SponsoredProductsCampaignMutationSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignMutationSuccessResponseItemWithDefaults() *SponsoredProductsCampaignMutationSuccessResponseItem {
	this := SponsoredProductsCampaignMutationSuccessResponseItem{}
	return &this
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationSuccessResponseItem) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationSuccessResponseItem) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationSuccessResponseItem) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *SponsoredProductsCampaignMutationSuccessResponseItem) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsCampaignMutationSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsCampaignMutationSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetCampaign returns the Campaign field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationSuccessResponseItem) GetCampaign() SponsoredProductsCampaign {
	if o == nil || IsNil(o.Campaign) {
		var ret SponsoredProductsCampaign
		return ret
	}
	return *o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationSuccessResponseItem) GetCampaignOk() (*SponsoredProductsCampaign, bool) {
	if o == nil || IsNil(o.Campaign) {
		return nil, false
	}
	return o.Campaign, true
}

// HasCampaign returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationSuccessResponseItem) HasCampaign() bool {
	if o != nil && !IsNil(o.Campaign) {
		return true
	}

	return false
}

// SetCampaign gets a reference to the given SponsoredProductsCampaign and assigns it to the Campaign field.
func (o *SponsoredProductsCampaignMutationSuccessResponseItem) SetCampaign(v SponsoredProductsCampaign) {
	o.Campaign = &v
}

func (o SponsoredProductsCampaignMutationSuccessResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	toSerialize["index"] = o.Index
	if !IsNil(o.Campaign) {
		toSerialize["campaign"] = o.Campaign
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignMutationSuccessResponseItem struct {
	value *SponsoredProductsCampaignMutationSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsCampaignMutationSuccessResponseItem) Get() *SponsoredProductsCampaignMutationSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsCampaignMutationSuccessResponseItem) Set(val *SponsoredProductsCampaignMutationSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignMutationSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignMutationSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignMutationSuccessResponseItem(val *SponsoredProductsCampaignMutationSuccessResponseItem) *NullableSponsoredProductsCampaignMutationSuccessResponseItem {
	return &NullableSponsoredProductsCampaignMutationSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignMutationSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignMutationSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
