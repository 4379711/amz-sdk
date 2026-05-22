package vendor_direct_fulfillment_shipping_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the GetCustomerInvoiceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCustomerInvoiceResponse{}

// GetCustomerInvoiceResponse The response schema for the getCustomerInvoice operation.
type GetCustomerInvoiceResponse struct {
	Payload *CustomerInvoice `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetCustomerInvoiceResponse instantiates a new GetCustomerInvoiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCustomerInvoiceResponse() *GetCustomerInvoiceResponse {
	this := GetCustomerInvoiceResponse{}
	return &this
}

// NewGetCustomerInvoiceResponseWithDefaults instantiates a new GetCustomerInvoiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCustomerInvoiceResponseWithDefaults() *GetCustomerInvoiceResponse {
	this := GetCustomerInvoiceResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetCustomerInvoiceResponse) GetPayload() CustomerInvoice {
	if o == nil || IsNil(o.Payload) {
		var ret CustomerInvoice
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCustomerInvoiceResponse) GetPayloadOk() (*CustomerInvoice, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetCustomerInvoiceResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given CustomerInvoice and assigns it to the Payload field.
func (o *GetCustomerInvoiceResponse) SetPayload(v CustomerInvoice) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetCustomerInvoiceResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCustomerInvoiceResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetCustomerInvoiceResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetCustomerInvoiceResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetCustomerInvoiceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetCustomerInvoiceResponse struct {
	value *GetCustomerInvoiceResponse
	isSet bool
}

func (v NullableGetCustomerInvoiceResponse) Get() *GetCustomerInvoiceResponse {
	return v.value
}

func (v *NullableGetCustomerInvoiceResponse) Set(val *GetCustomerInvoiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCustomerInvoiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCustomerInvoiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCustomerInvoiceResponse(val *GetCustomerInvoiceResponse) *NullableGetCustomerInvoiceResponse {
	return &NullableGetCustomerInvoiceResponse{value: val, isSet: true}
}

func (v NullableGetCustomerInvoiceResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetCustomerInvoiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
