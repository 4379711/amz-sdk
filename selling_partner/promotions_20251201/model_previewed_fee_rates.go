package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the PreviewedFeeRates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreviewedFeeRates{}

// PreviewedFeeRates Fee rates that were previewed to the selling partner. Example fee rates are `$100` or `5% of sales`.
type PreviewedFeeRates struct {
	UpfrontFee  *UpfrontFee  `json:"upfrontFee,omitempty"`
	VariableFee *VariableFee `json:"variableFee,omitempty"`
}

// NewPreviewedFeeRates instantiates a new PreviewedFeeRates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreviewedFeeRates() *PreviewedFeeRates {
	this := PreviewedFeeRates{}
	return &this
}

// NewPreviewedFeeRatesWithDefaults instantiates a new PreviewedFeeRates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreviewedFeeRatesWithDefaults() *PreviewedFeeRates {
	this := PreviewedFeeRates{}
	return &this
}

// GetUpfrontFee returns the UpfrontFee field value if set, zero value otherwise.
func (o *PreviewedFeeRates) GetUpfrontFee() UpfrontFee {
	if o == nil || IsNil(o.UpfrontFee) {
		var ret UpfrontFee
		return ret
	}
	return *o.UpfrontFee
}

// GetUpfrontFeeOk returns a tuple with the UpfrontFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreviewedFeeRates) GetUpfrontFeeOk() (*UpfrontFee, bool) {
	if o == nil || IsNil(o.UpfrontFee) {
		return nil, false
	}
	return o.UpfrontFee, true
}

// HasUpfrontFee returns a boolean if a field has been set.
func (o *PreviewedFeeRates) HasUpfrontFee() bool {
	if o != nil && !IsNil(o.UpfrontFee) {
		return true
	}

	return false
}

// SetUpfrontFee gets a reference to the given UpfrontFee and assigns it to the UpfrontFee field.
func (o *PreviewedFeeRates) SetUpfrontFee(v UpfrontFee) {
	o.UpfrontFee = &v
}

// GetVariableFee returns the VariableFee field value if set, zero value otherwise.
func (o *PreviewedFeeRates) GetVariableFee() VariableFee {
	if o == nil || IsNil(o.VariableFee) {
		var ret VariableFee
		return ret
	}
	return *o.VariableFee
}

// GetVariableFeeOk returns a tuple with the VariableFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreviewedFeeRates) GetVariableFeeOk() (*VariableFee, bool) {
	if o == nil || IsNil(o.VariableFee) {
		return nil, false
	}
	return o.VariableFee, true
}

// HasVariableFee returns a boolean if a field has been set.
func (o *PreviewedFeeRates) HasVariableFee() bool {
	if o != nil && !IsNil(o.VariableFee) {
		return true
	}

	return false
}

// SetVariableFee gets a reference to the given VariableFee and assigns it to the VariableFee field.
func (o *PreviewedFeeRates) SetVariableFee(v VariableFee) {
	o.VariableFee = &v
}

func (o PreviewedFeeRates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UpfrontFee) {
		toSerialize["upfrontFee"] = o.UpfrontFee
	}
	if !IsNil(o.VariableFee) {
		toSerialize["variableFee"] = o.VariableFee
	}
	return toSerialize, nil
}

type NullablePreviewedFeeRates struct {
	value *PreviewedFeeRates
	isSet bool
}

func (v NullablePreviewedFeeRates) Get() *PreviewedFeeRates {
	return v.value
}

func (v *NullablePreviewedFeeRates) Set(val *PreviewedFeeRates) {
	v.value = val
	v.isSet = true
}

func (v NullablePreviewedFeeRates) IsSet() bool {
	return v.isSet
}

func (v *NullablePreviewedFeeRates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreviewedFeeRates(val *PreviewedFeeRates) *NullablePreviewedFeeRates {
	return &NullablePreviewedFeeRates{value: val, isSet: true}
}

func (v NullablePreviewedFeeRates) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePreviewedFeeRates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
