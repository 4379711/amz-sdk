package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the GetDeliveryOffersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeliveryOffersResponse{}

// GetDeliveryOffersResponse The response schema for the getDeliveryOffers operation.
type GetDeliveryOffersResponse struct {
	Payload *GetDeliveryOffersResult `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetDeliveryOffersResponse instantiates a new GetDeliveryOffersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeliveryOffersResponse() *GetDeliveryOffersResponse {
	this := GetDeliveryOffersResponse{}
	return &this
}

// NewGetDeliveryOffersResponseWithDefaults instantiates a new GetDeliveryOffersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeliveryOffersResponseWithDefaults() *GetDeliveryOffersResponse {
	this := GetDeliveryOffersResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetDeliveryOffersResponse) GetPayload() GetDeliveryOffersResult {
	if o == nil || IsNil(o.Payload) {
		var ret GetDeliveryOffersResult
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeliveryOffersResponse) GetPayloadOk() (*GetDeliveryOffersResult, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetDeliveryOffersResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given GetDeliveryOffersResult and assigns it to the Payload field.
func (o *GetDeliveryOffersResponse) SetPayload(v GetDeliveryOffersResult) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetDeliveryOffersResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeliveryOffersResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetDeliveryOffersResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetDeliveryOffersResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetDeliveryOffersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetDeliveryOffersResponse struct {
	value *GetDeliveryOffersResponse
	isSet bool
}

func (v NullableGetDeliveryOffersResponse) Get() *GetDeliveryOffersResponse {
	return v.value
}

func (v *NullableGetDeliveryOffersResponse) Set(val *GetDeliveryOffersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeliveryOffersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeliveryOffersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeliveryOffersResponse(val *GetDeliveryOffersResponse) *NullableGetDeliveryOffersResponse {
	return &NullableGetDeliveryOffersResponse{value: val, isSet: true}
}

func (v NullableGetDeliveryOffersResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetDeliveryOffersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
