package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ListSponsoredBrandsAdGroupsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSponsoredBrandsAdGroupsRequestContent{}

// ListSponsoredBrandsAdGroupsRequestContent struct for ListSponsoredBrandsAdGroupsRequestContent
type ListSponsoredBrandsAdGroupsRequestContent struct {
	CampaignIdFilter *ObjectIdFilter    `json:"campaignIdFilter,omitempty"`
	StateFilter      *EntityStateFilter `json:"stateFilter,omitempty"`
	// Number of records to include in the paginated response. Defaults to max page size for given API.
	MaxResults *float32 `json:"maxResults,omitempty"`
	// Token value allowing to navigate to the next response page.
	NextToken       *string         `json:"nextToken,omitempty"`
	AdGroupIdFilter *ObjectIdFilter `json:"adGroupIdFilter,omitempty"`
	// Setting to true will slow down performance because the API needs to retrieve extra information for each campaign.
	IncludeExtendedDataFields *bool       `json:"includeExtendedDataFields,omitempty"`
	NameFilter                *NameFilter `json:"nameFilter,omitempty"`
}

// NewListSponsoredBrandsAdGroupsRequestContent instantiates a new ListSponsoredBrandsAdGroupsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSponsoredBrandsAdGroupsRequestContent() *ListSponsoredBrandsAdGroupsRequestContent {
	this := ListSponsoredBrandsAdGroupsRequestContent{}
	return &this
}

// NewListSponsoredBrandsAdGroupsRequestContentWithDefaults instantiates a new ListSponsoredBrandsAdGroupsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSponsoredBrandsAdGroupsRequestContentWithDefaults() *ListSponsoredBrandsAdGroupsRequestContent {
	this := ListSponsoredBrandsAdGroupsRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdGroupsRequestContent) GetCampaignIdFilter() ObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret ObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdGroupsRequestContent) GetCampaignIdFilterOk() (*ObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdGroupsRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given ObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *ListSponsoredBrandsAdGroupsRequestContent) SetCampaignIdFilter(v ObjectIdFilter) {
	o.CampaignIdFilter = &v
}

// GetStateFilter returns the StateFilter field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdGroupsRequestContent) GetStateFilter() EntityStateFilter {
	if o == nil || IsNil(o.StateFilter) {
		var ret EntityStateFilter
		return ret
	}
	return *o.StateFilter
}

// GetStateFilterOk returns a tuple with the StateFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdGroupsRequestContent) GetStateFilterOk() (*EntityStateFilter, bool) {
	if o == nil || IsNil(o.StateFilter) {
		return nil, false
	}
	return o.StateFilter, true
}

// HasStateFilter returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdGroupsRequestContent) HasStateFilter() bool {
	if o != nil && !IsNil(o.StateFilter) {
		return true
	}

	return false
}

// SetStateFilter gets a reference to the given EntityStateFilter and assigns it to the StateFilter field.
func (o *ListSponsoredBrandsAdGroupsRequestContent) SetStateFilter(v EntityStateFilter) {
	o.StateFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdGroupsRequestContent) GetMaxResults() float32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret float32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdGroupsRequestContent) GetMaxResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdGroupsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given float32 and assigns it to the MaxResults field.
func (o *ListSponsoredBrandsAdGroupsRequestContent) SetMaxResults(v float32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdGroupsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdGroupsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdGroupsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListSponsoredBrandsAdGroupsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdGroupsRequestContent) GetAdGroupIdFilter() ObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret ObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdGroupsRequestContent) GetAdGroupIdFilterOk() (*ObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdGroupsRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given ObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *ListSponsoredBrandsAdGroupsRequestContent) SetAdGroupIdFilter(v ObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdGroupsRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdGroupsRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdGroupsRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *ListSponsoredBrandsAdGroupsRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

// GetNameFilter returns the NameFilter field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdGroupsRequestContent) GetNameFilter() NameFilter {
	if o == nil || IsNil(o.NameFilter) {
		var ret NameFilter
		return ret
	}
	return *o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdGroupsRequestContent) GetNameFilterOk() (*NameFilter, bool) {
	if o == nil || IsNil(o.NameFilter) {
		return nil, false
	}
	return o.NameFilter, true
}

// HasNameFilter returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdGroupsRequestContent) HasNameFilter() bool {
	if o != nil && !IsNil(o.NameFilter) {
		return true
	}

	return false
}

// SetNameFilter gets a reference to the given NameFilter and assigns it to the NameFilter field.
func (o *ListSponsoredBrandsAdGroupsRequestContent) SetNameFilter(v NameFilter) {
	o.NameFilter = &v
}

func (o ListSponsoredBrandsAdGroupsRequestContent) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableListSponsoredBrandsAdGroupsRequestContent struct {
	value *ListSponsoredBrandsAdGroupsRequestContent
	isSet bool
}

func (v NullableListSponsoredBrandsAdGroupsRequestContent) Get() *ListSponsoredBrandsAdGroupsRequestContent {
	return v.value
}

func (v *NullableListSponsoredBrandsAdGroupsRequestContent) Set(val *ListSponsoredBrandsAdGroupsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListSponsoredBrandsAdGroupsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListSponsoredBrandsAdGroupsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSponsoredBrandsAdGroupsRequestContent(val *ListSponsoredBrandsAdGroupsRequestContent) *NullableListSponsoredBrandsAdGroupsRequestContent {
	return &NullableListSponsoredBrandsAdGroupsRequestContent{value: val, isSet: true}
}

func (v NullableListSponsoredBrandsAdGroupsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListSponsoredBrandsAdGroupsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
