package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the PrimeDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrimeDetails{}

// PrimeDetails Amazon Prime details.
type PrimeDetails struct {
	// Indicates whether the offer is an Amazon Prime offer.
	Eligibility string `json:"eligibility"`
}

type _PrimeDetails PrimeDetails

// NewPrimeDetails instantiates a new PrimeDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrimeDetails(eligibility string) *PrimeDetails {
	this := PrimeDetails{}
	this.Eligibility = eligibility
	return &this
}

// NewPrimeDetailsWithDefaults instantiates a new PrimeDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrimeDetailsWithDefaults() *PrimeDetails {
	this := PrimeDetails{}
	return &this
}

// GetEligibility returns the Eligibility field value
func (o *PrimeDetails) GetEligibility() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Eligibility
}

// GetEligibilityOk returns a tuple with the Eligibility field value
// and a boolean to check if the value has been set.
func (o *PrimeDetails) GetEligibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Eligibility, true
}

// SetEligibility sets field value
func (o *PrimeDetails) SetEligibility(v string) {
	o.Eligibility = v
}

func (o PrimeDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eligibility"] = o.Eligibility
	return toSerialize, nil
}

type NullablePrimeDetails struct {
	value *PrimeDetails
	isSet bool
}

func (v NullablePrimeDetails) Get() *PrimeDetails {
	return v.value
}

func (v *NullablePrimeDetails) Set(val *PrimeDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePrimeDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePrimeDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrimeDetails(val *PrimeDetails) *NullablePrimeDetails {
	return &NullablePrimeDetails{value: val, isSet: true}
}

func (v NullablePrimeDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePrimeDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
