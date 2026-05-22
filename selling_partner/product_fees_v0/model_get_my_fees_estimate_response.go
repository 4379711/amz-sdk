package product_fees_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetMyFeesEstimateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMyFeesEstimateResponse{}

// GetMyFeesEstimateResponse struct for GetMyFeesEstimateResponse
type GetMyFeesEstimateResponse struct {
	Payload *GetMyFeesEstimateResult `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetMyFeesEstimateResponse instantiates a new GetMyFeesEstimateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMyFeesEstimateResponse() *GetMyFeesEstimateResponse {
	this := GetMyFeesEstimateResponse{}
	return &this
}

// NewGetMyFeesEstimateResponseWithDefaults instantiates a new GetMyFeesEstimateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMyFeesEstimateResponseWithDefaults() *GetMyFeesEstimateResponse {
	this := GetMyFeesEstimateResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetMyFeesEstimateResponse) GetPayload() GetMyFeesEstimateResult {
	if o == nil || IsNil(o.Payload) {
		var ret GetMyFeesEstimateResult
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMyFeesEstimateResponse) GetPayloadOk() (*GetMyFeesEstimateResult, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetMyFeesEstimateResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given GetMyFeesEstimateResult and assigns it to the Payload field.
func (o *GetMyFeesEstimateResponse) SetPayload(v GetMyFeesEstimateResult) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetMyFeesEstimateResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMyFeesEstimateResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetMyFeesEstimateResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetMyFeesEstimateResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetMyFeesEstimateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetMyFeesEstimateResponse struct {
	value *GetMyFeesEstimateResponse
	isSet bool
}

func (v NullableGetMyFeesEstimateResponse) Get() *GetMyFeesEstimateResponse {
	return v.value
}

func (v *NullableGetMyFeesEstimateResponse) Set(val *GetMyFeesEstimateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMyFeesEstimateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMyFeesEstimateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMyFeesEstimateResponse(val *GetMyFeesEstimateResponse) *NullableGetMyFeesEstimateResponse {
	return &NullableGetMyFeesEstimateResponse{value: val, isSet: true}
}

func (v NullableGetMyFeesEstimateResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetMyFeesEstimateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
