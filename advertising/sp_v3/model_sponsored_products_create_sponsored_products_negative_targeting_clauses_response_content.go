package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent{}

// SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent struct for SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent
type SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent struct {
	NegativeTargetingClauses SponsoredProductsBulkNegativeTargetingClauseOperationResponse `json:"negativeTargetingClauses"`
}

type _SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent

// NewSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent(negativeTargetingClauses SponsoredProductsBulkNegativeTargetingClauseOperationResponse) *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent {
	this := SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent{}
	this.NegativeTargetingClauses = negativeTargetingClauses
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent {
	this := SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent{}
	return &this
}

// GetNegativeTargetingClauses returns the NegativeTargetingClauses field value
func (o *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent) GetNegativeTargetingClauses() SponsoredProductsBulkNegativeTargetingClauseOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkNegativeTargetingClauseOperationResponse
		return ret
	}

	return o.NegativeTargetingClauses
}

// GetNegativeTargetingClausesOk returns a tuple with the NegativeTargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent) GetNegativeTargetingClausesOk() (*SponsoredProductsBulkNegativeTargetingClauseOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeTargetingClauses, true
}

// SetNegativeTargetingClauses sets field value
func (o *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent) SetNegativeTargetingClauses(v SponsoredProductsBulkNegativeTargetingClauseOperationResponse) {
	o.NegativeTargetingClauses = v
}

func (o SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeTargetingClauses"] = o.NegativeTargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent) Get() *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent(val *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent) *NullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
