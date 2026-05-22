package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the EligibilityStatusDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EligibilityStatusDetail{}

// EligibilityStatusDetail Describes a single program's eligibility status
type EligibilityStatusDetail struct {
	// String identifier for the status.
	Reasons []ReasonItem `json:"reasons,omitempty"`
	// Boolean value where if true, advertiser is eligible to access the given program.
	Eligible *bool `json:"eligible,omitempty"`
}

// NewEligibilityStatusDetail instantiates a new EligibilityStatusDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEligibilityStatusDetail() *EligibilityStatusDetail {
	this := EligibilityStatusDetail{}
	return &this
}

// NewEligibilityStatusDetailWithDefaults instantiates a new EligibilityStatusDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEligibilityStatusDetailWithDefaults() *EligibilityStatusDetail {
	this := EligibilityStatusDetail{}
	return &this
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *EligibilityStatusDetail) GetReasons() []ReasonItem {
	if o == nil || IsNil(o.Reasons) {
		var ret []ReasonItem
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EligibilityStatusDetail) GetReasonsOk() ([]ReasonItem, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *EligibilityStatusDetail) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []ReasonItem and assigns it to the Reasons field.
func (o *EligibilityStatusDetail) SetReasons(v []ReasonItem) {
	o.Reasons = v
}

// GetEligible returns the Eligible field value if set, zero value otherwise.
func (o *EligibilityStatusDetail) GetEligible() bool {
	if o == nil || IsNil(o.Eligible) {
		var ret bool
		return ret
	}
	return *o.Eligible
}

// GetEligibleOk returns a tuple with the Eligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EligibilityStatusDetail) GetEligibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Eligible) {
		return nil, false
	}
	return o.Eligible, true
}

// HasEligible returns a boolean if a field has been set.
func (o *EligibilityStatusDetail) HasEligible() bool {
	if o != nil && !IsNil(o.Eligible) {
		return true
	}

	return false
}

// SetEligible gets a reference to the given bool and assigns it to the Eligible field.
func (o *EligibilityStatusDetail) SetEligible(v bool) {
	o.Eligible = &v
}

func (o EligibilityStatusDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	if !IsNil(o.Eligible) {
		toSerialize["eligible"] = o.Eligible
	}
	return toSerialize, nil
}

type NullableEligibilityStatusDetail struct {
	value *EligibilityStatusDetail
	isSet bool
}

func (v NullableEligibilityStatusDetail) Get() *EligibilityStatusDetail {
	return v.value
}

func (v *NullableEligibilityStatusDetail) Set(val *EligibilityStatusDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableEligibilityStatusDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableEligibilityStatusDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEligibilityStatusDetail(val *EligibilityStatusDetail) *NullableEligibilityStatusDetail {
	return &NullableEligibilityStatusDetail{value: val, isSet: true}
}

func (v NullableEligibilityStatusDetail) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableEligibilityStatusDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
