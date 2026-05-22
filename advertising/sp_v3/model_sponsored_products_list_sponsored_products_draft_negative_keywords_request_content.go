package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent{}

// SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent struct for SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent
type SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent struct {
	CampaignIdFilter          *SponsoredProductsReducedObjectIdFilter `json:"campaignIdFilter,omitempty"`
	NegativeKeywordTextFilter *SponsoredProductsKeywordTextFilter     `json:"negativeKeywordTextFilter,omitempty"`
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken       *string                                 `json:"nextToken,omitempty"`
	AdGroupIdFilter *SponsoredProductsReducedObjectIdFilter `json:"adGroupIdFilter,omitempty"`
	// Whether to get entity with extended data fields such as creationDate, lastUpdateDate, servingStatus
	IncludeExtendedDataFields *bool `json:"includeExtendedDataFields,omitempty"`
	// The locale preference of the advertiser.
	Locale                  *string                          `json:"locale,omitempty"`
	NegativeKeywordIdFilter *SponsoredProductsObjectIdFilter `json:"negativeKeywordIdFilter,omitempty"`
	// Restricts results to resources with the selected matchType
	MatchTypeFilter []SponsoredProductsNegativeMatchType `json:"matchTypeFilter,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent instantiates a new SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent() *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent {
	this := SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent {
	this := SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) GetCampaignIdFilter() SponsoredProductsReducedObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret SponsoredProductsReducedObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsReducedObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given SponsoredProductsReducedObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) SetCampaignIdFilter(v SponsoredProductsReducedObjectIdFilter) {
	o.CampaignIdFilter = &v
}

// GetNegativeKeywordTextFilter returns the NegativeKeywordTextFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) GetNegativeKeywordTextFilter() SponsoredProductsKeywordTextFilter {
	if o == nil || IsNil(o.NegativeKeywordTextFilter) {
		var ret SponsoredProductsKeywordTextFilter
		return ret
	}
	return *o.NegativeKeywordTextFilter
}

// GetNegativeKeywordTextFilterOk returns a tuple with the NegativeKeywordTextFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) GetNegativeKeywordTextFilterOk() (*SponsoredProductsKeywordTextFilter, bool) {
	if o == nil || IsNil(o.NegativeKeywordTextFilter) {
		return nil, false
	}
	return o.NegativeKeywordTextFilter, true
}

// HasNegativeKeywordTextFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) HasNegativeKeywordTextFilter() bool {
	if o != nil && !IsNil(o.NegativeKeywordTextFilter) {
		return true
	}

	return false
}

// SetNegativeKeywordTextFilter gets a reference to the given SponsoredProductsKeywordTextFilter and assigns it to the NegativeKeywordTextFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) SetNegativeKeywordTextFilter(v SponsoredProductsKeywordTextFilter) {
	o.NegativeKeywordTextFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) GetAdGroupIdFilter() SponsoredProductsReducedObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret SponsoredProductsReducedObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) GetAdGroupIdFilterOk() (*SponsoredProductsReducedObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given SponsoredProductsReducedObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) SetAdGroupIdFilter(v SponsoredProductsReducedObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) SetLocale(v string) {
	o.Locale = &v
}

// GetNegativeKeywordIdFilter returns the NegativeKeywordIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) GetNegativeKeywordIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.NegativeKeywordIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.NegativeKeywordIdFilter
}

// GetNegativeKeywordIdFilterOk returns a tuple with the NegativeKeywordIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) GetNegativeKeywordIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.NegativeKeywordIdFilter) {
		return nil, false
	}
	return o.NegativeKeywordIdFilter, true
}

// HasNegativeKeywordIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) HasNegativeKeywordIdFilter() bool {
	if o != nil && !IsNil(o.NegativeKeywordIdFilter) {
		return true
	}

	return false
}

// SetNegativeKeywordIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the NegativeKeywordIdFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) SetNegativeKeywordIdFilter(v SponsoredProductsObjectIdFilter) {
	o.NegativeKeywordIdFilter = &v
}

// GetMatchTypeFilter returns the MatchTypeFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) GetMatchTypeFilter() []SponsoredProductsNegativeMatchType {
	if o == nil || IsNil(o.MatchTypeFilter) {
		var ret []SponsoredProductsNegativeMatchType
		return ret
	}
	return o.MatchTypeFilter
}

// GetMatchTypeFilterOk returns a tuple with the MatchTypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) GetMatchTypeFilterOk() ([]SponsoredProductsNegativeMatchType, bool) {
	if o == nil || IsNil(o.MatchTypeFilter) {
		return nil, false
	}
	return o.MatchTypeFilter, true
}

// HasMatchTypeFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) HasMatchTypeFilter() bool {
	if o != nil && !IsNil(o.MatchTypeFilter) {
		return true
	}

	return false
}

// SetMatchTypeFilter gets a reference to the given []SponsoredProductsNegativeMatchType and assigns it to the MatchTypeFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) SetMatchTypeFilter(v []SponsoredProductsNegativeMatchType) {
	o.MatchTypeFilter = v
}

func (o SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignIdFilter) {
		toSerialize["campaignIdFilter"] = o.CampaignIdFilter
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

type NullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent struct {
	value *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) Get() *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) Set(val *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent(val *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) *NullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
