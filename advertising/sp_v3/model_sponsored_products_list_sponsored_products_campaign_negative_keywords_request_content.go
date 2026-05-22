package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent{}

// SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent struct for SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent
type SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent struct {
	CampaignIdFilter                *SponsoredProductsReducedObjectIdFilter `json:"campaignIdFilter,omitempty"`
	CampaignNegativeKeywordIdFilter *SponsoredProductsObjectIdFilter        `json:"campaignNegativeKeywordIdFilter,omitempty"`
	StateFilter                     *SponsoredProductsEntityStateFilter     `json:"stateFilter,omitempty"`
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
	// Whether to get entity with extended data fields such as creationDate, lastUpdateDate, servingStatus
	IncludeExtendedDataFields         *bool                               `json:"includeExtendedDataFields,omitempty"`
	CampaignNegativeKeywordTextFilter *SponsoredProductsKeywordTextFilter `json:"campaignNegativeKeywordTextFilter,omitempty"`
	// Restricts results to resources with the selected matchType
	MatchTypeFilter []SponsoredProductsNegativeMatchType `json:"matchTypeFilter,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent instantiates a new SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent() *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent {
	this := SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent {
	this := SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) GetCampaignIdFilter() SponsoredProductsReducedObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret SponsoredProductsReducedObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsReducedObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given SponsoredProductsReducedObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) SetCampaignIdFilter(v SponsoredProductsReducedObjectIdFilter) {
	o.CampaignIdFilter = &v
}

// GetCampaignNegativeKeywordIdFilter returns the CampaignNegativeKeywordIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywordIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.CampaignNegativeKeywordIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.CampaignNegativeKeywordIdFilter
}

// GetCampaignNegativeKeywordIdFilterOk returns a tuple with the CampaignNegativeKeywordIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywordIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignNegativeKeywordIdFilter) {
		return nil, false
	}
	return o.CampaignNegativeKeywordIdFilter, true
}

// HasCampaignNegativeKeywordIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) HasCampaignNegativeKeywordIdFilter() bool {
	if o != nil && !IsNil(o.CampaignNegativeKeywordIdFilter) {
		return true
	}

	return false
}

// SetCampaignNegativeKeywordIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the CampaignNegativeKeywordIdFilter field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) SetCampaignNegativeKeywordIdFilter(v SponsoredProductsObjectIdFilter) {
	o.CampaignNegativeKeywordIdFilter = &v
}

// GetStateFilter returns the StateFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) GetStateFilter() SponsoredProductsEntityStateFilter {
	if o == nil || IsNil(o.StateFilter) {
		var ret SponsoredProductsEntityStateFilter
		return ret
	}
	return *o.StateFilter
}

// GetStateFilterOk returns a tuple with the StateFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) GetStateFilterOk() (*SponsoredProductsEntityStateFilter, bool) {
	if o == nil || IsNil(o.StateFilter) {
		return nil, false
	}
	return o.StateFilter, true
}

// HasStateFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) HasStateFilter() bool {
	if o != nil && !IsNil(o.StateFilter) {
		return true
	}

	return false
}

// SetStateFilter gets a reference to the given SponsoredProductsEntityStateFilter and assigns it to the StateFilter field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) SetStateFilter(v SponsoredProductsEntityStateFilter) {
	o.StateFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

// GetCampaignNegativeKeywordTextFilter returns the CampaignNegativeKeywordTextFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywordTextFilter() SponsoredProductsKeywordTextFilter {
	if o == nil || IsNil(o.CampaignNegativeKeywordTextFilter) {
		var ret SponsoredProductsKeywordTextFilter
		return ret
	}
	return *o.CampaignNegativeKeywordTextFilter
}

// GetCampaignNegativeKeywordTextFilterOk returns a tuple with the CampaignNegativeKeywordTextFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywordTextFilterOk() (*SponsoredProductsKeywordTextFilter, bool) {
	if o == nil || IsNil(o.CampaignNegativeKeywordTextFilter) {
		return nil, false
	}
	return o.CampaignNegativeKeywordTextFilter, true
}

// HasCampaignNegativeKeywordTextFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) HasCampaignNegativeKeywordTextFilter() bool {
	if o != nil && !IsNil(o.CampaignNegativeKeywordTextFilter) {
		return true
	}

	return false
}

// SetCampaignNegativeKeywordTextFilter gets a reference to the given SponsoredProductsKeywordTextFilter and assigns it to the CampaignNegativeKeywordTextFilter field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) SetCampaignNegativeKeywordTextFilter(v SponsoredProductsKeywordTextFilter) {
	o.CampaignNegativeKeywordTextFilter = &v
}

// GetMatchTypeFilter returns the MatchTypeFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) GetMatchTypeFilter() []SponsoredProductsNegativeMatchType {
	if o == nil || IsNil(o.MatchTypeFilter) {
		var ret []SponsoredProductsNegativeMatchType
		return ret
	}
	return o.MatchTypeFilter
}

// GetMatchTypeFilterOk returns a tuple with the MatchTypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) GetMatchTypeFilterOk() ([]SponsoredProductsNegativeMatchType, bool) {
	if o == nil || IsNil(o.MatchTypeFilter) {
		return nil, false
	}
	return o.MatchTypeFilter, true
}

// HasMatchTypeFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) HasMatchTypeFilter() bool {
	if o != nil && !IsNil(o.MatchTypeFilter) {
		return true
	}

	return false
}

// SetMatchTypeFilter gets a reference to the given []SponsoredProductsNegativeMatchType and assigns it to the MatchTypeFilter field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) SetMatchTypeFilter(v []SponsoredProductsNegativeMatchType) {
	o.MatchTypeFilter = v
}

func (o SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignIdFilter) {
		toSerialize["campaignIdFilter"] = o.CampaignIdFilter
	}
	if !IsNil(o.CampaignNegativeKeywordIdFilter) {
		toSerialize["campaignNegativeKeywordIdFilter"] = o.CampaignNegativeKeywordIdFilter
	}
	if !IsNil(o.StateFilter) {
		toSerialize["stateFilter"] = o.StateFilter
	}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.IncludeExtendedDataFields) {
		toSerialize["includeExtendedDataFields"] = o.IncludeExtendedDataFields
	}
	if !IsNil(o.CampaignNegativeKeywordTextFilter) {
		toSerialize["campaignNegativeKeywordTextFilter"] = o.CampaignNegativeKeywordTextFilter
	}
	if !IsNil(o.MatchTypeFilter) {
		toSerialize["matchTypeFilter"] = o.MatchTypeFilter
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent struct {
	value *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) Get() *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) Set(val *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent(val *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) *NullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
