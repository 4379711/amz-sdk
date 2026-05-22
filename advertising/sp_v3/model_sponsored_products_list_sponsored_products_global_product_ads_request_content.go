package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent{}

// SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent struct for SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent
type SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent struct {
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken  *string                          `json:"nextToken,omitempty"`
	AdIdFilter *SponsoredProductsObjectIdFilter `json:"adIdFilter,omitempty"`
	// Whether to get entity with extended data fields such as creationDate, lastUpdateDate, servingStatus
	IncludeExtendedDataFields *bool `json:"includeExtendedDataFields,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent instantiates a new SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent() *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent {
	this := SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsGlobalProductAdsRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsGlobalProductAdsRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent {
	this := SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent{}
	return &this
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetAdIdFilter returns the AdIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) GetAdIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.AdIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.AdIdFilter
}

// GetAdIdFilterOk returns a tuple with the AdIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) GetAdIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdIdFilter) {
		return nil, false
	}
	return o.AdIdFilter, true
}

// HasAdIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) HasAdIdFilter() bool {
	if o != nil && !IsNil(o.AdIdFilter) {
		return true
	}

	return false
}

// SetAdIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the AdIdFilter field.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) SetAdIdFilter(v SponsoredProductsObjectIdFilter) {
	o.AdIdFilter = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

func (o SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.AdIdFilter) {
		toSerialize["adIdFilter"] = o.AdIdFilter
	}
	if !IsNil(o.IncludeExtendedDataFields) {
		toSerialize["includeExtendedDataFields"] = o.IncludeExtendedDataFields
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent struct {
	value *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) Get() *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) Set(val *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent(val *SponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) *NullableSponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalProductAdsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
