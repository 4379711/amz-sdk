package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent{}

// SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent struct for SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent
type SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent struct {
	TargetingClauses SponsoredProductsBulkDraftTargetingClauseOperationResponse `json:"targetingClauses"`
}

type _SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent

// NewSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent(targetingClauses SponsoredProductsBulkDraftTargetingClauseOperationResponse) *SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent{}
	this.TargetingClauses = targetingClauses
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent{}
	return &this
}

// GetTargetingClauses returns the TargetingClauses field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent) GetTargetingClauses() SponsoredProductsBulkDraftTargetingClauseOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkDraftTargetingClauseOperationResponse
		return ret
	}

	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent) GetTargetingClausesOk() (*SponsoredProductsBulkDraftTargetingClauseOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingClauses, true
}

// SetTargetingClauses sets field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent) SetTargetingClauses(v SponsoredProductsBulkDraftTargetingClauseOperationResponse) {
	o.TargetingClauses = v
}

func (o SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetingClauses"] = o.TargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent(val *SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
