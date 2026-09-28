package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the UpfrontFee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpfrontFee{}

// UpfrontFee Upfront fee (or a flat fee) for each promotion submitted.
type UpfrontFee struct {
	Rate      Currency     `json:"rate"`
	Frequency FeeFrequency `json:"frequency"`
}

// NewUpfrontFee instantiates a new UpfrontFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpfrontFee(rate Currency, frequency FeeFrequency) *UpfrontFee {
	this := UpfrontFee{}
	this.Rate = rate
	this.Frequency = frequency
	return &this
}

// NewUpfrontFeeWithDefaults instantiates a new UpfrontFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpfrontFeeWithDefaults() *UpfrontFee {
	this := UpfrontFee{}
	return &this
}

// GetRate returns the Rate field value
func (o *UpfrontFee) GetRate() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *UpfrontFee) GetRateOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value
func (o *UpfrontFee) SetRate(v Currency) {
	o.Rate = v
}

// GetFrequency returns the Frequency field value
func (o *UpfrontFee) GetFrequency() FeeFrequency {
	if o == nil {
		var ret FeeFrequency
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *UpfrontFee) GetFrequencyOk() (*FeeFrequency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *UpfrontFee) SetFrequency(v FeeFrequency) {
	o.Frequency = v
}

func (o UpfrontFee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rate"] = o.Rate
	toSerialize["frequency"] = o.Frequency
	return toSerialize, nil
}

type NullableUpfrontFee struct {
	value *UpfrontFee
	isSet bool
}

func (v NullableUpfrontFee) Get() *UpfrontFee {
	return v.value
}

func (v *NullableUpfrontFee) Set(val *UpfrontFee) {
	v.value = val
	v.isSet = true
}

func (v NullableUpfrontFee) IsSet() bool {
	return v.isSet
}

func (v *NullableUpfrontFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpfrontFee(val *UpfrontFee) *NullableUpfrontFee {
	return &NullableUpfrontFee{value: val, isSet: true}
}

func (v NullableUpfrontFee) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpfrontFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
