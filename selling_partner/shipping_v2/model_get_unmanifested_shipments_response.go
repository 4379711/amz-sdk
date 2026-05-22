package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the GetUnmanifestedShipmentsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUnmanifestedShipmentsResponse{}

// GetUnmanifestedShipmentsResponse The Response  for the GetUnmanifestedShipmentsResponse operation.
type GetUnmanifestedShipmentsResponse struct {
	// A list of UnmanifestedCarrierInformation
	UnmanifestedCarrierInformationList []UnmanifestedCarrierInformation `json:"unmanifestedCarrierInformationList,omitempty"`
}

// NewGetUnmanifestedShipmentsResponse instantiates a new GetUnmanifestedShipmentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUnmanifestedShipmentsResponse() *GetUnmanifestedShipmentsResponse {
	this := GetUnmanifestedShipmentsResponse{}
	return &this
}

// NewGetUnmanifestedShipmentsResponseWithDefaults instantiates a new GetUnmanifestedShipmentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUnmanifestedShipmentsResponseWithDefaults() *GetUnmanifestedShipmentsResponse {
	this := GetUnmanifestedShipmentsResponse{}
	return &this
}

// GetUnmanifestedCarrierInformationList returns the UnmanifestedCarrierInformationList field value if set, zero value otherwise.
func (o *GetUnmanifestedShipmentsResponse) GetUnmanifestedCarrierInformationList() []UnmanifestedCarrierInformation {
	if o == nil || IsNil(o.UnmanifestedCarrierInformationList) {
		var ret []UnmanifestedCarrierInformation
		return ret
	}
	return o.UnmanifestedCarrierInformationList
}

// GetUnmanifestedCarrierInformationListOk returns a tuple with the UnmanifestedCarrierInformationList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUnmanifestedShipmentsResponse) GetUnmanifestedCarrierInformationListOk() ([]UnmanifestedCarrierInformation, bool) {
	if o == nil || IsNil(o.UnmanifestedCarrierInformationList) {
		return nil, false
	}
	return o.UnmanifestedCarrierInformationList, true
}

// HasUnmanifestedCarrierInformationList returns a boolean if a field has been set.
func (o *GetUnmanifestedShipmentsResponse) HasUnmanifestedCarrierInformationList() bool {
	if o != nil && !IsNil(o.UnmanifestedCarrierInformationList) {
		return true
	}

	return false
}

// SetUnmanifestedCarrierInformationList gets a reference to the given []UnmanifestedCarrierInformation and assigns it to the UnmanifestedCarrierInformationList field.
func (o *GetUnmanifestedShipmentsResponse) SetUnmanifestedCarrierInformationList(v []UnmanifestedCarrierInformation) {
	o.UnmanifestedCarrierInformationList = v
}

func (o GetUnmanifestedShipmentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UnmanifestedCarrierInformationList) {
		toSerialize["unmanifestedCarrierInformationList"] = o.UnmanifestedCarrierInformationList
	}
	return toSerialize, nil
}

type NullableGetUnmanifestedShipmentsResponse struct {
	value *GetUnmanifestedShipmentsResponse
	isSet bool
}

func (v NullableGetUnmanifestedShipmentsResponse) Get() *GetUnmanifestedShipmentsResponse {
	return v.value
}

func (v *NullableGetUnmanifestedShipmentsResponse) Set(val *GetUnmanifestedShipmentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUnmanifestedShipmentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUnmanifestedShipmentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUnmanifestedShipmentsResponse(val *GetUnmanifestedShipmentsResponse) *NullableGetUnmanifestedShipmentsResponse {
	return &NullableGetUnmanifestedShipmentsResponse{value: val, isSet: true}
}

func (v NullableGetUnmanifestedShipmentsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetUnmanifestedShipmentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
