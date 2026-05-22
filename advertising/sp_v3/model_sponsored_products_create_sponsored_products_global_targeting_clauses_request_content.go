package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent{}

// SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent struct for SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent
type SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent struct {
	// An array of targetingClauses.
	TargetingClauses []SponsoredProductsCreateGlobalTargetingClause `json:"targetingClauses,omitempty"`
}

// NewSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent() *SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent{}
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent{}
	return &this
}

// GetTargetingClauses returns the TargetingClauses field value if set, zero value otherwise.
func (o *SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent) GetTargetingClauses() []SponsoredProductsCreateGlobalTargetingClause {
	if o == nil || IsNil(o.TargetingClauses) {
		var ret []SponsoredProductsCreateGlobalTargetingClause
		return ret
	}
	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent) GetTargetingClausesOk() ([]SponsoredProductsCreateGlobalTargetingClause, bool) {
	if o == nil || IsNil(o.TargetingClauses) {
		return nil, false
	}
	return o.TargetingClauses, true
}

// HasTargetingClauses returns a boolean if a field has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent) HasTargetingClauses() bool {
	if o != nil && !IsNil(o.TargetingClauses) {
		return true
	}

	return false
}

// SetTargetingClauses gets a reference to the given []SponsoredProductsCreateGlobalTargetingClause and assigns it to the TargetingClauses field.
func (o *SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent) SetTargetingClauses(v []SponsoredProductsCreateGlobalTargetingClause) {
	o.TargetingClauses = v
}

func (o SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetingClauses) {
		toSerialize["targetingClauses"] = o.TargetingClauses
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent) Get() *SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent(val *SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent) *NullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
