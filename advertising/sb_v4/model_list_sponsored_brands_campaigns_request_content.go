package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ListSponsoredBrandsCampaignsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSponsoredBrandsCampaignsRequestContent{}

// ListSponsoredBrandsCampaignsRequestContent struct for ListSponsoredBrandsCampaignsRequestContent
type ListSponsoredBrandsCampaignsRequestContent struct {
	CampaignIdFilter  *ObjectIdFilter    `json:"campaignIdFilter,omitempty"`
	PortfolioIdFilter *ObjectIdFilter    `json:"portfolioIdFilter,omitempty"`
	StateFilter       *EntityStateFilter `json:"stateFilter,omitempty"`
	// Number of records to include in the paginated response. Defaults to max page size for given API.
	MaxResults *float32 `json:"maxResults,omitempty"`
	// Token value allowing to navigate to the next response page.
	NextToken      *string         `json:"nextToken,omitempty"`
	GoalTypeFilter *GoalTypeFilter `json:"goalTypeFilter,omitempty"`
	// Setting to true will slow down performance because the API needs to retrieve extra information for each campaign.
	IncludeExtendedDataFields *bool       `json:"includeExtendedDataFields,omitempty"`
	NameFilter                *NameFilter `json:"nameFilter,omitempty"`
}

// NewListSponsoredBrandsCampaignsRequestContent instantiates a new ListSponsoredBrandsCampaignsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSponsoredBrandsCampaignsRequestContent() *ListSponsoredBrandsCampaignsRequestContent {
	this := ListSponsoredBrandsCampaignsRequestContent{}
	return &this
}

// NewListSponsoredBrandsCampaignsRequestContentWithDefaults instantiates a new ListSponsoredBrandsCampaignsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSponsoredBrandsCampaignsRequestContentWithDefaults() *ListSponsoredBrandsCampaignsRequestContent {
	this := ListSponsoredBrandsCampaignsRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *ListSponsoredBrandsCampaignsRequestContent) GetCampaignIdFilter() ObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret ObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsCampaignsRequestContent) GetCampaignIdFilterOk() (*ObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *ListSponsoredBrandsCampaignsRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given ObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *ListSponsoredBrandsCampaignsRequestContent) SetCampaignIdFilter(v ObjectIdFilter) {
	o.CampaignIdFilter = &v
}

// GetPortfolioIdFilter returns the PortfolioIdFilter field value if set, zero value otherwise.
func (o *ListSponsoredBrandsCampaignsRequestContent) GetPortfolioIdFilter() ObjectIdFilter {
	if o == nil || IsNil(o.PortfolioIdFilter) {
		var ret ObjectIdFilter
		return ret
	}
	return *o.PortfolioIdFilter
}

// GetPortfolioIdFilterOk returns a tuple with the PortfolioIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsCampaignsRequestContent) GetPortfolioIdFilterOk() (*ObjectIdFilter, bool) {
	if o == nil || IsNil(o.PortfolioIdFilter) {
		return nil, false
	}
	return o.PortfolioIdFilter, true
}

// HasPortfolioIdFilter returns a boolean if a field has been set.
func (o *ListSponsoredBrandsCampaignsRequestContent) HasPortfolioIdFilter() bool {
	if o != nil && !IsNil(o.PortfolioIdFilter) {
		return true
	}

	return false
}

// SetPortfolioIdFilter gets a reference to the given ObjectIdFilter and assigns it to the PortfolioIdFilter field.
func (o *ListSponsoredBrandsCampaignsRequestContent) SetPortfolioIdFilter(v ObjectIdFilter) {
	o.PortfolioIdFilter = &v
}

// GetStateFilter returns the StateFilter field value if set, zero value otherwise.
func (o *ListSponsoredBrandsCampaignsRequestContent) GetStateFilter() EntityStateFilter {
	if o == nil || IsNil(o.StateFilter) {
		var ret EntityStateFilter
		return ret
	}
	return *o.StateFilter
}

// GetStateFilterOk returns a tuple with the StateFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsCampaignsRequestContent) GetStateFilterOk() (*EntityStateFilter, bool) {
	if o == nil || IsNil(o.StateFilter) {
		return nil, false
	}
	return o.StateFilter, true
}

// HasStateFilter returns a boolean if a field has been set.
func (o *ListSponsoredBrandsCampaignsRequestContent) HasStateFilter() bool {
	if o != nil && !IsNil(o.StateFilter) {
		return true
	}

	return false
}

