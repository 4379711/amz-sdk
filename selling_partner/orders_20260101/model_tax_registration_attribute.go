package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the TaxRegistrationAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxRegistrationAttribute{}

// TaxRegistrationAttribute An additional attribute associated with a tax registration.
type TaxRegistrationAttribute struct {
	// The name of the tax registration attribute.  **Possible values**: `TAX_OFFICE`
	Key *string `json:"key,omitempty"`
	// The value of the tax registration attribute.
	Value *string `json:"value,omitempty"`
}

// NewTaxRegistrationAttribute instantiates a new TaxRegistrationAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxRegistrationAttribute() *TaxRegistrationAttribute {
	this := TaxRegistrationAttribute{}
	return &this
}

// NewTaxRegistrationAttributeWithDefaults instantiates a new TaxRegistrationAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxRegistrationAttributeWithDefaults() *TaxRegistrationAttribute {
	this := TaxRegistrationAttribute{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *TaxRegistrationAttribute) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRegistrationAttribute) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *TaxRegistrationAttribute) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *TaxRegistrationAttribute) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TaxRegistrationAttribute) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRegistrationAttribute) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TaxRegistrationAttribute) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TaxRegistrationAttribute) SetValue(v string) {
	o.Value = &v
}

func (o TaxRegistrationAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTaxRegistrationAttribute struct {
	value *TaxRegistrationAttribute
	isSet bool
}

func (v NullableTaxRegistrationAttribute) Get() *TaxRegistrationAttribute {
	return v.value
}

func (v *NullableTaxRegistrationAttribute) Set(val *TaxRegistrationAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxRegistrationAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxRegistrationAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxRegistrationAttribute(val *TaxRegistrationAttribute) *NullableTaxRegistrationAttribute {
	return &NullableTaxRegistrationAttribute{value: val, isSet: true}
}

func (v NullableTaxRegistrationAttribute) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTaxRegistrationAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
