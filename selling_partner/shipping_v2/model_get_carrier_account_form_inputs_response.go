package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the GetCarrierAccountFormInputsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCarrierAccountFormInputsResponse{}

// GetCarrierAccountFormInputsResponse The Response  for the GetCarrierAccountFormInputsResponse operation.
type GetCarrierAccountFormInputsResponse struct {
	// A list of LinkableCarrier
	LinkableCarriersList []LinkableCarrier `json:"linkableCarriersList,omitempty"`
}

// NewGetCarrierAccountFormInputsResponse instantiates a new GetCarrierAccountFormInputsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCarrierAccountFormInputsResponse() *GetCarrierAccountFormInputsResponse {
	this := GetCarrierAccountFormInputsResponse{}
	return &this
}

// NewGetCarrierAccountFormInputsResponseWithDefaults instantiates a new GetCarrierAccountFormInputsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCarrierAccountFormInputsResponseWithDefaults() *GetCarrierAccountFormInputsResponse {
	this := GetCarrierAccountFormInputsResponse{}
	return &this
}

// GetLinkableCarriersList returns the LinkableCarriersList field value if set, zero value otherwise.
func (o *GetCarrierAccountFormInputsResponse) GetLinkableCarriersList() []LinkableCarrier {
	if o == nil || IsNil(o.LinkableCarriersList) {
		var ret []LinkableCarrier
		return ret
	}
	return o.LinkableCarriersList
}

// GetLinkableCarriersListOk returns a tuple with the LinkableCarriersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCarrierAccountFormInputsResponse) GetLinkableCarriersListOk() ([]LinkableCarrier, bool) {
	if o == nil || IsNil(o.LinkableCarriersList) {
		return nil, false
	}
	return o.LinkableCarriersList, true
}

// HasLinkableCarriersList returns a boolean if a field has been set.
func (o *GetCarrierAccountFormInputsResponse) HasLinkableCarriersList() bool {
	if o != nil && !IsNil(o.LinkableCarriersList) {
		return true
	}

	return false
}

// SetLinkableCarriersList gets a reference to the given []LinkableCarrier and assigns it to the LinkableCarriersList field.
func (o *GetCarrierAccountFormInputsResponse) SetLinkableCarriersList(v []LinkableCarrier) {
	o.LinkableCarriersList = v
}

func (o GetCarrierAccountFormInputsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LinkableCarriersList) {
		toSerialize["linkableCarriersList"] = o.LinkableCarriersList
	}
	return toSerialize, nil
}

type NullableGetCarrierAccountFormInputsResponse struct {
	value *GetCarrierAccountFormInputsResponse
	isSet bool
}

func (v NullableGetCarrierAccountFormInputsResponse) Get() *GetCarrierAccountFormInputsResponse {
	return v.value
}

func (v *NullableGetCarrierAccountFormInputsResponse) Set(val *GetCarrierAccountFormInputsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCarrierAccountFormInputsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCarrierAccountFormInputsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCarrierAccountFormInputsResponse(val *GetCarrierAccountFormInputsResponse) *NullableGetCarrierAccountFormInputsResponse {
	return &NullableGetCarrierAccountFormInputsResponse{value: val, isSet: true}
}

func (v NullableGetCarrierAccountFormInputsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetCarrierAccountFormInputsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
