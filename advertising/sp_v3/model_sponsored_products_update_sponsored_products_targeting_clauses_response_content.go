package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent{}

// SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent struct for SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent
type SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent struct {
	TargetingClauses SponsoredProductsBulkTargetingClauseOperationResponse `json:"targetingClauses"`
}

type _SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent

// NewSponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent instantiates a new SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent(targetingClauses SponsoredProductsBulkTargetingClauseOperationResponse) *SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent{}
	this.TargetingClauses = targetingClauses
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent{}
	return &this
}

// GetTargetingClauses returns the TargetingClauses field value
func (o *SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent) GetTargetingClauses() SponsoredProductsBulkTargetingClauseOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkTargetingClauseOperationResponse
		return ret
	}

	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent) GetTargetingClausesOk() (*SponsoredProductsBulkTargetingClauseOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingClauses, true
}

// SetTargetingClauses sets field value
func (o *SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent) SetTargetingClauses(v SponsoredProductsBulkTargetingClauseOperationResponse) {
	o.TargetingClauses = v
}

func (o SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetingClauses"] = o.TargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent struct {
	value *SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent) Get() *SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent) Set(val *SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent(val *SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent) *NullableSponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
