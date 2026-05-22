package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the CancelShipmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelShipmentResponse{}

// CancelShipmentResponse The response schema for the cancelShipment operation.
type CancelShipmentResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCancelShipmentResponse instantiates a new CancelShipmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelShipmentResponse() *CancelShipmentResponse {
	this := CancelShipmentResponse{}
	return &this
}

// NewCancelShipmentResponseWithDefaults instantiates a new CancelShipmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelShipmentResponseWithDefaults() *CancelShipmentResponse {
	this := CancelShipmentResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CancelShipmentResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelShipmentResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CancelShipmentResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CancelShipmentResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CancelShipmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCancelShipmentResponse struct {
	value *CancelShipmentResponse
	isSet bool
}

func (v NullableCancelShipmentResponse) Get() *CancelShipmentResponse {
	return v.value
}

func (v *NullableCancelShipmentResponse) Set(val *CancelShipmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelShipmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelShipmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelShipmentResponse(val *CancelShipmentResponse) *NullableCancelShipmentResponse {
	return &NullableCancelShipmentResponse{value: val, isSet: true}
}

func (v NullableCancelShipmentResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCancelShipmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
