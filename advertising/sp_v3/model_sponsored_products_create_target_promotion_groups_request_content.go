package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateTargetPromotionGroupsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateTargetPromotionGroupsRequestContent{}

// SponsoredProductsCreateTargetPromotionGroupsRequestContent Request object for creating a Target Promotion Group.
type SponsoredProductsCreateTargetPromotionGroupsRequestContent struct {
	// The list of adIds (optional) of the Ad Group of the Auto-Targeting campaign, that will be part of the Target Promotion Group. If this     list is empty, all the Product Ads under the Ad Group will be part of the Target Promotion Group.
	AdIds                   []string                                  `json:"adIds,omitempty"`
	ExistingCampaignDetails *SponsoredProductsExistingCampaignDetails `json:"existingCampaignDetails,omitempty"`
	// The adGroupId of the Ad Group of an Auto-Targeting campaign that will be part of the Target Promotion Group.
	AdGroupId          string                               `json:"adGroupId"`
	NewCampaignDetails *SponsoredProductsNewCampaignDetails `json:"newCampaignDetails,omitempty"`
}

type _SponsoredProductsCreateTargetPromotionGroupsRequestContent SponsoredProductsCreateTargetPromotionGroupsRequestContent

// NewSponsoredProductsCreateTargetPromotionGroupsRequestContent instantiates a new SponsoredProductsCreateTargetPromotionGroupsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateTargetPromotionGroupsRequestContent(adGroupId string) *SponsoredProductsCreateTargetPromotionGroupsRequestContent {
	this := SponsoredProductsCreateTargetPromotionGroupsRequestContent{}
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsCreateTargetPromotionGroupsRequestContentWithDefaults instantiates a new SponsoredProductsCreateTargetPromotionGroupsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateTargetPromotionGroupsRequestContentWithDefaults() *SponsoredProductsCreateTargetPromotionGroupsRequestContent {
	this := SponsoredProductsCreateTargetPromotionGroupsRequestContent{}
	return &this
}

// GetAdIds returns the AdIds field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupsRequestContent) GetAdIds() []string {
	if o == nil || IsNil(o.AdIds) {
		var ret []string
		return ret
	}
	return o.AdIds
}

// GetAdIdsOk returns a tuple with the AdIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupsRequestContent) GetAdIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AdIds) {
		return nil, false
	}
	return o.AdIds, true
}

// HasAdIds returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupsRequestContent) HasAdIds() bool {
	if o != nil && !IsNil(o.AdIds) {
		return true
	}

	return false
}

// SetAdIds gets a reference to the given []string and assigns it to the AdIds field.
func (o *SponsoredProductsCreateTargetPromotionGroupsRequestContent) SetAdIds(v []string) {
	o.AdIds = v
}

// GetExistingCampaignDetails returns the ExistingCampaignDetails field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupsRequestContent) GetExistingCampaignDetails() SponsoredProductsExistingCampaignDetails {
	if o == nil || IsNil(o.ExistingCampaignDetails) {
		var ret SponsoredProductsExistingCampaignDetails
		return ret
	}
	return *o.ExistingCampaignDetails
}

// GetExistingCampaignDetailsOk returns a tuple with the ExistingCampaignDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupsRequestContent) GetExistingCampaignDetailsOk() (*SponsoredProductsExistingCampaignDetails, bool) {
	if o == nil || IsNil(o.ExistingCampaignDetails) {
		return nil, false
	}
	return o.ExistingCampaignDetails, true
}

// HasExistingCampaignDetails returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupsRequestContent) HasExistingCampaignDetails() bool {
	if o != nil && !IsNil(o.ExistingCampaignDetails) {
		return true
	}

	return false
}

// SetExistingCampaignDetails gets a reference to the given SponsoredProductsExistingCampaignDetails and assigns it to the ExistingCampaignDetails field.
func (o *SponsoredProductsCreateTargetPromotionGroupsRequestContent) SetExistingCampaignDetails(v SponsoredProductsExistingCampaignDetails) {
	o.ExistingCampaignDetails = &v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsCreateTargetPromotionGroupsRequestContent) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupsRequestContent) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsCreateTargetPromotionGroupsRequestContent) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetNewCampaignDetails returns the NewCampaignDetails field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupsRequestContent) GetNewCampaignDetails() SponsoredProductsNewCampaignDetails {
	if o == nil || IsNil(o.NewCampaignDetails) {
		var ret SponsoredProductsNewCampaignDetails
		return ret
	}
	return *o.NewCampaignDetails
}

// GetNewCampaignDetailsOk returns a tuple with the NewCampaignDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupsRequestContent) GetNewCampaignDetailsOk() (*SponsoredProductsNewCampaignDetails, bool) {
	if o == nil || IsNil(o.NewCampaignDetails) {
		return nil, false
	}
	return o.NewCampaignDetails, true
}

// HasNewCampaignDetails returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupsRequestContent) HasNewCampaignDetails() bool {
	if o != nil && !IsNil(o.NewCampaignDetails) {
		return true
	}

	return false
}

// SetNewCampaignDetails gets a reference to the given SponsoredProductsNewCampaignDetails and assigns it to the NewCampaignDetails field.
func (o *SponsoredProductsCreateTargetPromotionGroupsRequestContent) SetNewCampaignDetails(v SponsoredProductsNewCampaignDetails) {
	o.NewCampaignDetails = &v
}

func (o SponsoredProductsCreateTargetPromotionGroupsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdIds) {
		toSerialize["adIds"] = o.AdIds
	}
	if !IsNil(o.ExistingCampaignDetails) {
		toSerialize["existingCampaignDetails"] = o.ExistingCampaignDetails
	}
	toSerialize["adGroupId"] = o.AdGroupId
	if !IsNil(o.NewCampaignDetails) {
		toSerialize["newCampaignDetails"] = o.NewCampaignDetails
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateTargetPromotionGroupsRequestContent struct {
	value *SponsoredProductsCreateTargetPromotionGroupsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupsRequestContent) Get() *SponsoredProductsCreateTargetPromotionGroupsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupsRequestContent) Set(val *SponsoredProductsCreateTargetPromotionGroupsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateTargetPromotionGroupsRequestContent(val *SponsoredProductsCreateTargetPromotionGroupsRequestContent) *NullableSponsoredProductsCreateTargetPromotionGroupsRequestContent {
	return &NullableSponsoredProductsCreateTargetPromotionGroupsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
