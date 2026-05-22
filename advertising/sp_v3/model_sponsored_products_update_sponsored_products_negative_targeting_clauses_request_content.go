package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent{}

// SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent struct for SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent
type SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent struct {
	// An array of negativeTargeting with updated values.
	NegativeTargetingClauses []SponsoredProductsUpdateNegativeTargetingClause `json:"negativeTargetingClauses"`
}

type _SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent

// NewSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent instantiates a new SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent(negativeTargetingClauses []SponsoredProductsUpdateNegativeTargetingClause) *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent{}
	this.NegativeTargetingClauses = negativeTargetingClauses
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent{}
	return &this
}

// GetNegativeTargetingClauses returns the NegativeTargetingClauses field value
func (o *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent) GetNegativeTargetingClauses() []SponsoredProductsUpdateNegativeTargetingClause {
	if o == nil {
		var ret []SponsoredProductsUpdateNegativeTargetingClause
		return ret
	}

	return o.NegativeTargetingClauses
}

// GetNegativeTargetingClausesOk returns a tuple with the NegativeTargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent) GetNegativeTargetingClausesOk() ([]SponsoredProductsUpdateNegativeTargetingClause, bool) {
	if o == nil {
		return nil, false
	}
	return o.NegativeTargetingClauses, true
}

// SetNegativeTargetingClauses sets field value
func (o *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent) SetNegativeTargetingClauses(v []SponsoredProductsUpdateNegativeTargetingClause) {
	o.NegativeTargetingClauses = v
}

func (o SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeTargetingClauses"] = o.NegativeTargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent struct {
	value *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent) Get() *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent) Set(val *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent(val *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent) *NullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
