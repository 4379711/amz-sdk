package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the ProgramEligibilityResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProgramEligibilityResponseContent{}

// ProgramEligibilityResponseContent An object of program eligibility responses for an advertiser.
type ProgramEligibilityResponseContent struct {
	// This is a map that will be key'd on the ad program (SB/SD/DTC/MAAS/SPOT); the value will be an eligibility object.
	EligibilityStatusMap *map[string]EligibilityStatusDetail `json:"eligibilityStatusMap,omitempty"`
}

// NewProgramEligibilityResponseContent instantiates a new ProgramEligibilityResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgramEligibilityResponseContent() *ProgramEligibilityResponseContent {
	this := ProgramEligibilityResponseContent{}
	return &this
}

// NewProgramEligibilityResponseContentWithDefaults instantiates a new ProgramEligibilityResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramEligibilityResponseContentWithDefaults() *ProgramEligibilityResponseContent {
	this := ProgramEligibilityResponseContent{}
	return &this
}

// GetEligibilityStatusMap returns the EligibilityStatusMap field value if set, zero value otherwise.
func (o *ProgramEligibilityResponseContent) GetEligibilityStatusMap() map[string]EligibilityStatusDetail {
	if o == nil || IsNil(o.EligibilityStatusMap) {
		var ret map[string]EligibilityStatusDetail
		return ret
	}
	return *o.EligibilityStatusMap
}

// GetEligibilityStatusMapOk returns a tuple with the EligibilityStatusMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramEligibilityResponseContent) GetEligibilityStatusMapOk() (*map[string]EligibilityStatusDetail, bool) {
	if o == nil || IsNil(o.EligibilityStatusMap) {
		return nil, false
	}
	return o.EligibilityStatusMap, true
}

// HasEligibilityStatusMap returns a boolean if a field has been set.
func (o *ProgramEligibilityResponseContent) HasEligibilityStatusMap() bool {
	if o != nil && !IsNil(o.EligibilityStatusMap) {
		return true
	}

	return false
}

// SetEligibilityStatusMap gets a reference to the given map[string]EligibilityStatusDetail and assigns it to the EligibilityStatusMap field.
func (o *ProgramEligibilityResponseContent) SetEligibilityStatusMap(v map[string]EligibilityStatusDetail) {
	o.EligibilityStatusMap = &v
}

func (o ProgramEligibilityResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EligibilityStatusMap) {
		toSerialize["eligibilityStatusMap"] = o.EligibilityStatusMap
	}
	return toSerialize, nil
}

type NullableProgramEligibilityResponseContent struct {
	value *ProgramEligibilityResponseContent
	isSet bool
}

func (v NullableProgramEligibilityResponseContent) Get() *ProgramEligibilityResponseContent {
	return v.value
}

func (v *NullableProgramEligibilityResponseContent) Set(val *ProgramEligibilityResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramEligibilityResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramEligibilityResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramEligibilityResponseContent(val *ProgramEligibilityResponseContent) *NullableProgramEligibilityResponseContent {
	return &NullableProgramEligibilityResponseContent{value: val, isSet: true}
}

func (v NullableProgramEligibilityResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProgramEligibilityResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
