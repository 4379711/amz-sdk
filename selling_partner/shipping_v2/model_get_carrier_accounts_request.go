package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the GetCarrierAccountsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCarrierAccountsRequest{}

// GetCarrierAccountsRequest The request schema for the GetCarrierAccounts operation.
type GetCarrierAccountsRequest struct {
	// Object to pass additional information about the MCI Integrator shipperType: List of ClientReferenceDetail
	ClientReferenceDetails []ClientReferenceDetail `json:"clientReferenceDetails,omitempty"`
}

// NewGetCarrierAccountsRequest instantiates a new GetCarrierAccountsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCarrierAccountsRequest() *GetCarrierAccountsRequest {
	this := GetCarrierAccountsRequest{}
	return &this
}

// NewGetCarrierAccountsRequestWithDefaults instantiates a new GetCarrierAccountsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCarrierAccountsRequestWithDefaults() *GetCarrierAccountsRequest {
	this := GetCarrierAccountsRequest{}
	return &this
}

// GetClientReferenceDetails returns the ClientReferenceDetails field value if set, zero value otherwise.
func (o *GetCarrierAccountsRequest) GetClientReferenceDetails() []ClientReferenceDetail {
	if o == nil || IsNil(o.ClientReferenceDetails) {
		var ret []ClientReferenceDetail
		return ret
	}
	return o.ClientReferenceDetails
}

// GetClientReferenceDetailsOk returns a tuple with the ClientReferenceDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCarrierAccountsRequest) GetClientReferenceDetailsOk() ([]ClientReferenceDetail, bool) {
	if o == nil || IsNil(o.ClientReferenceDetails) {
		return nil, false
	}
	return o.ClientReferenceDetails, true
}

// HasClientReferenceDetails returns a boolean if a field has been set.
func (o *GetCarrierAccountsRequest) HasClientReferenceDetails() bool {
	if o != nil && !IsNil(o.ClientReferenceDetails) {
		return true
	}

	return false
}

// SetClientReferenceDetails gets a reference to the given []ClientReferenceDetail and assigns it to the ClientReferenceDetails field.
func (o *GetCarrierAccountsRequest) SetClientReferenceDetails(v []ClientReferenceDetail) {
	o.ClientReferenceDetails = v
}

func (o GetCarrierAccountsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientReferenceDetails) {
		toSerialize["clientReferenceDetails"] = o.ClientReferenceDetails
	}
	return toSerialize, nil
}

type NullableGetCarrierAccountsRequest struct {
	value *GetCarrierAccountsRequest
	isSet bool
}

func (v NullableGetCarrierAccountsRequest) Get() *GetCarrierAccountsRequest {
	return v.value
}

func (v *NullableGetCarrierAccountsRequest) Set(val *GetCarrierAccountsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCarrierAccountsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCarrierAccountsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCarrierAccountsRequest(val *GetCarrierAccountsRequest) *NullableGetCarrierAccountsRequest {
	return &NullableGetCarrierAccountsRequest{value: val, isSet: true}
}

func (v NullableGetCarrierAccountsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetCarrierAccountsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
