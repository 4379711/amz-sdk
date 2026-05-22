package product_metadata

import (
	"github.com/bytedance/sonic"
)

// checks if the LegacyEligibility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LegacyEligibility{}

// LegacyEligibility struct for LegacyEligibility
type LegacyEligibility struct {
	// List of ineligible status identifier
	IneligibilityCodes []string `json:"ineligibilityCodes,omitempty"`
	// List of reasons that made this item ineligible to be advertised
	IneligibilityReasons []string `json:"ineligibilityReasons,omitempty"`
}

// NewLegacyEligibility instantiates a new LegacyEligibility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacyEligibility() *LegacyEligibility {
	this := LegacyEligibility{}
	return &this
}

// NewLegacyEligibilityWithDefaults instantiates a new LegacyEligibility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacyEligibilityWithDefaults() *LegacyEligibility {
	this := LegacyEligibility{}
	return &this
}

// GetIneligibilityCodes returns the IneligibilityCodes field value if set, zero value otherwise.
func (o *LegacyEligibility) GetIneligibilityCodes() []string {
	if o == nil || IsNil(o.IneligibilityCodes) {
		var ret []string
		return ret
	}
	return o.IneligibilityCodes
}

// GetIneligibilityCodesOk returns a tuple with the IneligibilityCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyEligibility) GetIneligibilityCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.IneligibilityCodes) {
		return nil, false
	}
	return o.IneligibilityCodes, true
}

// HasIneligibilityCodes returns a boolean if a field has been set.
func (o *LegacyEligibility) HasIneligibilityCodes() bool {
	if o != nil && !IsNil(o.IneligibilityCodes) {
		return true
	}

	return false
}

// SetIneligibilityCodes gets a reference to the given []string and assigns it to the IneligibilityCodes field.
func (o *LegacyEligibility) SetIneligibilityCodes(v []string) {
	o.IneligibilityCodes = v
}

// GetIneligibilityReasons returns the IneligibilityReasons field value if set, zero value otherwise.
func (o *LegacyEligibility) GetIneligibilityReasons() []string {
	if o == nil || IsNil(o.IneligibilityReasons) {
		var ret []string
		return ret
	}
	return o.IneligibilityReasons
}

// GetIneligibilityReasonsOk returns a tuple with the IneligibilityReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyEligibility) GetIneligibilityReasonsOk() ([]string, bool) {
	if o == nil || IsNil(o.IneligibilityReasons) {
		return nil, false
	}
	return o.IneligibilityReasons, true
}

// HasIneligibilityReasons returns a boolean if a field has been set.
func (o *LegacyEligibility) HasIneligibilityReasons() bool {
	if o != nil && !IsNil(o.IneligibilityReasons) {
		return true
	}

	return false
}

// SetIneligibilityReasons gets a reference to the given []string and assigns it to the IneligibilityReasons field.
func (o *LegacyEligibility) SetIneligibilityReasons(v []string) {
	o.IneligibilityReasons = v
}

func (o LegacyEligibility) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IneligibilityCodes) {
		toSerialize["ineligibilityCodes"] = o.IneligibilityCodes
	}
	if !IsNil(o.IneligibilityReasons) {
		toSerialize["ineligibilityReasons"] = o.IneligibilityReasons
	}
	return toSerialize, nil
}

type NullableLegacyEligibility struct {
	value *LegacyEligibility
	isSet bool
}

func (v NullableLegacyEligibility) Get() *LegacyEligibility {
	return v.value
}

func (v *NullableLegacyEligibility) Set(val *LegacyEligibility) {
	v.value = val
	v.isSet = true
}

func (v NullableLegacyEligibility) IsSet() bool {
	return v.isSet
}

func (v *NullableLegacyEligibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegacyEligibility(val *LegacyEligibility) *NullableLegacyEligibility {
	return &NullableLegacyEligibility{value: val, isSet: true}
}

func (v NullableLegacyEligibility) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLegacyEligibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
