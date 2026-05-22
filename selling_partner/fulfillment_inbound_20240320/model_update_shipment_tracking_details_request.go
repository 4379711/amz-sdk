package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateShipmentTrackingDetailsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateShipmentTrackingDetailsRequest{}

// UpdateShipmentTrackingDetailsRequest The `updateShipmentTrackingDetails` request.
type UpdateShipmentTrackingDetailsRequest struct {
	TrackingDetails TrackingDetailsInput `json:"trackingDetails"`
}

type _UpdateShipmentTrackingDetailsRequest UpdateShipmentTrackingDetailsRequest

// NewUpdateShipmentTrackingDetailsRequest instantiates a new UpdateShipmentTrackingDetailsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateShipmentTrackingDetailsRequest(trackingDetails TrackingDetailsInput) *UpdateShipmentTrackingDetailsRequest {
	this := UpdateShipmentTrackingDetailsRequest{}
	this.TrackingDetails = trackingDetails
	return &this
}

// NewUpdateShipmentTrackingDetailsRequestWithDefaults instantiates a new UpdateShipmentTrackingDetailsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateShipmentTrackingDetailsRequestWithDefaults() *UpdateShipmentTrackingDetailsRequest {
	this := UpdateShipmentTrackingDetailsRequest{}
	return &this
}

// GetTrackingDetails returns the TrackingDetails field value
func (o *UpdateShipmentTrackingDetailsRequest) GetTrackingDetails() TrackingDetailsInput {
	if o == nil {
		var ret TrackingDetailsInput
		return ret
	}

	return o.TrackingDetails
}

// GetTrackingDetailsOk returns a tuple with the TrackingDetails field value
// and a boolean to check if the value has been set.
func (o *UpdateShipmentTrackingDetailsRequest) GetTrackingDetailsOk() (*TrackingDetailsInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrackingDetails, true
}

// SetTrackingDetails sets field value
func (o *UpdateShipmentTrackingDetailsRequest) SetTrackingDetails(v TrackingDetailsInput) {
	o.TrackingDetails = v
}

func (o UpdateShipmentTrackingDetailsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["trackingDetails"] = o.TrackingDetails
	return toSerialize, nil
}

type NullableUpdateShipmentTrackingDetailsRequest struct {
	value *UpdateShipmentTrackingDetailsRequest
	isSet bool
}

func (v NullableUpdateShipmentTrackingDetailsRequest) Get() *UpdateShipmentTrackingDetailsRequest {
	return v.value
}

func (v *NullableUpdateShipmentTrackingDetailsRequest) Set(val *UpdateShipmentTrackingDetailsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateShipmentTrackingDetailsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateShipmentTrackingDetailsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateShipmentTrackingDetailsRequest(val *UpdateShipmentTrackingDetailsRequest) *NullableUpdateShipmentTrackingDetailsRequest {
	return &NullableUpdateShipmentTrackingDetailsRequest{value: val, isSet: true}
}

func (v NullableUpdateShipmentTrackingDetailsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateShipmentTrackingDetailsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
