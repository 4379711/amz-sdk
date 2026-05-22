package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent{}

// SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent struct for SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent
type SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent struct {
	TargetingClauses SponsoredProductsBulkGlobalTargetingClauseOperationResponse `json:"targetingClauses"`
}

type _SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent

// NewSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent(targetingClauses SponsoredProductsBulkGlobalTargetingClauseOperationResponse) *SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent{}
	this.TargetingClauses = targetingClauses
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent{}
	return &this
}

// GetTargetingClauses returns the TargetingClauses field value
func (o *SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent) GetTargetingClauses() SponsoredProductsBulkGlobalTargetingClauseOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalTargetingClauseOperationResponse
		return ret
	}

	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent) GetTargetingClausesOk() (*SponsoredProductsBulkGlobalTargetingClauseOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingClauses, true
}

// SetTargetingClauses sets field value
func (o *SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent) SetTargetingClauses(v SponsoredProductsBulkGlobalTargetingClauseOperationResponse) {
	o.TargetingClauses = v
}

func (o SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetingClauses"] = o.TargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent(val *SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
