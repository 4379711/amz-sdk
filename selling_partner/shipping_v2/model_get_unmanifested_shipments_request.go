package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the GetUnmanifestedShipmentsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUnmanifestedShipmentsRequest{}

// GetUnmanifestedShipmentsRequest The request schema for the GetUnmanifestedShipmentsRequest operation.
type GetUnmanifestedShipmentsRequest struct {
	// Object to pass additional information about the MCI Integrator shipperType: List of ClientReferenceDetail
	ClientReferenceDetails []ClientReferenceDetail `json:"clientReferenceDetails,omitempty"`
}

// NewGetUnmanifestedShipmentsRequest instantiates a new GetUnmanifestedShipmentsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUnmanifestedShipmentsRequest() *GetUnmanifestedShipmentsRequest {
	this := GetUnmanifestedShipmentsRequest{}
	return &this
}

// NewGetUnmanifestedShipmentsRequestWithDefaults instantiates a new GetUnmanifestedShipmentsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUnmanifestedShipmentsRequestWithDefaults() *GetUnmanifestedShipmentsRequest {
	this := GetUnmanifestedShipmentsRequest{}
	return &this
}

// GetClientReferenceDetails returns the ClientReferenceDetails field value if set, zero value otherwise.
func (o *GetUnmanifestedShipmentsRequest) GetClientReferenceDetails() []ClientReferenceDetail {
	if o == nil || IsNil(o.ClientReferenceDetails) {
		var ret []ClientReferenceDetail
		return ret
	}
	return o.ClientReferenceDetails
}

// GetClientReferenceDetailsOk returns a tuple with the ClientReferenceDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUnmanifestedShipmentsRequest) GetClientReferenceDetailsOk() ([]ClientReferenceDetail, bool) {
	if o == nil || IsNil(o.ClientReferenceDetails) {
		return nil, false
	}
	return o.ClientReferenceDetails, true
}

// HasClientReferenceDetails returns a boolean if a field has been set.
func (o *GetUnmanifestedShipmentsRequest) HasClientReferenceDetails() bool {
	if o != nil && !IsNil(o.ClientReferenceDetails) {
		return true
	}

	return false
}

// SetClientReferenceDetails gets a reference to the given []ClientReferenceDetail and assigns it to the ClientReferenceDetails field.
func (o *GetUnmanifestedShipmentsRequest) SetClientReferenceDetails(v []ClientReferenceDetail) {
	o.ClientReferenceDetails = v
}

func (o GetUnmanifestedShipmentsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientReferenceDetails) {
		toSerialize["clientReferenceDetails"] = o.ClientReferenceDetails
	}
	return toSerialize, nil
}

type NullableGetUnmanifestedShipmentsRequest struct {
	value *GetUnmanifestedShipmentsRequest
	isSet bool
}

func (v NullableGetUnmanifestedShipmentsRequest) Get() *GetUnmanifestedShipmentsRequest {
	return v.value
}

func (v *NullableGetUnmanifestedShipmentsRequest) Set(val *GetUnmanifestedShipmentsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUnmanifestedShipmentsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUnmanifestedShipmentsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUnmanifestedShipmentsRequest(val *GetUnmanifestedShipmentsRequest) *NullableGetUnmanifestedShipmentsRequest {
	return &NullableGetUnmanifestedShipmentsRequest{value: val, isSet: true}
}

func (v NullableGetUnmanifestedShipmentsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetUnmanifestedShipmentsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
