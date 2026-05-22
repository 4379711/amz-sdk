package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateWarrantyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWarrantyResponse{}

// CreateWarrantyResponse The response schema for the createWarranty operation.
type CreateWarrantyResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCreateWarrantyResponse instantiates a new CreateWarrantyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWarrantyResponse() *CreateWarrantyResponse {
	this := CreateWarrantyResponse{}
	return &this
}

// NewCreateWarrantyResponseWithDefaults instantiates a new CreateWarrantyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWarrantyResponseWithDefaults() *CreateWarrantyResponse {
	this := CreateWarrantyResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateWarrantyResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWarrantyResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateWarrantyResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CreateWarrantyResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CreateWarrantyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCreateWarrantyResponse struct {
	value *CreateWarrantyResponse
	isSet bool
}

func (v NullableCreateWarrantyResponse) Get() *CreateWarrantyResponse {
	return v.value
}

func (v *NullableCreateWarrantyResponse) Set(val *CreateWarrantyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWarrantyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWarrantyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWarrantyResponse(val *CreateWarrantyResponse) *NullableCreateWarrantyResponse {
	return &NullableCreateWarrantyResponse{value: val, isSet: true}
}

func (v NullableCreateWarrantyResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateWarrantyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
