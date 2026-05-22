package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateFulfillmentOrderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFulfillmentOrderResponse{}

// UpdateFulfillmentOrderResponse The response schema for the `updateFulfillmentOrder` operation.
type UpdateFulfillmentOrderResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewUpdateFulfillmentOrderResponse instantiates a new UpdateFulfillmentOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFulfillmentOrderResponse() *UpdateFulfillmentOrderResponse {
	this := UpdateFulfillmentOrderResponse{}
	return &this
}

// NewUpdateFulfillmentOrderResponseWithDefaults instantiates a new UpdateFulfillmentOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFulfillmentOrderResponseWithDefaults() *UpdateFulfillmentOrderResponse {
	this := UpdateFulfillmentOrderResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *UpdateFulfillmentOrderResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFulfillmentOrderResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *UpdateFulfillmentOrderResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *UpdateFulfillmentOrderResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o UpdateFulfillmentOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableUpdateFulfillmentOrderResponse struct {
	value *UpdateFulfillmentOrderResponse
	isSet bool
}

func (v NullableUpdateFulfillmentOrderResponse) Get() *UpdateFulfillmentOrderResponse {
	return v.value
}

func (v *NullableUpdateFulfillmentOrderResponse) Set(val *UpdateFulfillmentOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFulfillmentOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFulfillmentOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFulfillmentOrderResponse(val *UpdateFulfillmentOrderResponse) *NullableUpdateFulfillmentOrderResponse {
	return &NullableUpdateFulfillmentOrderResponse{value: val, isSet: true}
}

func (v NullableUpdateFulfillmentOrderResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateFulfillmentOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
