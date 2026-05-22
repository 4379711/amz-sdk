package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateAllSPTargetsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateAllSPTargetsRequestContent{}

// SponsoredProductsUpdateAllSPTargetsRequestContent struct for SponsoredProductsUpdateAllSPTargetsRequestContent
type SponsoredProductsUpdateAllSPTargetsRequestContent struct {
	// An array of targets.
	Targets []SponsoredProductsUpdateAllTargets `json:"targets"`
}

type _SponsoredProductsUpdateAllSPTargetsRequestContent SponsoredProductsUpdateAllSPTargetsRequestContent

// NewSponsoredProductsUpdateAllSPTargetsRequestContent instantiates a new SponsoredProductsUpdateAllSPTargetsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateAllSPTargetsRequestContent(targets []SponsoredProductsUpdateAllTargets) *SponsoredProductsUpdateAllSPTargetsRequestContent {
	this := SponsoredProductsUpdateAllSPTargetsRequestContent{}
	this.Targets = targets
	return &this
}

// NewSponsoredProductsUpdateAllSPTargetsRequestContentWithDefaults instantiates a new SponsoredProductsUpdateAllSPTargetsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateAllSPTargetsRequestContentWithDefaults() *SponsoredProductsUpdateAllSPTargetsRequestContent {
	this := SponsoredProductsUpdateAllSPTargetsRequestContent{}
	return &this
}

// GetTargets returns the Targets field value
func (o *SponsoredProductsUpdateAllSPTargetsRequestContent) GetTargets() []SponsoredProductsUpdateAllTargets {
	if o == nil {
		var ret []SponsoredProductsUpdateAllTargets
		return ret
	}

	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateAllSPTargetsRequestContent) GetTargetsOk() ([]SponsoredProductsUpdateAllTargets, bool) {
	if o == nil {
		return nil, false
	}
	return o.Targets, true
}

// SetTargets sets field value
func (o *SponsoredProductsUpdateAllSPTargetsRequestContent) SetTargets(v []SponsoredProductsUpdateAllTargets) {
	o.Targets = v
}

func (o SponsoredProductsUpdateAllSPTargetsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targets"] = o.Targets
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateAllSPTargetsRequestContent struct {
	value *SponsoredProductsUpdateAllSPTargetsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateAllSPTargetsRequestContent) Get() *SponsoredProductsUpdateAllSPTargetsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateAllSPTargetsRequestContent) Set(val *SponsoredProductsUpdateAllSPTargetsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateAllSPTargetsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateAllSPTargetsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateAllSPTargetsRequestContent(val *SponsoredProductsUpdateAllSPTargetsRequestContent) *NullableSponsoredProductsUpdateAllSPTargetsRequestContent {
	return &NullableSponsoredProductsUpdateAllSPTargetsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateAllSPTargetsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateAllSPTargetsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
