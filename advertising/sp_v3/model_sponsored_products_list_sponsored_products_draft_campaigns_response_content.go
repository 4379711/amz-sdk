package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent{}

// SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent struct for SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent
type SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent struct {
	// The total number of entities
	TotalResults *int64                           `json:"totalResults,omitempty"`
	Campaigns    []SponsoredProductsDraftCampaign `json:"campaigns,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsDraftCampaignsResponseContent instantiates a new SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsDraftCampaignsResponseContent() *SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent {
	this := SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsDraftCampaignsResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsDraftCampaignsResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent {
	this := SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) GetCampaigns() []SponsoredProductsDraftCampaign {
	if o == nil || IsNil(o.Campaigns) {
		var ret []SponsoredProductsDraftCampaign
		return ret
	}
	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) GetCampaignsOk() ([]SponsoredProductsDraftCampaign, bool) {
	if o == nil || IsNil(o.Campaigns) {
		return nil, false
	}
	return o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) HasCampaigns() bool {
	if o != nil && !IsNil(o.Campaigns) {
		return true
	}

	return false
}

// SetCampaigns gets a reference to the given []SponsoredProductsDraftCampaign and assigns it to the Campaigns field.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) SetCampaigns(v []SponsoredProductsDraftCampaign) {
	o.Campaigns = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsListSponsoredProductsDraftCampaignsResponseContent struct {
	value *SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) Get() *SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) Set(val *SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsDraftCampaignsResponseContent(val *SponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) *NullableSponsoredProductsListSponsoredProductsDraftCampaignsResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsDraftCampaignsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftCampaignsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
