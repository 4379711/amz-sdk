package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent{}

// SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent struct for SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent
type SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent struct {
	// The total number of entities
	TotalResults *int64 `json:"totalResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken                *string                                    `json:"nextToken,omitempty"`
	CampaignNegativeKeywords []SponsoredProductsCampaignNegativeKeyword `json:"campaignNegativeKeywords,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent instantiates a new SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent() *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent {
	this := SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent {
	this := SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetCampaignNegativeKeywords returns the CampaignNegativeKeywords field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) GetCampaignNegativeKeywords() []SponsoredProductsCampaignNegativeKeyword {
	if o == nil || IsNil(o.CampaignNegativeKeywords) {
		var ret []SponsoredProductsCampaignNegativeKeyword
		return ret
	}
	return o.CampaignNegativeKeywords
}

// GetCampaignNegativeKeywordsOk returns a tuple with the CampaignNegativeKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) GetCampaignNegativeKeywordsOk() ([]SponsoredProductsCampaignNegativeKeyword, bool) {
	if o == nil || IsNil(o.CampaignNegativeKeywords) {
		return nil, false
	}
	return o.CampaignNegativeKeywords, true
}

// HasCampaignNegativeKeywords returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) HasCampaignNegativeKeywords() bool {
	if o != nil && !IsNil(o.CampaignNegativeKeywords) {
		return true
	}

	return false
}

// SetCampaignNegativeKeywords gets a reference to the given []SponsoredProductsCampaignNegativeKeyword and assigns it to the CampaignNegativeKeywords field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) SetCampaignNegativeKeywords(v []SponsoredProductsCampaignNegativeKeyword) {
	o.CampaignNegativeKeywords = v
}

func (o SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent struct {
	value *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) Get() *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) Set(val *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent(val *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) *NullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
