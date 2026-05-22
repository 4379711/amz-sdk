package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the GetDestinationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDestinationsResponse{}

// GetDestinationsResponse The response schema for the `getDestinations` operation.
type GetDestinationsResponse struct {
	// A list of destinations.
	Payload []Destination `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetDestinationsResponse instantiates a new GetDestinationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDestinationsResponse() *GetDestinationsResponse {
	this := GetDestinationsResponse{}
	return &this
}

// NewGetDestinationsResponseWithDefaults instantiates a new GetDestinationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDestinationsResponseWithDefaults() *GetDestinationsResponse {
	this := GetDestinationsResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetDestinationsResponse) GetPayload() []Destination {
	if o == nil || IsNil(o.Payload) {
		var ret []Destination
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDestinationsResponse) GetPayloadOk() ([]Destination, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetDestinationsResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given []Destination and assigns it to the Payload field.
func (o *GetDestinationsResponse) SetPayload(v []Destination) {
	o.Payload = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetDestinationsResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDestinationsResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetDestinationsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetDestinationsResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetDestinationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetDestinationsResponse struct {
	value *GetDestinationsResponse
	isSet bool
}

func (v NullableGetDestinationsResponse) Get() *GetDestinationsResponse {
	return v.value
}

func (v *NullableGetDestinationsResponse) Set(val *GetDestinationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDestinationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDestinationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDestinationsResponse(val *GetDestinationsResponse) *NullableGetDestinationsResponse {
	return &NullableGetDestinationsResponse{value: val, isSet: true}
}

func (v NullableGetDestinationsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetDestinationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
