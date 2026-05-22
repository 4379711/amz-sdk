package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the DeleteSubscriptionByIdResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSubscriptionByIdResponse{}

// DeleteSubscriptionByIdResponse The response schema for the `deleteSubscriptionById` operation.
type DeleteSubscriptionByIdResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewDeleteSubscriptionByIdResponse instantiates a new DeleteSubscriptionByIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSubscriptionByIdResponse() *DeleteSubscriptionByIdResponse {
	this := DeleteSubscriptionByIdResponse{}
	return &this
}

// NewDeleteSubscriptionByIdResponseWithDefaults instantiates a new DeleteSubscriptionByIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSubscriptionByIdResponseWithDefaults() *DeleteSubscriptionByIdResponse {
	this := DeleteSubscriptionByIdResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *DeleteSubscriptionByIdResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSubscriptionByIdResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *DeleteSubscriptionByIdResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *DeleteSubscriptionByIdResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o DeleteSubscriptionByIdResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableDeleteSubscriptionByIdResponse struct {
	value *DeleteSubscriptionByIdResponse
	isSet bool
}

func (v NullableDeleteSubscriptionByIdResponse) Get() *DeleteSubscriptionByIdResponse {
	return v.value
}

func (v *NullableDeleteSubscriptionByIdResponse) Set(val *DeleteSubscriptionByIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSubscriptionByIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSubscriptionByIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSubscriptionByIdResponse(val *DeleteSubscriptionByIdResponse) *NullableDeleteSubscriptionByIdResponse {
	return &NullableDeleteSubscriptionByIdResponse{value: val, isSet: true}
}

func (v NullableDeleteSubscriptionByIdResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeleteSubscriptionByIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
