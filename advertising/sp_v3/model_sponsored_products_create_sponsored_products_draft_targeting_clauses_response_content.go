package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent{}

// SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent struct for SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent
type SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent struct {
	TargetingClauses SponsoredProductsBulkDraftTargetingClauseOperationResponse `json:"targetingClauses"`
}

type _SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent

// NewSponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent(targetingClauses SponsoredProductsBulkDraftTargetingClauseOperationResponse) *SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent {
	this := SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent{}
	this.TargetingClauses = targetingClauses
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent {
	this := SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent{}
	return &this
}

// GetTargetingClauses returns the TargetingClauses field value
func (o *SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent) GetTargetingClauses() SponsoredProductsBulkDraftTargetingClauseOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkDraftTargetingClauseOperationResponse
		return ret
	}

	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent) GetTargetingClausesOk() (*SponsoredProductsBulkDraftTargetingClauseOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingClauses, true
}

// SetTargetingClauses sets field value
func (o *SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent) SetTargetingClauses(v SponsoredProductsBulkDraftTargetingClauseOperationResponse) {
	o.TargetingClauses = v
}

func (o SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetingClauses"] = o.TargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent) Get() *SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent(val *SponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent) *NullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
