package fulfillment_inbound_20240320

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the GenerateSelfShipAppointmentSlotsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateSelfShipAppointmentSlotsRequest{}

// GenerateSelfShipAppointmentSlotsRequest The `generateSelfShipAppointmentSlots` request.
type GenerateSelfShipAppointmentSlotsRequest struct {
	// The desired end date. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime format.
	DesiredEndDate *time.Time `json:"desiredEndDate,omitempty"`
	// The desired start date. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime format.
	DesiredStartDate *time.Time `json:"desiredStartDate,omitempty"`
}

// NewGenerateSelfShipAppointmentSlotsRequest instantiates a new GenerateSelfShipAppointmentSlotsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateSelfShipAppointmentSlotsRequest() *GenerateSelfShipAppointmentSlotsRequest {
	this := GenerateSelfShipAppointmentSlotsRequest{}
	return &this
}

// NewGenerateSelfShipAppointmentSlotsRequestWithDefaults instantiates a new GenerateSelfShipAppointmentSlotsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateSelfShipAppointmentSlotsRequestWithDefaults() *GenerateSelfShipAppointmentSlotsRequest {
	this := GenerateSelfShipAppointmentSlotsRequest{}
	return &this
}

// GetDesiredEndDate returns the DesiredEndDate field value if set, zero value otherwise.
func (o *GenerateSelfShipAppointmentSlotsRequest) GetDesiredEndDate() time.Time {
	if o == nil || IsNil(o.DesiredEndDate) {
		var ret time.Time
		return ret
	}
	return *o.DesiredEndDate
}

// GetDesiredEndDateOk returns a tuple with the DesiredEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateSelfShipAppointmentSlotsRequest) GetDesiredEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DesiredEndDate) {
		return nil, false
	}
	return o.DesiredEndDate, true
}

// HasDesiredEndDate returns a boolean if a field has been set.
func (o *GenerateSelfShipAppointmentSlotsRequest) HasDesiredEndDate() bool {
	if o != nil && !IsNil(o.DesiredEndDate) {
		return true
	}

	return false
}

// SetDesiredEndDate gets a reference to the given time.Time and assigns it to the DesiredEndDate field.
func (o *GenerateSelfShipAppointmentSlotsRequest) SetDesiredEndDate(v time.Time) {
	o.DesiredEndDate = &v
}

// GetDesiredStartDate returns the DesiredStartDate field value if set, zero value otherwise.
func (o *GenerateSelfShipAppointmentSlotsRequest) GetDesiredStartDate() time.Time {
	if o == nil || IsNil(o.DesiredStartDate) {
		var ret time.Time
		return ret
	}
	return *o.DesiredStartDate
}

// GetDesiredStartDateOk returns a tuple with the DesiredStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateSelfShipAppointmentSlotsRequest) GetDesiredStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DesiredStartDate) {
		return nil, false
	}
	return o.DesiredStartDate, true
}

// HasDesiredStartDate returns a boolean if a field has been set.
func (o *GenerateSelfShipAppointmentSlotsRequest) HasDesiredStartDate() bool {
	if o != nil && !IsNil(o.DesiredStartDate) {
		return true
	}

	return false
}

// SetDesiredStartDate gets a reference to the given time.Time and assigns it to the DesiredStartDate field.
func (o *GenerateSelfShipAppointmentSlotsRequest) SetDesiredStartDate(v time.Time) {
	o.DesiredStartDate = &v
}

func (o GenerateSelfShipAppointmentSlotsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DesiredEndDate) {
		toSerialize["desiredEndDate"] = o.DesiredEndDate
	}
	if !IsNil(o.DesiredStartDate) {
		toSerialize["desiredStartDate"] = o.DesiredStartDate
	}
	return toSerialize, nil
}

type NullableGenerateSelfShipAppointmentSlotsRequest struct {
	value *GenerateSelfShipAppointmentSlotsRequest
	isSet bool
}

func (v NullableGenerateSelfShipAppointmentSlotsRequest) Get() *GenerateSelfShipAppointmentSlotsRequest {
	return v.value
}

func (v *NullableGenerateSelfShipAppointmentSlotsRequest) Set(val *GenerateSelfShipAppointmentSlotsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateSelfShipAppointmentSlotsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateSelfShipAppointmentSlotsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateSelfShipAppointmentSlotsRequest(val *GenerateSelfShipAppointmentSlotsRequest) *NullableGenerateSelfShipAppointmentSlotsRequest {
	return &NullableGenerateSelfShipAppointmentSlotsRequest{value: val, isSet: true}
}

func (v NullableGenerateSelfShipAppointmentSlotsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGenerateSelfShipAppointmentSlotsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
