package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateShipmentStatusErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateShipmentStatusErrorResponse{}

// UpdateShipmentStatusErrorResponse The error response schema for the `UpdateShipmentStatus` operation.
type UpdateShipmentStatusErrorResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewUpdateShipmentStatusErrorResponse instantiates a new UpdateShipmentStatusErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateShipmentStatusErrorResponse() *UpdateShipmentStatusErrorResponse {
	this := UpdateShipmentStatusErrorResponse{}
	return &this
}

// NewUpdateShipmentStatusErrorResponseWithDefaults instantiates a new UpdateShipmentStatusErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateShipmentStatusErrorResponseWithDefaults() *UpdateShipmentStatusErrorResponse {
	this := UpdateShipmentStatusErrorResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *UpdateShipmentStatusErrorResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateShipmentStatusErrorResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *UpdateShipmentStatusErrorResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *UpdateShipmentStatusErrorResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o UpdateShipmentStatusErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableUpdateShipmentStatusErrorResponse struct {
	value *UpdateShipmentStatusErrorResponse
	isSet bool
}

func (v NullableUpdateShipmentStatusErrorResponse) Get() *UpdateShipmentStatusErrorResponse {
	return v.value
}

func (v *NullableUpdateShipmentStatusErrorResponse) Set(val *UpdateShipmentStatusErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateShipmentStatusErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateShipmentStatusErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateShipmentStatusErrorResponse(val *UpdateShipmentStatusErrorResponse) *NullableUpdateShipmentStatusErrorResponse {
	return &NullableUpdateShipmentStatusErrorResponse{value: val, isSet: true}
}

func (v NullableUpdateShipmentStatusErrorResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateShipmentStatusErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
