package seller_wallet_20240301

import (
	"github.com/bytedance/sonic"
)

// checks if the BankAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BankAccount{}

// BankAccount Details of an Amazon Seller Wallet bank account. This account is used to hold the money that a Seller Wallet customer earns by selling items.
type BankAccount struct {
	// The unique identifier provided by Amazon to identify the account.
	AccountId *string `json:"accountId,omitempty"`
	// The bank account holder's name (expected to be an Amazon customer).
	AccountHolderName       string                  `json:"accountHolderName"`
	BankAccountNumberFormat BankAccountNumberFormat `json:"bankAccountNumberFormat"`
	// The name of the bank. This value is Amazon Seller Wallet for Amazon Seller Wallet accounts.
	BankName                 *string                  `json:"bankName,omitempty"`
	BankAccountOwnershipType BankAccountOwnershipType `json:"bankAccountOwnershipType"`
	// Routing number for automated clearing house transfers. This value is nine consecutive zeros for Amazon Seller Wallet accounts.
	RoutingNumber    string           `json:"routingNumber"`
	BankNumberFormat BankNumberFormat `json:"bankNumberFormat"`
	// The two-digit country code in ISO 3166 format.
	AccountCountryCode string `json:"accountCountryCode"`
	// Bank account currency code in ISO 4217 format.
	AccountCurrency string `json:"accountCurrency"`
	// The last 3 digit of the bank account number. This value is three consecutive zeros for Amazon Seller Wallet accounts.
	BankAccountNumberTail   string                   `json:"bankAccountNumberTail"`
	BankAccountHolderStatus *BankAccountHolderStatus `json:"bankAccountHolderStatus,omitempty"`
}

type _BankAccount BankAccount

// NewBankAccount instantiates a new BankAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankAccount(accountHolderName string, bankAccountNumberFormat BankAccountNumberFormat, bankAccountOwnershipType BankAccountOwnershipType, routingNumber string, bankNumberFormat BankNumberFormat, accountCountryCode string, accountCurrency string, bankAccountNumberTail string) *BankAccount {
	this := BankAccount{}
	this.AccountHolderName = accountHolderName
	this.BankAccountNumberFormat = bankAccountNumberFormat
	this.BankAccountOwnershipType = bankAccountOwnershipType
	this.RoutingNumber = routingNumber
	this.BankNumberFormat = bankNumberFormat
	this.AccountCountryCode = accountCountryCode
	this.AccountCurrency = accountCurrency
	this.BankAccountNumberTail = bankAccountNumberTail
	return &this
}

// NewBankAccountWithDefaults instantiates a new BankAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankAccountWithDefaults() *BankAccount {
	this := BankAccount{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *BankAccount) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *BankAccount) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *BankAccount) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAccountHolderName returns the AccountHolderName field value
func (o *BankAccount) GetAccountHolderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountHolderName
}

// GetAccountHolderNameOk returns a tuple with the AccountHolderName field value
// and a boolean to check if the value has been set.
func (o *BankAccount) GetAccountHolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountHolderName, true
}

// SetAccountHolderName sets field value
func (o *BankAccount) SetAccountHolderName(v string) {
	o.AccountHolderName = v
}

// GetBankAccountNumberFormat returns the BankAccountNumberFormat field value
func (o *BankAccount) GetBankAccountNumberFormat() BankAccountNumberFormat {
	if o == nil {
		var ret BankAccountNumberFormat
		return ret
	}

	return o.BankAccountNumberFormat
}

// GetBankAccountNumberFormatOk returns a tuple with the BankAccountNumberFormat field value
// and a boolean to check if the value has been set.
func (o *BankAccount) GetBankAccountNumberFormatOk() (*BankAccountNumberFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankAccountNumberFormat, true
}

// SetBankAccountNumberFormat sets field value
func (o *BankAccount) SetBankAccountNumberFormat(v BankAccountNumberFormat) {
	o.BankAccountNumberFormat = v
}

// GetBankName returns the BankName field value if set, zero value otherwise.
func (o *BankAccount) GetBankName() string {
	if o == nil || IsNil(o.BankName) {
		var ret string
		return ret
	}
	return *o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetBankNameOk() (*string, bool) {
	if o == nil || IsNil(o.BankName) {
		return nil, false
	}
	return o.BankName, true
}

// HasBankName returns a boolean if a field has been set.
func (o *BankAccount) HasBankName() bool {
	if o != nil && !IsNil(o.BankName) {
		return true
	}

	return false
}

// SetBankName gets a reference to the given string and assigns it to the BankName field.
func (o *BankAccount) SetBankName(v string) {
	o.BankName = &v
}

// GetBankAccountOwnershipType returns the BankAccountOwnershipType field value
func (o *BankAccount) GetBankAccountOwnershipType() BankAccountOwnershipType {
	if o == nil {
		var ret BankAccountOwnershipType
		return ret
	}

	return o.BankAccountOwnershipType
}

// GetBankAccountOwnershipTypeOk returns a tuple with the BankAccountOwnershipType field value
// and a boolean to check if the value has been set.
func (o *BankAccount) GetBankAccountOwnershipTypeOk() (*BankAccountOwnershipType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankAccountOwnershipType, true
}

// SetBankAccountOwnershipType sets field value
func (o *BankAccount) SetBankAccountOwnershipType(v BankAccountOwnershipType) {
	o.BankAccountOwnershipType = v
}

// GetRoutingNumber returns the RoutingNumber field value
func (o *BankAccount) GetRoutingNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value
// and a boolean to check if the value has been set.
func (o *BankAccount) GetRoutingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoutingNumber, true
}

