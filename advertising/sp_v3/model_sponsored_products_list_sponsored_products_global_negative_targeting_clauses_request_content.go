package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent{}

// SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent struct for SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent
type SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent struct {
	NegativeTargetIdFilter *SponsoredProductsObjectIdFilter `json:"negativeTargetIdFilter,omitempty"`
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
	// Whether to get entity with extended data fields such as creationDate, lastUpdateDate, servingStatus
	IncludeExtendedDataFields *bool `json:"includeExtendedDataFields,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent instantiates a new SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent() *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent {
	this := SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent {
	this := SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent{}
	return &this
}

// GetNegativeTargetIdFilter returns the NegativeTargetIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) GetNegativeTargetIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.NegativeTargetIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.NegativeTargetIdFilter
}

// GetNegativeTargetIdFilterOk returns a tuple with the NegativeTargetIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) GetNegativeTargetIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.NegativeTargetIdFilter) {
		return nil, false
	}
	return o.NegativeTargetIdFilter, true
}

// HasNegativeTargetIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) HasNegativeTargetIdFilter() bool {
	if o != nil && !IsNil(o.NegativeTargetIdFilter) {
		return true
	}

	return false
}

// SetNegativeTargetIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the NegativeTargetIdFilter field.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) SetNegativeTargetIdFilter(v SponsoredProductsObjectIdFilter) {
	o.NegativeTargetIdFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

func (o SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NegativeTargetIdFilter) {
		toSerialize["negativeTargetIdFilter"] = o.NegativeTargetIdFilter
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

type NullableSponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent struct {
	value *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) Get() *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) Set(val *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent(val *SponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) *NullableSponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalNegativeTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
