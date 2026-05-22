package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ScheduleSelfShipAppointmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduleSelfShipAppointmentRequest{}

// ScheduleSelfShipAppointmentRequest The `scheduleSelfShipAppointment` request.
type ScheduleSelfShipAppointmentRequest struct {
	ReasonComment *ReasonComment `json:"reasonComment,omitempty"`
}

// NewScheduleSelfShipAppointmentRequest instantiates a new ScheduleSelfShipAppointmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleSelfShipAppointmentRequest() *ScheduleSelfShipAppointmentRequest {
	this := ScheduleSelfShipAppointmentRequest{}
	return &this
}

// NewScheduleSelfShipAppointmentRequestWithDefaults instantiates a new ScheduleSelfShipAppointmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleSelfShipAppointmentRequestWithDefaults() *ScheduleSelfShipAppointmentRequest {
	this := ScheduleSelfShipAppointmentRequest{}
	return &this
}

// GetReasonComment returns the ReasonComment field value if set, zero value otherwise.
func (o *ScheduleSelfShipAppointmentRequest) GetReasonComment() ReasonComment {
	if o == nil || IsNil(o.ReasonComment) {
		var ret ReasonComment
		return ret
	}
	return *o.ReasonComment
}

// GetReasonCommentOk returns a tuple with the ReasonComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleSelfShipAppointmentRequest) GetReasonCommentOk() (*ReasonComment, bool) {
	if o == nil || IsNil(o.ReasonComment) {
		return nil, false
	}
	return o.ReasonComment, true
}

// HasReasonComment returns a boolean if a field has been set.
func (o *ScheduleSelfShipAppointmentRequest) HasReasonComment() bool {
	if o != nil && !IsNil(o.ReasonComment) {
		return true
	}

	return false
}

// SetReasonComment gets a reference to the given ReasonComment and assigns it to the ReasonComment field.
func (o *ScheduleSelfShipAppointmentRequest) SetReasonComment(v ReasonComment) {
	o.ReasonComment = &v
}

func (o ScheduleSelfShipAppointmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReasonComment) {
		toSerialize["reasonComment"] = o.ReasonComment
	}
	return toSerialize, nil
}

type NullableScheduleSelfShipAppointmentRequest struct {
	value *ScheduleSelfShipAppointmentRequest
	isSet bool
}

func (v NullableScheduleSelfShipAppointmentRequest) Get() *ScheduleSelfShipAppointmentRequest {
	return v.value
}

func (v *NullableScheduleSelfShipAppointmentRequest) Set(val *ScheduleSelfShipAppointmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleSelfShipAppointmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleSelfShipAppointmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleSelfShipAppointmentRequest(val *ScheduleSelfShipAppointmentRequest) *NullableScheduleSelfShipAppointmentRequest {
	return &NullableScheduleSelfShipAppointmentRequest{value: val, isSet: true}
}

func (v NullableScheduleSelfShipAppointmentRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableScheduleSelfShipAppointmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
