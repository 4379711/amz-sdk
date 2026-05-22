package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the Benefits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Benefits{}

// Benefits Benefits that are included and excluded for each shipping offer. Benefits represents services provided by Amazon (eg. CLAIMS_PROTECTED, etc.) when sellers purchase shipping through Amazon. Benefit details will be made available for any shipment placed on or after January 1st 2024 00:00 UTC.
type Benefits struct {
	// A list of included benefits.
	IncludedBenefits []string `json:"includedBenefits"`
	// A list of excluded benefit. Refer to the ExcludeBenefit object for further documentation
	ExcludedBenefits []ExcludedBenefit `json:"excludedBenefits"`
}

type _Benefits Benefits

// NewBenefits instantiates a new Benefits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBenefits(includedBenefits []string, excludedBenefits []ExcludedBenefit) *Benefits {
	this := Benefits{}
	this.IncludedBenefits = includedBenefits
	this.ExcludedBenefits = excludedBenefits
	return &this
}

// NewBenefitsWithDefaults instantiates a new Benefits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBenefitsWithDefaults() *Benefits {
	this := Benefits{}
	return &this
}

// GetIncludedBenefits returns the IncludedBenefits field value
func (o *Benefits) GetIncludedBenefits() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IncludedBenefits
}

// GetIncludedBenefitsOk returns a tuple with the IncludedBenefits field value
// and a boolean to check if the value has been set.
func (o *Benefits) GetIncludedBenefitsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncludedBenefits, true
}

// SetIncludedBenefits sets field value
func (o *Benefits) SetIncludedBenefits(v []string) {
	o.IncludedBenefits = v
}

// GetExcludedBenefits returns the ExcludedBenefits field value
func (o *Benefits) GetExcludedBenefits() []ExcludedBenefit {
	if o == nil {
		var ret []ExcludedBenefit
		return ret
	}

	return o.ExcludedBenefits
}

// GetExcludedBenefitsOk returns a tuple with the ExcludedBenefits field value
// and a boolean to check if the value has been set.
func (o *Benefits) GetExcludedBenefitsOk() ([]ExcludedBenefit, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExcludedBenefits, true
}

// SetExcludedBenefits sets field value
func (o *Benefits) SetExcludedBenefits(v []ExcludedBenefit) {
	o.ExcludedBenefits = v
}

func (o Benefits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["includedBenefits"] = o.IncludedBenefits
	toSerialize["excludedBenefits"] = o.ExcludedBenefits
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
