package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent{}

// SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent struct for SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent
type SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent struct {
	NegativeTargetingClauses SponsoredProductsBulkNegativeTargetingClauseOperationResponse `json:"negativeTargetingClauses"`
}

type _SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent

// NewSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent instantiates a new SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent(negativeTargetingClauses SponsoredProductsBulkNegativeTargetingClauseOperationResponse) *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent{}
	this.NegativeTargetingClauses = negativeTargetingClauses
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent{}
	return &this
}

// GetNegativeTargetingClauses returns the NegativeTargetingClauses field value
func (o *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent) GetNegativeTargetingClauses() SponsoredProductsBulkNegativeTargetingClauseOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkNegativeTargetingClauseOperationResponse
		return ret
	}

	return o.NegativeTargetingClauses
}

// GetNegativeTargetingClausesOk returns a tuple with the NegativeTargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent) GetNegativeTargetingClausesOk() (*SponsoredProductsBulkNegativeTargetingClauseOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeTargetingClauses, true
}

// SetNegativeTargetingClauses sets field value
func (o *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent) SetNegativeTargetingClauses(v SponsoredProductsBulkNegativeTargetingClauseOperationResponse) {
	o.NegativeTargetingClauses = v
}

func (o SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeTargetingClauses"] = o.NegativeTargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent struct {
	value *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent) Get() *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent) Set(val *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent(val *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent) *NullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
