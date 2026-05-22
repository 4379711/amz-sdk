package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the DeliveryMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryMessage{}

// DeliveryMessage Localized messaging for a delivery offering.
type DeliveryMessage struct {
	// The message content for a delivery offering.
	Text *string `json:"text,omitempty"`
	// The locale for the message (for example, en_US).
	Locale *string `json:"locale,omitempty"`
}

// NewDeliveryMessage instantiates a new DeliveryMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryMessage() *DeliveryMessage {
	this := DeliveryMessage{}
	return &this
}

// NewDeliveryMessageWithDefaults instantiates a new DeliveryMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryMessageWithDefaults() *DeliveryMessage {
	this := DeliveryMessage{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *DeliveryMessage) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryMessage) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *DeliveryMessage) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *DeliveryMessage) SetText(v string) {
	o.Text = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *DeliveryMessage) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryMessage) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *DeliveryMessage) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *DeliveryMessage) SetLocale(v string) {
	o.Locale = &v
}

func (o DeliveryMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	return toSerialize, nil
}

type NullableDeliveryMessage struct {
	value *DeliveryMessage
	isSet bool
}

func (v NullableDeliveryMessage) Get() *DeliveryMessage {
	return v.value
}

func (v *NullableDeliveryMessage) Set(val *DeliveryMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryMessage(val *DeliveryMessage) *NullableDeliveryMessage {
	return &NullableDeliveryMessage{value: val, isSet: true}
}

func (v NullableDeliveryMessage) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeliveryMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
