package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent{}

// SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent struct for SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent
type SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent struct {
	AdIds                   []string                                  `json:"adIds,omitempty"`
	ExistingCampaignDetails *SponsoredProductsExistingCampaignDetails `json:"existingCampaignDetails,omitempty"`
	ApiGatewayContext       SponsoredProductsApiGatewayContext        `json:"apiGatewayContext"`
	// Entity object identifier
	AdGroupId          string                               `json:"adGroupId"`
	NewCampaignDetails *SponsoredProductsNewCampaignDetails `json:"newCampaignDetails,omitempty"`
}

type _SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent

// NewSponsoredProductsCreateTargetPromotionGroupsInternalRequestContent instantiates a new SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateTargetPromotionGroupsInternalRequestContent(apiGatewayContext SponsoredProductsApiGatewayContext, adGroupId string) *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent {
	this := SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent{}
	this.ApiGatewayContext = apiGatewayContext
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsCreateTargetPromotionGroupsInternalRequestContentWithDefaults instantiates a new SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateTargetPromotionGroupsInternalRequestContentWithDefaults() *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent {
	this := SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent{}
	return &this
}

// GetAdIds returns the AdIds field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) GetAdIds() []string {
	if o == nil || IsNil(o.AdIds) {
		var ret []string
		return ret
	}
	return o.AdIds
}

// GetAdIdsOk returns a tuple with the AdIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) GetAdIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AdIds) {
		return nil, false
	}
	return o.AdIds, true
}

// HasAdIds returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) HasAdIds() bool {
	if o != nil && !IsNil(o.AdIds) {
		return true
	}

	return false
}

// SetAdIds gets a reference to the given []string and assigns it to the AdIds field.
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) SetAdIds(v []string) {
	o.AdIds = v
}

// GetExistingCampaignDetails returns the ExistingCampaignDetails field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) GetExistingCampaignDetails() SponsoredProductsExistingCampaignDetails {
	if o == nil || IsNil(o.ExistingCampaignDetails) {
		var ret SponsoredProductsExistingCampaignDetails
		return ret
	}
	return *o.ExistingCampaignDetails
}

// GetExistingCampaignDetailsOk returns a tuple with the ExistingCampaignDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) GetExistingCampaignDetailsOk() (*SponsoredProductsExistingCampaignDetails, bool) {
	if o == nil || IsNil(o.ExistingCampaignDetails) {
		return nil, false
	}
	return o.ExistingCampaignDetails, true
}

// HasExistingCampaignDetails returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) HasExistingCampaignDetails() bool {
	if o != nil && !IsNil(o.ExistingCampaignDetails) {
		return true
	}

	return false
}

// SetExistingCampaignDetails gets a reference to the given SponsoredProductsExistingCampaignDetails and assigns it to the ExistingCampaignDetails field.
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) SetExistingCampaignDetails(v SponsoredProductsExistingCampaignDetails) {
	o.ExistingCampaignDetails = &v
}

// GetApiGatewayContext returns the ApiGatewayContext field value
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) GetApiGatewayContext() SponsoredProductsApiGatewayContext {
	if o == nil {
		var ret SponsoredProductsApiGatewayContext
		return ret
	}

	return o.ApiGatewayContext
}

// GetApiGatewayContextOk returns a tuple with the ApiGatewayContext field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) GetApiGatewayContextOk() (*SponsoredProductsApiGatewayContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiGatewayContext, true
}

// SetApiGatewayContext sets field value
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) SetApiGatewayContext(v SponsoredProductsApiGatewayContext) {
	o.ApiGatewayContext = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetNewCampaignDetails returns the NewCampaignDetails field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) GetNewCampaignDetails() SponsoredProductsNewCampaignDetails {
	if o == nil || IsNil(o.NewCampaignDetails) {
		var ret SponsoredProductsNewCampaignDetails
		return ret
	}
	return *o.NewCampaignDetails
}

// GetNewCampaignDetailsOk returns a tuple with the NewCampaignDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) GetNewCampaignDetailsOk() (*SponsoredProductsNewCampaignDetails, bool) {
	if o == nil || IsNil(o.NewCampaignDetails) {
		return nil, false
	}
	return o.NewCampaignDetails, true
}

// HasNewCampaignDetails returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) HasNewCampaignDetails() bool {
	if o != nil && !IsNil(o.NewCampaignDetails) {
		return true
	}

	return false
}

// SetNewCampaignDetails gets a reference to the given SponsoredProductsNewCampaignDetails and assigns it to the NewCampaignDetails field.
func (o *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) SetNewCampaignDetails(v SponsoredProductsNewCampaignDetails) {
	o.NewCampaignDetails = &v
}

func (o SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdIds) {
		toSerialize["adIds"] = o.AdIds
	}
	if !IsNil(o.ExistingCampaignDetails) {
		toSerialize["existingCampaignDetails"] = o.ExistingCampaignDetails
	}
	toSerialize["apiGatewayContext"] = o.ApiGatewayContext
	toSerialize["adGroupId"] = o.AdGroupId
	if !IsNil(o.NewCampaignDetails) {
		toSerialize["newCampaignDetails"] = o.NewCampaignDetails
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateTargetPromotionGroupsInternalRequestContent struct {
	value *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) Get() *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) Set(val *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateTargetPromotionGroupsInternalRequestContent(val *SponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) *NullableSponsoredProductsCreateTargetPromotionGroupsInternalRequestContent {
	return &NullableSponsoredProductsCreateTargetPromotionGroupsInternalRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupsInternalRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
