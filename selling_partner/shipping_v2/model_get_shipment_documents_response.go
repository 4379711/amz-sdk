package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the GetShipmentDocumentsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetShipmentDocumentsResponse{}

// GetShipmentDocumentsResponse The response schema for the the getShipmentDocuments operation.
type GetShipmentDocumentsResponse struct {
	Payload *GetShipmentDocumentsResult `json:"payload,omitempty"`
}

// NewGetShipmentDocumentsResponse instantiates a new GetShipmentDocumentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetShipmentDocumentsResponse() *GetShipmentDocumentsResponse {
	this := GetShipmentDocumentsResponse{}
	return &this
}

// NewGetShipmentDocumentsResponseWithDefaults instantiates a new GetShipmentDocumentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetShipmentDocumentsResponseWithDefaults() *GetShipmentDocumentsResponse {
	this := GetShipmentDocumentsResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetShipmentDocumentsResponse) GetPayload() GetShipmentDocumentsResult {
	if o == nil || IsNil(o.Payload) {
		var ret GetShipmentDocumentsResult
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentDocumentsResponse) GetPayloadOk() (*GetShipmentDocumentsResult, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetShipmentDocumentsResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given GetShipmentDocumentsResult and assigns it to the Payload field.
func (o *GetShipmentDocumentsResponse) SetPayload(v GetShipmentDocumentsResult) {
	o.Payload = &v
}

func (o GetShipmentDocumentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	return toSerialize, nil
}

type NullableGetShipmentDocumentsResponse struct {
	value *GetShipmentDocumentsResponse
	isSet bool
}

func (v NullableGetShipmentDocumentsResponse) Get() *GetShipmentDocumentsResponse {
	return v.value
}

func (v *NullableGetShipmentDocumentsResponse) Set(val *GetShipmentDocumentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetShipmentDocumentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetShipmentDocumentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetShipmentDocumentsResponse(val *GetShipmentDocumentsResponse) *NullableGetShipmentDocumentsResponse {
	return &NullableGetShipmentDocumentsResponse{value: val, isSet: true}
}

func (v NullableGetShipmentDocumentsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetShipmentDocumentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
