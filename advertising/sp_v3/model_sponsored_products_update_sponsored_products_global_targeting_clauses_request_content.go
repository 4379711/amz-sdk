package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent{}

// SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent struct for SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent
type SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent struct {
	// An array of targetingClauses with updated values.
	TargetingClauses []SponsoredProductsUpdateGlobalTargetingClause `json:"targetingClauses"`
}

type _SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent

// NewSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent(targetingClauses []SponsoredProductsUpdateGlobalTargetingClause) *SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent{}
	this.TargetingClauses = targetingClauses
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent{}
	return &this
}

// GetTargetingClauses returns the TargetingClauses field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent) GetTargetingClauses() []SponsoredProductsUpdateGlobalTargetingClause {
	if o == nil {
		var ret []SponsoredProductsUpdateGlobalTargetingClause
		return ret
	}

	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent) GetTargetingClausesOk() ([]SponsoredProductsUpdateGlobalTargetingClause, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetingClauses, true
}

// SetTargetingClauses sets field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent) SetTargetingClauses(v []SponsoredProductsUpdateGlobalTargetingClause) {
	o.TargetingClauses = v
}

func (o SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetingClauses"] = o.TargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent struct {
	value *SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent) Get() *SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent) Set(val *SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent(val *SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent) *NullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
