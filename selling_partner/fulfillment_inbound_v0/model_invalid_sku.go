package fulfillment_inbound_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the InvalidSKU type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvalidSKU{}

// InvalidSKU Contains detail about an invalid SKU
type InvalidSKU struct {
	// The seller SKU of the item.
	SellerSKU   *string      `json:"SellerSKU,omitempty"`
	ErrorReason *ErrorReason `json:"ErrorReason,omitempty"`
}

// NewInvalidSKU instantiates a new InvalidSKU object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidSKU() *InvalidSKU {
	this := InvalidSKU{}
	return &this
}

// NewInvalidSKUWithDefaults instantiates a new InvalidSKU object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidSKUWithDefaults() *InvalidSKU {
	this := InvalidSKU{}
	return &this
}

// GetSellerSKU returns the SellerSKU field value if set, zero value otherwise.
func (o *InvalidSKU) GetSellerSKU() string {
	if o == nil || IsNil(o.SellerSKU) {
		var ret string
		return ret
	}
	return *o.SellerSKU
}

// GetSellerSKUOk returns a tuple with the SellerSKU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidSKU) GetSellerSKUOk() (*string, bool) {
	if o == nil || IsNil(o.SellerSKU) {
		return nil, false
	}
	return o.SellerSKU, true
}

// HasSellerSKU returns a boolean if a field has been set.
func (o *InvalidSKU) HasSellerSKU() bool {
	if o != nil && !IsNil(o.SellerSKU) {
		return true
	}

	return false
}

// SetSellerSKU gets a reference to the given string and assigns it to the SellerSKU field.
func (o *InvalidSKU) SetSellerSKU(v string) {
	o.SellerSKU = &v
}

// GetErrorReason returns the ErrorReason field value if set, zero value otherwise.
func (o *InvalidSKU) GetErrorReason() ErrorReason {
	if o == nil || IsNil(o.ErrorReason) {
		var ret ErrorReason
		return ret
	}
	return *o.ErrorReason
}

// GetErrorReasonOk returns a tuple with the ErrorReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidSKU) GetErrorReasonOk() (*ErrorReason, bool) {
	if o == nil || IsNil(o.ErrorReason) {
		return nil, false
	}
	return o.ErrorReason, true
}

// HasErrorReason returns a boolean if a field has been set.
func (o *InvalidSKU) HasErrorReason() bool {
	if o != nil && !IsNil(o.ErrorReason) {
		return true
	}

	return false
}

// SetErrorReason gets a reference to the given ErrorReason and assigns it to the ErrorReason field.
func (o *InvalidSKU) SetErrorReason(v ErrorReason) {
	o.ErrorReason = &v
}

func (o InvalidSKU) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SellerSKU) {
		toSerialize["SellerSKU"] = o.SellerSKU
	}
	if !IsNil(o.ErrorReason) {
		toSerialize["ErrorReason"] = o.ErrorReason
	}
	return toSerialize, nil
}

type NullableInvalidSKU struct {
	value *InvalidSKU
	isSet bool
}

func (v NullableInvalidSKU) Get() *InvalidSKU {
	return v.value
}

func (v *NullableInvalidSKU) Set(val *InvalidSKU) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidSKU) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidSKU) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidSKU(val *InvalidSKU) *NullableInvalidSKU {
	return &NullableInvalidSKU{value: val, isSet: true}
}

func (v NullableInvalidSKU) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInvalidSKU) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
