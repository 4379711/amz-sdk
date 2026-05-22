package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetEligibleShipmentServicesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEligibleShipmentServicesResponse{}

// GetEligibleShipmentServicesResponse Response schema.
type GetEligibleShipmentServicesResponse struct {
	Payload *GetEligibleShipmentServicesResult `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetEligibleShipmentServicesResponse instantiates a new GetEligibleShipmentServicesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEligibleShipmentServicesResponse() *GetEligibleShipmentServicesResponse {
	this := GetEligibleShipmentServicesResponse{}
	return &this
}

// NewGetEligibleShipmentServicesResponseWithDefaults instantiates a new GetEligibleShipmentServicesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEligibleShipmentServicesResponseWithDefaults() *GetEligibleShipmentServicesResponse {
	this := GetEligibleShipmentServicesResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetEligibleShipmentServicesResponse) GetPayload() GetEligibleShipmentServicesResult {
	if o == nil || IsNil(o.Payload) {
		var ret GetEligibleShipmentServicesResult
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEligibleShipmentServicesResponse) GetPayloadOk() (*GetEligibleShipmentServicesResult, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetEligibleShipmentServicesResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given GetEligibleShipmentServicesResult and assigns it to the Payload field.
func (o *GetEligibleShipmentServicesResponse) SetPayload(v GetEligibleShipmentServicesResult) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetEligibleShipmentServicesResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEligibleShipmentServicesResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetEligibleShipmentServicesResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetEligibleShipmentServicesResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetEligibleShipmentServicesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetEligibleShipmentServicesResponse struct {
	value *GetEligibleShipmentServicesResponse
	isSet bool
}

func (v NullableGetEligibleShipmentServicesResponse) Get() *GetEligibleShipmentServicesResponse {
	return v.value
}

func (v *NullableGetEligibleShipmentServicesResponse) Set(val *GetEligibleShipmentServicesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEligibleShipmentServicesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEligibleShipmentServicesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEligibleShipmentServicesResponse(val *GetEligibleShipmentServicesResponse) *NullableGetEligibleShipmentServicesResponse {
	return &NullableGetEligibleShipmentServicesResponse{value: val, isSet: true}
}

func (v NullableGetEligibleShipmentServicesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetEligibleShipmentServicesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
