package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent{}

// SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent struct for SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent
type SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent struct {
	TargetingClauses SponsoredProductsBulkGlobalTargetingClauseOperationResponse `json:"targetingClauses"`
}

type _SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent

// NewSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent(targetingClauses SponsoredProductsBulkGlobalTargetingClauseOperationResponse) *SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent{}
	this.TargetingClauses = targetingClauses
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent{}
	return &this
}

// GetTargetingClauses returns the TargetingClauses field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent) GetTargetingClauses() SponsoredProductsBulkGlobalTargetingClauseOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalTargetingClauseOperationResponse
		return ret
	}

	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent) GetTargetingClausesOk() (*SponsoredProductsBulkGlobalTargetingClauseOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingClauses, true
}

// SetTargetingClauses sets field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent) SetTargetingClauses(v SponsoredProductsBulkGlobalTargetingClauseOperationResponse) {
	o.TargetingClauses = v
}

func (o SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetingClauses"] = o.TargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent struct {
	value *SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent) Get() *SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent) Set(val *SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent(val *SponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent) *NullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
