package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent{}

// SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent struct for SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent
type SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent struct {
	TargetingClauses SponsoredProductsBulkTargetingClauseOperationResponse `json:"targetingClauses"`
}

type _SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent

// NewSponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent(targetingClauses SponsoredProductsBulkTargetingClauseOperationResponse) *SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent{}
	this.TargetingClauses = targetingClauses
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent{}
	return &this
}

// GetTargetingClauses returns the TargetingClauses field value
func (o *SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent) GetTargetingClauses() SponsoredProductsBulkTargetingClauseOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkTargetingClauseOperationResponse
		return ret
	}

	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent) GetTargetingClausesOk() (*SponsoredProductsBulkTargetingClauseOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingClauses, true
}

// SetTargetingClauses sets field value
func (o *SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent) SetTargetingClauses(v SponsoredProductsBulkTargetingClauseOperationResponse) {
	o.TargetingClauses = v
}

func (o SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetingClauses"] = o.TargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent(val *SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
