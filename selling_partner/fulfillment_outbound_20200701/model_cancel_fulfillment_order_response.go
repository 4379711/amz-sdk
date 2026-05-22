package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the CancelFulfillmentOrderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelFulfillmentOrderResponse{}

// CancelFulfillmentOrderResponse The response schema for the `cancelFulfillmentOrder` operation.
type CancelFulfillmentOrderResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCancelFulfillmentOrderResponse instantiates a new CancelFulfillmentOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelFulfillmentOrderResponse() *CancelFulfillmentOrderResponse {
	this := CancelFulfillmentOrderResponse{}
	return &this
}

// NewCancelFulfillmentOrderResponseWithDefaults instantiates a new CancelFulfillmentOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelFulfillmentOrderResponseWithDefaults() *CancelFulfillmentOrderResponse {
	this := CancelFulfillmentOrderResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CancelFulfillmentOrderResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelFulfillmentOrderResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CancelFulfillmentOrderResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CancelFulfillmentOrderResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CancelFulfillmentOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCancelFulfillmentOrderResponse struct {
	value *CancelFulfillmentOrderResponse
	isSet bool
}

func (v NullableCancelFulfillmentOrderResponse) Get() *CancelFulfillmentOrderResponse {
	return v.value
}

func (v *NullableCancelFulfillmentOrderResponse) Set(val *CancelFulfillmentOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelFulfillmentOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelFulfillmentOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelFulfillmentOrderResponse(val *CancelFulfillmentOrderResponse) *NullableCancelFulfillmentOrderResponse {
	return &NullableCancelFulfillmentOrderResponse{value: val, isSet: true}
}

func (v NullableCancelFulfillmentOrderResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCancelFulfillmentOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
