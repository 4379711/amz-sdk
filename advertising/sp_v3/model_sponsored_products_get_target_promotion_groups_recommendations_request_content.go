package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent{}

// SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent struct for SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent
type SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent struct {
	CampaignIdFilter *SponsoredProductsObjectIdFilter `json:"campaignIdFilter,omitempty"`
	// Number of records to include in the paginated response. Defaults to 1000.
	MaxResults *int32 `json:"maxResults,omitempty"`
	// Token for fetching the next page
	NextToken       *string                          `json:"nextToken,omitempty"`
	AdIdFilter      *SponsoredProductsObjectIdFilter `json:"adIdFilter,omitempty"`
	AdGroupIdFilter *SponsoredProductsObjectIdFilter `json:"adGroupIdFilter,omitempty"`
}

// NewSponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent instantiates a new SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent() *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent {
	this := SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent{}
	return &this
}

// NewSponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContentWithDefaults instantiates a new SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContentWithDefaults() *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent {
	this := SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) GetCampaignIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) SetCampaignIdFilter(v SponsoredProductsObjectIdFilter) {
	o.CampaignIdFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetAdIdFilter returns the AdIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) GetAdIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.AdIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.AdIdFilter
}

// GetAdIdFilterOk returns a tuple with the AdIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) GetAdIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdIdFilter) {
		return nil, false
	}
	return o.AdIdFilter, true
}

// HasAdIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) HasAdIdFilter() bool {
	if o != nil && !IsNil(o.AdIdFilter) {
		return true
	}

	return false
}

// SetAdIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the AdIdFilter field.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) SetAdIdFilter(v SponsoredProductsObjectIdFilter) {
	o.AdIdFilter = &v
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) GetAdGroupIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) GetAdGroupIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) SetAdGroupIdFilter(v SponsoredProductsObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

func (o SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignIdFilter) {
		toSerialize["campaignIdFilter"] = o.CampaignIdFilter
	}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.AdIdFilter) {
		toSerialize["adIdFilter"] = o.AdIdFilter
	}
	if !IsNil(o.AdGroupIdFilter) {
		toSerialize["adGroupIdFilter"] = o.AdGroupIdFilter
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent struct {
	value *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) Get() *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) Set(val *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent(val *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) *NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent {
	return &NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
