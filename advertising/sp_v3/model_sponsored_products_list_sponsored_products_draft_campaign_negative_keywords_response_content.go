package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent{}

// SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent struct for SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent
type SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent struct {
	// The total number of entities
	TotalResults *int64 `json:"totalResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken                *string                                         `json:"nextToken,omitempty"`
	CampaignNegativeKeywords []SponsoredProductsDraftCampaignNegativeKeyword `json:"campaignNegativeKeywords,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent instantiates a new SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent() *SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent {
	this := SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent {
	this := SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetCampaignNegativeKeywords returns the CampaignNegativeKeywords field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) GetCampaignNegativeKeywords() []SponsoredProductsDraftCampaignNegativeKeyword {
	if o == nil || IsNil(o.CampaignNegativeKeywords) {
		var ret []SponsoredProductsDraftCampaignNegativeKeyword
		return ret
	}
	return o.CampaignNegativeKeywords
}

// GetCampaignNegativeKeywordsOk returns a tuple with the CampaignNegativeKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) GetCampaignNegativeKeywordsOk() ([]SponsoredProductsDraftCampaignNegativeKeyword, bool) {
	if o == nil || IsNil(o.CampaignNegativeKeywords) {
		return nil, false
	}
	return o.CampaignNegativeKeywords, true
}

// HasCampaignNegativeKeywords returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) HasCampaignNegativeKeywords() bool {
	if o != nil && !IsNil(o.CampaignNegativeKeywords) {
		return true
	}

	return false
}

// SetCampaignNegativeKeywords gets a reference to the given []SponsoredProductsDraftCampaignNegativeKeyword and assigns it to the CampaignNegativeKeywords field.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) SetCampaignNegativeKeywords(v []SponsoredProductsDraftCampaignNegativeKeyword) {
	o.CampaignNegativeKeywords = v
}

func (o SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.CampaignNegativeKeywords) {
		toSerialize["campaignNegativeKeywords"] = o.CampaignNegativeKeywords
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent struct {
	value *SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) Get() *SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) Set(val *SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent(val *SponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) *NullableSponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
