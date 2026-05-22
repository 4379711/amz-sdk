package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsRecommendedTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsRecommendedTarget{}

// SponsoredProductsRecommendedTarget struct for SponsoredProductsRecommendedTarget
type SponsoredProductsRecommendedTarget struct {
	// The ID of an ad for which the targets are recommended
	AdId *string `json:"adId,omitempty"`
	// The ID of a campaign for which the targets are recommended
	CampaignId *string                      `json:"campaignId,omitempty"`
	TargetType *SponsoredProductsTargetType `json:"targetType,omitempty"`
	// The ASIN of the product being advertised
	AdAsin *string `json:"adAsin,omitempty"`
	// The keyword or ASIN that is being targeted
	RecommendedTarget *string `json:"recommendedTarget,omitempty"`
	// The ID of an ad group for which the targets are recommended
	AdGroupId *string `json:"adGroupId,omitempty"`
	// Provides a list of reasons for why this target is being recommended for harvesting
	RecommendationReasons []SponsoredProductsRecommendationReason `json:"recommendationReasons,omitempty"`
}

// NewSponsoredProductsRecommendedTarget instantiates a new SponsoredProductsRecommendedTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsRecommendedTarget() *SponsoredProductsRecommendedTarget {
	this := SponsoredProductsRecommendedTarget{}
	return &this
}

// NewSponsoredProductsRecommendedTargetWithDefaults instantiates a new SponsoredProductsRecommendedTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsRecommendedTargetWithDefaults() *SponsoredProductsRecommendedTarget {
	this := SponsoredProductsRecommendedTarget{}
	return &this
}

// GetAdId returns the AdId field value if set, zero value otherwise.
func (o *SponsoredProductsRecommendedTarget) GetAdId() string {
	if o == nil || IsNil(o.AdId) {
		var ret string
		return ret
	}
	return *o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsRecommendedTarget) GetAdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdId) {
		return nil, false
	}
	return o.AdId, true
}

// HasAdId returns a boolean if a field has been set.
func (o *SponsoredProductsRecommendedTarget) HasAdId() bool {
	if o != nil && !IsNil(o.AdId) {
		return true
	}

	return false
}

