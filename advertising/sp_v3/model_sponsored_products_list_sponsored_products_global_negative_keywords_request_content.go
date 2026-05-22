package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent{}

// SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent struct for SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent
type SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent struct {
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
	// Whether to get entity with extended data fields such as creationDate, lastUpdateDate, servingStatus
	IncludeExtendedDataFields *bool                            `json:"includeExtendedDataFields,omitempty"`
	NegativeKeywordIdFilter   *SponsoredProductsObjectIdFilter `json:"negativeKeywordIdFilter,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent instantiates a new SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent() *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent {
	this := SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent {
	this := SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent{}
	return &this
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

// GetNegativeKeywordIdFilter returns the NegativeKeywordIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) GetNegativeKeywordIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.NegativeKeywordIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.NegativeKeywordIdFilter
}

// GetNegativeKeywordIdFilterOk returns a tuple with the NegativeKeywordIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) GetNegativeKeywordIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.NegativeKeywordIdFilter) {
		return nil, false
	}
	return o.NegativeKeywordIdFilter, true
}

// HasNegativeKeywordIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) HasNegativeKeywordIdFilter() bool {
	if o != nil && !IsNil(o.NegativeKeywordIdFilter) {
		return true
	}

	return false
}

// SetNegativeKeywordIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the NegativeKeywordIdFilter field.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) SetNegativeKeywordIdFilter(v SponsoredProductsObjectIdFilter) {
	o.NegativeKeywordIdFilter = &v
}

func (o SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.NegativeKeywordIdFilter) {
		toSerialize["negativeKeywordIdFilter"] = o.NegativeKeywordIdFilter
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent struct {
	value *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) Get() *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) Set(val *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent(val *SponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) *NullableSponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalNegativeKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
