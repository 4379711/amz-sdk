package fba_inventory

import (
	"github.com/bytedance/sonic"
)

// checks if the AddInventoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddInventoryResponse{}

// AddInventoryResponse The response schema for the AddInventory operation.
type AddInventoryResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewAddInventoryResponse instantiates a new AddInventoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddInventoryResponse() *AddInventoryResponse {
	this := AddInventoryResponse{}
	return &this
}

// NewAddInventoryResponseWithDefaults instantiates a new AddInventoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddInventoryResponseWithDefaults() *AddInventoryResponse {
	this := AddInventoryResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AddInventoryResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddInventoryResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AddInventoryResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *AddInventoryResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o AddInventoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableAddInventoryResponse struct {
	value *AddInventoryResponse
	isSet bool
}

func (v NullableAddInventoryResponse) Get() *AddInventoryResponse {
	return v.value
}

func (v *NullableAddInventoryResponse) Set(val *AddInventoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddInventoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddInventoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddInventoryResponse(val *AddInventoryResponse) *NullableAddInventoryResponse {
	return &NullableAddInventoryResponse{value: val, isSet: true}
}

func (v NullableAddInventoryResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAddInventoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
