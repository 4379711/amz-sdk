package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent{}

// SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent struct for SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent
type SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent struct {
	CampaignNegativeTargetingClauses SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse `json:"campaignNegativeTargetingClauses"`
}

type _SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent

// NewSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent(campaignNegativeTargetingClauses SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent{}
	this.CampaignNegativeTargetingClauses = campaignNegativeTargetingClauses
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent{}
	return &this
}

// GetCampaignNegativeTargetingClauses returns the CampaignNegativeTargetingClauses field value
func (o *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent) GetCampaignNegativeTargetingClauses() SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse
		return ret
	}

	return o.CampaignNegativeTargetingClauses
}

// GetCampaignNegativeTargetingClausesOk returns a tuple with the CampaignNegativeTargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent) GetCampaignNegativeTargetingClausesOk() (*SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignNegativeTargetingClauses, true
}

// SetCampaignNegativeTargetingClauses sets field value
func (o *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent) SetCampaignNegativeTargetingClauses(v SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) {
	o.CampaignNegativeTargetingClauses = v
}

func (o SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignNegativeTargetingClauses"] = o.CampaignNegativeTargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent(val *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
