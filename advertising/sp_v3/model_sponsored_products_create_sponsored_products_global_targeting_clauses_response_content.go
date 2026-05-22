package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent{}

// SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent struct for SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent
type SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent struct {
	TargetingClauses SponsoredProductsBulkGlobalTargetingClauseOperationResponse `json:"targetingClauses"`
}

type _SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent

// NewSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent(targetingClauses SponsoredProductsBulkGlobalTargetingClauseOperationResponse) *SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent{}
	this.TargetingClauses = targetingClauses
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent{}
	return &this
}

// GetTargetingClauses returns the TargetingClauses field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent) GetTargetingClauses() SponsoredProductsBulkGlobalTargetingClauseOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalTargetingClauseOperationResponse
		return ret
	}

	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent) GetTargetingClausesOk() (*SponsoredProductsBulkGlobalTargetingClauseOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingClauses, true
}

// SetTargetingClauses sets field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent) SetTargetingClauses(v SponsoredProductsBulkGlobalTargetingClauseOperationResponse) {
	o.TargetingClauses = v
}

func (o SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetingClauses"] = o.TargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent) Get() *SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent(val *SponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent) *NullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
