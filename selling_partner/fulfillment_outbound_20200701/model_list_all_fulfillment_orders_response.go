package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the ListAllFulfillmentOrdersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAllFulfillmentOrdersResponse{}

// ListAllFulfillmentOrdersResponse The response schema for the `listAllFulfillmentOrders` operation.
type ListAllFulfillmentOrdersResponse struct {
	Payload *ListAllFulfillmentOrdersResult `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewListAllFulfillmentOrdersResponse instantiates a new ListAllFulfillmentOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllFulfillmentOrdersResponse() *ListAllFulfillmentOrdersResponse {
	this := ListAllFulfillmentOrdersResponse{}
	return &this
}

// NewListAllFulfillmentOrdersResponseWithDefaults instantiates a new ListAllFulfillmentOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllFulfillmentOrdersResponseWithDefaults() *ListAllFulfillmentOrdersResponse {
	this := ListAllFulfillmentOrdersResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ListAllFulfillmentOrdersResponse) GetPayload() ListAllFulfillmentOrdersResult {
	if o == nil || IsNil(o.Payload) {
		var ret ListAllFulfillmentOrdersResult
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAllFulfillmentOrdersResponse) GetPayloadOk() (*ListAllFulfillmentOrdersResult, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ListAllFulfillmentOrdersResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given ListAllFulfillmentOrdersResult and assigns it to the Payload field.
func (o *ListAllFulfillmentOrdersResponse) SetPayload(v ListAllFulfillmentOrdersResult) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ListAllFulfillmentOrdersResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAllFulfillmentOrdersResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ListAllFulfillmentOrdersResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *ListAllFulfillmentOrdersResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o ListAllFulfillmentOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableListAllFulfillmentOrdersResponse struct {
	value *ListAllFulfillmentOrdersResponse
	isSet bool
}

func (v NullableListAllFulfillmentOrdersResponse) Get() *ListAllFulfillmentOrdersResponse {
	return v.value
}

func (v *NullableListAllFulfillmentOrdersResponse) Set(val *ListAllFulfillmentOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListAllFulfillmentOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListAllFulfillmentOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAllFulfillmentOrdersResponse(val *ListAllFulfillmentOrdersResponse) *NullableListAllFulfillmentOrdersResponse {
	return &NullableListAllFulfillmentOrdersResponse{value: val, isSet: true}
}

func (v NullableListAllFulfillmentOrdersResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListAllFulfillmentOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
