package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the ActiveAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActiveAccount{}

// ActiveAccount Active Account Details
type ActiveAccount struct {
	// Identifier for the seller's carrier account.
	AccountId *string `json:"accountId,omitempty"`
	// The carrier identifier for the offering, provided by the carrier.
	CarrierId *string `json:"carrierId,omitempty"`
}

// NewActiveAccount instantiates a new ActiveAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveAccount() *ActiveAccount {
	this := ActiveAccount{}
	return &this
}

// NewActiveAccountWithDefaults instantiates a new ActiveAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveAccountWithDefaults() *ActiveAccount {
	this := ActiveAccount{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ActiveAccount) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveAccount) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ActiveAccount) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ActiveAccount) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCarrierId returns the CarrierId field value if set, zero value otherwise.
func (o *ActiveAccount) GetCarrierId() string {
	if o == nil || IsNil(o.CarrierId) {
		var ret string
		return ret
	}
	return *o.CarrierId
}

// GetCarrierIdOk returns a tuple with the CarrierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveAccount) GetCarrierIdOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierId) {
		return nil, false
	}
	return o.CarrierId, true
}

// HasCarrierId returns a boolean if a field has been set.
func (o *ActiveAccount) HasCarrierId() bool {
	if o != nil && !IsNil(o.CarrierId) {
		return true
	}

	return false
}

// SetCarrierId gets a reference to the given string and assigns it to the CarrierId field.
func (o *ActiveAccount) SetCarrierId(v string) {
	o.CarrierId = &v
}

func (o ActiveAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.CarrierId) {
		toSerialize["carrierId"] = o.CarrierId
	}
	return toSerialize, nil
}

type NullableActiveAccount struct {
	value *ActiveAccount
	isSet bool
}

func (v NullableActiveAccount) Get() *ActiveAccount {
	return v.value
}

func (v *NullableActiveAccount) Set(val *ActiveAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveAccount(val *ActiveAccount) *NullableActiveAccount {
	return &NullableActiveAccount{value: val, isSet: true}
}

func (v NullableActiveAccount) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableActiveAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
