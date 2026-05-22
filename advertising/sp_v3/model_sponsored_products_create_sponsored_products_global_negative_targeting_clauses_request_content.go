package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent{}

// SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent struct for SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent
type SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent struct {
	// An array of negativeTargeting.
	NegativeTargetingClauses []SponsoredProductsCreateGlobalNegativeTargetingClause `json:"negativeTargetingClauses,omitempty"`
}

// NewSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent() *SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent{}
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent{}
	return &this
}

// GetNegativeTargetingClauses returns the NegativeTargetingClauses field value if set, zero value otherwise.
func (o *SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) GetNegativeTargetingClauses() []SponsoredProductsCreateGlobalNegativeTargetingClause {
	if o == nil || IsNil(o.NegativeTargetingClauses) {
		var ret []SponsoredProductsCreateGlobalNegativeTargetingClause
		return ret
	}
	return o.NegativeTargetingClauses
}

// GetNegativeTargetingClausesOk returns a tuple with the NegativeTargetingClauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) GetNegativeTargetingClausesOk() ([]SponsoredProductsCreateGlobalNegativeTargetingClause, bool) {
	if o == nil || IsNil(o.NegativeTargetingClauses) {
		return nil, false
	}
	return o.NegativeTargetingClauses, true
}

// HasNegativeTargetingClauses returns a boolean if a field has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) HasNegativeTargetingClauses() bool {
	if o != nil && !IsNil(o.NegativeTargetingClauses) {
		return true
	}

	return false
}

// SetNegativeTargetingClauses gets a reference to the given []SponsoredProductsCreateGlobalNegativeTargetingClause and assigns it to the NegativeTargetingClauses field.
func (o *SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) SetNegativeTargetingClauses(v []SponsoredProductsCreateGlobalNegativeTargetingClause) {
	o.NegativeTargetingClauses = v
}

func (o SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NegativeTargetingClauses) {
		toSerialize["negativeTargetingClauses"] = o.NegativeTargetingClauses
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) Get() *SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent(val *SponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) *NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
