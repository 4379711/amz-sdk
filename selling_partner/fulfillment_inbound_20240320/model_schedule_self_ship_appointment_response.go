package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ScheduleSelfShipAppointmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduleSelfShipAppointmentResponse{}

// ScheduleSelfShipAppointmentResponse The `scheduleSelfShipAppointment` response.
type ScheduleSelfShipAppointmentResponse struct {
	SelfShipAppointmentDetails SelfShipAppointmentDetails `json:"selfShipAppointmentDetails"`
}

type _ScheduleSelfShipAppointmentResponse ScheduleSelfShipAppointmentResponse

// NewScheduleSelfShipAppointmentResponse instantiates a new ScheduleSelfShipAppointmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleSelfShipAppointmentResponse(selfShipAppointmentDetails SelfShipAppointmentDetails) *ScheduleSelfShipAppointmentResponse {
	this := ScheduleSelfShipAppointmentResponse{}
	this.SelfShipAppointmentDetails = selfShipAppointmentDetails
	return &this
}

// NewScheduleSelfShipAppointmentResponseWithDefaults instantiates a new ScheduleSelfShipAppointmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleSelfShipAppointmentResponseWithDefaults() *ScheduleSelfShipAppointmentResponse {
	this := ScheduleSelfShipAppointmentResponse{}
	return &this
}

// GetSelfShipAppointmentDetails returns the SelfShipAppointmentDetails field value
func (o *ScheduleSelfShipAppointmentResponse) GetSelfShipAppointmentDetails() SelfShipAppointmentDetails {
	if o == nil {
		var ret SelfShipAppointmentDetails
		return ret
	}

	return o.SelfShipAppointmentDetails
}

// GetSelfShipAppointmentDetailsOk returns a tuple with the SelfShipAppointmentDetails field value
// and a boolean to check if the value has been set.
func (o *ScheduleSelfShipAppointmentResponse) GetSelfShipAppointmentDetailsOk() (*SelfShipAppointmentDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfShipAppointmentDetails, true
}

// SetSelfShipAppointmentDetails sets field value
func (o *ScheduleSelfShipAppointmentResponse) SetSelfShipAppointmentDetails(v SelfShipAppointmentDetails) {
	o.SelfShipAppointmentDetails = v
}

func (o ScheduleSelfShipAppointmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["selfShipAppointmentDetails"] = o.SelfShipAppointmentDetails
	return toSerialize, nil
}

type NullableScheduleSelfShipAppointmentResponse struct {
	value *ScheduleSelfShipAppointmentResponse
	isSet bool
}

func (v NullableScheduleSelfShipAppointmentResponse) Get() *ScheduleSelfShipAppointmentResponse {
	return v.value
}

func (v *NullableScheduleSelfShipAppointmentResponse) Set(val *ScheduleSelfShipAppointmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleSelfShipAppointmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleSelfShipAppointmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleSelfShipAppointmentResponse(val *ScheduleSelfShipAppointmentResponse) *NullableScheduleSelfShipAppointmentResponse {
	return &NullableScheduleSelfShipAppointmentResponse{value: val, isSet: true}
}

func (v NullableScheduleSelfShipAppointmentResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableScheduleSelfShipAppointmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
