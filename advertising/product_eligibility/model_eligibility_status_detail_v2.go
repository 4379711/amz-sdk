package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the EligibilityStatusDetailV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EligibilityStatusDetailV2{}

// EligibilityStatusDetailV2 Describes a single program's eligibility status
type EligibilityStatusDetailV2 struct {
	// String identifier for the status.
	Reasons []ReasonItem `json:"reasons,omitempty"`
	// Boolean value where if true, advertiser is eligible to access the given program.
	Eligible  *bool      `json:"eligible,omitempty"`
	AdProgram *AdProgram `json:"adProgram,omitempty"`
}

// NewEligibilityStatusDetailV2 instantiates a new EligibilityStatusDetailV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEligibilityStatusDetailV2() *EligibilityStatusDetailV2 {
	this := EligibilityStatusDetailV2{}
	return &this
}

// NewEligibilityStatusDetailV2WithDefaults instantiates a new EligibilityStatusDetailV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEligibilityStatusDetailV2WithDefaults() *EligibilityStatusDetailV2 {
	this := EligibilityStatusDetailV2{}
	return &this
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *EligibilityStatusDetailV2) GetReasons() []ReasonItem {
	if o == nil || IsNil(o.Reasons) {
		var ret []ReasonItem
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EligibilityStatusDetailV2) GetReasonsOk() ([]ReasonItem, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *EligibilityStatusDetailV2) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []ReasonItem and assigns it to the Reasons field.
func (o *EligibilityStatusDetailV2) SetReasons(v []ReasonItem) {
	o.Reasons = v
}

// GetEligible returns the Eligible field value if set, zero value otherwise.
func (o *EligibilityStatusDetailV2) GetEligible() bool {
	if o == nil || IsNil(o.Eligible) {
		var ret bool
		return ret
	}
	return *o.Eligible
}

// GetEligibleOk returns a tuple with the Eligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EligibilityStatusDetailV2) GetEligibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Eligible) {
		return nil, false
	}
	return o.Eligible, true
}

// HasEligible returns a boolean if a field has been set.
func (o *EligibilityStatusDetailV2) HasEligible() bool {
	if o != nil && !IsNil(o.Eligible) {
		return true
	}

	return false
}

// SetEligible gets a reference to the given bool and assigns it to the Eligible field.
func (o *EligibilityStatusDetailV2) SetEligible(v bool) {
	o.Eligible = &v
}

// GetAdProgram returns the AdProgram field value if set, zero value otherwise.
func (o *EligibilityStatusDetailV2) GetAdProgram() AdProgram {
	if o == nil || IsNil(o.AdProgram) {
		var ret AdProgram
		return ret
	}
	return *o.AdProgram
}

// GetAdProgramOk returns a tuple with the AdProgram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EligibilityStatusDetailV2) GetAdProgramOk() (*AdProgram, bool) {
	if o == nil || IsNil(o.AdProgram) {
		return nil, false
	}
	return o.AdProgram, true
}

// HasAdProgram returns a boolean if a field has been set.
func (o *EligibilityStatusDetailV2) HasAdProgram() bool {
	if o != nil && !IsNil(o.AdProgram) {
		return true
	}

	return false
}

// SetAdProgram gets a reference to the given AdProgram and assigns it to the AdProgram field.
func (o *EligibilityStatusDetailV2) SetAdProgram(v AdProgram) {
	o.AdProgram = &v
}

func (o EligibilityStatusDetailV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	if !IsNil(o.Eligible) {
		toSerialize["eligible"] = o.Eligible
	}
	if !IsNil(o.AdProgram) {
		toSerialize["adProgram"] = o.AdProgram
	}
	return toSerialize, nil
}

type NullableEligibilityStatusDetailV2 struct {
	value *EligibilityStatusDetailV2
	isSet bool
}

func (v NullableEligibilityStatusDetailV2) Get() *EligibilityStatusDetailV2 {
	return v.value
}

func (v *NullableEligibilityStatusDetailV2) Set(val *EligibilityStatusDetailV2) {
	v.value = val
	v.isSet = true
}

func (v NullableEligibilityStatusDetailV2) IsSet() bool {
	return v.isSet
}

func (v *NullableEligibilityStatusDetailV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEligibilityStatusDetailV2(val *EligibilityStatusDetailV2) *NullableEligibilityStatusDetailV2 {
	return &NullableEligibilityStatusDetailV2{value: val, isSet: true}
}

func (v NullableEligibilityStatusDetailV2) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableEligibilityStatusDetailV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