// SetStateFilter gets a reference to the given EntityStateFilter and assigns it to the StateFilter field.
func (o *ListSponsoredBrandsCampaignsRequestContent) SetStateFilter(v EntityStateFilter) {
	o.StateFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *ListSponsoredBrandsCampaignsRequestContent) GetMaxResults() float32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret float32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsCampaignsRequestContent) GetMaxResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *ListSponsoredBrandsCampaignsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given float32 and assigns it to the MaxResults field.
func (o *ListSponsoredBrandsCampaignsRequestContent) SetMaxResults(v float32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListSponsoredBrandsCampaignsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsCampaignsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListSponsoredBrandsCampaignsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListSponsoredBrandsCampaignsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetGoalTypeFilter returns the GoalTypeFilter field value if set, zero value otherwise.
func (o *ListSponsoredBrandsCampaignsRequestContent) GetGoalTypeFilter() GoalTypeFilter {
	if o == nil || IsNil(o.GoalTypeFilter) {
		var ret GoalTypeFilter
		return ret
	}
	return *o.GoalTypeFilter
}

// GetGoalTypeFilterOk returns a tuple with the GoalTypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsCampaignsRequestContent) GetGoalTypeFilterOk() (*GoalTypeFilter, bool) {
	if o == nil || IsNil(o.GoalTypeFilter) {
		return nil, false
	}
	return o.GoalTypeFilter, true
}

// HasGoalTypeFilter returns a boolean if a field has been set.
func (o *ListSponsoredBrandsCampaignsRequestContent) HasGoalTypeFilter() bool {
	if o != nil && !IsNil(o.GoalTypeFilter) {
		return true
	}

	return false
}

// SetGoalTypeFilter gets a reference to the given GoalTypeFilter and assigns it to the GoalTypeFilter field.
func (o *ListSponsoredBrandsCampaignsRequestContent) SetGoalTypeFilter(v GoalTypeFilter) {
	o.GoalTypeFilter = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *ListSponsoredBrandsCampaignsRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsCampaignsRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *ListSponsoredBrandsCampaignsRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *ListSponsoredBrandsCampaignsRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

// GetNameFilter returns the NameFilter field value if set, zero value otherwise.
func (o *ListSponsoredBrandsCampaignsRequestContent) GetNameFilter() NameFilter {
	if o == nil || IsNil(o.NameFilter) {
		var ret NameFilter
		return ret
	}
	return *o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsCampaignsRequestContent) GetNameFilterOk() (*NameFilter, bool) {
	if o == nil || IsNil(o.NameFilter) {
		return nil, false
	}
	return o.NameFilter, true
}

// HasNameFilter returns a boolean if a field has been set.
func (o *ListSponsoredBrandsCampaignsRequestContent) HasNameFilter() bool {
	if o != nil && !IsNil(o.NameFilter) {
		return true
	}

	return false
}

// SetNameFilter gets a reference to the given NameFilter and assigns it to the NameFilter field.
func (o *ListSponsoredBrandsCampaignsRequestContent) SetNameFilter(v NameFilter) {
	o.NameFilter = &v
}

func (o ListSponsoredBrandsCampaignsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignIdFilter) {
		toSerialize["campaignIdFilter"] = o.CampaignIdFilter
	}
	if !IsNil(o.PortfolioIdFilter) {
		toSerialize["portfolioIdFilter"] = o.PortfolioIdFilter
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
	if !IsNil(o.GoalTypeFilter) {
		toSerialize["goalTypeFilter"] = o.GoalTypeFilter
	}
	if !IsNil(o.IncludeExtendedDataFields) {
		toSerialize["includeExtendedDataFields"] = o.IncludeExtendedDataFields
	}
	if !IsNil(o.NameFilter) {
		toSerialize["nameFilter"] = o.NameFilter
	}
	return toSerialize, nil
}

type NullableListSponsoredBrandsCampaignsRequestContent struct {
	value *ListSponsoredBrandsCampaignsRequestContent
	isSet bool
}

func (v NullableListSponsoredBrandsCampaignsRequestContent) Get() *ListSponsoredBrandsCampaignsRequestContent {
	return v.value
}

func (v *NullableListSponsoredBrandsCampaignsRequestContent) Set(val *ListSponsoredBrandsCampaignsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListSponsoredBrandsCampaignsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListSponsoredBrandsCampaignsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSponsoredBrandsCampaignsRequestContent(val *ListSponsoredBrandsCampaignsRequestContent) *NullableListSponsoredBrandsCampaignsRequestContent {
	return &NullableListSponsoredBrandsCampaignsRequestContent{value: val, isSet: true}
}

func (v NullableListSponsoredBrandsCampaignsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListSponsoredBrandsCampaignsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
