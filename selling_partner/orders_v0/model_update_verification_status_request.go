package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateVerificationStatusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVerificationStatusRequest{}

// UpdateVerificationStatusRequest The request body for the `updateVerificationStatus` operation.
type UpdateVerificationStatusRequest struct {
	RegulatedOrderVerificationStatus UpdateVerificationStatusRequestBody `json:"regulatedOrderVerificationStatus"`
}

type _UpdateVerificationStatusRequest UpdateVerificationStatusRequest

// NewUpdateVerificationStatusRequest instantiates a new UpdateVerificationStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVerificationStatusRequest(regulatedOrderVerificationStatus UpdateVerificationStatusRequestBody) *UpdateVerificationStatusRequest {
	this := UpdateVerificationStatusRequest{}
	this.RegulatedOrderVerificationStatus = regulatedOrderVerificationStatus
	return &this
}

// NewUpdateVerificationStatusRequestWithDefaults instantiates a new UpdateVerificationStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVerificationStatusRequestWithDefaults() *UpdateVerificationStatusRequest {
	this := UpdateVerificationStatusRequest{}
	return &this
}

// GetRegulatedOrderVerificationStatus returns the RegulatedOrderVerificationStatus field value
func (o *UpdateVerificationStatusRequest) GetRegulatedOrderVerificationStatus() UpdateVerificationStatusRequestBody {
	if o == nil {
		var ret UpdateVerificationStatusRequestBody
		return ret
	}

	return o.RegulatedOrderVerificationStatus
}

// GetRegulatedOrderVerificationStatusOk returns a tuple with the RegulatedOrderVerificationStatus field value
// and a boolean to check if the value has been set.
func (o *UpdateVerificationStatusRequest) GetRegulatedOrderVerificationStatusOk() (*UpdateVerificationStatusRequestBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegulatedOrderVerificationStatus, true
}

// SetRegulatedOrderVerificationStatus sets field value
func (o *UpdateVerificationStatusRequest) SetRegulatedOrderVerificationStatus(v UpdateVerificationStatusRequestBody) {
	o.RegulatedOrderVerificationStatus = v
}

func (o UpdateVerificationStatusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["regulatedOrderVerificationStatus"] = o.RegulatedOrderVerificationStatus
	return toSerialize, nil
}

type NullableUpdateVerificationStatusRequest struct {
	value *UpdateVerificationStatusRequest
	isSet bool
}

func (v NullableUpdateVerificationStatusRequest) Get() *UpdateVerificationStatusRequest {
	return v.value
}

func (v *NullableUpdateVerificationStatusRequest) Set(val *UpdateVerificationStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVerificationStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVerificationStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVerificationStatusRequest(val *UpdateVerificationStatusRequest) *NullableUpdateVerificationStatusRequest {
	return &NullableUpdateVerificationStatusRequest{value: val, isSet: true}
}

func (v NullableUpdateVerificationStatusRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateVerificationStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
