package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent{}

// SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent struct for SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent
type SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent struct {
	NegativeTargetingClauses SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse `json:"negativeTargetingClauses"`
}

type _SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent

// NewSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent(negativeTargetingClauses SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) *SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent{}
	this.NegativeTargetingClauses = negativeTargetingClauses
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent{}
	return &this
}

// GetNegativeTargetingClauses returns the NegativeTargetingClauses field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent) GetNegativeTargetingClauses() SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse
		return ret
	}

	return o.NegativeTargetingClauses
}

// GetNegativeTargetingClausesOk returns a tuple with the NegativeTargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent) GetNegativeTargetingClausesOk() (*SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeTargetingClauses, true
}

// SetNegativeTargetingClauses sets field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent) SetNegativeTargetingClauses(v SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) {
	o.NegativeTargetingClauses = v
}

func (o SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeTargetingClauses"] = o.NegativeTargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent(val *SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
