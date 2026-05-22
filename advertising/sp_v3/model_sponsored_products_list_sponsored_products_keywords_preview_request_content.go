package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent{}

// SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent struct for SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent
type SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent struct {
	CampaignIdFilter *SponsoredProductsReducedObjectIdFilter `json:"campaignIdFilter,omitempty"`
	StateFilter      *SponsoredProductsEntityStateFilter     `json:"stateFilter,omitempty"`
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken       *string                                 `json:"nextToken,omitempty"`
	AdGroupIdFilter *SponsoredProductsReducedObjectIdFilter `json:"adGroupIdFilter,omitempty"`
	// Whether to get entity with extended data fields such as creationDate, lastUpdateDate, servingStatus
	IncludeExtendedDataFields *bool `json:"includeExtendedDataFields,omitempty"`
	// Restricts results to keywords associated with locale
	Locale            *string                             `json:"locale,omitempty"`
	KeywordTextFilter *SponsoredProductsKeywordTextFilter `json:"keywordTextFilter,omitempty"`
	KeywordIdFilter   *SponsoredProductsObjectIdFilter    `json:"keywordIdFilter,omitempty"`
	// Only the keyword with match type that is in this list will be listed
	MatchTypeFilter []SponsoredProductsMatchType `json:"matchTypeFilter,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent instantiates a new SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent() *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent {
	this := SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsKeywordsPreviewRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsKeywordsPreviewRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent {
	this := SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetCampaignIdFilter() SponsoredProductsReducedObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret SponsoredProductsReducedObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsReducedObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given SponsoredProductsReducedObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) SetCampaignIdFilter(v SponsoredProductsReducedObjectIdFilter) {
	o.CampaignIdFilter = &v
}

// GetStateFilter returns the StateFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetStateFilter() SponsoredProductsEntityStateFilter {
	if o == nil || IsNil(o.StateFilter) {
		var ret SponsoredProductsEntityStateFilter
		return ret
	}
	return *o.StateFilter
}

// GetStateFilterOk returns a tuple with the StateFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetStateFilterOk() (*SponsoredProductsEntityStateFilter, bool) {
	if o == nil || IsNil(o.StateFilter) {
		return nil, false
	}
	return o.StateFilter, true
}

// HasStateFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) HasStateFilter() bool {
	if o != nil && !IsNil(o.StateFilter) {
		return true
	}

	return false
}

// SetStateFilter gets a reference to the given SponsoredProductsEntityStateFilter and assigns it to the StateFilter field.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) SetStateFilter(v SponsoredProductsEntityStateFilter) {
	o.StateFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetAdGroupIdFilter() SponsoredProductsReducedObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret SponsoredProductsReducedObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetAdGroupIdFilterOk() (*SponsoredProductsReducedObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given SponsoredProductsReducedObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) SetAdGroupIdFilter(v SponsoredProductsReducedObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) SetLocale(v string) {
	o.Locale = &v
}

// GetKeywordTextFilter returns the KeywordTextFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetKeywordTextFilter() SponsoredProductsKeywordTextFilter {
	if o == nil || IsNil(o.KeywordTextFilter) {
		var ret SponsoredProductsKeywordTextFilter
		return ret
	}
	return *o.KeywordTextFilter
}

// GetKeywordTextFilterOk returns a tuple with the KeywordTextFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetKeywordTextFilterOk() (*SponsoredProductsKeywordTextFilter, bool) {
	if o == nil || IsNil(o.KeywordTextFilter) {
		return nil, false
	}
	return o.KeywordTextFilter, true
}

// HasKeywordTextFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) HasKeywordTextFilter() bool {
	if o != nil && !IsNil(o.KeywordTextFilter) {
		return true
	}

	return false
}

// SetKeywordTextFilter gets a reference to the given SponsoredProductsKeywordTextFilter and assigns it to the KeywordTextFilter field.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) SetKeywordTextFilter(v SponsoredProductsKeywordTextFilter) {
	o.KeywordTextFilter = &v
}

// GetKeywordIdFilter returns the KeywordIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetKeywordIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.KeywordIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.KeywordIdFilter
}

// GetKeywordIdFilterOk returns a tuple with the KeywordIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetKeywordIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.KeywordIdFilter) {
		return nil, false
	}
	return o.KeywordIdFilter, true
}

// HasKeywordIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) HasKeywordIdFilter() bool {
	if o != nil && !IsNil(o.KeywordIdFilter) {
		return true
	}

	return false
}

// SetKeywordIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the KeywordIdFilter field.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) SetKeywordIdFilter(v SponsoredProductsObjectIdFilter) {
	o.KeywordIdFilter = &v
}

// GetMatchTypeFilter returns the MatchTypeFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetMatchTypeFilter() []SponsoredProductsMatchType {
	if o == nil || IsNil(o.MatchTypeFilter) {
		var ret []SponsoredProductsMatchType
		return ret
	}
	return o.MatchTypeFilter
}

// GetMatchTypeFilterOk returns a tuple with the MatchTypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) GetMatchTypeFilterOk() ([]SponsoredProductsMatchType, bool) {
	if o == nil || IsNil(o.MatchTypeFilter) {
		return nil, false
	}
	return o.MatchTypeFilter, true
}

// HasMatchTypeFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) HasMatchTypeFilter() bool {
	if o != nil && !IsNil(o.MatchTypeFilter) {
		return true
	}

	return false
}

// SetMatchTypeFilter gets a reference to the given []SponsoredProductsMatchType and assigns it to the MatchTypeFilter field.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) SetMatchTypeFilter(v []SponsoredProductsMatchType) {
	o.MatchTypeFilter = v
}

func (o SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignIdFilter) {
		toSerialize["campaignIdFilter"] = o.CampaignIdFilter
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
	if !IsNil(o.AdGroupIdFilter) {
		toSerialize["adGroupIdFilter"] = o.AdGroupIdFilter
	}
	if !IsNil(o.IncludeExtendedDataFields) {
		toSerialize["includeExtendedDataFields"] = o.IncludeExtendedDataFields
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.KeywordTextFilter) {
		toSerialize["keywordTextFilter"] = o.KeywordTextFilter
	}
	if !IsNil(o.KeywordIdFilter) {
		toSerialize["keywordIdFilter"] = o.KeywordIdFilter
	}
	if !IsNil(o.MatchTypeFilter) {
		toSerialize["matchTypeFilter"] = o.MatchTypeFilter
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent struct {
	value *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) Get() *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) Set(val *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent(val *SponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) *NullableSponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsKeywordsPreviewRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
