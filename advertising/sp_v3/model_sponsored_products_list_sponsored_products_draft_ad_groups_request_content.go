package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent{}

// SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent struct for SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent
type SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent struct {
	CampaignIdFilter *SponsoredProductsReducedObjectIdFilter `json:"campaignIdFilter,omitempty"`
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

// NewSponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent instantiates a new SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent() *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent {
	this := SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsDraftAdGroupsRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsDraftAdGroupsRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent {
	this := SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) GetCampaignIdFilter() SponsoredProductsReducedObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret SponsoredProductsReducedObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsReducedObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given SponsoredProductsReducedObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) SetCampaignIdFilter(v SponsoredProductsReducedObjectIdFilter) {
	o.CampaignIdFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) GetAdGroupIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) GetAdGroupIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) SetAdGroupIdFilter(v SponsoredProductsObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

// GetNameFilter returns the NameFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) GetNameFilter() SponsoredProductsNameFilter {
	if o == nil || IsNil(o.NameFilter) {
		var ret SponsoredProductsNameFilter
		return ret
	}
	return *o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) GetNameFilterOk() (*SponsoredProductsNameFilter, bool) {
	if o == nil || IsNil(o.NameFilter) {
		return nil, false
	}
	return o.NameFilter, true
}

// HasNameFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) HasNameFilter() bool {
	if o != nil && !IsNil(o.NameFilter) {
		return true
	}

	return false
}

// SetNameFilter gets a reference to the given SponsoredProductsNameFilter and assigns it to the NameFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) SetNameFilter(v SponsoredProductsNameFilter) {
	o.NameFilter = &v
}

// GetCampaignTargetingTypeFilter returns the CampaignTargetingTypeFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) GetCampaignTargetingTypeFilter() SponsoredProductsTargetingType {
	if o == nil || IsNil(o.CampaignTargetingTypeFilter) {
		var ret SponsoredProductsTargetingType
		return ret
	}
	return *o.CampaignTargetingTypeFilter
}

// GetCampaignTargetingTypeFilterOk returns a tuple with the CampaignTargetingTypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) GetCampaignTargetingTypeFilterOk() (*SponsoredProductsTargetingType, bool) {
	if o == nil || IsNil(o.CampaignTargetingTypeFilter) {
		return nil, false
	}
	return o.CampaignTargetingTypeFilter, true
}

// HasCampaignTargetingTypeFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) HasCampaignTargetingTypeFilter() bool {
	if o != nil && !IsNil(o.CampaignTargetingTypeFilter) {
		return true
	}

	return false
}

// SetCampaignTargetingTypeFilter gets a reference to the given SponsoredProductsTargetingType and assigns it to the CampaignTargetingTypeFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) SetCampaignTargetingTypeFilter(v SponsoredProductsTargetingType) {
	o.CampaignTargetingTypeFilter = &v
}

func (o SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignIdFilter) {
		toSerialize["campaignIdFilter"] = o.CampaignIdFilter
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

type NullableSponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent struct {
	value *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) Get() *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) Set(val *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent(val *SponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) *NullableSponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftAdGroupsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
