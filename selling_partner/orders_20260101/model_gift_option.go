package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the GiftOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftOption{}

// GiftOption Gift wrapping and personalization options selected by the customer for an order item.
type GiftOption struct {
	// Personal message from the buyer to be included with the gift-wrapped item.
	GiftMessage *string `json:"giftMessage,omitempty"`
	// Type or quality level of gift wrapping service selected by the customer.
	GiftWrapLevel *string `json:"giftWrapLevel,omitempty"`
}

// NewGiftOption instantiates a new GiftOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftOption() *GiftOption {
	this := GiftOption{}
	return &this
}

// NewGiftOptionWithDefaults instantiates a new GiftOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftOptionWithDefaults() *GiftOption {
	this := GiftOption{}
	return &this
}

// GetGiftMessage returns the GiftMessage field value if set, zero value otherwise.
func (o *GiftOption) GetGiftMessage() string {
	if o == nil || IsNil(o.GiftMessage) {
		var ret string
		return ret
	}
	return *o.GiftMessage
}

// GetGiftMessageOk returns a tuple with the GiftMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftOption) GetGiftMessageOk() (*string, bool) {
	if o == nil || IsNil(o.GiftMessage) {
		return nil, false
	}
	return o.GiftMessage, true
}

// HasGiftMessage returns a boolean if a field has been set.
func (o *GiftOption) HasGiftMessage() bool {
	if o != nil && !IsNil(o.GiftMessage) {
		return true
	}

	return false
}

// SetGiftMessage gets a reference to the given string and assigns it to the GiftMessage field.
func (o *GiftOption) SetGiftMessage(v string) {
	o.GiftMessage = &v
}

// GetGiftWrapLevel returns the GiftWrapLevel field value if set, zero value otherwise.
func (o *GiftOption) GetGiftWrapLevel() string {
	if o == nil || IsNil(o.GiftWrapLevel) {
		var ret string
		return ret
	}
	return *o.GiftWrapLevel
}

// GetGiftWrapLevelOk returns a tuple with the GiftWrapLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftOption) GetGiftWrapLevelOk() (*string, bool) {
	if o == nil || IsNil(o.GiftWrapLevel) {
		return nil, false
	}
	return o.GiftWrapLevel, true
}

// HasGiftWrapLevel returns a boolean if a field has been set.
func (o *GiftOption) HasGiftWrapLevel() bool {
	if o != nil && !IsNil(o.GiftWrapLevel) {
		return true
	}

	return false
}

// SetGiftWrapLevel gets a reference to the given string and assigns it to the GiftWrapLevel field.
func (o *GiftOption) SetGiftWrapLevel(v string) {
	o.GiftWrapLevel = &v
}

func (o GiftOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GiftMessage) {
		toSerialize["giftMessage"] = o.GiftMessage
	}
	if !IsNil(o.GiftWrapLevel) {
		toSerialize["giftWrapLevel"] = o.GiftWrapLevel
	}
	return toSerialize, nil
}

type NullableGiftOption struct {
	value *GiftOption
	isSet bool
}

func (v NullableGiftOption) Get() *GiftOption {
	return v.value
}

func (v *NullableGiftOption) Set(val *GiftOption) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftOption) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftOption(val *GiftOption) *NullableGiftOption {
	return &NullableGiftOption{value: val, isSet: true}
}

func (v NullableGiftOption) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGiftOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
