package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateConfirmCustomizationDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateConfirmCustomizationDetailsResponse{}

// CreateConfirmCustomizationDetailsResponse The response schema for the confirmCustomizationDetails operation.
type CreateConfirmCustomizationDetailsResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCreateConfirmCustomizationDetailsResponse instantiates a new CreateConfirmCustomizationDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConfirmCustomizationDetailsResponse() *CreateConfirmCustomizationDetailsResponse {
	this := CreateConfirmCustomizationDetailsResponse{}
	return &this
}

// NewCreateConfirmCustomizationDetailsResponseWithDefaults instantiates a new CreateConfirmCustomizationDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConfirmCustomizationDetailsResponseWithDefaults() *CreateConfirmCustomizationDetailsResponse {
	this := CreateConfirmCustomizationDetailsResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateConfirmCustomizationDetailsResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConfirmCustomizationDetailsResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateConfirmCustomizationDetailsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CreateConfirmCustomizationDetailsResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CreateConfirmCustomizationDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCreateConfirmCustomizationDetailsResponse struct {
	value *CreateConfirmCustomizationDetailsResponse
	isSet bool
}

func (v NullableCreateConfirmCustomizationDetailsResponse) Get() *CreateConfirmCustomizationDetailsResponse {
	return v.value
}

func (v *NullableCreateConfirmCustomizationDetailsResponse) Set(val *CreateConfirmCustomizationDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConfirmCustomizationDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConfirmCustomizationDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConfirmCustomizationDetailsResponse(val *CreateConfirmCustomizationDetailsResponse) *NullableCreateConfirmCustomizationDetailsResponse {
	return &NullableCreateConfirmCustomizationDetailsResponse{value: val, isSet: true}
}

func (v NullableCreateConfirmCustomizationDetailsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateConfirmCustomizationDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
