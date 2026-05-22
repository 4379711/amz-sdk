package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the CancelSelfShipAppointmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelSelfShipAppointmentRequest{}

// CancelSelfShipAppointmentRequest The `cancelSelfShipAppointment` request.
type CancelSelfShipAppointmentRequest struct {
	ReasonComment *ReasonComment `json:"reasonComment,omitempty"`
}

// NewCancelSelfShipAppointmentRequest instantiates a new CancelSelfShipAppointmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelSelfShipAppointmentRequest() *CancelSelfShipAppointmentRequest {
	this := CancelSelfShipAppointmentRequest{}
	return &this
}

// NewCancelSelfShipAppointmentRequestWithDefaults instantiates a new CancelSelfShipAppointmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelSelfShipAppointmentRequestWithDefaults() *CancelSelfShipAppointmentRequest {
	this := CancelSelfShipAppointmentRequest{}
	return &this
}

// GetReasonComment returns the ReasonComment field value if set, zero value otherwise.
func (o *CancelSelfShipAppointmentRequest) GetReasonComment() ReasonComment {
	if o == nil || IsNil(o.ReasonComment) {
		var ret ReasonComment
		return ret
	}
	return *o.ReasonComment
}

// GetReasonCommentOk returns a tuple with the ReasonComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelSelfShipAppointmentRequest) GetReasonCommentOk() (*ReasonComment, bool) {
	if o == nil || IsNil(o.ReasonComment) {
		return nil, false
	}
	return o.ReasonComment, true
}

// HasReasonComment returns a boolean if a field has been set.
func (o *CancelSelfShipAppointmentRequest) HasReasonComment() bool {
	if o != nil && !IsNil(o.ReasonComment) {
		return true
	}

	return false
}

// SetReasonComment gets a reference to the given ReasonComment and assigns it to the ReasonComment field.
func (o *CancelSelfShipAppointmentRequest) SetReasonComment(v ReasonComment) {
	o.ReasonComment = &v
}

func (o CancelSelfShipAppointmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReasonComment) {
		toSerialize["reasonComment"] = o.ReasonComment
	}
	return toSerialize, nil
}

type NullableCancelSelfShipAppointmentRequest struct {
	value *CancelSelfShipAppointmentRequest
	isSet bool
}

func (v NullableCancelSelfShipAppointmentRequest) Get() *CancelSelfShipAppointmentRequest {
	return v.value
}

func (v *NullableCancelSelfShipAppointmentRequest) Set(val *CancelSelfShipAppointmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelSelfShipAppointmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelSelfShipAppointmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelSelfShipAppointmentRequest(val *CancelSelfShipAppointmentRequest) *NullableCancelSelfShipAppointmentRequest {
	return &NullableCancelSelfShipAppointmentRequest{value: val, isSet: true}
}

func (v NullableCancelSelfShipAppointmentRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCancelSelfShipAppointmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
