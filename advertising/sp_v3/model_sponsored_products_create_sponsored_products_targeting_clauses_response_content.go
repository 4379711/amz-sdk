package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent{}

// SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent struct for SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent
type SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent struct {
	TargetingClauses SponsoredProductsBulkTargetingClauseOperationResponse `json:"targetingClauses"`
}

type _SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent

// NewSponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent(targetingClauses SponsoredProductsBulkTargetingClauseOperationResponse) *SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent {
	this := SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent{}
	this.TargetingClauses = targetingClauses
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsTargetingClausesResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent {
	this := SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent{}
	return &this
}

// GetTargetingClauses returns the TargetingClauses field value
func (o *SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent) GetTargetingClauses() SponsoredProductsBulkTargetingClauseOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkTargetingClauseOperationResponse
		return ret
	}

	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent) GetTargetingClausesOk() (*SponsoredProductsBulkTargetingClauseOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingClauses, true
}

// SetTargetingClauses sets field value
func (o *SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent) SetTargetingClauses(v SponsoredProductsBulkTargetingClauseOperationResponse) {
	o.TargetingClauses = v
}

func (o SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetingClauses"] = o.TargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent) Get() *SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent(val *SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent) *NullableSponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
