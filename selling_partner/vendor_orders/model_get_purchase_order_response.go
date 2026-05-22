package vendor_orders

import (
	"github.com/bytedance/sonic"
)

// checks if the GetPurchaseOrderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPurchaseOrderResponse{}

// GetPurchaseOrderResponse The response schema for the getPurchaseOrder operation.
type GetPurchaseOrderResponse struct {
	Payload *Order `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetPurchaseOrderResponse instantiates a new GetPurchaseOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPurchaseOrderResponse() *GetPurchaseOrderResponse {
	this := GetPurchaseOrderResponse{}
	return &this
}

// NewGetPurchaseOrderResponseWithDefaults instantiates a new GetPurchaseOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPurchaseOrderResponseWithDefaults() *GetPurchaseOrderResponse {
	this := GetPurchaseOrderResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetPurchaseOrderResponse) GetPayload() Order {
	if o == nil || IsNil(o.Payload) {
		var ret Order
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPurchaseOrderResponse) GetPayloadOk() (*Order, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetPurchaseOrderResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given Order and assigns it to the Payload field.
func (o *GetPurchaseOrderResponse) SetPayload(v Order) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetPurchaseOrderResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPurchaseOrderResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetPurchaseOrderResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetPurchaseOrderResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetPurchaseOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetPurchaseOrderResponse struct {
	value *GetPurchaseOrderResponse
	isSet bool
}

func (v NullableGetPurchaseOrderResponse) Get() *GetPurchaseOrderResponse {
	return v.value
}

func (v *NullableGetPurchaseOrderResponse) Set(val *GetPurchaseOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPurchaseOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPurchaseOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPurchaseOrderResponse(val *GetPurchaseOrderResponse) *NullableGetPurchaseOrderResponse {
	return &NullableGetPurchaseOrderResponse{value: val, isSet: true}
}

func (v NullableGetPurchaseOrderResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetPurchaseOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
