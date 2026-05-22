package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent{}

// SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent struct for SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent
type SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent struct {
	NegativeTargetingClauses SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse `json:"negativeTargetingClauses"`
}

type _SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent

// NewSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent(negativeTargetingClauses SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) *SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent {
	this := SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent{}
	this.NegativeTargetingClauses = negativeTargetingClauses
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent {
	this := SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent{}
	return &this
}

// GetNegativeTargetingClauses returns the NegativeTargetingClauses field value
func (o *SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent) GetNegativeTargetingClauses() SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse
		return ret
	}

	return o.NegativeTargetingClauses
}

// GetNegativeTargetingClausesOk returns a tuple with the NegativeTargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent) GetNegativeTargetingClausesOk() (*SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeTargetingClauses, true
}

// SetNegativeTargetingClauses sets field value
func (o *SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent) SetNegativeTargetingClauses(v SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) {
	o.NegativeTargetingClauses = v
}

func (o SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeTargetingClauses"] = o.NegativeTargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent) Get() *SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent(val *SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent) *NullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
