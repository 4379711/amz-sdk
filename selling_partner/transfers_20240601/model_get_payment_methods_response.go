package transfers_20240601

import (
	"github.com/bytedance/sonic"
)

// checks if the GetPaymentMethodsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPaymentMethodsResponse{}

// GetPaymentMethodsResponse The response schema for the `getPaymentMethods` operation.
type GetPaymentMethodsResponse struct {
	// The list of payment methods with payment method details.
	PaymentMethods []PaymentMethodDetails `json:"paymentMethods,omitempty"`
}

// NewGetPaymentMethodsResponse instantiates a new GetPaymentMethodsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPaymentMethodsResponse() *GetPaymentMethodsResponse {
	this := GetPaymentMethodsResponse{}
	return &this
}

// NewGetPaymentMethodsResponseWithDefaults instantiates a new GetPaymentMethodsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPaymentMethodsResponseWithDefaults() *GetPaymentMethodsResponse {
	this := GetPaymentMethodsResponse{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *GetPaymentMethodsResponse) GetPaymentMethods() []PaymentMethodDetails {
	if o == nil || IsNil(o.PaymentMethods) {
		var ret []PaymentMethodDetails
		return ret
	}
	return o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentMethodsResponse) GetPaymentMethodsOk() ([]PaymentMethodDetails, bool) {
	if o == nil || IsNil(o.PaymentMethods) {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *GetPaymentMethodsResponse) HasPaymentMethods() bool {
	if o != nil && !IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given []PaymentMethodDetails and assigns it to the PaymentMethods field.
func (o *GetPaymentMethodsResponse) SetPaymentMethods(v []PaymentMethodDetails) {
	o.PaymentMethods = v
}

func (o GetPaymentMethodsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethods) {
		toSerialize["paymentMethods"] = o.PaymentMethods
	}
	return toSerialize, nil
}

type NullableGetPaymentMethodsResponse struct {
	value *GetPaymentMethodsResponse
	isSet bool
}

func (v NullableGetPaymentMethodsResponse) Get() *GetPaymentMethodsResponse {
	return v.value
}

func (v *NullableGetPaymentMethodsResponse) Set(val *GetPaymentMethodsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPaymentMethodsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPaymentMethodsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPaymentMethodsResponse(val *GetPaymentMethodsResponse) *NullableGetPaymentMethodsResponse {
	return &NullableGetPaymentMethodsResponse{value: val, isSet: true}
}

func (v NullableGetPaymentMethodsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetPaymentMethodsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
