package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent{}

// SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent struct for SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent
type SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent struct {
	// An array of negativeTargeting.
	NegativeTargetingClauses []SponsoredProductsCreateNegativeTargetingClause `json:"negativeTargetingClauses"`
}

type _SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent

// NewSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent(negativeTargetingClauses []SponsoredProductsCreateNegativeTargetingClause) *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent {
	this := SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent{}
	this.NegativeTargetingClauses = negativeTargetingClauses
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent {
	this := SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent{}
	return &this
}

// GetNegativeTargetingClauses returns the NegativeTargetingClauses field value
func (o *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent) GetNegativeTargetingClauses() []SponsoredProductsCreateNegativeTargetingClause {
	if o == nil {
		var ret []SponsoredProductsCreateNegativeTargetingClause
		return ret
	}

	return o.NegativeTargetingClauses
}

// GetNegativeTargetingClausesOk returns a tuple with the NegativeTargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent) GetNegativeTargetingClausesOk() ([]SponsoredProductsCreateNegativeTargetingClause, bool) {
	if o == nil {
		return nil, false
	}
	return o.NegativeTargetingClauses, true
}

// SetNegativeTargetingClauses sets field value
func (o *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent) SetNegativeTargetingClauses(v []SponsoredProductsCreateNegativeTargetingClause) {
	o.NegativeTargetingClauses = v
}

func (o SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeTargetingClauses"] = o.NegativeTargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent) Get() *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent(val *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent) *NullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
