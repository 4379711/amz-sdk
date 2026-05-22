package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the DeleteDestinationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteDestinationResponse{}

// DeleteDestinationResponse The response schema for the `deleteDestination` operation.
type DeleteDestinationResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewDeleteDestinationResponse instantiates a new DeleteDestinationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteDestinationResponse() *DeleteDestinationResponse {
	this := DeleteDestinationResponse{}
	return &this
}

// NewDeleteDestinationResponseWithDefaults instantiates a new DeleteDestinationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteDestinationResponseWithDefaults() *DeleteDestinationResponse {
	this := DeleteDestinationResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *DeleteDestinationResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDestinationResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *DeleteDestinationResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *DeleteDestinationResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o DeleteDestinationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableDeleteDestinationResponse struct {
	value *DeleteDestinationResponse
	isSet bool
}

func (v NullableDeleteDestinationResponse) Get() *DeleteDestinationResponse {
	return v.value
}

func (v *NullableDeleteDestinationResponse) Set(val *DeleteDestinationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteDestinationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteDestinationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteDestinationResponse(val *DeleteDestinationResponse) *NullableDeleteDestinationResponse {
	return &NullableDeleteDestinationResponse{value: val, isSet: true}
}

func (v NullableDeleteDestinationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeleteDestinationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
