package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListMatchedAdGroupsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListMatchedAdGroupsRequestContent{}

// SponsoredProductsListMatchedAdGroupsRequestContent struct for SponsoredProductsListMatchedAdGroupsRequestContent
type SponsoredProductsListMatchedAdGroupsRequestContent struct {
	// List of ad Ids for ASIN/SKU matches.
	AdIds []string `json:"adIds,omitempty"`
	// The maximum number of results requested.
	MaxResults *int32 `json:"maxResults,omitempty"`
	// Token value allowing to navigate to the next or previous response page
	NextToken *string `json:"nextToken,omitempty"`
	// The id of an adGroup. ASIN/SKU match will be based on all the ads this adGroup id contains if adIds are not provided.
	AdGroupId *string `json:"adGroupId,omitempty"`
}

// NewSponsoredProductsListMatchedAdGroupsRequestContent instantiates a new SponsoredProductsListMatchedAdGroupsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListMatchedAdGroupsRequestContent() *SponsoredProductsListMatchedAdGroupsRequestContent {
	this := SponsoredProductsListMatchedAdGroupsRequestContent{}
	return &this
}

// NewSponsoredProductsListMatchedAdGroupsRequestContentWithDefaults instantiates a new SponsoredProductsListMatchedAdGroupsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListMatchedAdGroupsRequestContentWithDefaults() *SponsoredProductsListMatchedAdGroupsRequestContent {
	this := SponsoredProductsListMatchedAdGroupsRequestContent{}
	return &this
}

// GetAdIds returns the AdIds field value if set, zero value otherwise.
func (o *SponsoredProductsListMatchedAdGroupsRequestContent) GetAdIds() []string {
	if o == nil || IsNil(o.AdIds) {
		var ret []string
		return ret
	}
	return o.AdIds
}

// GetAdIdsOk returns a tuple with the AdIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListMatchedAdGroupsRequestContent) GetAdIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AdIds) {
		return nil, false
	}
	return o.AdIds, true
}

// HasAdIds returns a boolean if a field has been set.
func (o *SponsoredProductsListMatchedAdGroupsRequestContent) HasAdIds() bool {
	if o != nil && !IsNil(o.AdIds) {
		return true
	}

	return false
}

// SetAdIds gets a reference to the given []string and assigns it to the AdIds field.
func (o *SponsoredProductsListMatchedAdGroupsRequestContent) SetAdIds(v []string) {
	o.AdIds = v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListMatchedAdGroupsRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListMatchedAdGroupsRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListMatchedAdGroupsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListMatchedAdGroupsRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListMatchedAdGroupsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListMatchedAdGroupsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListMatchedAdGroupsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListMatchedAdGroupsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *SponsoredProductsListMatchedAdGroupsRequestContent) GetAdGroupId() string {
	if o == nil || IsNil(o.AdGroupId) {
		var ret string
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListMatchedAdGroupsRequestContent) GetAdGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *SponsoredProductsListMatchedAdGroupsRequestContent) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given string and assigns it to the AdGroupId field.
func (o *SponsoredProductsListMatchedAdGroupsRequestContent) SetAdGroupId(v string) {
	o.AdGroupId = &v
}

func (o SponsoredProductsListMatchedAdGroupsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdIds) {
		toSerialize["adIds"] = o.AdIds
	}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListMatchedAdGroupsRequestContent struct {
	value *SponsoredProductsListMatchedAdGroupsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListMatchedAdGroupsRequestContent) Get() *SponsoredProductsListMatchedAdGroupsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListMatchedAdGroupsRequestContent) Set(val *SponsoredProductsListMatchedAdGroupsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListMatchedAdGroupsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListMatchedAdGroupsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListMatchedAdGroupsRequestContent(val *SponsoredProductsListMatchedAdGroupsRequestContent) *NullableSponsoredProductsListMatchedAdGroupsRequestContent {
	return &NullableSponsoredProductsListMatchedAdGroupsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListMatchedAdGroupsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListMatchedAdGroupsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
