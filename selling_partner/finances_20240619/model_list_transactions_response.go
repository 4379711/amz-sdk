package finances_20240619

import (
	"github.com/bytedance/sonic"
)

// checks if the ListTransactionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTransactionsResponse{}

// ListTransactionsResponse The Response schema.
type ListTransactionsResponseWrapper struct {
	Payload    ListTransactionsResponse `json:"payload,omitempty"`
	StatusCode int32                    `json:"statusCode,omitempty"`
}

// ListTransactionsResponse The response schema for the `listTransactions` operation.
type ListTransactionsResponse struct {
	// The response includes `nextToken` when the number of results exceeds the specified `pageSize` value. To get the next page of results, call the operation with this token and include the same arguments as the call that produced the token. To get a complete list, call this operation until `nextToken` is null. Note that this operation can return empty pages.
	NextToken *string `json:"nextToken,omitempty"`
	// A list of transactions within the specified time period.
	Transactions []Transaction `json:"transactions,omitempty"`
}

// NewListTransactionsResponse instantiates a new ListTransactionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsResponse() *ListTransactionsResponse {
	this := ListTransactionsResponse{}
	return &this
}

// NewListTransactionsResponseWithDefaults instantiates a new ListTransactionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsResponseWithDefaults() *ListTransactionsResponse {
	this := ListTransactionsResponse{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListTransactionsResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransactionsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListTransactionsResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListTransactionsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *ListTransactionsResponse) GetTransactions() []Transaction {
	if o == nil || IsNil(o.Transactions) {
		var ret []Transaction
		return ret
	}
	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransactionsResponse) GetTransactionsOk() ([]Transaction, bool) {
	if o == nil || IsNil(o.Transactions) {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *ListTransactionsResponse) HasTransactions() bool {
	if o != nil && !IsNil(o.Transactions) {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given []Transaction and assigns it to the Transactions field.
func (o *ListTransactionsResponse) SetTransactions(v []Transaction) {
	o.Transactions = v
}

func (o ListTransactionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.Transactions) {
		toSerialize["transactions"] = o.Transactions
	}
	return toSerialize, nil
}

type NullableListTransactionsResponse struct {
	value *ListTransactionsResponse
	isSet bool
}

func (v NullableListTransactionsResponse) Get() *ListTransactionsResponse {
	return v.value
}

func (v *NullableListTransactionsResponse) Set(val *ListTransactionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsResponse(val *ListTransactionsResponse) *NullableListTransactionsResponse {
	return &NullableListTransactionsResponse{value: val, isSet: true}
}

func (v NullableListTransactionsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListTransactionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
