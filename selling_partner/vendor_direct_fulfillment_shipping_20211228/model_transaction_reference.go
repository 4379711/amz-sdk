package vendor_direct_fulfillment_shipping_20211228

import (
	"github.com/bytedance/sonic"
)

// checks if the TransactionReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionReference{}

// TransactionReference Response containing the transaction ID.
type TransactionReference struct {
	// GUID to identify this transaction. This value can be used with the Transaction Status API to return the status of this transaction.
	TransactionId *string `json:"transactionId,omitempty"`
}

// NewTransactionReference instantiates a new TransactionReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionReference() *TransactionReference {
	this := TransactionReference{}
	return &this
}

// NewTransactionReferenceWithDefaults instantiates a new TransactionReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionReferenceWithDefaults() *TransactionReference {
	this := TransactionReference{}
	return &this
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *TransactionReference) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionReference) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *TransactionReference) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *TransactionReference) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o TransactionReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransactionId) {
		toSerialize["transactionId"] = o.TransactionId
	}
	return toSerialize, nil
}

type NullableTransactionReference struct {
	value *TransactionReference
	isSet bool
}

func (v NullableTransactionReference) Get() *TransactionReference {
	return v.value
}

func (v *NullableTransactionReference) Set(val *TransactionReference) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionReference) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionReference(val *TransactionReference) *NullableTransactionReference {
	return &NullableTransactionReference{value: val, isSet: true}
}

func (v NullableTransactionReference) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTransactionReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
