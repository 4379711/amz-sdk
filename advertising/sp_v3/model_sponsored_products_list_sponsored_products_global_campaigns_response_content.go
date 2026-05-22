package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent{}

// SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent struct for SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent
type SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent struct {
	// The total number of entities
	TotalResults *int64                            `json:"totalResults,omitempty"`
	Campaigns    []SponsoredProductsGlobalCampaign `json:"campaigns,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent instantiates a new SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent() *SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent {
	this := SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsGlobalCampaignsResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsGlobalCampaignsResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent {
	this := SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) GetCampaigns() []SponsoredProductsGlobalCampaign {
	if o == nil || IsNil(o.Campaigns) {
		var ret []SponsoredProductsGlobalCampaign
		return ret
	}
	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) GetCampaignsOk() ([]SponsoredProductsGlobalCampaign, bool) {
	if o == nil || IsNil(o.Campaigns) {
		return nil, false
	}
	return o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) HasCampaigns() bool {
	if o != nil && !IsNil(o.Campaigns) {
		return true
	}

	return false
}

// SetCampaigns gets a reference to the given []SponsoredProductsGlobalCampaign and assigns it to the Campaigns field.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) SetCampaigns(v []SponsoredProductsGlobalCampaign) {
	o.Campaigns = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent struct {
	value *SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) Get() *SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) Set(val *SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent(val *SponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) *NullableSponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalCampaignsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
