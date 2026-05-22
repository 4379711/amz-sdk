package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsAdGroupsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsAdGroupsRequestContent{}

// SponsoredProductsListSponsoredProductsAdGroupsRequestContent struct for SponsoredProductsListSponsoredProductsAdGroupsRequestContent
type SponsoredProductsListSponsoredProductsAdGroupsRequestContent struct {
	CampaignIdFilter *SponsoredProductsReducedObjectIdFilter `json:"campaignIdFilter,omitempty"`
	StateFilter      *SponsoredProductsEntityStateFilter     `json:"stateFilter,omitempty"`
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken       *string                          `json:"nextToken,omitempty"`
	AdGroupIdFilter *SponsoredProductsObjectIdFilter `json:"adGroupIdFilter,omitempty"`
	// Whether to get entity with extended data fields such as creationDate, lastUpdateDate, servingStatus
	IncludeExtendedDataFields   *bool                           `json:"includeExtendedDataFields,omitempty"`
	NameFilter                  *SponsoredProductsNameFilter    `json:"nameFilter,omitempty"`
	CampaignTargetingTypeFilter *SponsoredProductsTargetingType `json:"campaignTargetingTypeFilter,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsAdGroupsRequestContent instantiates a new SponsoredProductsListSponsoredProductsAdGroupsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsAdGroupsRequestContent() *SponsoredProductsListSponsoredProductsAdGroupsRequestContent {
	this := SponsoredProductsListSponsoredProductsAdGroupsRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsAdGroupsRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsAdGroupsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsAdGroupsRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsAdGroupsRequestContent {
	this := SponsoredProductsListSponsoredProductsAdGroupsRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) GetCampaignIdFilter() SponsoredProductsReducedObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret SponsoredProductsReducedObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsReducedObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given SponsoredProductsReducedObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) SetCampaignIdFilter(v SponsoredProductsReducedObjectIdFilter) {
	o.CampaignIdFilter = &v
}

// GetStateFilter returns the StateFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) GetStateFilter() SponsoredProductsEntityStateFilter {
	if o == nil || IsNil(o.StateFilter) {
		var ret SponsoredProductsEntityStateFilter
		return ret
	}
	return *o.StateFilter
}

// GetStateFilterOk returns a tuple with the StateFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) GetStateFilterOk() (*SponsoredProductsEntityStateFilter, bool) {
	if o == nil || IsNil(o.StateFilter) {
		return nil, false
	}
	return o.StateFilter, true
}

// HasStateFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) HasStateFilter() bool {
	if o != nil && !IsNil(o.StateFilter) {
		return true
	}

	return false
}

// SetStateFilter gets a reference to the given SponsoredProductsEntityStateFilter and assigns it to the StateFilter field.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) SetStateFilter(v SponsoredProductsEntityStateFilter) {
	o.StateFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) GetAdGroupIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) GetAdGroupIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) SetAdGroupIdFilter(v SponsoredProductsObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

// GetNameFilter returns the NameFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) GetNameFilter() SponsoredProductsNameFilter {
	if o == nil || IsNil(o.NameFilter) {
		var ret SponsoredProductsNameFilter
		return ret
	}
	return *o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) GetNameFilterOk() (*SponsoredProductsNameFilter, bool) {
	if o == nil || IsNil(o.NameFilter) {
		return nil, false
	}
	return o.NameFilter, true
}

// HasNameFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) HasNameFilter() bool {
	if o != nil && !IsNil(o.NameFilter) {
		return true
	}

	return false
}

// SetNameFilter gets a reference to the given SponsoredProductsNameFilter and assigns it to the NameFilter field.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) SetNameFilter(v SponsoredProductsNameFilter) {
	o.NameFilter = &v
}

// GetCampaignTargetingTypeFilter returns the CampaignTargetingTypeFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) GetCampaignTargetingTypeFilter() SponsoredProductsTargetingType {
	if o == nil || IsNil(o.CampaignTargetingTypeFilter) {
		var ret SponsoredProductsTargetingType
		return ret
	}
	return *o.CampaignTargetingTypeFilter
}

// GetCampaignTargetingTypeFilterOk returns a tuple with the CampaignTargetingTypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) GetCampaignTargetingTypeFilterOk() (*SponsoredProductsTargetingType, bool) {
	if o == nil || IsNil(o.CampaignTargetingTypeFilter) {
		return nil, false
	}
	return o.CampaignTargetingTypeFilter, true
}

// HasCampaignTargetingTypeFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) HasCampaignTargetingTypeFilter() bool {
	if o != nil && !IsNil(o.CampaignTargetingTypeFilter) {
		return true
	}

	return false
}

// SetCampaignTargetingTypeFilter gets a reference to the given SponsoredProductsTargetingType and assigns it to the CampaignTargetingTypeFilter field.
func (o *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) SetCampaignTargetingTypeFilter(v SponsoredProductsTargetingType) {
	o.CampaignTargetingTypeFilter = &v
}

func (o SponsoredProductsListSponsoredProductsAdGroupsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.AdGroupIdFilter) {
		toSerialize["adGroupIdFilter"] = o.AdGroupIdFilter
	}
	if !IsNil(o.IncludeExtendedDataFields) {
		toSerialize["includeExtendedDataFields"] = o.IncludeExtendedDataFields
	}
	if !IsNil(o.NameFilter) {
		toSerialize["nameFilter"] = o.NameFilter
	}
	if !IsNil(o.CampaignTargetingTypeFilter) {
		toSerialize["campaignTargetingTypeFilter"] = o.CampaignTargetingTypeFilter
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsAdGroupsRequestContent struct {
	value *SponsoredProductsListSponsoredProductsAdGroupsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsAdGroupsRequestContent) Get() *SponsoredProductsListSponsoredProductsAdGroupsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsAdGroupsRequestContent) Set(val *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsAdGroupsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsAdGroupsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsAdGroupsRequestContent(val *SponsoredProductsListSponsoredProductsAdGroupsRequestContent) *NullableSponsoredProductsListSponsoredProductsAdGroupsRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsAdGroupsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsAdGroupsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsAdGroupsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
