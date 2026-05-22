package shipment_invoicing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetInvoiceStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInvoiceStatusResponse{}

// GetInvoiceStatusResponse The response schema for the getInvoiceStatus operation.
type GetInvoiceStatusResponse struct {
	Payload *ShipmentInvoiceStatusResponse `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetInvoiceStatusResponse instantiates a new GetInvoiceStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInvoiceStatusResponse() *GetInvoiceStatusResponse {
	this := GetInvoiceStatusResponse{}
	return &this
}

// NewGetInvoiceStatusResponseWithDefaults instantiates a new GetInvoiceStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInvoiceStatusResponseWithDefaults() *GetInvoiceStatusResponse {
	this := GetInvoiceStatusResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetInvoiceStatusResponse) GetPayload() ShipmentInvoiceStatusResponse {
	if o == nil || IsNil(o.Payload) {
		var ret ShipmentInvoiceStatusResponse
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvoiceStatusResponse) GetPayloadOk() (*ShipmentInvoiceStatusResponse, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetInvoiceStatusResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given ShipmentInvoiceStatusResponse and assigns it to the Payload field.
func (o *GetInvoiceStatusResponse) SetPayload(v ShipmentInvoiceStatusResponse) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetInvoiceStatusResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvoiceStatusResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetInvoiceStatusResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetInvoiceStatusResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetInvoiceStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetInvoiceStatusResponse struct {
	value *GetInvoiceStatusResponse
	isSet bool
}

func (v NullableGetInvoiceStatusResponse) Get() *GetInvoiceStatusResponse {
	return v.value
}

func (v *NullableGetInvoiceStatusResponse) Set(val *GetInvoiceStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInvoiceStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInvoiceStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInvoiceStatusResponse(val *GetInvoiceStatusResponse) *NullableGetInvoiceStatusResponse {
	return &NullableGetInvoiceStatusResponse{value: val, isSet: true}
}

func (v NullableGetInvoiceStatusResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetInvoiceStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