// SetRoutingNumber sets field value
func (o *BankAccount) SetRoutingNumber(v string) {
	o.RoutingNumber = v
}

// GetBankNumberFormat returns the BankNumberFormat field value
func (o *BankAccount) GetBankNumberFormat() BankNumberFormat {
	if o == nil {
		var ret BankNumberFormat
		return ret
	}

	return o.BankNumberFormat
}

// GetBankNumberFormatOk returns a tuple with the BankNumberFormat field value
// and a boolean to check if the value has been set.
func (o *BankAccount) GetBankNumberFormatOk() (*BankNumberFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankNumberFormat, true
}

// SetBankNumberFormat sets field value
func (o *BankAccount) SetBankNumberFormat(v BankNumberFormat) {
	o.BankNumberFormat = v
}

// GetAccountCountryCode returns the AccountCountryCode field value
func (o *BankAccount) GetAccountCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountCountryCode
}

// GetAccountCountryCodeOk returns a tuple with the AccountCountryCode field value
// and a boolean to check if the value has been set.
func (o *BankAccount) GetAccountCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountCountryCode, true
}

// SetAccountCountryCode sets field value
func (o *BankAccount) SetAccountCountryCode(v string) {
	o.AccountCountryCode = v
}

// GetAccountCurrency returns the AccountCurrency field value
func (o *BankAccount) GetAccountCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountCurrency
}

// GetAccountCurrencyOk returns a tuple with the AccountCurrency field value
// and a boolean to check if the value has been set.
func (o *BankAccount) GetAccountCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountCurrency, true
}

// SetAccountCurrency sets field value
func (o *BankAccount) SetAccountCurrency(v string) {
	o.AccountCurrency = v
}

// GetBankAccountNumberTail returns the BankAccountNumberTail field value
func (o *BankAccount) GetBankAccountNumberTail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankAccountNumberTail
}

// GetBankAccountNumberTailOk returns a tuple with the BankAccountNumberTail field value
// and a boolean to check if the value has been set.
func (o *BankAccount) GetBankAccountNumberTailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankAccountNumberTail, true
}

// SetBankAccountNumberTail sets field value
func (o *BankAccount) SetBankAccountNumberTail(v string) {
	o.BankAccountNumberTail = v
}

// GetBankAccountHolderStatus returns the BankAccountHolderStatus field value if set, zero value otherwise.
func (o *BankAccount) GetBankAccountHolderStatus() BankAccountHolderStatus {
	if o == nil || IsNil(o.BankAccountHolderStatus) {
		var ret BankAccountHolderStatus
		return ret
	}
	return *o.BankAccountHolderStatus
}

// GetBankAccountHolderStatusOk returns a tuple with the BankAccountHolderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetBankAccountHolderStatusOk() (*BankAccountHolderStatus, bool) {
	if o == nil || IsNil(o.BankAccountHolderStatus) {
		return nil, false
	}
	return o.BankAccountHolderStatus, true
}

// HasBankAccountHolderStatus returns a boolean if a field has been set.
func (o *BankAccount) HasBankAccountHolderStatus() bool {
	if o != nil && !IsNil(o.BankAccountHolderStatus) {
		return true
	}

	return false
}

// SetBankAccountHolderStatus gets a reference to the given BankAccountHolderStatus and assigns it to the BankAccountHolderStatus field.
func (o *BankAccount) SetBankAccountHolderStatus(v BankAccountHolderStatus) {
	o.BankAccountHolderStatus = &v
}

func (o BankAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	toSerialize["accountHolderName"] = o.AccountHolderName
	toSerialize["bankAccountNumberFormat"] = o.BankAccountNumberFormat
	if !IsNil(o.BankName) {
		toSerialize["bankName"] = o.BankName
	}
	toSerialize["bankAccountOwnershipType"] = o.BankAccountOwnershipType
	toSerialize["routingNumber"] = o.RoutingNumber
	toSerialize["bankNumberFormat"] = o.BankNumberFormat
	toSerialize["accountCountryCode"] = o.AccountCountryCode
	toSerialize["accountCurrency"] = o.AccountCurrency
	toSerialize["bankAccountNumberTail"] = o.BankAccountNumberTail
	if !IsNil(o.BankAccountHolderStatus) {
		toSerialize["bankAccountHolderStatus"] = o.BankAccountHolderStatus
	}
	return toSerialize, nil
}

type NullableBankAccount struct {
	value *BankAccount
	isSet bool
}

func (v NullableBankAccount) Get() *BankAccount {
	return v.value
}

func (v *NullableBankAccount) Set(val *BankAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableBankAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableBankAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankAccount(val *BankAccount) *NullableBankAccount {
	return &NullableBankAccount{value: val, isSet: true}
}

func (v NullableBankAccount) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBankAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
