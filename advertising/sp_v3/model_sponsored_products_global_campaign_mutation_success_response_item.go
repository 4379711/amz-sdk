package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalCampaignMutationSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalCampaignMutationSuccessResponseItem{}

// SponsoredProductsGlobalCampaignMutationSuccessResponseItem struct for SponsoredProductsGlobalCampaignMutationSuccessResponseItem
type SponsoredProductsGlobalCampaignMutationSuccessResponseItem struct {
	// the campaign ID
	CampaignId *string `json:"campaignId,omitempty"`
	// the index of the campaign in the array from the request body
	Index    int32                            `json:"index"`
	Campaign *SponsoredProductsGlobalCampaign `json:"campaign,omitempty"`
}

type _SponsoredProductsGlobalCampaignMutationSuccessResponseItem SponsoredProductsGlobalCampaignMutationSuccessResponseItem

// NewSponsoredProductsGlobalCampaignMutationSuccessResponseItem instantiates a new SponsoredProductsGlobalCampaignMutationSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalCampaignMutationSuccessResponseItem(index int32) *SponsoredProductsGlobalCampaignMutationSuccessResponseItem {
	this := SponsoredProductsGlobalCampaignMutationSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsGlobalCampaignMutationSuccessResponseItemWithDefaults instantiates a new SponsoredProductsGlobalCampaignMutationSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalCampaignMutationSuccessResponseItemWithDefaults() *SponsoredProductsGlobalCampaignMutationSuccessResponseItem {
	this := SponsoredProductsGlobalCampaignMutationSuccessResponseItem{}
	return &this
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalCampaignMutationSuccessResponseItem) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignMutationSuccessResponseItem) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalCampaignMutationSuccessResponseItem) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *SponsoredProductsGlobalCampaignMutationSuccessResponseItem) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsGlobalCampaignMutationSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignMutationSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsGlobalCampaignMutationSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetCampaign returns the Campaign field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalCampaignMutationSuccessResponseItem) GetCampaign() SponsoredProductsGlobalCampaign {
	if o == nil || IsNil(o.Campaign) {
		var ret SponsoredProductsGlobalCampaign
		return ret
	}
	return *o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignMutationSuccessResponseItem) GetCampaignOk() (*SponsoredProductsGlobalCampaign, bool) {
	if o == nil || IsNil(o.Campaign) {
		return nil, false
	}
	return o.Campaign, true
}

// HasCampaign returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalCampaignMutationSuccessResponseItem) HasCampaign() bool {
	if o != nil && !IsNil(o.Campaign) {
		return true
	}

	return false
}

// SetCampaign gets a reference to the given SponsoredProductsGlobalCampaign and assigns it to the Campaign field.
func (o *SponsoredProductsGlobalCampaignMutationSuccessResponseItem) SetCampaign(v SponsoredProductsGlobalCampaign) {
	o.Campaign = &v
}

func (o SponsoredProductsGlobalCampaignMutationSuccessResponseItem) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsGlobalCampaignMutationSuccessResponseItem struct {
	value *SponsoredProductsGlobalCampaignMutationSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsGlobalCampaignMutationSuccessResponseItem) Get() *SponsoredProductsGlobalCampaignMutationSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsGlobalCampaignMutationSuccessResponseItem) Set(val *SponsoredProductsGlobalCampaignMutationSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalCampaignMutationSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalCampaignMutationSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalCampaignMutationSuccessResponseItem(val *SponsoredProductsGlobalCampaignMutationSuccessResponseItem) *NullableSponsoredProductsGlobalCampaignMutationSuccessResponseItem {
	return &NullableSponsoredProductsGlobalCampaignMutationSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalCampaignMutationSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalCampaignMutationSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
