package seller_wallet_20240301

import (
	"github.com/bytedance/sonic"
)

// checks if the TransactionInstrumentDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionInstrumentDetails{}

// TransactionInstrumentDetails Details of the destination bank account in the transaction request.
type TransactionInstrumentDetails struct {
	BankAccount BankAccount `json:"bankAccount"`
	// The bank account number of the destination payment method.  **Note:** This field is encrypted before Amazon receives it, so should not be used to generate `destAccountDigitalSignature`, and should not be included in the request signature.
	BankAccountNumber string `json:"bankAccountNumber"`
}

type _TransactionInstrumentDetails TransactionInstrumentDetails

// NewTransactionInstrumentDetails instantiates a new TransactionInstrumentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionInstrumentDetails(bankAccount BankAccount, bankAccountNumber string) *TransactionInstrumentDetails {
	this := TransactionInstrumentDetails{}
	this.BankAccount = bankAccount
	this.BankAccountNumber = bankAccountNumber
	return &this
}

// NewTransactionInstrumentDetailsWithDefaults instantiates a new TransactionInstrumentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionInstrumentDetailsWithDefaults() *TransactionInstrumentDetails {
	this := TransactionInstrumentDetails{}
	return &this
}

// GetBankAccount returns the BankAccount field value
func (o *TransactionInstrumentDetails) GetBankAccount() BankAccount {
	if o == nil {
		var ret BankAccount
		return ret
	}

	return o.BankAccount
}

// GetBankAccountOk returns a tuple with the BankAccount field value
// and a boolean to check if the value has been set.
func (o *TransactionInstrumentDetails) GetBankAccountOk() (*BankAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankAccount, true
}

// SetBankAccount sets field value
func (o *TransactionInstrumentDetails) SetBankAccount(v BankAccount) {
	o.BankAccount = v
}

// GetBankAccountNumber returns the BankAccountNumber field value
func (o *TransactionInstrumentDetails) GetBankAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankAccountNumber
}

// GetBankAccountNumberOk returns a tuple with the BankAccountNumber field value
// and a boolean to check if the value has been set.
func (o *TransactionInstrumentDetails) GetBankAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankAccountNumber, true
}

// SetBankAccountNumber sets field value
func (o *TransactionInstrumentDetails) SetBankAccountNumber(v string) {
	o.BankAccountNumber = v
}

func (o TransactionInstrumentDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bankAccount"] = o.BankAccount
	toSerialize["bankAccountNumber"] = o.BankAccountNumber
	return toSerialize, nil
}

type NullableTransactionInstrumentDetails struct {
	value *TransactionInstrumentDetails
	isSet bool
}

func (v NullableTransactionInstrumentDetails) Get() *TransactionInstrumentDetails {
	return v.value
}

func (v *NullableTransactionInstrumentDetails) Set(val *TransactionInstrumentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionInstrumentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionInstrumentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionInstrumentDetails(val *TransactionInstrumentDetails) *NullableTransactionInstrumentDetails {
	return &NullableTransactionInstrumentDetails{value: val, isSet: true}
}

func (v NullableTransactionInstrumentDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTransactionInstrumentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
