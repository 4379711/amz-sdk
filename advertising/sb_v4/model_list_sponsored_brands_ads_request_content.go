package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ListSponsoredBrandsAdsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSponsoredBrandsAdsRequestContent{}

// ListSponsoredBrandsAdsRequestContent struct for ListSponsoredBrandsAdsRequestContent
type ListSponsoredBrandsAdsRequestContent struct {
	CampaignIdFilter *ObjectIdFilter    `json:"campaignIdFilter,omitempty"`
	StateFilter      *EntityStateFilter `json:"stateFilter,omitempty"`
	// Number of records to include in the paginated response. Defaults to max page size for given API.
	MaxResults *float32 `json:"maxResults,omitempty"`
	// Token value allowing to navigate to the next response page.
	NextToken       *string         `json:"nextToken,omitempty"`
	AdIdFilter      *ObjectIdFilter `json:"adIdFilter,omitempty"`
	AdGroupIdFilter *ObjectIdFilter `json:"adGroupIdFilter,omitempty"`
	NameFilter      *NameFilter     `json:"nameFilter,omitempty"`
}

// NewListSponsoredBrandsAdsRequestContent instantiates a new ListSponsoredBrandsAdsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSponsoredBrandsAdsRequestContent() *ListSponsoredBrandsAdsRequestContent {
	this := ListSponsoredBrandsAdsRequestContent{}
	return &this
}

// NewListSponsoredBrandsAdsRequestContentWithDefaults instantiates a new ListSponsoredBrandsAdsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSponsoredBrandsAdsRequestContentWithDefaults() *ListSponsoredBrandsAdsRequestContent {
	this := ListSponsoredBrandsAdsRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdsRequestContent) GetCampaignIdFilter() ObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret ObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdsRequestContent) GetCampaignIdFilterOk() (*ObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdsRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given ObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *ListSponsoredBrandsAdsRequestContent) SetCampaignIdFilter(v ObjectIdFilter) {
	o.CampaignIdFilter = &v
}

// GetStateFilter returns the StateFilter field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdsRequestContent) GetStateFilter() EntityStateFilter {
	if o == nil || IsNil(o.StateFilter) {
		var ret EntityStateFilter
		return ret
	}
	return *o.StateFilter
}

// GetStateFilterOk returns a tuple with the StateFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdsRequestContent) GetStateFilterOk() (*EntityStateFilter, bool) {
	if o == nil || IsNil(o.StateFilter) {
		return nil, false
	}
	return o.StateFilter, true
}

// HasStateFilter returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdsRequestContent) HasStateFilter() bool {
	if o != nil && !IsNil(o.StateFilter) {
		return true
	}

	return false
}

// SetStateFilter gets a reference to the given EntityStateFilter and assigns it to the StateFilter field.
func (o *ListSponsoredBrandsAdsRequestContent) SetStateFilter(v EntityStateFilter) {
	o.StateFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdsRequestContent) GetMaxResults() float32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret float32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdsRequestContent) GetMaxResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given float32 and assigns it to the MaxResults field.
func (o *ListSponsoredBrandsAdsRequestContent) SetMaxResults(v float32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListSponsoredBrandsAdsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetAdIdFilter returns the AdIdFilter field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdsRequestContent) GetAdIdFilter() ObjectIdFilter {
	if o == nil || IsNil(o.AdIdFilter) {
		var ret ObjectIdFilter
		return ret
	}
	return *o.AdIdFilter
}

// GetAdIdFilterOk returns a tuple with the AdIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdsRequestContent) GetAdIdFilterOk() (*ObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdIdFilter) {
		return nil, false
	}
	return o.AdIdFilter, true
}

// HasAdIdFilter returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdsRequestContent) HasAdIdFilter() bool {
	if o != nil && !IsNil(o.AdIdFilter) {
		return true
	}

	return false
}

// SetAdIdFilter gets a reference to the given ObjectIdFilter and assigns it to the AdIdFilter field.
func (o *ListSponsoredBrandsAdsRequestContent) SetAdIdFilter(v ObjectIdFilter) {
	o.AdIdFilter = &v
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdsRequestContent) GetAdGroupIdFilter() ObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret ObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdsRequestContent) GetAdGroupIdFilterOk() (*ObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdsRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given ObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *ListSponsoredBrandsAdsRequestContent) SetAdGroupIdFilter(v ObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

// GetNameFilter returns the NameFilter field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdsRequestContent) GetNameFilter() NameFilter {
	if o == nil || IsNil(o.NameFilter) {
		var ret NameFilter
		return ret
	}
	return *o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdsRequestContent) GetNameFilterOk() (*NameFilter, bool) {
	if o == nil || IsNil(o.NameFilter) {
		return nil, false
	}
	return o.NameFilter, true
}

// HasNameFilter returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdsRequestContent) HasNameFilter() bool {
	if o != nil && !IsNil(o.NameFilter) {
		return true
	}

	return false
}

// SetNameFilter gets a reference to the given NameFilter and assigns it to the NameFilter field.
func (o *ListSponsoredBrandsAdsRequestContent) SetNameFilter(v NameFilter) {
	o.NameFilter = &v
}

func (o ListSponsoredBrandsAdsRequestContent) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AdIdFilter) {
		toSerialize["adIdFilter"] = o.AdIdFilter
	}
	if !IsNil(o.AdGroupIdFilter) {
		toSerialize["adGroupIdFilter"] = o.AdGroupIdFilter
	}
	if !IsNil(o.NameFilter) {
		toSerialize["nameFilter"] = o.NameFilter
	}
	return toSerialize, nil
}

type NullableListSponsoredBrandsAdsRequestContent struct {
	value *ListSponsoredBrandsAdsRequestContent
	isSet bool
}

func (v NullableListSponsoredBrandsAdsRequestContent) Get() *ListSponsoredBrandsAdsRequestContent {
	return v.value
}

func (v *NullableListSponsoredBrandsAdsRequestContent) Set(val *ListSponsoredBrandsAdsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListSponsoredBrandsAdsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListSponsoredBrandsAdsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSponsoredBrandsAdsRequestContent(val *ListSponsoredBrandsAdsRequestContent) *NullableListSponsoredBrandsAdsRequestContent {
	return &NullableListSponsoredBrandsAdsRequestContent{value: val, isSet: true}
}

func (v NullableListSponsoredBrandsAdsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListSponsoredBrandsAdsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
