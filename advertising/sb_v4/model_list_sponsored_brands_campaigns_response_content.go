package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ListSponsoredBrandsCampaignsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSponsoredBrandsCampaignsResponseContent{}

// ListSponsoredBrandsCampaignsResponseContent struct for ListSponsoredBrandsCampaignsResponseContent
type ListSponsoredBrandsCampaignsResponseContent struct {
	Campaigns []Campaign `json:"campaigns,omitempty"`
	// Token value allowing to navigate to the next response page.
	NextToken *string `json:"nextToken,omitempty"`
	// The total number of entities.
	TotalCount *float32 `json:"totalCount,omitempty"`
}

// NewListSponsoredBrandsCampaignsResponseContent instantiates a new ListSponsoredBrandsCampaignsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSponsoredBrandsCampaignsResponseContent() *ListSponsoredBrandsCampaignsResponseContent {
	this := ListSponsoredBrandsCampaignsResponseContent{}
	return &this
}

// NewListSponsoredBrandsCampaignsResponseContentWithDefaults instantiates a new ListSponsoredBrandsCampaignsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSponsoredBrandsCampaignsResponseContentWithDefaults() *ListSponsoredBrandsCampaignsResponseContent {
	this := ListSponsoredBrandsCampaignsResponseContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise.
func (o *ListSponsoredBrandsCampaignsResponseContent) GetCampaigns() []Campaign {
	if o == nil || IsNil(o.Campaigns) {
		var ret []Campaign
		return ret
	}
	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsCampaignsResponseContent) GetCampaignsOk() ([]Campaign, bool) {
	if o == nil || IsNil(o.Campaigns) {
		return nil, false
	}
	return o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *ListSponsoredBrandsCampaignsResponseContent) HasCampaigns() bool {
	if o != nil && !IsNil(o.Campaigns) {
		return true
	}

	return false
}

// SetCampaigns gets a reference to the given []Campaign and assigns it to the Campaigns field.
func (o *ListSponsoredBrandsCampaignsResponseContent) SetCampaigns(v []Campaign) {
	o.Campaigns = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListSponsoredBrandsCampaignsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsCampaignsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListSponsoredBrandsCampaignsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListSponsoredBrandsCampaignsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ListSponsoredBrandsCampaignsResponseContent) GetTotalCount() float32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret float32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsCampaignsResponseContent) GetTotalCountOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ListSponsoredBrandsCampaignsResponseContent) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given float32 and assigns it to the TotalCount field.
func (o *ListSponsoredBrandsCampaignsResponseContent) SetTotalCount(v float32) {
	o.TotalCount = &v
}

func (o ListSponsoredBrandsCampaignsResponseContent) ToMap() (map[string]interface{}, error) {
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

type NullableListSponsoredBrandsCampaignsResponseContent struct {
	value *ListSponsoredBrandsCampaignsResponseContent
	isSet bool
}

func (v NullableListSponsoredBrandsCampaignsResponseContent) Get() *ListSponsoredBrandsCampaignsResponseContent {
	return v.value
}

func (v *NullableListSponsoredBrandsCampaignsResponseContent) Set(val *ListSponsoredBrandsCampaignsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListSponsoredBrandsCampaignsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListSponsoredBrandsCampaignsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSponsoredBrandsCampaignsResponseContent(val *ListSponsoredBrandsCampaignsResponseContent) *NullableListSponsoredBrandsCampaignsResponseContent {
	return &NullableListSponsoredBrandsCampaignsResponseContent{value: val, isSet: true}
}

func (v NullableListSponsoredBrandsCampaignsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListSponsoredBrandsCampaignsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
