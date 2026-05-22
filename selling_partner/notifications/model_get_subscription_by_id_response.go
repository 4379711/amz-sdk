package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the GetSubscriptionByIdResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSubscriptionByIdResponse{}

// GetSubscriptionByIdResponse The response schema for the `getSubscriptionById` operation.
type GetSubscriptionByIdResponse struct {
	Payload *Subscription `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetSubscriptionByIdResponse instantiates a new GetSubscriptionByIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSubscriptionByIdResponse() *GetSubscriptionByIdResponse {
	this := GetSubscriptionByIdResponse{}
	return &this
}

// NewGetSubscriptionByIdResponseWithDefaults instantiates a new GetSubscriptionByIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSubscriptionByIdResponseWithDefaults() *GetSubscriptionByIdResponse {
	this := GetSubscriptionByIdResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetSubscriptionByIdResponse) GetPayload() Subscription {
	if o == nil || IsNil(o.Payload) {
		var ret Subscription
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubscriptionByIdResponse) GetPayloadOk() (*Subscription, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetSubscriptionByIdResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given Subscription and assigns it to the Payload field.
func (o *GetSubscriptionByIdResponse) SetPayload(v Subscription) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetSubscriptionByIdResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubscriptionByIdResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetSubscriptionByIdResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetSubscriptionByIdResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetSubscriptionByIdResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetSubscriptionByIdResponse struct {
	value *GetSubscriptionByIdResponse
	isSet bool
}

func (v NullableGetSubscriptionByIdResponse) Get() *GetSubscriptionByIdResponse {
	return v.value
}

func (v *NullableGetSubscriptionByIdResponse) Set(val *GetSubscriptionByIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSubscriptionByIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSubscriptionByIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSubscriptionByIdResponse(val *GetSubscriptionByIdResponse) *NullableGetSubscriptionByIdResponse {
	return &NullableGetSubscriptionByIdResponse{value: val, isSet: true}
}

func (v NullableGetSubscriptionByIdResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetSubscriptionByIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
