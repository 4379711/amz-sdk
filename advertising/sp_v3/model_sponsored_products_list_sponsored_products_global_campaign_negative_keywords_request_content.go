package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent{}

// SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent struct for SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent
type SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent struct {
	CampaignNegativeKeywordIdFilter *SponsoredProductsObjectIdFilter `json:"campaignNegativeKeywordIdFilter,omitempty"`
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
	// Whether to get entity with extended data fields such as creationDate, lastUpdateDate, servingStatus
	IncludeExtendedDataFields *bool `json:"includeExtendedDataFields,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent instantiates a new SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent() *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent {
	this := SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent {
	this := SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent{}
	return &this
}

// GetCampaignNegativeKeywordIdFilter returns the CampaignNegativeKeywordIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywordIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.CampaignNegativeKeywordIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.CampaignNegativeKeywordIdFilter
}

// GetCampaignNegativeKeywordIdFilterOk returns a tuple with the CampaignNegativeKeywordIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywordIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignNegativeKeywordIdFilter) {
		return nil, false
	}
	return o.CampaignNegativeKeywordIdFilter, true
}

// HasCampaignNegativeKeywordIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) HasCampaignNegativeKeywordIdFilter() bool {
	if o != nil && !IsNil(o.CampaignNegativeKeywordIdFilter) {
		return true
	}

	return false
}

// SetCampaignNegativeKeywordIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the CampaignNegativeKeywordIdFilter field.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) SetCampaignNegativeKeywordIdFilter(v SponsoredProductsObjectIdFilter) {
	o.CampaignNegativeKeywordIdFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

func (o SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignNegativeKeywordIdFilter) {
		toSerialize["campaignNegativeKeywordIdFilter"] = o.CampaignNegativeKeywordIdFilter
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
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent struct {
	value *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) Get() *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) Set(val *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent(val *SponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) *NullableSponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
