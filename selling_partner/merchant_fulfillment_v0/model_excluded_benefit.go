package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ExcludedBenefit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExcludedBenefit{}

// ExcludedBenefit An object representing an excluded benefit that is excluded for a shipping offer or rate.
type ExcludedBenefit struct {
	// A benefit that is being excluded from a shipment.
	Benefit *string `json:"Benefit,omitempty"`
	// List of reasons why a benefit is excluded for a shipping offer (for example, `LATE_DELIVERY_RISK`).
	ReasonCodes []string `json:"ReasonCodes,omitempty"`
}

// NewExcludedBenefit instantiates a new ExcludedBenefit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExcludedBenefit() *ExcludedBenefit {
	this := ExcludedBenefit{}
	return &this
}

// NewExcludedBenefitWithDefaults instantiates a new ExcludedBenefit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExcludedBenefitWithDefaults() *ExcludedBenefit {
	this := ExcludedBenefit{}
	return &this
}

// GetBenefit returns the Benefit field value if set, zero value otherwise.
func (o *ExcludedBenefit) GetBenefit() string {
	if o == nil || IsNil(o.Benefit) {
		var ret string
		return ret
	}
	return *o.Benefit
}

// GetBenefitOk returns a tuple with the Benefit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExcludedBenefit) GetBenefitOk() (*string, bool) {
	if o == nil || IsNil(o.Benefit) {
		return nil, false
	}
	return o.Benefit, true
}

// HasBenefit returns a boolean if a field has been set.
func (o *ExcludedBenefit) HasBenefit() bool {
	if o != nil && !IsNil(o.Benefit) {
		return true
	}

	return false
}

// SetBenefit gets a reference to the given string and assigns it to the Benefit field.
func (o *ExcludedBenefit) SetBenefit(v string) {
	o.Benefit = &v
}

// GetReasonCodes returns the ReasonCodes field value if set, zero value otherwise.
func (o *ExcludedBenefit) GetReasonCodes() []string {
	if o == nil || IsNil(o.ReasonCodes) {
		var ret []string
		return ret
	}
	return o.ReasonCodes
}

// GetReasonCodesOk returns a tuple with the ReasonCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExcludedBenefit) GetReasonCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.ReasonCodes) {
		return nil, false
	}
	return o.ReasonCodes, true
}

// HasReasonCodes returns a boolean if a field has been set.
func (o *ExcludedBenefit) HasReasonCodes() bool {
	if o != nil && !IsNil(o.ReasonCodes) {
		return true
	}

	return false
}

// SetReasonCodes gets a reference to the given []string and assigns it to the ReasonCodes field.
func (o *ExcludedBenefit) SetReasonCodes(v []string) {
	o.ReasonCodes = v
}

func (o ExcludedBenefit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Benefit) {
		toSerialize["Benefit"] = o.Benefit
	}
	if !IsNil(o.ReasonCodes) {
		toSerialize["ReasonCodes"] = o.ReasonCodes
	}
	return toSerialize, nil
}

type NullableExcludedBenefit struct {
	value *ExcludedBenefit
	isSet bool
}

func (v NullableExcludedBenefit) Get() *ExcludedBenefit {
	return v.value
}

func (v *NullableExcludedBenefit) Set(val *ExcludedBenefit) {
	v.value = val
	v.isSet = true
}

func (v NullableExcludedBenefit) IsSet() bool {
	return v.isSet
}

func (v *NullableExcludedBenefit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExcludedBenefit(val *ExcludedBenefit) *NullableExcludedBenefit {
	return &NullableExcludedBenefit{value: val, isSet: true}
}

func (v NullableExcludedBenefit) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableExcludedBenefit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
