package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent{}

// SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent struct for SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent
type SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent struct {
	CampaignNegativeTargetingClauses SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse `json:"campaignNegativeTargetingClauses"`
}

type _SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent

// NewSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent instantiates a new SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent(campaignNegativeTargetingClauses SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent{}
	this.CampaignNegativeTargetingClauses = campaignNegativeTargetingClauses
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent{}
	return &this
}

// GetCampaignNegativeTargetingClauses returns the CampaignNegativeTargetingClauses field value
func (o *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) GetCampaignNegativeTargetingClauses() SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse
		return ret
	}

	return o.CampaignNegativeTargetingClauses
}

// GetCampaignNegativeTargetingClausesOk returns a tuple with the CampaignNegativeTargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) GetCampaignNegativeTargetingClausesOk() (*SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignNegativeTargetingClauses, true
}

// SetCampaignNegativeTargetingClauses sets field value
func (o *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) SetCampaignNegativeTargetingClauses(v SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) {
	o.CampaignNegativeTargetingClauses = v
}

func (o SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignNegativeTargetingClauses"] = o.CampaignNegativeTargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent struct {
	value *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) Get() *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) Set(val *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent(val *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) *NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
