package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the GetRatesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRatesResponse{}

// GetRatesResponse The response schema for the getRates operation.
type GetRatesResponse struct {
	Payload *GetRatesResult `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetRatesResponse instantiates a new GetRatesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRatesResponse() *GetRatesResponse {
	this := GetRatesResponse{}
	return &this
}

// NewGetRatesResponseWithDefaults instantiates a new GetRatesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRatesResponseWithDefaults() *GetRatesResponse {
	this := GetRatesResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetRatesResponse) GetPayload() GetRatesResult {
	if o == nil || IsNil(o.Payload) {
		var ret GetRatesResult
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRatesResponse) GetPayloadOk() (*GetRatesResult, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetRatesResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given GetRatesResult and assigns it to the Payload field.
func (o *GetRatesResponse) SetPayload(v GetRatesResult) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetRatesResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRatesResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetRatesResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetRatesResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetRatesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetRatesResponse struct {
	value *GetRatesResponse
	isSet bool
}

func (v NullableGetRatesResponse) Get() *GetRatesResponse {
	return v.value
}

func (v *NullableGetRatesResponse) Set(val *GetRatesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRatesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRatesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRatesResponse(val *GetRatesResponse) *NullableGetRatesResponse {
	return &NullableGetRatesResponse{value: val, isSet: true}
}

func (v NullableGetRatesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetRatesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
