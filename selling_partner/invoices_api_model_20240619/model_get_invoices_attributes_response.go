package invoices_api_model_20240619

import (
	"github.com/bytedance/sonic"
)

// checks if the GetInvoicesAttributesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInvoicesAttributesResponse{}

// GetInvoicesAttributesResponse Success.
type GetInvoicesAttributesResponse struct {
	InvoicesAttributes *InvoicesAttributes `json:"invoicesAttributes,omitempty"`
}

// NewGetInvoicesAttributesResponse instantiates a new GetInvoicesAttributesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInvoicesAttributesResponse() *GetInvoicesAttributesResponse {
	this := GetInvoicesAttributesResponse{}
	return &this
}

// NewGetInvoicesAttributesResponseWithDefaults instantiates a new GetInvoicesAttributesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInvoicesAttributesResponseWithDefaults() *GetInvoicesAttributesResponse {
	this := GetInvoicesAttributesResponse{}
	return &this
}

// GetInvoicesAttributes returns the InvoicesAttributes field value if set, zero value otherwise.
func (o *GetInvoicesAttributesResponse) GetInvoicesAttributes() InvoicesAttributes {
	if o == nil || IsNil(o.InvoicesAttributes) {
		var ret InvoicesAttributes
		return ret
	}
	return *o.InvoicesAttributes
}

// GetInvoicesAttributesOk returns a tuple with the InvoicesAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvoicesAttributesResponse) GetInvoicesAttributesOk() (*InvoicesAttributes, bool) {
	if o == nil || IsNil(o.InvoicesAttributes) {
		return nil, false
	}
	return o.InvoicesAttributes, true
}

// HasInvoicesAttributes returns a boolean if a field has been set.
func (o *GetInvoicesAttributesResponse) HasInvoicesAttributes() bool {
	if o != nil && !IsNil(o.InvoicesAttributes) {
		return true
	}

	return false
}

// SetInvoicesAttributes gets a reference to the given InvoicesAttributes and assigns it to the InvoicesAttributes field.
func (o *GetInvoicesAttributesResponse) SetInvoicesAttributes(v InvoicesAttributes) {
	o.InvoicesAttributes = &v
}

func (o GetInvoicesAttributesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoicesAttributes) {
		toSerialize["invoicesAttributes"] = o.InvoicesAttributes
	}
	return toSerialize, nil
}

type NullableGetInvoicesAttributesResponse struct {
	value *GetInvoicesAttributesResponse
	isSet bool
}

func (v NullableGetInvoicesAttributesResponse) Get() *GetInvoicesAttributesResponse {
	return v.value
}

func (v *NullableGetInvoicesAttributesResponse) Set(val *GetInvoicesAttributesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInvoicesAttributesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInvoicesAttributesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInvoicesAttributesResponse(val *GetInvoicesAttributesResponse) *NullableGetInvoicesAttributesResponse {
	return &NullableGetInvoicesAttributesResponse{value: val, isSet: true}
}

func (v NullableGetInvoicesAttributesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetInvoicesAttributesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
