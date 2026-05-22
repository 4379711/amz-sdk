package finances_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ListFinancialEventsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFinancialEventsResponse{}

// ListFinancialEventsResponse The response schema for the listFinancialEvents operation.
type ListFinancialEventsResponse struct {
	Payload *ListFinancialEventsPayload `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewListFinancialEventsResponse instantiates a new ListFinancialEventsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFinancialEventsResponse() *ListFinancialEventsResponse {
	this := ListFinancialEventsResponse{}
	return &this
}

// NewListFinancialEventsResponseWithDefaults instantiates a new ListFinancialEventsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFinancialEventsResponseWithDefaults() *ListFinancialEventsResponse {
	this := ListFinancialEventsResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ListFinancialEventsResponse) GetPayload() ListFinancialEventsPayload {
	if o == nil || IsNil(o.Payload) {
		var ret ListFinancialEventsPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFinancialEventsResponse) GetPayloadOk() (*ListFinancialEventsPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ListFinancialEventsResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given ListFinancialEventsPayload and assigns it to the Payload field.
func (o *ListFinancialEventsResponse) SetPayload(v ListFinancialEventsPayload) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ListFinancialEventsResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFinancialEventsResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ListFinancialEventsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *ListFinancialEventsResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o ListFinancialEventsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableListFinancialEventsResponse struct {
	value *ListFinancialEventsResponse
	isSet bool
}

func (v NullableListFinancialEventsResponse) Get() *ListFinancialEventsResponse {
	return v.value
}

func (v *NullableListFinancialEventsResponse) Set(val *ListFinancialEventsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListFinancialEventsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListFinancialEventsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFinancialEventsResponse(val *ListFinancialEventsResponse) *NullableListFinancialEventsResponse {
	return &NullableListFinancialEventsResponse{value: val, isSet: true}
}

func (v NullableListFinancialEventsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListFinancialEventsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
