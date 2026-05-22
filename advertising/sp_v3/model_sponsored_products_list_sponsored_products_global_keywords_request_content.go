package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent{}

// SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent struct for SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent
type SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent struct {
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
	// Whether to get entity with extended data fields such as creationDate, lastUpdateDate, servingStatus
	IncludeExtendedDataFields *bool                            `json:"includeExtendedDataFields,omitempty"`
	KeywordIdFilter           *SponsoredProductsObjectIdFilter `json:"keywordIdFilter,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent instantiates a new SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent() *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent {
	this := SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsGlobalKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsGlobalKeywordsRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent {
	this := SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent{}
	return &this
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

// GetKeywordIdFilter returns the KeywordIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) GetKeywordIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.KeywordIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.KeywordIdFilter
}

// GetKeywordIdFilterOk returns a tuple with the KeywordIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) GetKeywordIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.KeywordIdFilter) {
		return nil, false
	}
	return o.KeywordIdFilter, true
}

// HasKeywordIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) HasKeywordIdFilter() bool {
	if o != nil && !IsNil(o.KeywordIdFilter) {
		return true
	}

	return false
}

// SetKeywordIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the KeywordIdFilter field.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) SetKeywordIdFilter(v SponsoredProductsObjectIdFilter) {
	o.KeywordIdFilter = &v
}

func (o SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.IncludeExtendedDataFields) {
		toSerialize["includeExtendedDataFields"] = o.IncludeExtendedDataFields
	}
	if !IsNil(o.KeywordIdFilter) {
		toSerialize["keywordIdFilter"] = o.KeywordIdFilter
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent struct {
	value *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) Get() *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) Set(val *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent(val *SponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) *NullableSponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
