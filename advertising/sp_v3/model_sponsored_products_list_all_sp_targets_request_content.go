package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListAllSPTargetsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListAllSPTargetsRequestContent{}

// SponsoredProductsListAllSPTargetsRequestContent struct for SponsoredProductsListAllSPTargetsRequestContent
type SponsoredProductsListAllSPTargetsRequestContent struct {
	CampaignIdFilter *SponsoredProductsObjectIdFilter    `json:"campaignIdFilter,omitempty"`
	StateFilter      *SponsoredProductsEntityStateFilter `json:"stateFilter,omitempty"`
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken       *string                          `json:"nextToken,omitempty"`
	TargetIdFilter  *SponsoredProductsObjectIdFilter `json:"targetIdFilter,omitempty"`
	AdGroupIdFilter *SponsoredProductsObjectIdFilter `json:"adGroupIdFilter,omitempty"`
}

// NewSponsoredProductsListAllSPTargetsRequestContent instantiates a new SponsoredProductsListAllSPTargetsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListAllSPTargetsRequestContent() *SponsoredProductsListAllSPTargetsRequestContent {
	this := SponsoredProductsListAllSPTargetsRequestContent{}
	return &this
}

// NewSponsoredProductsListAllSPTargetsRequestContentWithDefaults instantiates a new SponsoredProductsListAllSPTargetsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListAllSPTargetsRequestContentWithDefaults() *SponsoredProductsListAllSPTargetsRequestContent {
	this := SponsoredProductsListAllSPTargetsRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListAllSPTargetsRequestContent) GetCampaignIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListAllSPTargetsRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListAllSPTargetsRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *SponsoredProductsListAllSPTargetsRequestContent) SetCampaignIdFilter(v SponsoredProductsObjectIdFilter) {
	o.CampaignIdFilter = &v
}

// GetStateFilter returns the StateFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListAllSPTargetsRequestContent) GetStateFilter() SponsoredProductsEntityStateFilter {
	if o == nil || IsNil(o.StateFilter) {
		var ret SponsoredProductsEntityStateFilter
		return ret
	}
	return *o.StateFilter
}

// GetStateFilterOk returns a tuple with the StateFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListAllSPTargetsRequestContent) GetStateFilterOk() (*SponsoredProductsEntityStateFilter, bool) {
	if o == nil || IsNil(o.StateFilter) {
		return nil, false
	}
	return o.StateFilter, true
}

// HasStateFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListAllSPTargetsRequestContent) HasStateFilter() bool {
	if o != nil && !IsNil(o.StateFilter) {
		return true
	}

	return false
}

// SetStateFilter gets a reference to the given SponsoredProductsEntityStateFilter and assigns it to the StateFilter field.
func (o *SponsoredProductsListAllSPTargetsRequestContent) SetStateFilter(v SponsoredProductsEntityStateFilter) {
	o.StateFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListAllSPTargetsRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListAllSPTargetsRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListAllSPTargetsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListAllSPTargetsRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListAllSPTargetsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListAllSPTargetsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListAllSPTargetsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListAllSPTargetsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTargetIdFilter returns the TargetIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListAllSPTargetsRequestContent) GetTargetIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.TargetIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.TargetIdFilter
}

// GetTargetIdFilterOk returns a tuple with the TargetIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListAllSPTargetsRequestContent) GetTargetIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.TargetIdFilter) {
		return nil, false
	}
	return o.TargetIdFilter, true
}

// HasTargetIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListAllSPTargetsRequestContent) HasTargetIdFilter() bool {
	if o != nil && !IsNil(o.TargetIdFilter) {
		return true
	}

	return false
}

// SetTargetIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the TargetIdFilter field.
func (o *SponsoredProductsListAllSPTargetsRequestContent) SetTargetIdFilter(v SponsoredProductsObjectIdFilter) {
	o.TargetIdFilter = &v
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListAllSPTargetsRequestContent) GetAdGroupIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListAllSPTargetsRequestContent) GetAdGroupIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListAllSPTargetsRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *SponsoredProductsListAllSPTargetsRequestContent) SetAdGroupIdFilter(v SponsoredProductsObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

func (o SponsoredProductsListAllSPTargetsRequestContent) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.TargetIdFilter) {
		toSerialize["targetIdFilter"] = o.TargetIdFilter
	}
	if !IsNil(o.AdGroupIdFilter) {
		toSerialize["adGroupIdFilter"] = o.AdGroupIdFilter
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListAllSPTargetsRequestContent struct {
	value *SponsoredProductsListAllSPTargetsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListAllSPTargetsRequestContent) Get() *SponsoredProductsListAllSPTargetsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListAllSPTargetsRequestContent) Set(val *SponsoredProductsListAllSPTargetsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListAllSPTargetsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListAllSPTargetsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListAllSPTargetsRequestContent(val *SponsoredProductsListAllSPTargetsRequestContent) *NullableSponsoredProductsListAllSPTargetsRequestContent {
	return &NullableSponsoredProductsListAllSPTargetsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListAllSPTargetsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListAllSPTargetsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
