package seller_wallet_20240301

import (
	"github.com/bytedance/sonic"
)

// checks if the TransactionAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionAccount{}

// TransactionAccount Details of the bank account involved in transaction.
type TransactionAccount struct {
	// The unique identifier provided by Amazon to identify the account.
	AccountId *string `json:"accountId,omitempty"`
	// The account holder's name.
	BankAccountHolderName string `json:"bankAccountHolderName"`
	// The name of the bank.
	BankName                string                  `json:"bankName"`
	BankAccountNumberFormat BankAccountNumberFormat `json:"bankAccountNumberFormat"`
	// The last three digits of the bank account number.
	BankAccountNumberTail *string `json:"bankAccountNumberTail,omitempty"`
	// The two-digit country code, in ISO 3166 format. This field is optional for `transactionSourceAccount`, but is mandatory for `transactionDestinationAccount`.
	BankAccountCountryCode *string `json:"bankAccountCountryCode,omitempty"`
	// The currency code in ISO 4217 format.
	BankAccountCurrency string `json:"bankAccountCurrency"`
}

type _TransactionAccount TransactionAccount

// NewTransactionAccount instantiates a new TransactionAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionAccount(bankAccountHolderName string, bankName string, bankAccountNumberFormat BankAccountNumberFormat, bankAccountCurrency string) *TransactionAccount {
	this := TransactionAccount{}
	this.BankAccountHolderName = bankAccountHolderName
	this.BankName = bankName
	this.BankAccountNumberFormat = bankAccountNumberFormat
	this.BankAccountCurrency = bankAccountCurrency
	return &this
}

// NewTransactionAccountWithDefaults instantiates a new TransactionAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionAccountWithDefaults() *TransactionAccount {
	this := TransactionAccount{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *TransactionAccount) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionAccount) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *TransactionAccount) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *TransactionAccount) SetAccountId(v string) {
	o.AccountId = &v
}

// GetBankAccountHolderName returns the BankAccountHolderName field value
func (o *TransactionAccount) GetBankAccountHolderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankAccountHolderName
}

// GetBankAccountHolderNameOk returns a tuple with the BankAccountHolderName field value
// and a boolean to check if the value has been set.
func (o *TransactionAccount) GetBankAccountHolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankAccountHolderName, true
}

// SetBankAccountHolderName sets field value
func (o *TransactionAccount) SetBankAccountHolderName(v string) {
	o.BankAccountHolderName = v
}

// GetBankName returns the BankName field value
func (o *TransactionAccount) GetBankName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value
// and a boolean to check if the value has been set.
func (o *TransactionAccount) GetBankNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankName, true
}

// SetBankName sets field value
func (o *TransactionAccount) SetBankName(v string) {
	o.BankName = v
}

// GetBankAccountNumberFormat returns the BankAccountNumberFormat field value
func (o *TransactionAccount) GetBankAccountNumberFormat() BankAccountNumberFormat {
	if o == nil {
		var ret BankAccountNumberFormat
		return ret
	}

	return o.BankAccountNumberFormat
}

// GetBankAccountNumberFormatOk returns a tuple with the BankAccountNumberFormat field value
// and a boolean to check if the value has been set.
func (o *TransactionAccount) GetBankAccountNumberFormatOk() (*BankAccountNumberFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankAccountNumberFormat, true
}

// SetBankAccountNumberFormat sets field value
func (o *TransactionAccount) SetBankAccountNumberFormat(v BankAccountNumberFormat) {
	o.BankAccountNumberFormat = v
}

// GetBankAccountNumberTail returns the BankAccountNumberTail field value if set, zero value otherwise.
func (o *TransactionAccount) GetBankAccountNumberTail() string {
	if o == nil || IsNil(o.BankAccountNumberTail) {
		var ret string
		return ret
	}
	return *o.BankAccountNumberTail
}

// GetBankAccountNumberTailOk returns a tuple with the BankAccountNumberTail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionAccount) GetBankAccountNumberTailOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccountNumberTail) {
		return nil, false
	}
	return o.BankAccountNumberTail, true
}

// HasBankAccountNumberTail returns a boolean if a field has been set.
func (o *TransactionAccount) HasBankAccountNumberTail() bool {
	if o != nil && !IsNil(o.BankAccountNumberTail) {
		return true
	}

	return false
}

// SetBankAccountNumberTail gets a reference to the given string and assigns it to the BankAccountNumberTail field.
func (o *TransactionAccount) SetBankAccountNumberTail(v string) {
	o.BankAccountNumberTail = &v
}

// GetBankAccountCountryCode returns the BankAccountCountryCode field value if set, zero value otherwise.
func (o *TransactionAccount) GetBankAccountCountryCode() string {
	if o == nil || IsNil(o.BankAccountCountryCode) {
		var ret string
		return ret
	}
	return *o.BankAccountCountryCode
}

// GetBankAccountCountryCodeOk returns a tuple with the BankAccountCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionAccount) GetBankAccountCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccountCountryCode) {
		return nil, false
	}
	return o.BankAccountCountryCode, true
}

// HasBankAccountCountryCode returns a boolean if a field has been set.
func (o *TransactionAccount) HasBankAccountCountryCode() bool {
	if o != nil && !IsNil(o.BankAccountCountryCode) {
		return true
	}

	return false
}

// SetBankAccountCountryCode gets a reference to the given string and assigns it to the BankAccountCountryCode field.
func (o *TransactionAccount) SetBankAccountCountryCode(v string) {
	o.BankAccountCountryCode = &v
}

// GetBankAccountCurrency returns the BankAccountCurrency field value
func (o *TransactionAccount) GetBankAccountCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankAccountCurrency
}

// GetBankAccountCurrencyOk returns a tuple with the BankAccountCurrency field value
// and a boolean to check if the value has been set.
func (o *TransactionAccount) GetBankAccountCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankAccountCurrency, true
}

// SetBankAccountCurrency sets field value
func (o *TransactionAccount) SetBankAccountCurrency(v string) {
	o.BankAccountCurrency = v
}

func (o TransactionAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	toSerialize["bankAccountHolderName"] = o.BankAccountHolderName
	toSerialize["bankName"] = o.BankName
	toSerialize["bankAccountNumberFormat"] = o.BankAccountNumberFormat
	if !IsNil(o.BankAccountNumberTail) {
		toSerialize["bankAccountNumberTail"] = o.BankAccountNumberTail
	}
	if !IsNil(o.BankAccountCountryCode) {
		toSerialize["bankAccountCountryCode"] = o.BankAccountCountryCode
	}
	toSerialize["bankAccountCurrency"] = o.BankAccountCurrency
	return toSerialize, nil
}

type NullableTransactionAccount struct {
	value *TransactionAccount
	isSet bool
}

func (v NullableTransactionAccount) Get() *TransactionAccount {
	return v.value
}

func (v *NullableTransactionAccount) Set(val *TransactionAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionAccount(val *TransactionAccount) *NullableTransactionAccount {
	return &NullableTransactionAccount{value: val, isSet: true}
}

func (v NullableTransactionAccount) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTransactionAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
