package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent{}

// SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent struct for SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent
type SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent struct {
	// An array of targetingClauses with updated values.
	TargetingClauses []SponsoredProductsUpdateTargetingClause `json:"targetingClauses"`
}

type _SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent

// NewSponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent instantiates a new SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent(targetingClauses []SponsoredProductsUpdateTargetingClause) *SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent{}
	this.TargetingClauses = targetingClauses
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent{}
	return &this
}

// GetTargetingClauses returns the TargetingClauses field value
func (o *SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent) GetTargetingClauses() []SponsoredProductsUpdateTargetingClause {
	if o == nil {
		var ret []SponsoredProductsUpdateTargetingClause
		return ret
	}

	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent) GetTargetingClausesOk() ([]SponsoredProductsUpdateTargetingClause, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetingClauses, true
}

// SetTargetingClauses sets field value
func (o *SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent) SetTargetingClauses(v []SponsoredProductsUpdateTargetingClause) {
	o.TargetingClauses = v
}

func (o SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetingClauses"] = o.TargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent struct {
	value *SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent) Get() *SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent) Set(val *SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent(val *SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent) *NullableSponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
