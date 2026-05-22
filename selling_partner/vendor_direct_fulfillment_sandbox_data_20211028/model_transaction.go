package vendor_direct_fulfillment_sandbox_data_20211028

import (
	"github.com/bytedance/sonic"
)

// checks if the Transaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Transaction{}

// Transaction The transaction details including the status. If the transaction was successful, also includes the requested test order data.
type Transaction struct {
	// The unique identifier returned in the response to the generateOrderScenarios request.
	TransactionId string `json:"transactionId"`
	// The current processing status of the transaction.
	Status       string        `json:"status"`
	TestCaseData *TestCaseData `json:"testCaseData,omitempty"`
}

type _Transaction Transaction

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction(transactionId string, status string) *Transaction {
	this := Transaction{}
	this.TransactionId = transactionId
	this.Status = status
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	return &this
}

// GetTransactionId returns the TransactionId field value
func (o *Transaction) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *Transaction) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetStatus returns the Status field value
func (o *Transaction) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Transaction) SetStatus(v string) {
	o.Status = v
}

// GetTestCaseData returns the TestCaseData field value if set, zero value otherwise.
func (o *Transaction) GetTestCaseData() TestCaseData {
	if o == nil || IsNil(o.TestCaseData) {
		var ret TestCaseData
		return ret
	}
	return *o.TestCaseData
}

// GetTestCaseDataOk returns a tuple with the TestCaseData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTestCaseDataOk() (*TestCaseData, bool) {
	if o == nil || IsNil(o.TestCaseData) {
		return nil, false
	}
	return o.TestCaseData, true
}

// HasTestCaseData returns a boolean if a field has been set.
func (o *Transaction) HasTestCaseData() bool {
	if o != nil && !IsNil(o.TestCaseData) {
		return true
	}

	return false
}

// SetTestCaseData gets a reference to the given TestCaseData and assigns it to the TestCaseData field.
func (o *Transaction) SetTestCaseData(v TestCaseData) {
	o.TestCaseData = &v
}

func (o Transaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transactionId"] = o.TransactionId
	toSerialize["status"] = o.Status
	if !IsNil(o.TestCaseData) {
		toSerialize["testCaseData"] = o.TestCaseData
	}
	return toSerialize, nil
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
