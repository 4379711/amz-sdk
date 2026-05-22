package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent{}

// SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent struct for SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent
type SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent struct {
	// An array of negativeTargetingClauses.
	NegativeTargetingClauses []SponsoredProductsCreateDraftNegativeTargetingClause `json:"negativeTargetingClauses"`
}

type _SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent

// NewSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent(negativeTargetingClauses []SponsoredProductsCreateDraftNegativeTargetingClause) *SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent {
	this := SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent{}
	this.NegativeTargetingClauses = negativeTargetingClauses
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent {
	this := SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent{}
	return &this
}

// GetNegativeTargetingClauses returns the NegativeTargetingClauses field value
func (o *SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent) GetNegativeTargetingClauses() []SponsoredProductsCreateDraftNegativeTargetingClause {
	if o == nil {
		var ret []SponsoredProductsCreateDraftNegativeTargetingClause
		return ret
	}

	return o.NegativeTargetingClauses
}

// GetNegativeTargetingClausesOk returns a tuple with the NegativeTargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent) GetNegativeTargetingClausesOk() ([]SponsoredProductsCreateDraftNegativeTargetingClause, bool) {
	if o == nil {
		return nil, false
	}
	return o.NegativeTargetingClauses, true
}

// SetNegativeTargetingClauses sets field value
func (o *SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent) SetNegativeTargetingClauses(v []SponsoredProductsCreateDraftNegativeTargetingClause) {
	o.NegativeTargetingClauses = v
}

func (o SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeTargetingClauses"] = o.NegativeTargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent) Get() *SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent(val *SponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent) *NullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftNegativeTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
