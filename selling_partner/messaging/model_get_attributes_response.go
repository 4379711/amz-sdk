package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the GetAttributesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAttributesResponse{}

// GetAttributesResponse The response schema for the GetAttributes operation.
type GetAttributesResponse struct {
	Buyer *GetAttributesResponseBuyer `json:"buyer,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetAttributesResponse instantiates a new GetAttributesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAttributesResponse() *GetAttributesResponse {
	this := GetAttributesResponse{}
	return &this
}

// NewGetAttributesResponseWithDefaults instantiates a new GetAttributesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAttributesResponseWithDefaults() *GetAttributesResponse {
	this := GetAttributesResponse{}
	return &this
}

// GetBuyer returns the Buyer field value if set, zero value otherwise.
func (o *GetAttributesResponse) GetBuyer() GetAttributesResponseBuyer {
	if o == nil || IsNil(o.Buyer) {
		var ret GetAttributesResponseBuyer
		return ret
	}
	return *o.Buyer
}

// GetBuyerOk returns a tuple with the Buyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAttributesResponse) GetBuyerOk() (*GetAttributesResponseBuyer, bool) {
	if o == nil || IsNil(o.Buyer) {
		return nil, false
	}
	return o.Buyer, true
}

// HasBuyer returns a boolean if a field has been set.
func (o *GetAttributesResponse) HasBuyer() bool {
	if o != nil && !IsNil(o.Buyer) {
		return true
	}

	return false
}

// SetBuyer gets a reference to the given GetAttributesResponseBuyer and assigns it to the Buyer field.
func (o *GetAttributesResponse) SetBuyer(v GetAttributesResponseBuyer) {
	o.Buyer = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetAttributesResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAttributesResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetAttributesResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetAttributesResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetAttributesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Buyer) {
		toSerialize["buyer"] = o.Buyer
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetAttributesResponse struct {
	value *GetAttributesResponse
	isSet bool
}

func (v NullableGetAttributesResponse) Get() *GetAttributesResponse {
	return v.value
}

func (v *NullableGetAttributesResponse) Set(val *GetAttributesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAttributesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAttributesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAttributesResponse(val *GetAttributesResponse) *NullableGetAttributesResponse {
	return &NullableGetAttributesResponse{value: val, isSet: true}
}

func (v NullableGetAttributesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetAttributesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
