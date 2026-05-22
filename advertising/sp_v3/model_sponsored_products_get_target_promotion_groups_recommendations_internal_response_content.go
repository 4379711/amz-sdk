package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent{}

// SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent struct for SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent
type SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent struct {
	// Total number of records available
	TotalResults int32 `json:"totalResults"`
	// Token for fetching the next page
	NextToken *string `json:"nextToken,omitempty"`
	// List of optimized targets for the request, as recommended by Amazon heuristics
	Targets []SponsoredProductsRecommendedTarget `json:"targets"`
}

type _SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent

// NewSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent instantiates a new SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent(totalResults int32, targets []SponsoredProductsRecommendedTarget) *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent {
	this := SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent{}
	this.TotalResults = totalResults
	this.Targets = targets
	return &this
}

// NewSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContentWithDefaults instantiates a new SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContentWithDefaults() *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent {
	this := SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent) GetTotalResults() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent) GetTotalResultsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalResults, true
}

// SetTotalResults sets field value
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent) SetTotalResults(v int32) {
	o.TotalResults = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTargets returns the Targets field value
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent) GetTargets() []SponsoredProductsRecommendedTarget {
	if o == nil {
		var ret []SponsoredProductsRecommendedTarget
		return ret
	}

	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent) GetTargetsOk() ([]SponsoredProductsRecommendedTarget, bool) {
	if o == nil {
		return nil, false
	}
	return o.Targets, true
}

// SetTargets sets field value
func (o *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent) SetTargets(v []SponsoredProductsRecommendedTarget) {
	o.Targets = v
}

func (o SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["totalResults"] = o.TotalResults
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	toSerialize["targets"] = o.Targets
	return toSerialize, nil
}

type NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent struct {
	value *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent
	isSet bool
}

func (v NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent) Get() *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent) Set(val *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent(val *SponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent) *NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent {
	return &NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGetTargetPromotionGroupsRecommendationsInternalResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
