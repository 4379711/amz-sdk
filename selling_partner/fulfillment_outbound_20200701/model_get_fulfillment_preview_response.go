package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the GetFulfillmentPreviewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFulfillmentPreviewResponse{}

// GetFulfillmentPreviewResponse The response schema for the `getFulfillmentPreview` operation.
type GetFulfillmentPreviewResponse struct {
	Payload *GetFulfillmentPreviewResult `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetFulfillmentPreviewResponse instantiates a new GetFulfillmentPreviewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFulfillmentPreviewResponse() *GetFulfillmentPreviewResponse {
	this := GetFulfillmentPreviewResponse{}
	return &this
}

// NewGetFulfillmentPreviewResponseWithDefaults instantiates a new GetFulfillmentPreviewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFulfillmentPreviewResponseWithDefaults() *GetFulfillmentPreviewResponse {
	this := GetFulfillmentPreviewResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetFulfillmentPreviewResponse) GetPayload() GetFulfillmentPreviewResult {
	if o == nil || IsNil(o.Payload) {
		var ret GetFulfillmentPreviewResult
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFulfillmentPreviewResponse) GetPayloadOk() (*GetFulfillmentPreviewResult, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetFulfillmentPreviewResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given GetFulfillmentPreviewResult and assigns it to the Payload field.
func (o *GetFulfillmentPreviewResponse) SetPayload(v GetFulfillmentPreviewResult) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetFulfillmentPreviewResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFulfillmentPreviewResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetFulfillmentPreviewResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetFulfillmentPreviewResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetFulfillmentPreviewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetFulfillmentPreviewResponse struct {
	value *GetFulfillmentPreviewResponse
	isSet bool
}

func (v NullableGetFulfillmentPreviewResponse) Get() *GetFulfillmentPreviewResponse {
	return v.value
}

func (v *NullableGetFulfillmentPreviewResponse) Set(val *GetFulfillmentPreviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFulfillmentPreviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFulfillmentPreviewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFulfillmentPreviewResponse(val *GetFulfillmentPreviewResponse) *NullableGetFulfillmentPreviewResponse {
	return &NullableGetFulfillmentPreviewResponse{value: val, isSet: true}
}

func (v NullableGetFulfillmentPreviewResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetFulfillmentPreviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
