package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent{}

// SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent struct for SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent
type SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent struct {
	CampaignIdFilter *SponsoredProductsObjectIdFilter `json:"campaignIdFilter,omitempty"`
	ClientId         *string                          `json:"clientId,omitempty"`
	ApiScope         *string                          `json:"apiScope,omitempty"`
	MaxResults       *int32                           `json:"maxResults,omitempty"`
	// Token for fetching different pages in the response
	NextToken         *string                            `json:"nextToken,omitempty"`
	AdIdFilter        *SponsoredProductsObjectIdFilter   `json:"adIdFilter,omitempty"`
	AdGroupIdFilter   *SponsoredProductsObjectIdFilter   `json:"adGroupIdFilter,omitempty"`
	ApiGatewayContext SponsoredProductsApiGatewayContext `json:"apiGatewayContext"`
}

type _SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent

// NewSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent instantiates a new SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent(apiGatewayContext SponsoredProductsApiGatewayContext) *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent {
	this := SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent{}
	this.ApiGatewayContext = apiGatewayContext
	return &this
}

// NewSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContentWithDefaults instantiates a new SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContentWithDefaults() *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent {
	this := SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) GetCampaignIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) SetCampaignIdFilter(v SponsoredProductsObjectIdFilter) {
	o.CampaignIdFilter = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) SetClientId(v string) {
	o.ClientId = &v
}

// GetApiScope returns the ApiScope field value if set, zero value otherwise.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) GetApiScope() string {
	if o == nil || IsNil(o.ApiScope) {
		var ret string
		return ret
	}
	return *o.ApiScope
}

// GetApiScopeOk returns a tuple with the ApiScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) GetApiScopeOk() (*string, bool) {
	if o == nil || IsNil(o.ApiScope) {
		return nil, false
	}
	return o.ApiScope, true
}

// HasApiScope returns a boolean if a field has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) HasApiScope() bool {
	if o != nil && !IsNil(o.ApiScope) {
		return true
	}

	return false
}

// SetApiScope gets a reference to the given string and assigns it to the ApiScope field.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) SetApiScope(v string) {
	o.ApiScope = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetAdIdFilter returns the AdIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) GetAdIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.AdIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.AdIdFilter
}

// GetAdIdFilterOk returns a tuple with the AdIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) GetAdIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdIdFilter) {
		return nil, false
	}
	return o.AdIdFilter, true
}

// HasAdIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) HasAdIdFilter() bool {
	if o != nil && !IsNil(o.AdIdFilter) {
		return true
	}

	return false
}

// SetAdIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the AdIdFilter field.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) SetAdIdFilter(v SponsoredProductsObjectIdFilter) {
	o.AdIdFilter = &v
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) GetAdGroupIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) GetAdGroupIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) SetAdGroupIdFilter(v SponsoredProductsObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

// GetApiGatewayContext returns the ApiGatewayContext field value
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) GetApiGatewayContext() SponsoredProductsApiGatewayContext {
	if o == nil {
		var ret SponsoredProductsApiGatewayContext
		return ret
	}

	return o.ApiGatewayContext
}

// GetApiGatewayContextOk returns a tuple with the ApiGatewayContext field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) GetApiGatewayContextOk() (*SponsoredProductsApiGatewayContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiGatewayContext, true
}

// SetApiGatewayContext sets field value
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) SetApiGatewayContext(v SponsoredProductsApiGatewayContext) {
	o.ApiGatewayContext = v
}

func (o SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignIdFilter) {
		toSerialize["campaignIdFilter"] = o.CampaignIdFilter
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.ApiScope) {
		toSerialize["apiScope"] = o.ApiScope
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
	toSerialize["apiGatewayContext"] = o.ApiGatewayContext
	return toSerialize, nil
}

type NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent struct {
	value *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent
	isSet bool
}

func (v NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) Get() *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) Set(val *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent(val *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) *NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent {
	return &NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
