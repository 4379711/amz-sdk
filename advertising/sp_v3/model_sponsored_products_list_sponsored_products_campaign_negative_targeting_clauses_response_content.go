package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent{}

// SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent struct for SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent
type SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent struct {
	// The total number of entities
	TotalResults *int64 `json:"totalResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken                        *string                                            `json:"nextToken,omitempty"`
	CampaignNegativeTargetingClauses []SponsoredProductsCampaignNegativeTargetingClause `json:"campaignNegativeTargetingClauses,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent instantiates a new SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent() *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent {
	this := SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent {
	this := SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetCampaignNegativeTargetingClauses returns the CampaignNegativeTargetingClauses field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) GetCampaignNegativeTargetingClauses() []SponsoredProductsCampaignNegativeTargetingClause {
	if o == nil || IsNil(o.CampaignNegativeTargetingClauses) {
		var ret []SponsoredProductsCampaignNegativeTargetingClause
		return ret
	}
	return o.CampaignNegativeTargetingClauses
}

// GetCampaignNegativeTargetingClausesOk returns a tuple with the CampaignNegativeTargetingClauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) GetCampaignNegativeTargetingClausesOk() ([]SponsoredProductsCampaignNegativeTargetingClause, bool) {
	if o == nil || IsNil(o.CampaignNegativeTargetingClauses) {
		return nil, false
	}
	return o.CampaignNegativeTargetingClauses, true
}

// HasCampaignNegativeTargetingClauses returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) HasCampaignNegativeTargetingClauses() bool {
	if o != nil && !IsNil(o.CampaignNegativeTargetingClauses) {
		return true
	}

	return false
}

// SetCampaignNegativeTargetingClauses gets a reference to the given []SponsoredProductsCampaignNegativeTargetingClause and assigns it to the CampaignNegativeTargetingClauses field.
func (o *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) SetCampaignNegativeTargetingClauses(v []SponsoredProductsCampaignNegativeTargetingClause) {
	o.CampaignNegativeTargetingClauses = v
}

func (o SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.CampaignNegativeTargetingClauses) {
		toSerialize["campaignNegativeTargetingClauses"] = o.CampaignNegativeTargetingClauses
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent struct {
	value *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) Get() *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) Set(val *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent(val *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) *NullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
