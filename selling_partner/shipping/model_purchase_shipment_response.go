package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the PurchaseShipmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurchaseShipmentResponse{}

// PurchaseShipmentResponse The response schema for the purchaseShipment operation.
type PurchaseShipmentResponse struct {
	Payload *PurchaseShipmentResult `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewPurchaseShipmentResponse instantiates a new PurchaseShipmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseShipmentResponse() *PurchaseShipmentResponse {
	this := PurchaseShipmentResponse{}
	return &this
}

// NewPurchaseShipmentResponseWithDefaults instantiates a new PurchaseShipmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseShipmentResponseWithDefaults() *PurchaseShipmentResponse {
	this := PurchaseShipmentResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *PurchaseShipmentResponse) GetPayload() PurchaseShipmentResult {
	if o == nil || IsNil(o.Payload) {
		var ret PurchaseShipmentResult
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentResponse) GetPayloadOk() (*PurchaseShipmentResult, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *PurchaseShipmentResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given PurchaseShipmentResult and assigns it to the Payload field.
func (o *PurchaseShipmentResponse) SetPayload(v PurchaseShipmentResult) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *PurchaseShipmentResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *PurchaseShipmentResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *PurchaseShipmentResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o PurchaseShipmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullablePurchaseShipmentResponse struct {
	value *PurchaseShipmentResponse
	isSet bool
}

func (v NullablePurchaseShipmentResponse) Get() *PurchaseShipmentResponse {
	return v.value
}

func (v *NullablePurchaseShipmentResponse) Set(val *PurchaseShipmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseShipmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseShipmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseShipmentResponse(val *PurchaseShipmentResponse) *NullablePurchaseShipmentResponse {
	return &NullablePurchaseShipmentResponse{value: val, isSet: true}
}

func (v NullablePurchaseShipmentResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePurchaseShipmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
