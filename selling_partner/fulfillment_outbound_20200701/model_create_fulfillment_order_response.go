package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateFulfillmentOrderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFulfillmentOrderResponse{}

// CreateFulfillmentOrderResponse The response schema for the `createFulfillmentOrder` operation.
type CreateFulfillmentOrderResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCreateFulfillmentOrderResponse instantiates a new CreateFulfillmentOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFulfillmentOrderResponse() *CreateFulfillmentOrderResponse {
	this := CreateFulfillmentOrderResponse{}
	return &this
}

// NewCreateFulfillmentOrderResponseWithDefaults instantiates a new CreateFulfillmentOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFulfillmentOrderResponseWithDefaults() *CreateFulfillmentOrderResponse {
	this := CreateFulfillmentOrderResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateFulfillmentOrderResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFulfillmentOrderResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateFulfillmentOrderResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CreateFulfillmentOrderResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CreateFulfillmentOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCreateFulfillmentOrderResponse struct {
	value *CreateFulfillmentOrderResponse
	isSet bool
}

func (v NullableCreateFulfillmentOrderResponse) Get() *CreateFulfillmentOrderResponse {
	return v.value
}

func (v *NullableCreateFulfillmentOrderResponse) Set(val *CreateFulfillmentOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFulfillmentOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFulfillmentOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFulfillmentOrderResponse(val *CreateFulfillmentOrderResponse) *NullableCreateFulfillmentOrderResponse {
	return &NullableCreateFulfillmentOrderResponse{value: val, isSet: true}
}

func (v NullableCreateFulfillmentOrderResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateFulfillmentOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
