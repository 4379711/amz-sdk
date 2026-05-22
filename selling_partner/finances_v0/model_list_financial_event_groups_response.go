package finances_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ListFinancialEventGroupsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFinancialEventGroupsResponse{}

// ListFinancialEventGroupsResponse The response schema for the listFinancialEventGroups operation.
type ListFinancialEventGroupsResponse struct {
	Payload *ListFinancialEventGroupsPayload `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewListFinancialEventGroupsResponse instantiates a new ListFinancialEventGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFinancialEventGroupsResponse() *ListFinancialEventGroupsResponse {
	this := ListFinancialEventGroupsResponse{}
	return &this
}

// NewListFinancialEventGroupsResponseWithDefaults instantiates a new ListFinancialEventGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFinancialEventGroupsResponseWithDefaults() *ListFinancialEventGroupsResponse {
	this := ListFinancialEventGroupsResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ListFinancialEventGroupsResponse) GetPayload() ListFinancialEventGroupsPayload {
	if o == nil || IsNil(o.Payload) {
		var ret ListFinancialEventGroupsPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFinancialEventGroupsResponse) GetPayloadOk() (*ListFinancialEventGroupsPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ListFinancialEventGroupsResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given ListFinancialEventGroupsPayload and assigns it to the Payload field.
func (o *ListFinancialEventGroupsResponse) SetPayload(v ListFinancialEventGroupsPayload) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ListFinancialEventGroupsResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFinancialEventGroupsResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ListFinancialEventGroupsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *ListFinancialEventGroupsResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o ListFinancialEventGroupsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableListFinancialEventGroupsResponse struct {
	value *ListFinancialEventGroupsResponse
	isSet bool
}

func (v NullableListFinancialEventGroupsResponse) Get() *ListFinancialEventGroupsResponse {
	return v.value
}

func (v *NullableListFinancialEventGroupsResponse) Set(val *ListFinancialEventGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListFinancialEventGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListFinancialEventGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFinancialEventGroupsResponse(val *ListFinancialEventGroupsResponse) *NullableListFinancialEventGroupsResponse {
	return &NullableListFinancialEventGroupsResponse{value: val, isSet: true}
}

func (v NullableListFinancialEventGroupsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListFinancialEventGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
