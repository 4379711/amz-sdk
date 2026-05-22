package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent{}

// SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent struct for SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent
type SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent struct {
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken       *string                          `json:"nextToken,omitempty"`
	AdGroupIdFilter *SponsoredProductsObjectIdFilter `json:"adGroupIdFilter,omitempty"`
	// Whether to get entity with extended data fields such as creationDate, lastUpdateDate, servingStatus
	IncludeExtendedDataFields *bool `json:"includeExtendedDataFields,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent instantiates a new SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent() *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent {
	this := SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent {
	this := SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent{}
	return &this
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) GetAdGroupIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) GetAdGroupIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) SetAdGroupIdFilter(v SponsoredProductsObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

func (o SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent struct {
	value *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) Get() *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) Set(val *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent(val *SponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) *NullableSponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalAdGroupsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
