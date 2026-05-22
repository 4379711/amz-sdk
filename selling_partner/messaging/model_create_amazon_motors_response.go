package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateAmazonMotorsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAmazonMotorsResponse{}

// CreateAmazonMotorsResponse The response schema for the createAmazonMotors operation.
type CreateAmazonMotorsResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCreateAmazonMotorsResponse instantiates a new CreateAmazonMotorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAmazonMotorsResponse() *CreateAmazonMotorsResponse {
	this := CreateAmazonMotorsResponse{}
	return &this
}

// NewCreateAmazonMotorsResponseWithDefaults instantiates a new CreateAmazonMotorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAmazonMotorsResponseWithDefaults() *CreateAmazonMotorsResponse {
	this := CreateAmazonMotorsResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateAmazonMotorsResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAmazonMotorsResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateAmazonMotorsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CreateAmazonMotorsResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CreateAmazonMotorsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCreateAmazonMotorsResponse struct {
	value *CreateAmazonMotorsResponse
	isSet bool
}

func (v NullableCreateAmazonMotorsResponse) Get() *CreateAmazonMotorsResponse {
	return v.value
}

func (v *NullableCreateAmazonMotorsResponse) Set(val *CreateAmazonMotorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAmazonMotorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAmazonMotorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAmazonMotorsResponse(val *CreateAmazonMotorsResponse) *NullableCreateAmazonMotorsResponse {
	return &NullableCreateAmazonMotorsResponse{value: val, isSet: true}
}

func (v NullableCreateAmazonMotorsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateAmazonMotorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
