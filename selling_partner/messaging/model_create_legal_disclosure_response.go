package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateLegalDisclosureResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLegalDisclosureResponse{}

// CreateLegalDisclosureResponse The response schema for the createLegalDisclosure operation.
type CreateLegalDisclosureResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCreateLegalDisclosureResponse instantiates a new CreateLegalDisclosureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLegalDisclosureResponse() *CreateLegalDisclosureResponse {
	this := CreateLegalDisclosureResponse{}
	return &this
}

// NewCreateLegalDisclosureResponseWithDefaults instantiates a new CreateLegalDisclosureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLegalDisclosureResponseWithDefaults() *CreateLegalDisclosureResponse {
	this := CreateLegalDisclosureResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateLegalDisclosureResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLegalDisclosureResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateLegalDisclosureResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CreateLegalDisclosureResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CreateLegalDisclosureResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCreateLegalDisclosureResponse struct {
	value *CreateLegalDisclosureResponse
	isSet bool
}

func (v NullableCreateLegalDisclosureResponse) Get() *CreateLegalDisclosureResponse {
	return v.value
}

func (v *NullableCreateLegalDisclosureResponse) Set(val *CreateLegalDisclosureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLegalDisclosureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLegalDisclosureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLegalDisclosureResponse(val *CreateLegalDisclosureResponse) *NullableCreateLegalDisclosureResponse {
	return &NullableCreateLegalDisclosureResponse{value: val, isSet: true}
}

func (v NullableCreateLegalDisclosureResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateLegalDisclosureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
