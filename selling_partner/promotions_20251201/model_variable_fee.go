package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the VariableFee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableFee{}

// VariableFee Performance based fee, usually represented as a percentage of promotion sales with a cap amount.
type VariableFee struct {
	// Sales Percentage
	SalesPercentage float32   `json:"salesPercentage"`
	FeeCapAmount    *Currency `json:"feeCapAmount,omitempty"`
}

// NewVariableFee instantiates a new VariableFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableFee(salesPercentage float32) *VariableFee {
	this := VariableFee{}
	this.SalesPercentage = salesPercentage
	return &this
}

// NewVariableFeeWithDefaults instantiates a new VariableFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableFeeWithDefaults() *VariableFee {
	this := VariableFee{}
	return &this
}

// GetSalesPercentage returns the SalesPercentage field value
func (o *VariableFee) GetSalesPercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SalesPercentage
}

// GetSalesPercentageOk returns a tuple with the SalesPercentage field value
// and a boolean to check if the value has been set.
func (o *VariableFee) GetSalesPercentageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SalesPercentage, true
}

// SetSalesPercentage sets field value
func (o *VariableFee) SetSalesPercentage(v float32) {
	o.SalesPercentage = v
}

// GetFeeCapAmount returns the FeeCapAmount field value if set, zero value otherwise.
func (o *VariableFee) GetFeeCapAmount() Currency {
	if o == nil || IsNil(o.FeeCapAmount) {
		var ret Currency
		return ret
	}
	return *o.FeeCapAmount
}

// GetFeeCapAmountOk returns a tuple with the FeeCapAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableFee) GetFeeCapAmountOk() (*Currency, bool) {
	if o == nil || IsNil(o.FeeCapAmount) {
		return nil, false
	}
	return o.FeeCapAmount, true
}

// HasFeeCapAmount returns a boolean if a field has been set.
func (o *VariableFee) HasFeeCapAmount() bool {
	if o != nil && !IsNil(o.FeeCapAmount) {
		return true
	}

	return false
}

// SetFeeCapAmount gets a reference to the given Currency and assigns it to the FeeCapAmount field.
func (o *VariableFee) SetFeeCapAmount(v Currency) {
	o.FeeCapAmount = &v
}

func (o VariableFee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["salesPercentage"] = o.SalesPercentage
	if !IsNil(o.FeeCapAmount) {
		toSerialize["feeCapAmount"] = o.FeeCapAmount
	}
	return toSerialize, nil
}

type NullableVariableFee struct {
	value *VariableFee
	isSet bool
}

func (v NullableVariableFee) Get() *VariableFee {
	return v.value
}

func (v *NullableVariableFee) Set(val *VariableFee) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableFee) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableFee(val *VariableFee) *NullableVariableFee {
	return &NullableVariableFee{value: val, isSet: true}
}

func (v NullableVariableFee) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableVariableFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
