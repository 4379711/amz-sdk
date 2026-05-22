package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the ProgramEligibilityRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProgramEligibilityRequestContent{}

// ProgramEligibilityRequestContent A request to evaluate account level eligibility for Amazon ad programs (Sponsored Products, Sponsored Brands, Sponsored Display, Stores, DirectToConsumer, Amazon Attribution, etc).
type ProgramEligibilityRequestContent struct {
	SkipChecks *Check `json:"skipChecks,omitempty"`
}

// NewProgramEligibilityRequestContent instantiates a new ProgramEligibilityRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgramEligibilityRequestContent() *ProgramEligibilityRequestContent {
	this := ProgramEligibilityRequestContent{}
	return &this
}

// NewProgramEligibilityRequestContentWithDefaults instantiates a new ProgramEligibilityRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramEligibilityRequestContentWithDefaults() *ProgramEligibilityRequestContent {
	this := ProgramEligibilityRequestContent{}
	return &this
}

// GetSkipChecks returns the SkipChecks field value if set, zero value otherwise.
func (o *ProgramEligibilityRequestContent) GetSkipChecks() Check {
	if o == nil || IsNil(o.SkipChecks) {
		var ret Check
		return ret
	}
	return *o.SkipChecks
}

// GetSkipChecksOk returns a tuple with the SkipChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramEligibilityRequestContent) GetSkipChecksOk() (*Check, bool) {
	if o == nil || IsNil(o.SkipChecks) {
		return nil, false
	}
	return o.SkipChecks, true
}

// HasSkipChecks returns a boolean if a field has been set.
func (o *ProgramEligibilityRequestContent) HasSkipChecks() bool {
	if o != nil && !IsNil(o.SkipChecks) {
		return true
	}

	return false
}

// SetSkipChecks gets a reference to the given Check and assigns it to the SkipChecks field.
func (o *ProgramEligibilityRequestContent) SetSkipChecks(v Check) {
	o.SkipChecks = &v
}

func (o ProgramEligibilityRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SkipChecks) {
		toSerialize["skipChecks"] = o.SkipChecks
	}
	return toSerialize, nil
}

type NullableProgramEligibilityRequestContent struct {
	value *ProgramEligibilityRequestContent
	isSet bool
}

func (v NullableProgramEligibilityRequestContent) Get() *ProgramEligibilityRequestContent {
	return v.value
}

func (v *NullableProgramEligibilityRequestContent) Set(val *ProgramEligibilityRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramEligibilityRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramEligibilityRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramEligibilityRequestContent(val *ProgramEligibilityRequestContent) *NullableProgramEligibilityRequestContent {
	return &NullableProgramEligibilityRequestContent{value: val, isSet: true}
}

func (v NullableProgramEligibilityRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProgramEligibilityRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
