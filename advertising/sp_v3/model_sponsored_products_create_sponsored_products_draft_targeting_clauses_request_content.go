package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent{}

// SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent struct for SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent
type SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent struct {
	// An array of draftTargetingClauses.
	TargetingClauses []SponsoredProductsCreateDraftTargetingClause `json:"targetingClauses"`
}

type _SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent

// NewSponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent(targetingClauses []SponsoredProductsCreateDraftTargetingClause) *SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent {
	this := SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent{}
	this.TargetingClauses = targetingClauses
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent {
	this := SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent{}
	return &this
}

// GetTargetingClauses returns the TargetingClauses field value
func (o *SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent) GetTargetingClauses() []SponsoredProductsCreateDraftTargetingClause {
	if o == nil {
		var ret []SponsoredProductsCreateDraftTargetingClause
		return ret
	}

	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent) GetTargetingClausesOk() ([]SponsoredProductsCreateDraftTargetingClause, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetingClauses, true
}

// SetTargetingClauses sets field value
func (o *SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent) SetTargetingClauses(v []SponsoredProductsCreateDraftTargetingClause) {
	o.TargetingClauses = v
}

func (o SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetingClauses"] = o.TargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent) Get() *SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent(val *SponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent) *NullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
