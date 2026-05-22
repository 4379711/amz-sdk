package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListTargetPromotionGroupsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListTargetPromotionGroupsRequestContent{}

// SponsoredProductsListTargetPromotionGroupsRequestContent Request object for querying target promotion groups.
type SponsoredProductsListTargetPromotionGroupsRequestContent struct {
	// The maximum number of results requested.
	MaxResults *int32 `json:"maxResults,omitempty"`
	// Token value allowing to navigate to the next or previous response page
	NextToken                    *string                          `json:"nextToken,omitempty"`
	TargetPromotionGroupIdFilter *SponsoredProductsObjectIdFilter `json:"targetPromotionGroupIdFilter,omitempty"`
	AdGroupIdFilter              *SponsoredProductsObjectIdFilter `json:"adGroupIdFilter,omitempty"`
}

// NewSponsoredProductsListTargetPromotionGroupsRequestContent instantiates a new SponsoredProductsListTargetPromotionGroupsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListTargetPromotionGroupsRequestContent() *SponsoredProductsListTargetPromotionGroupsRequestContent {
	this := SponsoredProductsListTargetPromotionGroupsRequestContent{}
	return &this
}

// NewSponsoredProductsListTargetPromotionGroupsRequestContentWithDefaults instantiates a new SponsoredProductsListTargetPromotionGroupsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListTargetPromotionGroupsRequestContentWithDefaults() *SponsoredProductsListTargetPromotionGroupsRequestContent {
	this := SponsoredProductsListTargetPromotionGroupsRequestContent{}
	return &this
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListTargetPromotionGroupsRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListTargetPromotionGroupsRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListTargetPromotionGroupsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListTargetPromotionGroupsRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListTargetPromotionGroupsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListTargetPromotionGroupsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListTargetPromotionGroupsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListTargetPromotionGroupsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTargetPromotionGroupIdFilter returns the TargetPromotionGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListTargetPromotionGroupsRequestContent) GetTargetPromotionGroupIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.TargetPromotionGroupIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.TargetPromotionGroupIdFilter
}

// GetTargetPromotionGroupIdFilterOk returns a tuple with the TargetPromotionGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListTargetPromotionGroupsRequestContent) GetTargetPromotionGroupIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.TargetPromotionGroupIdFilter) {
		return nil, false
	}
	return o.TargetPromotionGroupIdFilter, true
}

// HasTargetPromotionGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListTargetPromotionGroupsRequestContent) HasTargetPromotionGroupIdFilter() bool {
	if o != nil && !IsNil(o.TargetPromotionGroupIdFilter) {
		return true
	}

	return false
}

// SetTargetPromotionGroupIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the TargetPromotionGroupIdFilter field.
func (o *SponsoredProductsListTargetPromotionGroupsRequestContent) SetTargetPromotionGroupIdFilter(v SponsoredProductsObjectIdFilter) {
	o.TargetPromotionGroupIdFilter = &v
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListTargetPromotionGroupsRequestContent) GetAdGroupIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListTargetPromotionGroupsRequestContent) GetAdGroupIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListTargetPromotionGroupsRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *SponsoredProductsListTargetPromotionGroupsRequestContent) SetAdGroupIdFilter(v SponsoredProductsObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

func (o SponsoredProductsListTargetPromotionGroupsRequestContent) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsListTargetPromotionGroupsRequestContent struct {
	value *SponsoredProductsListTargetPromotionGroupsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListTargetPromotionGroupsRequestContent) Get() *SponsoredProductsListTargetPromotionGroupsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListTargetPromotionGroupsRequestContent) Set(val *SponsoredProductsListTargetPromotionGroupsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListTargetPromotionGroupsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListTargetPromotionGroupsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListTargetPromotionGroupsRequestContent(val *SponsoredProductsListTargetPromotionGroupsRequestContent) *NullableSponsoredProductsListTargetPromotionGroupsRequestContent {
	return &NullableSponsoredProductsListTargetPromotionGroupsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListTargetPromotionGroupsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListTargetPromotionGroupsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
