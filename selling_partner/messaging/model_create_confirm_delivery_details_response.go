package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateConfirmDeliveryDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateConfirmDeliveryDetailsResponse{}

// CreateConfirmDeliveryDetailsResponse The response schema for the createConfirmDeliveryDetails operation.
type CreateConfirmDeliveryDetailsResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCreateConfirmDeliveryDetailsResponse instantiates a new CreateConfirmDeliveryDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConfirmDeliveryDetailsResponse() *CreateConfirmDeliveryDetailsResponse {
	this := CreateConfirmDeliveryDetailsResponse{}
	return &this
}

// NewCreateConfirmDeliveryDetailsResponseWithDefaults instantiates a new CreateConfirmDeliveryDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConfirmDeliveryDetailsResponseWithDefaults() *CreateConfirmDeliveryDetailsResponse {
	this := CreateConfirmDeliveryDetailsResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateConfirmDeliveryDetailsResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConfirmDeliveryDetailsResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateConfirmDeliveryDetailsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CreateConfirmDeliveryDetailsResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CreateConfirmDeliveryDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCreateConfirmDeliveryDetailsResponse struct {
	value *CreateConfirmDeliveryDetailsResponse
	isSet bool
}

func (v NullableCreateConfirmDeliveryDetailsResponse) Get() *CreateConfirmDeliveryDetailsResponse {
	return v.value
}

func (v *NullableCreateConfirmDeliveryDetailsResponse) Set(val *CreateConfirmDeliveryDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConfirmDeliveryDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConfirmDeliveryDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConfirmDeliveryDetailsResponse(val *CreateConfirmDeliveryDetailsResponse) *NullableCreateConfirmDeliveryDetailsResponse {
	return &NullableCreateConfirmDeliveryDetailsResponse{value: val, isSet: true}
}

func (v NullableCreateConfirmDeliveryDetailsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateConfirmDeliveryDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
