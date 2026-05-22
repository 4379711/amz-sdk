package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ListSponsoredBrandsCampaignsBetaResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSponsoredBrandsCampaignsBetaResponseContent{}

// ListSponsoredBrandsCampaignsBetaResponseContent struct for ListSponsoredBrandsCampaignsBetaResponseContent
type ListSponsoredBrandsCampaignsBetaResponseContent struct {
	Campaigns []Campaign `json:"campaigns,omitempty"`
	// Token value allowing to navigate to the next response page.
	NextToken *string `json:"nextToken,omitempty"`
	// The total number of entities.
	TotalCount *float32 `json:"totalCount,omitempty"`
}

// NewListSponsoredBrandsCampaignsBetaResponseContent instantiates a new ListSponsoredBrandsCampaignsBetaResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSponsoredBrandsCampaignsBetaResponseContent() *ListSponsoredBrandsCampaignsBetaResponseContent {
	this := ListSponsoredBrandsCampaignsBetaResponseContent{}
	return &this
}

// NewListSponsoredBrandsCampaignsBetaResponseContentWithDefaults instantiates a new ListSponsoredBrandsCampaignsBetaResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSponsoredBrandsCampaignsBetaResponseContentWithDefaults() *ListSponsoredBrandsCampaignsBetaResponseContent {
	this := ListSponsoredBrandsCampaignsBetaResponseContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise.
func (o *ListSponsoredBrandsCampaignsBetaResponseContent) GetCampaigns() []Campaign {
	if o == nil || IsNil(o.Campaigns) {
		var ret []Campaign
		return ret
	}
	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsCampaignsBetaResponseContent) GetCampaignsOk() ([]Campaign, bool) {
	if o == nil || IsNil(o.Campaigns) {
		return nil, false
	}
	return o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *ListSponsoredBrandsCampaignsBetaResponseContent) HasCampaigns() bool {
	if o != nil && !IsNil(o.Campaigns) {
		return true
	}

	return false
}

// SetCampaigns gets a reference to the given []Campaign and assigns it to the Campaigns field.
func (o *ListSponsoredBrandsCampaignsBetaResponseContent) SetCampaigns(v []Campaign) {
	o.Campaigns = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListSponsoredBrandsCampaignsBetaResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsCampaignsBetaResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListSponsoredBrandsCampaignsBetaResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListSponsoredBrandsCampaignsBetaResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ListSponsoredBrandsCampaignsBetaResponseContent) GetTotalCount() float32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret float32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsCampaignsBetaResponseContent) GetTotalCountOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ListSponsoredBrandsCampaignsBetaResponseContent) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given float32 and assigns it to the TotalCount field.
func (o *ListSponsoredBrandsCampaignsBetaResponseContent) SetTotalCount(v float32) {
	o.TotalCount = &v
}

func (o ListSponsoredBrandsCampaignsBetaResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Campaigns) {
		toSerialize["campaigns"] = o.Campaigns
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableListSponsoredBrandsCampaignsBetaResponseContent struct {
	value *ListSponsoredBrandsCampaignsBetaResponseContent
	isSet bool
}

func (v NullableListSponsoredBrandsCampaignsBetaResponseContent) Get() *ListSponsoredBrandsCampaignsBetaResponseContent {
	return v.value
}

func (v *NullableListSponsoredBrandsCampaignsBetaResponseContent) Set(val *ListSponsoredBrandsCampaignsBetaResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListSponsoredBrandsCampaignsBetaResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListSponsoredBrandsCampaignsBetaResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSponsoredBrandsCampaignsBetaResponseContent(val *ListSponsoredBrandsCampaignsBetaResponseContent) *NullableListSponsoredBrandsCampaignsBetaResponseContent {
	return &NullableListSponsoredBrandsCampaignsBetaResponseContent{value: val, isSet: true}
}

func (v NullableListSponsoredBrandsCampaignsBetaResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListSponsoredBrandsCampaignsBetaResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
