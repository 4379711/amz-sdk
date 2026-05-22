package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent{}

// SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent struct for SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent
type SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent struct {
	CampaignIdFilter          *SponsoredProductsReducedObjectIdFilter `json:"campaignIdFilter,omitempty"`
	StateFilter               *SponsoredProductsEntityStateFilter     `json:"stateFilter,omitempty"`
	NegativeKeywordTextFilter *SponsoredProductsKeywordTextFilter     `json:"negativeKeywordTextFilter,omitempty"`
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken       *string                                 `json:"nextToken,omitempty"`
	AdGroupIdFilter *SponsoredProductsReducedObjectIdFilter `json:"adGroupIdFilter,omitempty"`
	// Whether to get entity with extended data fields such as creationDate, lastUpdateDate, servingStatus
	IncludeExtendedDataFields *bool `json:"includeExtendedDataFields,omitempty"`
	// Restricts results to negativeKeywords that match the specified locale.
	Locale                  *string                          `json:"locale,omitempty"`
	NegativeKeywordIdFilter *SponsoredProductsObjectIdFilter `json:"negativeKeywordIdFilter,omitempty"`
	// Only the negativeKeyword with the match type that is in this list will be listed
	MatchTypeFilter []SponsoredProductsNegativeMatchType `json:"matchTypeFilter,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent instantiates a new SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent() *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent {
	this := SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent {
	this := SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetCampaignIdFilter() SponsoredProductsReducedObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret SponsoredProductsReducedObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsReducedObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given SponsoredProductsReducedObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) SetCampaignIdFilter(v SponsoredProductsReducedObjectIdFilter) {
	o.CampaignIdFilter = &v
}

// GetStateFilter returns the StateFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetStateFilter() SponsoredProductsEntityStateFilter {
	if o == nil || IsNil(o.StateFilter) {
		var ret SponsoredProductsEntityStateFilter
		return ret
	}
	return *o.StateFilter
}

// GetStateFilterOk returns a tuple with the StateFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetStateFilterOk() (*SponsoredProductsEntityStateFilter, bool) {
	if o == nil || IsNil(o.StateFilter) {
		return nil, false
	}
	return o.StateFilter, true
}

// HasStateFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) HasStateFilter() bool {
	if o != nil && !IsNil(o.StateFilter) {
		return true
	}

	return false
}

// SetStateFilter gets a reference to the given SponsoredProductsEntityStateFilter and assigns it to the StateFilter field.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) SetStateFilter(v SponsoredProductsEntityStateFilter) {
	o.StateFilter = &v
}

// GetNegativeKeywordTextFilter returns the NegativeKeywordTextFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetNegativeKeywordTextFilter() SponsoredProductsKeywordTextFilter {
	if o == nil || IsNil(o.NegativeKeywordTextFilter) {
		var ret SponsoredProductsKeywordTextFilter
		return ret
	}
	return *o.NegativeKeywordTextFilter
}

// GetNegativeKeywordTextFilterOk returns a tuple with the NegativeKeywordTextFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetNegativeKeywordTextFilterOk() (*SponsoredProductsKeywordTextFilter, bool) {
	if o == nil || IsNil(o.NegativeKeywordTextFilter) {
		return nil, false
	}
	return o.NegativeKeywordTextFilter, true
}

// HasNegativeKeywordTextFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) HasNegativeKeywordTextFilter() bool {
	if o != nil && !IsNil(o.NegativeKeywordTextFilter) {
		return true
	}

	return false
}

// SetNegativeKeywordTextFilter gets a reference to the given SponsoredProductsKeywordTextFilter and assigns it to the NegativeKeywordTextFilter field.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) SetNegativeKeywordTextFilter(v SponsoredProductsKeywordTextFilter) {
	o.NegativeKeywordTextFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetAdGroupIdFilter() SponsoredProductsReducedObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret SponsoredProductsReducedObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetAdGroupIdFilterOk() (*SponsoredProductsReducedObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given SponsoredProductsReducedObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) SetAdGroupIdFilter(v SponsoredProductsReducedObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) SetLocale(v string) {
	o.Locale = &v
}

// GetNegativeKeywordIdFilter returns the NegativeKeywordIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetNegativeKeywordIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.NegativeKeywordIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.NegativeKeywordIdFilter
}

// GetNegativeKeywordIdFilterOk returns a tuple with the NegativeKeywordIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetNegativeKeywordIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.NegativeKeywordIdFilter) {
		return nil, false
	}
	return o.NegativeKeywordIdFilter, true
}

// HasNegativeKeywordIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) HasNegativeKeywordIdFilter() bool {
	if o != nil && !IsNil(o.NegativeKeywordIdFilter) {
		return true
	}

	return false
}

// SetNegativeKeywordIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the NegativeKeywordIdFilter field.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) SetNegativeKeywordIdFilter(v SponsoredProductsObjectIdFilter) {
	o.NegativeKeywordIdFilter = &v
}

// GetMatchTypeFilter returns the MatchTypeFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetMatchTypeFilter() []SponsoredProductsNegativeMatchType {
	if o == nil || IsNil(o.MatchTypeFilter) {
		var ret []SponsoredProductsNegativeMatchType
		return ret
	}
	return o.MatchTypeFilter
}

// GetMatchTypeFilterOk returns a tuple with the MatchTypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) GetMatchTypeFilterOk() ([]SponsoredProductsNegativeMatchType, bool) {
	if o == nil || IsNil(o.MatchTypeFilter) {
		return nil, false
	}
	return o.MatchTypeFilter, true
}

// HasMatchTypeFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) HasMatchTypeFilter() bool {
	if o != nil && !IsNil(o.MatchTypeFilter) {
		return true
	}

	return false
}

// SetMatchTypeFilter gets a reference to the given []SponsoredProductsNegativeMatchType and assigns it to the MatchTypeFilter field.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) SetMatchTypeFilter(v []SponsoredProductsNegativeMatchType) {
	o.MatchTypeFilter = v
}

func (o SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignIdFilter) {
		toSerialize["campaignIdFilter"] = o.CampaignIdFilter
	}
	if !IsNil(o.StateFilter) {
		toSerialize["stateFilter"] = o.StateFilter
	}
	if !IsNil(o.NegativeKeywordTextFilter) {
		toSerialize["negativeKeywordTextFilter"] = o.NegativeKeywordTextFilter
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
	if !IsNil(o.NegativeKeywordIdFilter) {
		toSerialize["negativeKeywordIdFilter"] = o.NegativeKeywordIdFilter
	}
	if !IsNil(o.MatchTypeFilter) {
		toSerialize["matchTypeFilter"] = o.MatchTypeFilter
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent struct {
	value *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) Get() *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) Set(val *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent(val *SponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) *NullableSponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsNegativeKeywordsPreviewRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
