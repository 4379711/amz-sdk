package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent{}

// SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent struct for SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent
type SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent struct {
	CampaignIdFilter               *SponsoredProductsReducedObjectIdFilter `json:"campaignIdFilter,omitempty"`
	CampaignNegativeTargetIdFilter *SponsoredProductsObjectIdFilter        `json:"campaignNegativeTargetIdFilter,omitempty"`
	StateFilter                    *SponsoredProductsEntityStateFilter     `json:"stateFilter,omitempty"`
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken  *string                      `json:"nextToken,omitempty"`
	AsinFilter *SponsoredProductsAsinFilter `json:"asinFilter,omitempty"`
	// Whether to get entity with extended data fields such as creationDate, lastUpdateDate, servingStatus
	IncludeExtendedDataFields *bool `json:"includeExtendedDataFields,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent instantiates a new SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent() *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent {
	this := SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent {
	this := SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetCampaignIdFilter() SponsoredProductsReducedObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret SponsoredProductsReducedObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsReducedObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given SponsoredProductsReducedObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) SetCampaignIdFilter(v SponsoredProductsReducedObjectIdFilter) {
	o.CampaignIdFilter = &v
}

// GetCampaignNegativeTargetIdFilter returns the CampaignNegativeTargetIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetCampaignNegativeTargetIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.CampaignNegativeTargetIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.CampaignNegativeTargetIdFilter
}

// GetCampaignNegativeTargetIdFilterOk returns a tuple with the CampaignNegativeTargetIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetCampaignNegativeTargetIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignNegativeTargetIdFilter) {
		return nil, false
	}
	return o.CampaignNegativeTargetIdFilter, true
}

// HasCampaignNegativeTargetIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) HasCampaignNegativeTargetIdFilter() bool {
	if o != nil && !IsNil(o.CampaignNegativeTargetIdFilter) {
		return true
	}

	return false
}

// SetCampaignNegativeTargetIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the CampaignNegativeTargetIdFilter field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) SetCampaignNegativeTargetIdFilter(v SponsoredProductsObjectIdFilter) {
	o.CampaignNegativeTargetIdFilter = &v
}

// GetStateFilter returns the StateFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetStateFilter() SponsoredProductsEntityStateFilter {
	if o == nil || IsNil(o.StateFilter) {
		var ret SponsoredProductsEntityStateFilter
		return ret
	}
	return *o.StateFilter
}

// GetStateFilterOk returns a tuple with the StateFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetStateFilterOk() (*SponsoredProductsEntityStateFilter, bool) {
	if o == nil || IsNil(o.StateFilter) {
		return nil, false
	}
	return o.StateFilter, true
}

// HasStateFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) HasStateFilter() bool {
	if o != nil && !IsNil(o.StateFilter) {
		return true
	}

	return false
}

// SetStateFilter gets a reference to the given SponsoredProductsEntityStateFilter and assigns it to the StateFilter field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) SetStateFilter(v SponsoredProductsEntityStateFilter) {
	o.StateFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetAsinFilter returns the AsinFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetAsinFilter() SponsoredProductsAsinFilter {
	if o == nil || IsNil(o.AsinFilter) {
		var ret SponsoredProductsAsinFilter
		return ret
	}
	return *o.AsinFilter
}

// GetAsinFilterOk returns a tuple with the AsinFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetAsinFilterOk() (*SponsoredProductsAsinFilter, bool) {
	if o == nil || IsNil(o.AsinFilter) {
		return nil, false
	}
	return o.AsinFilter, true
}

// HasAsinFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) HasAsinFilter() bool {
	if o != nil && !IsNil(o.AsinFilter) {
		return true
	}

	return false
}

// SetAsinFilter gets a reference to the given SponsoredProductsAsinFilter and assigns it to the AsinFilter field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) SetAsinFilter(v SponsoredProductsAsinFilter) {
	o.AsinFilter = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

func (o SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignIdFilter) {
		toSerialize["campaignIdFilter"] = o.CampaignIdFilter
	}
	if !IsNil(o.CampaignNegativeTargetIdFilter) {
		toSerialize["campaignNegativeTargetIdFilter"] = o.CampaignNegativeTargetIdFilter
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
	if !IsNil(o.AsinFilter) {
		toSerialize["asinFilter"] = o.AsinFilter
	}
	if !IsNil(o.IncludeExtendedDataFields) {
		toSerialize["includeExtendedDataFields"] = o.IncludeExtendedDataFields
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent struct {
	value *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) Get() *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) Set(val *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent(val *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) *NullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
