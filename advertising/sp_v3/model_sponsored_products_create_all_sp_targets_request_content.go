package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateAllSPTargetsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateAllSPTargetsRequestContent{}

// SponsoredProductsCreateAllSPTargetsRequestContent struct for SponsoredProductsCreateAllSPTargetsRequestContent
type SponsoredProductsCreateAllSPTargetsRequestContent struct {
	// An array of targets.
	Targets []SponsoredProductsCreateAllTargets `json:"targets"`
}

type _SponsoredProductsCreateAllSPTargetsRequestContent SponsoredProductsCreateAllSPTargetsRequestContent

// NewSponsoredProductsCreateAllSPTargetsRequestContent instantiates a new SponsoredProductsCreateAllSPTargetsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateAllSPTargetsRequestContent(targets []SponsoredProductsCreateAllTargets) *SponsoredProductsCreateAllSPTargetsRequestContent {
	this := SponsoredProductsCreateAllSPTargetsRequestContent{}
	this.Targets = targets
	return &this
}

// NewSponsoredProductsCreateAllSPTargetsRequestContentWithDefaults instantiates a new SponsoredProductsCreateAllSPTargetsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateAllSPTargetsRequestContentWithDefaults() *SponsoredProductsCreateAllSPTargetsRequestContent {
	this := SponsoredProductsCreateAllSPTargetsRequestContent{}
	return &this
}

// GetTargets returns the Targets field value
func (o *SponsoredProductsCreateAllSPTargetsRequestContent) GetTargets() []SponsoredProductsCreateAllTargets {
	if o == nil {
		var ret []SponsoredProductsCreateAllTargets
		return ret
	}

	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateAllSPTargetsRequestContent) GetTargetsOk() ([]SponsoredProductsCreateAllTargets, bool) {
	if o == nil {
		return nil, false
	}
	return o.Targets, true
}

// SetTargets sets field value
func (o *SponsoredProductsCreateAllSPTargetsRequestContent) SetTargets(v []SponsoredProductsCreateAllTargets) {
	o.Targets = v
}

func (o SponsoredProductsCreateAllSPTargetsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targets"] = o.Targets
	return toSerialize, nil
}

type NullableSponsoredProductsCreateAllSPTargetsRequestContent struct {
	value *SponsoredProductsCreateAllSPTargetsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateAllSPTargetsRequestContent) Get() *SponsoredProductsCreateAllSPTargetsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateAllSPTargetsRequestContent) Set(val *SponsoredProductsCreateAllSPTargetsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateAllSPTargetsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateAllSPTargetsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateAllSPTargetsRequestContent(val *SponsoredProductsCreateAllSPTargetsRequestContent) *NullableSponsoredProductsCreateAllSPTargetsRequestContent {
	return &NullableSponsoredProductsCreateAllSPTargetsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateAllSPTargetsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateAllSPTargetsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