// SetAdId gets a reference to the given string and assigns it to the AdId field.
func (o *SponsoredProductsRecommendedTarget) SetAdId(v string) {
	o.AdId = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *SponsoredProductsRecommendedTarget) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsRecommendedTarget) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *SponsoredProductsRecommendedTarget) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *SponsoredProductsRecommendedTarget) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *SponsoredProductsRecommendedTarget) GetTargetType() SponsoredProductsTargetType {
	if o == nil || IsNil(o.TargetType) {
		var ret SponsoredProductsTargetType
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsRecommendedTarget) GetTargetTypeOk() (*SponsoredProductsTargetType, bool) {
	if o == nil || IsNil(o.TargetType) {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *SponsoredProductsRecommendedTarget) HasTargetType() bool {
	if o != nil && !IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given SponsoredProductsTargetType and assigns it to the TargetType field.
func (o *SponsoredProductsRecommendedTarget) SetTargetType(v SponsoredProductsTargetType) {
	o.TargetType = &v
}

// GetAdAsin returns the AdAsin field value if set, zero value otherwise.
func (o *SponsoredProductsRecommendedTarget) GetAdAsin() string {
	if o == nil || IsNil(o.AdAsin) {
		var ret string
		return ret
	}
	return *o.AdAsin
}

// GetAdAsinOk returns a tuple with the AdAsin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsRecommendedTarget) GetAdAsinOk() (*string, bool) {
	if o == nil || IsNil(o.AdAsin) {
		return nil, false
	}
	return o.AdAsin, true
}

// HasAdAsin returns a boolean if a field has been set.
func (o *SponsoredProductsRecommendedTarget) HasAdAsin() bool {
	if o != nil && !IsNil(o.AdAsin) {
		return true
	}

	return false
}

// SetAdAsin gets a reference to the given string and assigns it to the AdAsin field.
func (o *SponsoredProductsRecommendedTarget) SetAdAsin(v string) {
	o.AdAsin = &v
}

// GetRecommendedTarget returns the RecommendedTarget field value if set, zero value otherwise.
func (o *SponsoredProductsRecommendedTarget) GetRecommendedTarget() string {
	if o == nil || IsNil(o.RecommendedTarget) {
		var ret string
		return ret
	}
	return *o.RecommendedTarget
}

// GetRecommendedTargetOk returns a tuple with the RecommendedTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsRecommendedTarget) GetRecommendedTargetOk() (*string, bool) {
	if o == nil || IsNil(o.RecommendedTarget) {
		return nil, false
	}
	return o.RecommendedTarget, true
}

// HasRecommendedTarget returns a boolean if a field has been set.
func (o *SponsoredProductsRecommendedTarget) HasRecommendedTarget() bool {
	if o != nil && !IsNil(o.RecommendedTarget) {
		return true
	}

	return false
}

// SetRecommendedTarget gets a reference to the given string and assigns it to the RecommendedTarget field.
func (o *SponsoredProductsRecommendedTarget) SetRecommendedTarget(v string) {
	o.RecommendedTarget = &v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *SponsoredProductsRecommendedTarget) GetAdGroupId() string {
	if o == nil || IsNil(o.AdGroupId) {
		var ret string
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsRecommendedTarget) GetAdGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *SponsoredProductsRecommendedTarget) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given string and assigns it to the AdGroupId field.
func (o *SponsoredProductsRecommendedTarget) SetAdGroupId(v string) {
	o.AdGroupId = &v
}

// GetRecommendationReasons returns the RecommendationReasons field value if set, zero value otherwise.
func (o *SponsoredProductsRecommendedTarget) GetRecommendationReasons() []SponsoredProductsRecommendationReason {
	if o == nil || IsNil(o.RecommendationReasons) {
		var ret []SponsoredProductsRecommendationReason
		return ret
	}
	return o.RecommendationReasons
}

// GetRecommendationReasonsOk returns a tuple with the RecommendationReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsRecommendedTarget) GetRecommendationReasonsOk() ([]SponsoredProductsRecommendationReason, bool) {
	if o == nil || IsNil(o.RecommendationReasons) {
		return nil, false
	}
	return o.RecommendationReasons, true
}

// HasRecommendationReasons returns a boolean if a field has been set.
func (o *SponsoredProductsRecommendedTarget) HasRecommendationReasons() bool {
	if o != nil && !IsNil(o.RecommendationReasons) {
		return true
	}

	return false
}

// SetRecommendationReasons gets a reference to the given []SponsoredProductsRecommendationReason and assigns it to the RecommendationReasons field.
func (o *SponsoredProductsRecommendedTarget) SetRecommendationReasons(v []SponsoredProductsRecommendationReason) {
	o.RecommendationReasons = v
}

func (o SponsoredProductsRecommendedTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdId) {
		toSerialize["adId"] = o.AdId
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	if !IsNil(o.TargetType) {
		toSerialize["targetType"] = o.TargetType
	}
	if !IsNil(o.AdAsin) {
		toSerialize["adAsin"] = o.AdAsin
	}
	if !IsNil(o.RecommendedTarget) {
		toSerialize["recommendedTarget"] = o.RecommendedTarget
	}
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	if !IsNil(o.RecommendationReasons) {
		toSerialize["recommendationReasons"] = o.RecommendationReasons
	}
	return toSerialize, nil
}

type NullableSponsoredProductsRecommendedTarget struct {
	value *SponsoredProductsRecommendedTarget
	isSet bool
}

func (v NullableSponsoredProductsRecommendedTarget) Get() *SponsoredProductsRecommendedTarget {
	return v.value
}

func (v *NullableSponsoredProductsRecommendedTarget) Set(val *SponsoredProductsRecommendedTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsRecommendedTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsRecommendedTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsRecommendedTarget(val *SponsoredProductsRecommendedTarget) *NullableSponsoredProductsRecommendedTarget {
	return &NullableSponsoredProductsRecommendedTarget{value: val, isSet: true}
}

func (v NullableSponsoredProductsRecommendedTarget) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsRecommendedTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
