package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent{}

// SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent struct for SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent
type SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent struct {
	NegativeTargetIdFilter *SponsoredProductsObjectIdFilter        `json:"negativeTargetIdFilter,omitempty"`
	CampaignIdFilter       *SponsoredProductsReducedObjectIdFilter `json:"campaignIdFilter,omitempty"`
	StateFilter            *SponsoredProductsEntityStateFilter     `json:"stateFilter,omitempty"`
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken       *string                                 `json:"nextToken,omitempty"`
	AsinFilter      *SponsoredProductsAsinFilter            `json:"asinFilter,omitempty"`
	AdGroupIdFilter *SponsoredProductsReducedObjectIdFilter `json:"adGroupIdFilter,omitempty"`
	// Whether to get entity with extended data fields such as creationDate, lastUpdateDate, servingStatus
	IncludeExtendedDataFields *bool `json:"includeExtendedDataFields,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent instantiates a new SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent() *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent {
	this := SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent {
	this := SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent{}
	return &this
}

// GetNegativeTargetIdFilter returns the NegativeTargetIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) GetNegativeTargetIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.NegativeTargetIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.NegativeTargetIdFilter
}

// GetNegativeTargetIdFilterOk returns a tuple with the NegativeTargetIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) GetNegativeTargetIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.NegativeTargetIdFilter) {
		return nil, false
	}
	return o.NegativeTargetIdFilter, true
}

// HasNegativeTargetIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) HasNegativeTargetIdFilter() bool {
	if o != nil && !IsNil(o.NegativeTargetIdFilter) {
		return true
	}

	return false
}

// SetNegativeTargetIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the NegativeTargetIdFilter field.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) SetNegativeTargetIdFilter(v SponsoredProductsObjectIdFilter) {
	o.NegativeTargetIdFilter = &v
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) GetCampaignIdFilter() SponsoredProductsReducedObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret SponsoredProductsReducedObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsReducedObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given SponsoredProductsReducedObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) SetCampaignIdFilter(v SponsoredProductsReducedObjectIdFilter) {
	o.CampaignIdFilter = &v
}

// GetStateFilter returns the StateFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) GetStateFilter() SponsoredProductsEntityStateFilter {
	if o == nil || IsNil(o.StateFilter) {
		var ret SponsoredProductsEntityStateFilter
		return ret
	}
	return *o.StateFilter
}

// GetStateFilterOk returns a tuple with the StateFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) GetStateFilterOk() (*SponsoredProductsEntityStateFilter, bool) {
	if o == nil || IsNil(o.StateFilter) {
		return nil, false
	}
	return o.StateFilter, true
}

// HasStateFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) HasStateFilter() bool {
	if o != nil && !IsNil(o.StateFilter) {
		return true
	}

	return false
}

// SetStateFilter gets a reference to the given SponsoredProductsEntityStateFilter and assigns it to the StateFilter field.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) SetStateFilter(v SponsoredProductsEntityStateFilter) {
	o.StateFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetAsinFilter returns the AsinFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) GetAsinFilter() SponsoredProductsAsinFilter {
	if o == nil || IsNil(o.AsinFilter) {
		var ret SponsoredProductsAsinFilter
		return ret
	}
	return *o.AsinFilter
}

// GetAsinFilterOk returns a tuple with the AsinFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) GetAsinFilterOk() (*SponsoredProductsAsinFilter, bool) {
	if o == nil || IsNil(o.AsinFilter) {
		return nil, false
	}
	return o.AsinFilter, true
}

// HasAsinFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) HasAsinFilter() bool {
	if o != nil && !IsNil(o.AsinFilter) {
		return true
	}

	return false
}

// SetAsinFilter gets a reference to the given SponsoredProductsAsinFilter and assigns it to the AsinFilter field.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) SetAsinFilter(v SponsoredProductsAsinFilter) {
	o.AsinFilter = &v
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) GetAdGroupIdFilter() SponsoredProductsReducedObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret SponsoredProductsReducedObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) GetAdGroupIdFilterOk() (*SponsoredProductsReducedObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given SponsoredProductsReducedObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) SetAdGroupIdFilter(v SponsoredProductsReducedObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

func (o SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NegativeTargetIdFilter) {
		toSerialize["negativeTargetIdFilter"] = o.NegativeTargetIdFilter
	}
	if !IsNil(o.CampaignIdFilter) {
		toSerialize["campaignIdFilter"] = o.CampaignIdFilter
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
	if !IsNil(o.AdGroupIdFilter) {
		toSerialize["adGroupIdFilter"] = o.AdGroupIdFilter
	}
	if !IsNil(o.IncludeExtendedDataFields) {
		toSerialize["includeExtendedDataFields"] = o.IncludeExtendedDataFields
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent struct {
	value *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) Get() *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) Set(val *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent(val *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) *NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
