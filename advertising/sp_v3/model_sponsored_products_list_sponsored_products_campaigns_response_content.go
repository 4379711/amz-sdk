package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsCampaignsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsCampaignsResponseContent{}

// SponsoredProductsListSponsoredProductsCampaignsResponseContent struct for SponsoredProductsListSponsoredProductsCampaignsResponseContent
type SponsoredProductsListSponsoredProductsCampaignsResponseContent struct {
	// The total number of entities
	TotalResults *int64                      `json:"totalResults,omitempty"`
	Campaigns    []SponsoredProductsCampaign `json:"campaigns,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsCampaignsResponseContent instantiates a new SponsoredProductsListSponsoredProductsCampaignsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsCampaignsResponseContent() *SponsoredProductsListSponsoredProductsCampaignsResponseContent {
	this := SponsoredProductsListSponsoredProductsCampaignsResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsCampaignsResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsCampaignsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsCampaignsResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsCampaignsResponseContent {
	this := SponsoredProductsListSponsoredProductsCampaignsResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignsResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignsResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsCampaignsResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignsResponseContent) GetCampaigns() []SponsoredProductsCampaign {
	if o == nil || IsNil(o.Campaigns) {
		var ret []SponsoredProductsCampaign
		return ret
	}
	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignsResponseContent) GetCampaignsOk() ([]SponsoredProductsCampaign, bool) {
	if o == nil || IsNil(o.Campaigns) {
		return nil, false
	}
	return o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignsResponseContent) HasCampaigns() bool {
	if o != nil && !IsNil(o.Campaigns) {
		return true
	}

	return false
}

// SetCampaigns gets a reference to the given []SponsoredProductsCampaign and assigns it to the Campaigns field.
func (o *SponsoredProductsListSponsoredProductsCampaignsResponseContent) SetCampaigns(v []SponsoredProductsCampaign) {
	o.Campaigns = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsCampaignsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o SponsoredProductsListSponsoredProductsCampaignsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.Campaigns) {
		toSerialize["campaigns"] = o.Campaigns
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsCampaignsResponseContent struct {
	value *SponsoredProductsListSponsoredProductsCampaignsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsCampaignsResponseContent) Get() *SponsoredProductsListSponsoredProductsCampaignsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsCampaignsResponseContent) Set(val *SponsoredProductsListSponsoredProductsCampaignsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsCampaignsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsCampaignsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsCampaignsResponseContent(val *SponsoredProductsListSponsoredProductsCampaignsResponseContent) *NullableSponsoredProductsListSponsoredProductsCampaignsResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsCampaignsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsCampaignsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsCampaignsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
