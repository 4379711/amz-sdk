package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the Recipient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Recipient{}

// Recipient Information about the recipient to whom the order should be delivered.
type Recipient struct {
	DeliveryAddress    *CustomerAddress    `json:"deliveryAddress,omitempty"`
	DeliveryPreference *DeliveryPreference `json:"deliveryPreference,omitempty"`
}

// NewRecipient instantiates a new Recipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipient() *Recipient {
	this := Recipient{}
	return &this
}

// NewRecipientWithDefaults instantiates a new Recipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientWithDefaults() *Recipient {
	this := Recipient{}
	return &this
}

// GetDeliveryAddress returns the DeliveryAddress field value if set, zero value otherwise.
func (o *Recipient) GetDeliveryAddress() CustomerAddress {
	if o == nil || IsNil(o.DeliveryAddress) {
		var ret CustomerAddress
		return ret
	}
	return *o.DeliveryAddress
}

// GetDeliveryAddressOk returns a tuple with the DeliveryAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recipient) GetDeliveryAddressOk() (*CustomerAddress, bool) {
	if o == nil || IsNil(o.DeliveryAddress) {
		return nil, false
	}
	return o.DeliveryAddress, true
}

// HasDeliveryAddress returns a boolean if a field has been set.
func (o *Recipient) HasDeliveryAddress() bool {
	if o != nil && !IsNil(o.DeliveryAddress) {
		return true
	}

	return false
}

// SetDeliveryAddress gets a reference to the given CustomerAddress and assigns it to the DeliveryAddress field.
func (o *Recipient) SetDeliveryAddress(v CustomerAddress) {
	o.DeliveryAddress = &v
}

// GetDeliveryPreference returns the DeliveryPreference field value if set, zero value otherwise.
func (o *Recipient) GetDeliveryPreference() DeliveryPreference {
	if o == nil || IsNil(o.DeliveryPreference) {
		var ret DeliveryPreference
		return ret
	}
	return *o.DeliveryPreference
}

// GetDeliveryPreferenceOk returns a tuple with the DeliveryPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recipient) GetDeliveryPreferenceOk() (*DeliveryPreference, bool) {
	if o == nil || IsNil(o.DeliveryPreference) {
		return nil, false
	}
	return o.DeliveryPreference, true
}

// HasDeliveryPreference returns a boolean if a field has been set.
func (o *Recipient) HasDeliveryPreference() bool {
	if o != nil && !IsNil(o.DeliveryPreference) {
		return true
	}

	return false
}

// SetDeliveryPreference gets a reference to the given DeliveryPreference and assigns it to the DeliveryPreference field.
func (o *Recipient) SetDeliveryPreference(v DeliveryPreference) {
	o.DeliveryPreference = &v
}

func (o Recipient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeliveryAddress) {
		toSerialize["deliveryAddress"] = o.DeliveryAddress
	}
	if !IsNil(o.DeliveryPreference) {
		toSerialize["deliveryPreference"] = o.DeliveryPreference
	}
	return toSerialize, nil
}

type NullableRecipient struct {
	value *Recipient
	isSet bool
}

func (v NullableRecipient) Get() *Recipient {
	return v.value
}

func (v *NullableRecipient) Set(val *Recipient) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipient(val *Recipient) *NullableRecipient {
	return &NullableRecipient{value: val, isSet: true}
}

func (v NullableRecipient) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
