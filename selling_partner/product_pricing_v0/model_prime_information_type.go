package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the PrimeInformationType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrimeInformationType{}

// PrimeInformationType Amazon Prime information.
type PrimeInformationType struct {
	// Indicates whether the offer is an Amazon Prime offer.
	IsPrime bool `json:"IsPrime"`
	// Indicates whether the offer is an Amazon Prime offer throughout the entire marketplace where it is listed.
	IsNationalPrime bool `json:"IsNationalPrime"`
}

type _PrimeInformationType PrimeInformationType

// NewPrimeInformationType instantiates a new PrimeInformationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrimeInformationType(isPrime bool, isNationalPrime bool) *PrimeInformationType {
	this := PrimeInformationType{}
	this.IsPrime = isPrime
	this.IsNationalPrime = isNationalPrime
	return &this
}

// NewPrimeInformationTypeWithDefaults instantiates a new PrimeInformationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrimeInformationTypeWithDefaults() *PrimeInformationType {
	this := PrimeInformationType{}
	return &this
}

// GetIsPrime returns the IsPrime field value
func (o *PrimeInformationType) GetIsPrime() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPrime
}

// GetIsPrimeOk returns a tuple with the IsPrime field value
// and a boolean to check if the value has been set.
func (o *PrimeInformationType) GetIsPrimeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPrime, true
}

// SetIsPrime sets field value
func (o *PrimeInformationType) SetIsPrime(v bool) {
	o.IsPrime = v
}

// GetIsNationalPrime returns the IsNationalPrime field value
func (o *PrimeInformationType) GetIsNationalPrime() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsNationalPrime
}

// GetIsNationalPrimeOk returns a tuple with the IsNationalPrime field value
// and a boolean to check if the value has been set.
func (o *PrimeInformationType) GetIsNationalPrimeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsNationalPrime, true
}

// SetIsNationalPrime sets field value
func (o *PrimeInformationType) SetIsNationalPrime(v bool) {
	o.IsNationalPrime = v
}

func (o PrimeInformationType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["IsPrime"] = o.IsPrime
	toSerialize["IsNationalPrime"] = o.IsNationalPrime
	return toSerialize, nil
}

type NullablePrimeInformationType struct {
	value *PrimeInformationType
	isSet bool
}

func (v NullablePrimeInformationType) Get() *PrimeInformationType {
	return v.value
}

func (v *NullablePrimeInformationType) Set(val *PrimeInformationType) {
	v.value = val
	v.isSet = true
}

func (v NullablePrimeInformationType) IsSet() bool {
	return v.isSet
}

func (v *NullablePrimeInformationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrimeInformationType(val *PrimeInformationType) *NullablePrimeInformationType {
	return &NullablePrimeInformationType{value: val, isSet: true}
}

func (v NullablePrimeInformationType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePrimeInformationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
