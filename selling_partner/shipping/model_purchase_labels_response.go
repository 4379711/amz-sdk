package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the PurchaseLabelsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurchaseLabelsResponse{}

// PurchaseLabelsResponse The response schema for the purchaseLabels operation.
type PurchaseLabelsResponse struct {
	Payload *PurchaseLabelsResult `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewPurchaseLabelsResponse instantiates a new PurchaseLabelsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseLabelsResponse() *PurchaseLabelsResponse {
	this := PurchaseLabelsResponse{}
	return &this
}

// NewPurchaseLabelsResponseWithDefaults instantiates a new PurchaseLabelsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseLabelsResponseWithDefaults() *PurchaseLabelsResponse {
	this := PurchaseLabelsResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *PurchaseLabelsResponse) GetPayload() PurchaseLabelsResult {
	if o == nil || IsNil(o.Payload) {
		var ret PurchaseLabelsResult
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseLabelsResponse) GetPayloadOk() (*PurchaseLabelsResult, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *PurchaseLabelsResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given PurchaseLabelsResult and assigns it to the Payload field.
func (o *PurchaseLabelsResponse) SetPayload(v PurchaseLabelsResult) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *PurchaseLabelsResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseLabelsResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *PurchaseLabelsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *PurchaseLabelsResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o PurchaseLabelsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullablePurchaseLabelsResponse struct {
	value *PurchaseLabelsResponse
	isSet bool
}

func (v NullablePurchaseLabelsResponse) Get() *PurchaseLabelsResponse {
	return v.value
}

func (v *NullablePurchaseLabelsResponse) Set(val *PurchaseLabelsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseLabelsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseLabelsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseLabelsResponse(val *PurchaseLabelsResponse) *NullablePurchaseLabelsResponse {
	return &NullablePurchaseLabelsResponse{value: val, isSet: true}
}

func (v NullablePurchaseLabelsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePurchaseLabelsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
