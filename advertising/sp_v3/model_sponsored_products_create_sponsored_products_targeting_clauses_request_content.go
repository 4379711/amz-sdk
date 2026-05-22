package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent{}

// SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent struct for SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent
type SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent struct {
	// An array of targetingClauses.
	TargetingClauses []SponsoredProductsCreateTargetingClause `json:"targetingClauses"`
}

type _SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent

// NewSponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent(targetingClauses []SponsoredProductsCreateTargetingClause) *SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent {
	this := SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent{}
	this.TargetingClauses = targetingClauses
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsTargetingClausesRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent {
	this := SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent{}
	return &this
}

// GetTargetingClauses returns the TargetingClauses field value
func (o *SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent) GetTargetingClauses() []SponsoredProductsCreateTargetingClause {
	if o == nil {
		var ret []SponsoredProductsCreateTargetingClause
		return ret
	}

	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent) GetTargetingClausesOk() ([]SponsoredProductsCreateTargetingClause, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetingClauses, true
}

// SetTargetingClauses sets field value
func (o *SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent) SetTargetingClauses(v []SponsoredProductsCreateTargetingClause) {
	o.TargetingClauses = v
}

func (o SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetingClauses"] = o.TargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent) Get() *SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent(val *SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent) *NullableSponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
