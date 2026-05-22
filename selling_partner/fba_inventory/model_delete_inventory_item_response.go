package fba_inventory

import (
	"github.com/bytedance/sonic"
)

// checks if the DeleteInventoryItemResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteInventoryItemResponse{}

// DeleteInventoryItemResponse The response schema for the DeleteInventoryItem operation.
type DeleteInventoryItemResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewDeleteInventoryItemResponse instantiates a new DeleteInventoryItemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteInventoryItemResponse() *DeleteInventoryItemResponse {
	this := DeleteInventoryItemResponse{}
	return &this
}

// NewDeleteInventoryItemResponseWithDefaults instantiates a new DeleteInventoryItemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteInventoryItemResponseWithDefaults() *DeleteInventoryItemResponse {
	this := DeleteInventoryItemResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *DeleteInventoryItemResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteInventoryItemResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *DeleteInventoryItemResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *DeleteInventoryItemResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o DeleteInventoryItemResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableDeleteInventoryItemResponse struct {
	value *DeleteInventoryItemResponse
	isSet bool
}

func (v NullableDeleteInventoryItemResponse) Get() *DeleteInventoryItemResponse {
	return v.value
}

func (v *NullableDeleteInventoryItemResponse) Set(val *DeleteInventoryItemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteInventoryItemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteInventoryItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteInventoryItemResponse(val *DeleteInventoryItemResponse) *NullableDeleteInventoryItemResponse {
	return &NullableDeleteInventoryItemResponse{value: val, isSet: true}
}

func (v NullableDeleteInventoryItemResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeleteInventoryItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
