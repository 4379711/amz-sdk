package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListTargetPromotionGroupTargetsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListTargetPromotionGroupTargetsRequestContent{}

// SponsoredProductsListTargetPromotionGroupTargetsRequestContent Request object for querying target promotion group targets.
type SponsoredProductsListTargetPromotionGroupTargetsRequestContent struct {
	// The maximum number of results requested.
	MaxResults *int32 `json:"maxResults,omitempty"`
	// Token value allowing to navigate to the next or previous response page
	NextToken                    *string                          `json:"nextToken,omitempty"`
	TargetPromotionGroupIdFilter *SponsoredProductsObjectIdFilter `json:"targetPromotionGroupIdFilter,omitempty"`
	AdGroupIdFilter              *SponsoredProductsObjectIdFilter `json:"adGroupIdFilter,omitempty"`
}

// NewSponsoredProductsListTargetPromotionGroupTargetsRequestContent instantiates a new SponsoredProductsListTargetPromotionGroupTargetsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListTargetPromotionGroupTargetsRequestContent() *SponsoredProductsListTargetPromotionGroupTargetsRequestContent {
	this := SponsoredProductsListTargetPromotionGroupTargetsRequestContent{}
	return &this
}

// NewSponsoredProductsListTargetPromotionGroupTargetsRequestContentWithDefaults instantiates a new SponsoredProductsListTargetPromotionGroupTargetsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListTargetPromotionGroupTargetsRequestContentWithDefaults() *SponsoredProductsListTargetPromotionGroupTargetsRequestContent {
	this := SponsoredProductsListTargetPromotionGroupTargetsRequestContent{}
	return &this
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListTargetPromotionGroupTargetsRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListTargetPromotionGroupTargetsRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListTargetPromotionGroupTargetsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListTargetPromotionGroupTargetsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTargetPromotionGroupIdFilter returns the TargetPromotionGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListTargetPromotionGroupTargetsRequestContent) GetTargetPromotionGroupIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.TargetPromotionGroupIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.TargetPromotionGroupIdFilter
}

// GetTargetPromotionGroupIdFilterOk returns a tuple with the TargetPromotionGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsRequestContent) GetTargetPromotionGroupIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.TargetPromotionGroupIdFilter) {
		return nil, false
	}
	return o.TargetPromotionGroupIdFilter, true
}

// HasTargetPromotionGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsRequestContent) HasTargetPromotionGroupIdFilter() bool {
	if o != nil && !IsNil(o.TargetPromotionGroupIdFilter) {
		return true
	}

	return false
}

// SetTargetPromotionGroupIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the TargetPromotionGroupIdFilter field.
func (o *SponsoredProductsListTargetPromotionGroupTargetsRequestContent) SetTargetPromotionGroupIdFilter(v SponsoredProductsObjectIdFilter) {
	o.TargetPromotionGroupIdFilter = &v
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListTargetPromotionGroupTargetsRequestContent) GetAdGroupIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsRequestContent) GetAdGroupIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *SponsoredProductsListTargetPromotionGroupTargetsRequestContent) SetAdGroupIdFilter(v SponsoredProductsObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

func (o SponsoredProductsListTargetPromotionGroupTargetsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.TargetPromotionGroupIdFilter) {
		toSerialize["targetPromotionGroupIdFilter"] = o.TargetPromotionGroupIdFilter
	}
	if !IsNil(o.AdGroupIdFilter) {
		toSerialize["adGroupIdFilter"] = o.AdGroupIdFilter
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListTargetPromotionGroupTargetsRequestContent struct {
	value *SponsoredProductsListTargetPromotionGroupTargetsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListTargetPromotionGroupTargetsRequestContent) Get() *SponsoredProductsListTargetPromotionGroupTargetsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListTargetPromotionGroupTargetsRequestContent) Set(val *SponsoredProductsListTargetPromotionGroupTargetsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListTargetPromotionGroupTargetsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListTargetPromotionGroupTargetsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListTargetPromotionGroupTargetsRequestContent(val *SponsoredProductsListTargetPromotionGroupTargetsRequestContent) *NullableSponsoredProductsListTargetPromotionGroupTargetsRequestContent {
	return &NullableSponsoredProductsListTargetPromotionGroupTargetsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListTargetPromotionGroupTargetsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListTargetPromotionGroupTargetsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
