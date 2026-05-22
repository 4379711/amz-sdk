package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent{}

// SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent struct for SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent
type SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent struct {
	// The total number of entities
	TotalResults     *int64                                          `json:"totalResults,omitempty"`
	CampaignStatuses []SponsoredProductsDraftCampaignPromotionStatus `json:"campaignStatuses,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent instantiates a new SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent() *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent {
	this := SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent {
	this := SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetCampaignStatuses returns the CampaignStatuses field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) GetCampaignStatuses() []SponsoredProductsDraftCampaignPromotionStatus {
	if o == nil || IsNil(o.CampaignStatuses) {
		var ret []SponsoredProductsDraftCampaignPromotionStatus
		return ret
	}
	return o.CampaignStatuses
}

// GetCampaignStatusesOk returns a tuple with the CampaignStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) GetCampaignStatusesOk() ([]SponsoredProductsDraftCampaignPromotionStatus, bool) {
	if o == nil || IsNil(o.CampaignStatuses) {
		return nil, false
	}
	return o.CampaignStatuses, true
}

// HasCampaignStatuses returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) HasCampaignStatuses() bool {
	if o != nil && !IsNil(o.CampaignStatuses) {
		return true
	}

	return false
}

// SetCampaignStatuses gets a reference to the given []SponsoredProductsDraftCampaignPromotionStatus and assigns it to the CampaignStatuses field.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) SetCampaignStatuses(v []SponsoredProductsDraftCampaignPromotionStatus) {
	o.CampaignStatuses = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.CampaignStatuses) {
		toSerialize["campaignStatuses"] = o.CampaignStatuses
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent struct {
	value *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) Get() *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) Set(val *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent(val *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) *NullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
