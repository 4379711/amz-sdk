package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent{}

// SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent struct for SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent
type SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent struct {
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken      *string                          `json:"nextToken,omitempty"`
	TargetIdFilter *SponsoredProductsObjectIdFilter `json:"targetIdFilter,omitempty"`
	// Whether to get entity with extended data fields such as creationDate, lastUpdateDate, servingStatus
	IncludeExtendedDataFields *bool `json:"includeExtendedDataFields,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent instantiates a new SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent() *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent {
	this := SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent {
	this := SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent{}
	return &this
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTargetIdFilter returns the TargetIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) GetTargetIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.TargetIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.TargetIdFilter
}

// GetTargetIdFilterOk returns a tuple with the TargetIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) GetTargetIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.TargetIdFilter) {
		return nil, false
	}
	return o.TargetIdFilter, true
}

// HasTargetIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) HasTargetIdFilter() bool {
	if o != nil && !IsNil(o.TargetIdFilter) {
		return true
	}

	return false
}

// SetTargetIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the TargetIdFilter field.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) SetTargetIdFilter(v SponsoredProductsObjectIdFilter) {
	o.TargetIdFilter = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

func (o SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.TargetIdFilter) {
		toSerialize["targetIdFilter"] = o.TargetIdFilter
	}
	if !IsNil(o.IncludeExtendedDataFields) {
		toSerialize["includeExtendedDataFields"] = o.IncludeExtendedDataFields
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent struct {
	value *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) Get() *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) Set(val *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent(val *SponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) *NullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
