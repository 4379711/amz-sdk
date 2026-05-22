package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the CarrierCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CarrierCode{}

// CarrierCode Identifies the carrier that will deliver the shipment.
type CarrierCode struct {
	CarrierCodeType *CarrierCodeType `json:"carrierCodeType,omitempty"`
	// Value of the carrier code.
	CarrierCodeValue *string `json:"carrierCodeValue,omitempty"`
}

// NewCarrierCode instantiates a new CarrierCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCarrierCode() *CarrierCode {
	this := CarrierCode{}
	return &this
}

// NewCarrierCodeWithDefaults instantiates a new CarrierCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCarrierCodeWithDefaults() *CarrierCode {
	this := CarrierCode{}
	return &this
}

// GetCarrierCodeType returns the CarrierCodeType field value if set, zero value otherwise.
func (o *CarrierCode) GetCarrierCodeType() CarrierCodeType {
	if o == nil || IsNil(o.CarrierCodeType) {
		var ret CarrierCodeType
		return ret
	}
	return *o.CarrierCodeType
}

// GetCarrierCodeTypeOk returns a tuple with the CarrierCodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CarrierCode) GetCarrierCodeTypeOk() (*CarrierCodeType, bool) {
	if o == nil || IsNil(o.CarrierCodeType) {
		return nil, false
	}
	return o.CarrierCodeType, true
}

// HasCarrierCodeType returns a boolean if a field has been set.
func (o *CarrierCode) HasCarrierCodeType() bool {
	if o != nil && !IsNil(o.CarrierCodeType) {
		return true
	}

	return false
}

// SetCarrierCodeType gets a reference to the given CarrierCodeType and assigns it to the CarrierCodeType field.
func (o *CarrierCode) SetCarrierCodeType(v CarrierCodeType) {
	o.CarrierCodeType = &v
}

// GetCarrierCodeValue returns the CarrierCodeValue field value if set, zero value otherwise.
func (o *CarrierCode) GetCarrierCodeValue() string {
	if o == nil || IsNil(o.CarrierCodeValue) {
		var ret string
		return ret
	}
	return *o.CarrierCodeValue
}

// GetCarrierCodeValueOk returns a tuple with the CarrierCodeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CarrierCode) GetCarrierCodeValueOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierCodeValue) {
		return nil, false
	}
	return o.CarrierCodeValue, true
}

// HasCarrierCodeValue returns a boolean if a field has been set.
func (o *CarrierCode) HasCarrierCodeValue() bool {
	if o != nil && !IsNil(o.CarrierCodeValue) {
		return true
	}

	return false
}

// SetCarrierCodeValue gets a reference to the given string and assigns it to the CarrierCodeValue field.
func (o *CarrierCode) SetCarrierCodeValue(v string) {
	o.CarrierCodeValue = &v
}

func (o CarrierCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CarrierCodeType) {
		toSerialize["carrierCodeType"] = o.CarrierCodeType
	}
	if !IsNil(o.CarrierCodeValue) {
		toSerialize["carrierCodeValue"] = o.CarrierCodeValue
	}
	return toSerialize, nil
}

type NullableCarrierCode struct {
	value *CarrierCode
	isSet bool
}

func (v NullableCarrierCode) Get() *CarrierCode {
	return v.value
}

func (v *NullableCarrierCode) Set(val *CarrierCode) {
	v.value = val
	v.isSet = true
}

func (v NullableCarrierCode) IsSet() bool {
	return v.isSet
}

func (v *NullableCarrierCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarrierCode(val *CarrierCode) *NullableCarrierCode {
	return &NullableCarrierCode{value: val, isSet: true}
}

func (v NullableCarrierCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCarrierCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
