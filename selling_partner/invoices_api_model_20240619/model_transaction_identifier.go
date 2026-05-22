package invoices_api_model_20240619

import (
	"github.com/bytedance/sonic"
)

// checks if the TransactionIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionIdentifier{}

// TransactionIdentifier The identifier for a transaction.
type TransactionIdentifier struct {
	// The transaction identifier name. Use the `getInvoicesAttributes` operation to check `transactionIdentifierName` options.
	Name *string `json:"name,omitempty"`
	// The transaction identifier.
	Id *string `json:"id,omitempty"`
}

// NewTransactionIdentifier instantiates a new TransactionIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionIdentifier() *TransactionIdentifier {
	this := TransactionIdentifier{}
	return &this
}

// NewTransactionIdentifierWithDefaults instantiates a new TransactionIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionIdentifierWithDefaults() *TransactionIdentifier {
	this := TransactionIdentifier{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TransactionIdentifier) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionIdentifier) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TransactionIdentifier) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TransactionIdentifier) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransactionIdentifier) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionIdentifier) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransactionIdentifier) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TransactionIdentifier) SetId(v string) {
	o.Id = &v
}

func (o TransactionIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableTransactionIdentifier struct {
	value *TransactionIdentifier
	isSet bool
}

func (v NullableTransactionIdentifier) Get() *TransactionIdentifier {
	return v.value
}

func (v *NullableTransactionIdentifier) Set(val *TransactionIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionIdentifier(val *TransactionIdentifier) *NullableTransactionIdentifier {
	return &NullableTransactionIdentifier{value: val, isSet: true}
}

func (v NullableTransactionIdentifier) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTransactionIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
