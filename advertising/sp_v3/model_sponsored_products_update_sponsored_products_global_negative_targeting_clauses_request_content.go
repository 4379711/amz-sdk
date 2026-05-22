package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent{}

// SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent struct for SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent
type SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent struct {
	// An array of negativeTargeting with updated values.
	NegativeTargetingClauses []SponsoredProductsUpdateGlobalNegativeTargetingClause `json:"negativeTargetingClauses"`
}

type _SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent

// NewSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent(negativeTargetingClauses []SponsoredProductsUpdateGlobalNegativeTargetingClause) *SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent{}
	this.NegativeTargetingClauses = negativeTargetingClauses
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent{}
	return &this
}

// GetNegativeTargetingClauses returns the NegativeTargetingClauses field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) GetNegativeTargetingClauses() []SponsoredProductsUpdateGlobalNegativeTargetingClause {
	if o == nil {
		var ret []SponsoredProductsUpdateGlobalNegativeTargetingClause
		return ret
	}

	return o.NegativeTargetingClauses
}

// GetNegativeTargetingClausesOk returns a tuple with the NegativeTargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) GetNegativeTargetingClausesOk() ([]SponsoredProductsUpdateGlobalNegativeTargetingClause, bool) {
	if o == nil {
		return nil, false
	}
	return o.NegativeTargetingClauses, true
}

// SetNegativeTargetingClauses sets field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) SetNegativeTargetingClauses(v []SponsoredProductsUpdateGlobalNegativeTargetingClause) {
	o.NegativeTargetingClauses = v
}

func (o SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeTargetingClauses"] = o.NegativeTargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent struct {
	value *SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) Get() *SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) Set(val *SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent(val *SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) *NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
