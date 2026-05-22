package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignMutationSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignMutationSuccessResponseItem{}

// SponsoredProductsDraftCampaignMutationSuccessResponseItem struct for SponsoredProductsDraftCampaignMutationSuccessResponseItem
type SponsoredProductsDraftCampaignMutationSuccessResponseItem struct {
	// the draft ID
	CampaignId *string `json:"campaignId,omitempty"`
	// the index of the draft in the array from the request body
	Index    int32                           `json:"index"`
	Campaign *SponsoredProductsDraftCampaign `json:"campaign,omitempty"`
}

type _SponsoredProductsDraftCampaignMutationSuccessResponseItem SponsoredProductsDraftCampaignMutationSuccessResponseItem

// NewSponsoredProductsDraftCampaignMutationSuccessResponseItem instantiates a new SponsoredProductsDraftCampaignMutationSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignMutationSuccessResponseItem(index int32) *SponsoredProductsDraftCampaignMutationSuccessResponseItem {
	this := SponsoredProductsDraftCampaignMutationSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsDraftCampaignMutationSuccessResponseItemWithDefaults instantiates a new SponsoredProductsDraftCampaignMutationSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignMutationSuccessResponseItemWithDefaults() *SponsoredProductsDraftCampaignMutationSuccessResponseItem {
	this := SponsoredProductsDraftCampaignMutationSuccessResponseItem{}
	return &this
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignMutationSuccessResponseItem) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignMutationSuccessResponseItem) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignMutationSuccessResponseItem) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *SponsoredProductsDraftCampaignMutationSuccessResponseItem) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsDraftCampaignMutationSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignMutationSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsDraftCampaignMutationSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetCampaign returns the Campaign field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignMutationSuccessResponseItem) GetCampaign() SponsoredProductsDraftCampaign {
	if o == nil || IsNil(o.Campaign) {
		var ret SponsoredProductsDraftCampaign
		return ret
	}
	return *o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignMutationSuccessResponseItem) GetCampaignOk() (*SponsoredProductsDraftCampaign, bool) {
	if o == nil || IsNil(o.Campaign) {
		return nil, false
	}
	return o.Campaign, true
}

// HasCampaign returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignMutationSuccessResponseItem) HasCampaign() bool {
	if o != nil && !IsNil(o.Campaign) {
		return true
	}

	return false
}

// SetCampaign gets a reference to the given SponsoredProductsDraftCampaign and assigns it to the Campaign field.
func (o *SponsoredProductsDraftCampaignMutationSuccessResponseItem) SetCampaign(v SponsoredProductsDraftCampaign) {
	o.Campaign = &v
}

func (o SponsoredProductsDraftCampaignMutationSuccessResponseItem) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsDraftCampaignMutationSuccessResponseItem struct {
	value *SponsoredProductsDraftCampaignMutationSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignMutationSuccessResponseItem) Get() *SponsoredProductsDraftCampaignMutationSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignMutationSuccessResponseItem) Set(val *SponsoredProductsDraftCampaignMutationSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignMutationSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignMutationSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignMutationSuccessResponseItem(val *SponsoredProductsDraftCampaignMutationSuccessResponseItem) *NullableSponsoredProductsDraftCampaignMutationSuccessResponseItem {
	return &NullableSponsoredProductsDraftCampaignMutationSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignMutationSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignMutationSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
