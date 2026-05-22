package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ConfirmShipmentErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfirmShipmentErrorResponse{}

// ConfirmShipmentErrorResponse The error response schema for the `confirmShipment` operation.
type ConfirmShipmentErrorResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewConfirmShipmentErrorResponse instantiates a new ConfirmShipmentErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmShipmentErrorResponse() *ConfirmShipmentErrorResponse {
	this := ConfirmShipmentErrorResponse{}
	return &this
}

// NewConfirmShipmentErrorResponseWithDefaults instantiates a new ConfirmShipmentErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmShipmentErrorResponseWithDefaults() *ConfirmShipmentErrorResponse {
	this := ConfirmShipmentErrorResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ConfirmShipmentErrorResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfirmShipmentErrorResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ConfirmShipmentErrorResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *ConfirmShipmentErrorResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o ConfirmShipmentErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableConfirmShipmentErrorResponse struct {
	value *ConfirmShipmentErrorResponse
	isSet bool
}

func (v NullableConfirmShipmentErrorResponse) Get() *ConfirmShipmentErrorResponse {
	return v.value
}

func (v *NullableConfirmShipmentErrorResponse) Set(val *ConfirmShipmentErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmShipmentErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmShipmentErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmShipmentErrorResponse(val *ConfirmShipmentErrorResponse) *NullableConfirmShipmentErrorResponse {
	return &NullableConfirmShipmentErrorResponse{value: val, isSet: true}
}

func (v NullableConfirmShipmentErrorResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableConfirmShipmentErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
