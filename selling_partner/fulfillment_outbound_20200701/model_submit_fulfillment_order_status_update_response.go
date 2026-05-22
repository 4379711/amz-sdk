package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the SubmitFulfillmentOrderStatusUpdateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitFulfillmentOrderStatusUpdateResponse{}

// SubmitFulfillmentOrderStatusUpdateResponse The response schema for the `SubmitFulfillmentOrderStatusUpdate` operation.
type SubmitFulfillmentOrderStatusUpdateResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewSubmitFulfillmentOrderStatusUpdateResponse instantiates a new SubmitFulfillmentOrderStatusUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitFulfillmentOrderStatusUpdateResponse() *SubmitFulfillmentOrderStatusUpdateResponse {
	this := SubmitFulfillmentOrderStatusUpdateResponse{}
	return &this
}

// NewSubmitFulfillmentOrderStatusUpdateResponseWithDefaults instantiates a new SubmitFulfillmentOrderStatusUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitFulfillmentOrderStatusUpdateResponseWithDefaults() *SubmitFulfillmentOrderStatusUpdateResponse {
	this := SubmitFulfillmentOrderStatusUpdateResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SubmitFulfillmentOrderStatusUpdateResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitFulfillmentOrderStatusUpdateResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SubmitFulfillmentOrderStatusUpdateResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *SubmitFulfillmentOrderStatusUpdateResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o SubmitFulfillmentOrderStatusUpdateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSubmitFulfillmentOrderStatusUpdateResponse struct {
	value *SubmitFulfillmentOrderStatusUpdateResponse
	isSet bool
}

func (v NullableSubmitFulfillmentOrderStatusUpdateResponse) Get() *SubmitFulfillmentOrderStatusUpdateResponse {
	return v.value
}

func (v *NullableSubmitFulfillmentOrderStatusUpdateResponse) Set(val *SubmitFulfillmentOrderStatusUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitFulfillmentOrderStatusUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitFulfillmentOrderStatusUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitFulfillmentOrderStatusUpdateResponse(val *SubmitFulfillmentOrderStatusUpdateResponse) *NullableSubmitFulfillmentOrderStatusUpdateResponse {
	return &NullableSubmitFulfillmentOrderStatusUpdateResponse{value: val, isSet: true}
}

func (v NullableSubmitFulfillmentOrderStatusUpdateResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubmitFulfillmentOrderStatusUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
