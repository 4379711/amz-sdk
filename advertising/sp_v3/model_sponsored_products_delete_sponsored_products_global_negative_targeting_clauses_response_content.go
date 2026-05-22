package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent{}

// SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent struct for SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent
type SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent struct {
	NegativeTargetingClauses SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse `json:"negativeTargetingClauses"`
}

type _SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent

// NewSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent(negativeTargetingClauses SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) *SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent{}
	this.NegativeTargetingClauses = negativeTargetingClauses
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent{}
	return &this
}

// GetNegativeTargetingClauses returns the NegativeTargetingClauses field value
func (o *SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent) GetNegativeTargetingClauses() SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse
		return ret
	}

	return o.NegativeTargetingClauses
}

// GetNegativeTargetingClausesOk returns a tuple with the NegativeTargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent) GetNegativeTargetingClausesOk() (*SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeTargetingClauses, true
}

// SetNegativeTargetingClauses sets field value
func (o *SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent) SetNegativeTargetingClauses(v SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) {
	o.NegativeTargetingClauses = v
}

func (o SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeTargetingClauses"] = o.NegativeTargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent(val *SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
