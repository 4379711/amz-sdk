package vendor_direct_fulfillment_orders_20211228

import (
	"github.com/bytedance/sonic"
)

// checks if the TransactionId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionId{}

// TransactionId Response containing the transaction ID.
type TransactionId struct {
	// GUID assigned by Amazon to identify this transaction. This value can be used with the Transaction Status API to return the status of this transaction.
	TransactionId *string `json:"transactionId,omitempty"`
}

// NewTransactionId instantiates a new TransactionId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionId() *TransactionId {
	this := TransactionId{}
	return &this
}

// NewTransactionIdWithDefaults instantiates a new TransactionId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionIdWithDefaults() *TransactionId {
	this := TransactionId{}
	return &this
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *TransactionId) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionId) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *TransactionId) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *TransactionId) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o TransactionId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransactionId) {
		toSerialize["transactionId"] = o.TransactionId
	}
	return toSerialize, nil
}

type NullableTransactionId struct {
	value *TransactionId
	isSet bool
}

func (v NullableTransactionId) Get() *TransactionId {
	return v.value
}

func (v *NullableTransactionId) Set(val *TransactionId) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionId) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionId(val *TransactionId) *NullableTransactionId {
	return &NullableTransactionId{value: val, isSet: true}
}

func (v NullableTransactionId) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTransactionId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
