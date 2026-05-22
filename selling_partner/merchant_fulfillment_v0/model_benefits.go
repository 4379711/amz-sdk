package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the Benefits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Benefits{}

// Benefits Benefits that are included and excluded for each shipping offer. Benefits represents services provided by Amazon (for example, `CLAIMS_PROTECTED`) when sellers purchase shipping through Amazon. Benefit details are made available for any shipment placed on or after January 1st 2024 00:00 UTC.
type Benefits struct {
	// A list of included benefits.
	IncludedBenefits []string `json:"IncludedBenefits,omitempty"`
	// A list of excluded benefits. Refer to the `ExcludeBenefit` object for further documentation.
	ExcludedBenefits []ExcludedBenefit `json:"ExcludedBenefits,omitempty"`
}

// NewBenefits instantiates a new Benefits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBenefits() *Benefits {
	this := Benefits{}
	return &this
}

// NewBenefitsWithDefaults instantiates a new Benefits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBenefitsWithDefaults() *Benefits {
	this := Benefits{}
	return &this
}

// GetIncludedBenefits returns the IncludedBenefits field value if set, zero value otherwise.
func (o *Benefits) GetIncludedBenefits() []string {
	if o == nil || IsNil(o.IncludedBenefits) {
		var ret []string
		return ret
	}
	return o.IncludedBenefits
}

// GetIncludedBenefitsOk returns a tuple with the IncludedBenefits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benefits) GetIncludedBenefitsOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludedBenefits) {
		return nil, false
	}
	return o.IncludedBenefits, true
}

// HasIncludedBenefits returns a boolean if a field has been set.
func (o *Benefits) HasIncludedBenefits() bool {
	if o != nil && !IsNil(o.IncludedBenefits) {
		return true
	}

	return false
}

// SetIncludedBenefits gets a reference to the given []string and assigns it to the IncludedBenefits field.
func (o *Benefits) SetIncludedBenefits(v []string) {
	o.IncludedBenefits = v
}

// GetExcludedBenefits returns the ExcludedBenefits field value if set, zero value otherwise.
func (o *Benefits) GetExcludedBenefits() []ExcludedBenefit {
	if o == nil || IsNil(o.ExcludedBenefits) {
		var ret []ExcludedBenefit
		return ret
	}
	return o.ExcludedBenefits
}

// GetExcludedBenefitsOk returns a tuple with the ExcludedBenefits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benefits) GetExcludedBenefitsOk() ([]ExcludedBenefit, bool) {
	if o == nil || IsNil(o.ExcludedBenefits) {
		return nil, false
	}
	return o.ExcludedBenefits, true
}

// HasExcludedBenefits returns a boolean if a field has been set.
func (o *Benefits) HasExcludedBenefits() bool {
	if o != nil && !IsNil(o.ExcludedBenefits) {
		return true
	}

	return false
}

// SetExcludedBenefits gets a reference to the given []ExcludedBenefit and assigns it to the ExcludedBenefits field.
func (o *Benefits) SetExcludedBenefits(v []ExcludedBenefit) {
	o.ExcludedBenefits = v
}

func (o Benefits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncludedBenefits) {
		toSerialize["IncludedBenefits"] = o.IncludedBenefits
	}
	if !IsNil(o.ExcludedBenefits) {
		toSerialize["ExcludedBenefits"] = o.ExcludedBenefits
	}
	return toSerialize, nil
}

type NullableBenefits struct {
	value *Benefits
	isSet bool
}

func (v NullableBenefits) Get() *Benefits {
	return v.value
}

func (v *NullableBenefits) Set(val *Benefits) {
	v.value = val
	v.isSet = true
}

func (v NullableBenefits) IsSet() bool {
	return v.isSet
}

func (v *NullableBenefits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBenefits(val *Benefits) *NullableBenefits {
	return &NullableBenefits{value: val, isSet: true}
}

func (v NullableBenefits) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBenefits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
