package fulfillment_inbound_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the InvalidASIN type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvalidASIN{}

// InvalidASIN Contains details about an invalid ASIN
type InvalidASIN struct {
	// The Amazon Standard Identification Number (ASIN) of the item.
	ASIN        *string      `json:"ASIN,omitempty"`
	ErrorReason *ErrorReason `json:"ErrorReason,omitempty"`
}

// NewInvalidASIN instantiates a new InvalidASIN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidASIN() *InvalidASIN {
	this := InvalidASIN{}
	return &this
}

// NewInvalidASINWithDefaults instantiates a new InvalidASIN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidASINWithDefaults() *InvalidASIN {
	this := InvalidASIN{}
	return &this
}

// GetASIN returns the ASIN field value if set, zero value otherwise.
func (o *InvalidASIN) GetASIN() string {
	if o == nil || IsNil(o.ASIN) {
		var ret string
		return ret
	}
	return *o.ASIN
}

// GetASINOk returns a tuple with the ASIN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidASIN) GetASINOk() (*string, bool) {
	if o == nil || IsNil(o.ASIN) {
		return nil, false
	}
	return o.ASIN, true
}

// HasASIN returns a boolean if a field has been set.
func (o *InvalidASIN) HasASIN() bool {
	if o != nil && !IsNil(o.ASIN) {
		return true
	}

	return false
}

// SetASIN gets a reference to the given string and assigns it to the ASIN field.
func (o *InvalidASIN) SetASIN(v string) {
	o.ASIN = &v
}

// GetErrorReason returns the ErrorReason field value if set, zero value otherwise.
func (o *InvalidASIN) GetErrorReason() ErrorReason {
	if o == nil || IsNil(o.ErrorReason) {
		var ret ErrorReason
		return ret
	}
	return *o.ErrorReason
}

// GetErrorReasonOk returns a tuple with the ErrorReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidASIN) GetErrorReasonOk() (*ErrorReason, bool) {
	if o == nil || IsNil(o.ErrorReason) {
		return nil, false
	}
	return o.ErrorReason, true
}

// HasErrorReason returns a boolean if a field has been set.
func (o *InvalidASIN) HasErrorReason() bool {
	if o != nil && !IsNil(o.ErrorReason) {
		return true
	}

	return false
}

// SetErrorReason gets a reference to the given ErrorReason and assigns it to the ErrorReason field.
func (o *InvalidASIN) SetErrorReason(v ErrorReason) {
	o.ErrorReason = &v
}

func (o InvalidASIN) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ASIN) {
		toSerialize["ASIN"] = o.ASIN
	}
	if !IsNil(o.ErrorReason) {
		toSerialize["ErrorReason"] = o.ErrorReason
	}
	return toSerialize, nil
}

type NullableInvalidASIN struct {
	value *InvalidASIN
	isSet bool
}

func (v NullableInvalidASIN) Get() *InvalidASIN {
	return v.value
}

func (v *NullableInvalidASIN) Set(val *InvalidASIN) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidASIN) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidASIN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidASIN(val *InvalidASIN) *NullableInvalidASIN {
	return &NullableInvalidASIN{value: val, isSet: true}
}

func (v NullableInvalidASIN) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInvalidASIN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
