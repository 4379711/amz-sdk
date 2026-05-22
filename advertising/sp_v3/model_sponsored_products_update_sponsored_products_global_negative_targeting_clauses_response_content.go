package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent{}

// SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent struct for SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent
type SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent struct {
	NegativeTargetingClauses SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse `json:"negativeTargetingClauses"`
}

type _SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent

// NewSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent(negativeTargetingClauses SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) *SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent{}
	this.NegativeTargetingClauses = negativeTargetingClauses
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent{}
	return &this
}

// GetNegativeTargetingClauses returns the NegativeTargetingClauses field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) GetNegativeTargetingClauses() SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse
		return ret
	}

	return o.NegativeTargetingClauses
}

// GetNegativeTargetingClausesOk returns a tuple with the NegativeTargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) GetNegativeTargetingClausesOk() (*SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeTargetingClauses, true
}

// SetNegativeTargetingClauses sets field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) SetNegativeTargetingClauses(v SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) {
	o.NegativeTargetingClauses = v
}

func (o SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeTargetingClauses"] = o.NegativeTargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent struct {
	value *SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) Get() *SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) Set(val *SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent(val *SponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) *NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
