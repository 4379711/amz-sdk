package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the Amount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Amount{}

// Amount A quantity based on unit of measure.
type Amount struct {
	// The unit of measure for the amount.
	UnitOfMeasure string `json:"unitOfMeasure"`
	// A decimal number with no loss of precision. Useful when precision loss is unacceptable, as with currencies. Follows RFC7159 for number representation.
	Value string `json:"value"`
}

type _Amount Amount

// NewAmount instantiates a new Amount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmount(unitOfMeasure string, value string) *Amount {
	this := Amount{}
	this.UnitOfMeasure = unitOfMeasure
	this.Value = value
	return &this
}

// NewAmountWithDefaults instantiates a new Amount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmountWithDefaults() *Amount {
	this := Amount{}
	return &this
}

// GetUnitOfMeasure returns the UnitOfMeasure field value
func (o *Amount) GetUnitOfMeasure() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value
// and a boolean to check if the value has been set.
func (o *Amount) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitOfMeasure, true
}

// SetUnitOfMeasure sets field value
func (o *Amount) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure = v
}

// GetValue returns the Value field value
func (o *Amount) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Amount) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Amount) SetValue(v string) {
	o.Value = v
}

func (o Amount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["unitOfMeasure"] = o.UnitOfMeasure
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableAmount struct {
	value *Amount
	isSet bool
}

func (v NullableAmount) Get() *Amount {
	return v.value
}

func (v *NullableAmount) Set(val *Amount) {
	v.value = val
	v.isSet = true
}

func (v NullableAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmount(val *Amount) *NullableAmount {
	return &NullableAmount{value: val, isSet: true}
}

func (v NullableAmount) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
