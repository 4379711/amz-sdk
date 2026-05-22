package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent{}

// SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent struct for SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent
type SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent struct {
	NegativeTargetingClauses SponsoredProductsBulkNegativeTargetingClauseOperationResponse `json:"negativeTargetingClauses"`
}

type _SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent

// NewSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent(negativeTargetingClauses SponsoredProductsBulkNegativeTargetingClauseOperationResponse) *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent{}
	this.NegativeTargetingClauses = negativeTargetingClauses
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent{}
	return &this
}

// GetNegativeTargetingClauses returns the NegativeTargetingClauses field value
func (o *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent) GetNegativeTargetingClauses() SponsoredProductsBulkNegativeTargetingClauseOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkNegativeTargetingClauseOperationResponse
		return ret
	}

	return o.NegativeTargetingClauses
}

// GetNegativeTargetingClausesOk returns a tuple with the NegativeTargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent) GetNegativeTargetingClausesOk() (*SponsoredProductsBulkNegativeTargetingClauseOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeTargetingClauses, true
}

// SetNegativeTargetingClauses sets field value
func (o *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent) SetNegativeTargetingClauses(v SponsoredProductsBulkNegativeTargetingClauseOperationResponse) {
	o.NegativeTargetingClauses = v
}

func (o SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeTargetingClauses"] = o.NegativeTargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent(val *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
