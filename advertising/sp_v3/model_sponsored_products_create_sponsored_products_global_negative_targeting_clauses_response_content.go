package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent{}

// SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent struct for SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent
type SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent struct {
	NegativeTargetingClauses SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse `json:"negativeTargetingClauses"`
}

type _SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent

// NewSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent(negativeTargetingClauses SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) *SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent{}
	this.NegativeTargetingClauses = negativeTargetingClauses
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent{}
	return &this
}

// GetNegativeTargetingClauses returns the NegativeTargetingClauses field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) GetNegativeTargetingClauses() SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse
		return ret
	}

	return o.NegativeTargetingClauses
}

// GetNegativeTargetingClausesOk returns a tuple with the NegativeTargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) GetNegativeTargetingClausesOk() (*SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeTargetingClauses, true
}

// SetNegativeTargetingClauses sets field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) SetNegativeTargetingClauses(v SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) {
	o.NegativeTargetingClauses = v
}

func (o SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeTargetingClauses"] = o.NegativeTargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) Get() *SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent(val *SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) *NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
