package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the ListAllFulfillmentOrdersResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAllFulfillmentOrdersResult{}

// ListAllFulfillmentOrdersResult The request for the listAllFulfillmentOrders operation.
type ListAllFulfillmentOrdersResult struct {
	// When present and not empty, pass this string token in the next request to return the next response page.
	NextToken *string `json:"nextToken,omitempty"`
	// An array of fulfillment order information.
	FulfillmentOrders []FulfillmentOrder `json:"fulfillmentOrders,omitempty"`
}

// NewListAllFulfillmentOrdersResult instantiates a new ListAllFulfillmentOrdersResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllFulfillmentOrdersResult() *ListAllFulfillmentOrdersResult {
	this := ListAllFulfillmentOrdersResult{}
	return &this
}

// NewListAllFulfillmentOrdersResultWithDefaults instantiates a new ListAllFulfillmentOrdersResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllFulfillmentOrdersResultWithDefaults() *ListAllFulfillmentOrdersResult {
	this := ListAllFulfillmentOrdersResult{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListAllFulfillmentOrdersResult) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAllFulfillmentOrdersResult) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListAllFulfillmentOrdersResult) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListAllFulfillmentOrdersResult) SetNextToken(v string) {
	o.NextToken = &v
}

// GetFulfillmentOrders returns the FulfillmentOrders field value if set, zero value otherwise.
func (o *ListAllFulfillmentOrdersResult) GetFulfillmentOrders() []FulfillmentOrder {
	if o == nil || IsNil(o.FulfillmentOrders) {
		var ret []FulfillmentOrder
		return ret
	}
	return o.FulfillmentOrders
}

// GetFulfillmentOrdersOk returns a tuple with the FulfillmentOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAllFulfillmentOrdersResult) GetFulfillmentOrdersOk() ([]FulfillmentOrder, bool) {
	if o == nil || IsNil(o.FulfillmentOrders) {
		return nil, false
	}
	return o.FulfillmentOrders, true
}

// HasFulfillmentOrders returns a boolean if a field has been set.
func (o *ListAllFulfillmentOrdersResult) HasFulfillmentOrders() bool {
	if o != nil && !IsNil(o.FulfillmentOrders) {
		return true
	}

	return false
}

// SetFulfillmentOrders gets a reference to the given []FulfillmentOrder and assigns it to the FulfillmentOrders field.
func (o *ListAllFulfillmentOrdersResult) SetFulfillmentOrders(v []FulfillmentOrder) {
	o.FulfillmentOrders = v
}

func (o ListAllFulfillmentOrdersResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.FulfillmentOrders) {
		toSerialize["fulfillmentOrders"] = o.FulfillmentOrders
	}
	return toSerialize, nil
}

type NullableListAllFulfillmentOrdersResult struct {
	value *ListAllFulfillmentOrdersResult
	isSet bool
}

func (v NullableListAllFulfillmentOrdersResult) Get() *ListAllFulfillmentOrdersResult {
	return v.value
}

func (v *NullableListAllFulfillmentOrdersResult) Set(val *ListAllFulfillmentOrdersResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListAllFulfillmentOrdersResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListAllFulfillmentOrdersResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAllFulfillmentOrdersResult(val *ListAllFulfillmentOrdersResult) *NullableListAllFulfillmentOrdersResult {
	return &NullableListAllFulfillmentOrdersResult{value: val, isSet: true}
}

func (v NullableListAllFulfillmentOrdersResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListAllFulfillmentOrdersResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
