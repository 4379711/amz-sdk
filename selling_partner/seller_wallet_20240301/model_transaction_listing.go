package seller_wallet_20240301

import (
	"github.com/bytedance/sonic"
)

// checks if the TransactionListing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionListing{}

// TransactionListing A list of transactions.
type TransactionListing struct {
	// A token that you use to retrieve the next page of results. The response includes `nextPageToken` when the number of results exceeds 100. To get the next page of results, call the operation with this token and include the same arguments as the call that produced the token. To get a complete list, call this operation until `nextPageToken` is null. Note that this operation can return empty pages.
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// A list of transactions.
	Transactions []Transaction `json:"transactions"`
}

type _TransactionListing TransactionListing

// NewTransactionListing instantiates a new TransactionListing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionListing(transactions []Transaction) *TransactionListing {
	this := TransactionListing{}
	this.Transactions = transactions
	return &this
}

// NewTransactionListingWithDefaults instantiates a new TransactionListing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionListingWithDefaults() *TransactionListing {
	this := TransactionListing{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *TransactionListing) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionListing) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *TransactionListing) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *TransactionListing) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetTransactions returns the Transactions field value
func (o *TransactionListing) GetTransactions() []Transaction {
	if o == nil {
		var ret []Transaction
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *TransactionListing) GetTransactionsOk() ([]Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transactions, true
}

// SetTransactions sets field value
func (o *TransactionListing) SetTransactions(v []Transaction) {
	o.Transactions = v
}

func (o TransactionListing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	toSerialize["transactions"] = o.Transactions
	return toSerialize, nil
}

type NullableTransactionListing struct {
	value *TransactionListing
	isSet bool
}

func (v NullableTransactionListing) Get() *TransactionListing {
	return v.value
}

func (v *NullableTransactionListing) Set(val *TransactionListing) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionListing) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionListing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionListing(val *TransactionListing) *NullableTransactionListing {
	return &NullableTransactionListing{value: val, isSet: true}
}

func (v NullableTransactionListing) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTransactionListing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
