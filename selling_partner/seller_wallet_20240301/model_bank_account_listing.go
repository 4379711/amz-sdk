package seller_wallet_20240301

import (
	"github.com/bytedance/sonic"
)

// checks if the BankAccountListing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BankAccountListing{}

// BankAccountListing A list of bank accounts.
type BankAccountListing struct {
	// A list of bank accounts.
	Accounts []BankAccount `json:"accounts"`
}

type _BankAccountListing BankAccountListing

// NewBankAccountListing instantiates a new BankAccountListing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankAccountListing(accounts []BankAccount) *BankAccountListing {
	this := BankAccountListing{}
	this.Accounts = accounts
	return &this
}

// NewBankAccountListingWithDefaults instantiates a new BankAccountListing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankAccountListingWithDefaults() *BankAccountListing {
	this := BankAccountListing{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *BankAccountListing) GetAccounts() []BankAccount {
	if o == nil {
		var ret []BankAccount
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *BankAccountListing) GetAccountsOk() ([]BankAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Accounts, true
}

// SetAccounts sets field value
func (o *BankAccountListing) SetAccounts(v []BankAccount) {
	o.Accounts = v
}

func (o BankAccountListing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accounts"] = o.Accounts
	return toSerialize, nil
}

type NullableBankAccountListing struct {
	value *BankAccountListing
	isSet bool
}

func (v NullableBankAccountListing) Get() *BankAccountListing {
	return v.value
}

func (v *NullableBankAccountListing) Set(val *BankAccountListing) {
	v.value = val
	v.isSet = true
}

func (v NullableBankAccountListing) IsSet() bool {
	return v.isSet
}

func (v *NullableBankAccountListing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankAccountListing(val *BankAccountListing) *NullableBankAccountListing {
	return &NullableBankAccountListing{value: val, isSet: true}
}

func (v NullableBankAccountListing) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBankAccountListing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
