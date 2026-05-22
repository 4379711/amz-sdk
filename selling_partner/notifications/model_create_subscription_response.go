package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSubscriptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSubscriptionResponse{}

// CreateSubscriptionResponse The response schema for the `createSubscription` operation.
type CreateSubscriptionResponse struct {
	Payload *Subscription `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCreateSubscriptionResponse instantiates a new CreateSubscriptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSubscriptionResponse() *CreateSubscriptionResponse {
	this := CreateSubscriptionResponse{}
	return &this
}

// NewCreateSubscriptionResponseWithDefaults instantiates a new CreateSubscriptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSubscriptionResponseWithDefaults() *CreateSubscriptionResponse {
	this := CreateSubscriptionResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *CreateSubscriptionResponse) GetPayload() Subscription {
	if o == nil || IsNil(o.Payload) {
		var ret Subscription
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionResponse) GetPayloadOk() (*Subscription, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *CreateSubscriptionResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given Subscription and assigns it to the Payload field.
func (o *CreateSubscriptionResponse) SetPayload(v Subscription) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateSubscriptionResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateSubscriptionResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CreateSubscriptionResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CreateSubscriptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCreateSubscriptionResponse struct {
	value *CreateSubscriptionResponse
	isSet bool
}

func (v NullableCreateSubscriptionResponse) Get() *CreateSubscriptionResponse {
	return v.value
}

func (v *NullableCreateSubscriptionResponse) Set(val *CreateSubscriptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSubscriptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSubscriptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSubscriptionResponse(val *CreateSubscriptionResponse) *NullableCreateSubscriptionResponse {
	return &NullableCreateSubscriptionResponse{value: val, isSet: true}
}

func (v NullableCreateSubscriptionResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSubscriptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
