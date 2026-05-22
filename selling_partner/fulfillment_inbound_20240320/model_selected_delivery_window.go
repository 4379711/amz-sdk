package fulfillment_inbound_20240320

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the SelectedDeliveryWindow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectedDeliveryWindow{}

// SelectedDeliveryWindow Selected delivery window attributes.
type SelectedDeliveryWindow struct {
	// Identifies type of Delivery Window Availability. Values: `AVAILABLE`, `CONGESTED`
	AvailabilityType string `json:"availabilityType"`
	// Identifier of a delivery window option. A delivery window option represent one option for when a shipment is expected to be delivered.
	DeliveryWindowOptionId string `json:"deliveryWindowOptionId" validate:"regexp=^[a-zA-Z0-9-]*$"`
	// The timestamp at which this Window can no longer be edited.
	EditableUntil *time.Time `json:"editableUntil,omitempty"`
	// The end timestamp of the window.
	EndDate time.Time `json:"endDate"`
	// The start timestamp of the window.
	StartDate time.Time `json:"startDate"`
}

type _SelectedDeliveryWindow SelectedDeliveryWindow

// NewSelectedDeliveryWindow instantiates a new SelectedDeliveryWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectedDeliveryWindow(availabilityType string, deliveryWindowOptionId string, endDate time.Time, startDate time.Time) *SelectedDeliveryWindow {
	this := SelectedDeliveryWindow{}
	this.AvailabilityType = availabilityType
	this.DeliveryWindowOptionId = deliveryWindowOptionId
	this.EndDate = endDate
	this.StartDate = startDate
	return &this
}

// NewSelectedDeliveryWindowWithDefaults instantiates a new SelectedDeliveryWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectedDeliveryWindowWithDefaults() *SelectedDeliveryWindow {
	this := SelectedDeliveryWindow{}
	return &this
}

// GetAvailabilityType returns the AvailabilityType field value
func (o *SelectedDeliveryWindow) GetAvailabilityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvailabilityType
}

// GetAvailabilityTypeOk returns a tuple with the AvailabilityType field value
// and a boolean to check if the value has been set.
func (o *SelectedDeliveryWindow) GetAvailabilityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailabilityType, true
}

// SetAvailabilityType sets field value
func (o *SelectedDeliveryWindow) SetAvailabilityType(v string) {
	o.AvailabilityType = v
}

// GetDeliveryWindowOptionId returns the DeliveryWindowOptionId field value
func (o *SelectedDeliveryWindow) GetDeliveryWindowOptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliveryWindowOptionId
}

// GetDeliveryWindowOptionIdOk returns a tuple with the DeliveryWindowOptionId field value
// and a boolean to check if the value has been set.
func (o *SelectedDeliveryWindow) GetDeliveryWindowOptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryWindowOptionId, true
}

// SetDeliveryWindowOptionId sets field value
func (o *SelectedDeliveryWindow) SetDeliveryWindowOptionId(v string) {
	o.DeliveryWindowOptionId = v
}

// GetEditableUntil returns the EditableUntil field value if set, zero value otherwise.
func (o *SelectedDeliveryWindow) GetEditableUntil() time.Time {
	if o == nil || IsNil(o.EditableUntil) {
		var ret time.Time
		return ret
	}
	return *o.EditableUntil
}

// GetEditableUntilOk returns a tuple with the EditableUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectedDeliveryWindow) GetEditableUntilOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EditableUntil) {
		return nil, false
	}
	return o.EditableUntil, true
}

// HasEditableUntil returns a boolean if a field has been set.
func (o *SelectedDeliveryWindow) HasEditableUntil() bool {
	if o != nil && !IsNil(o.EditableUntil) {
		return true
	}

	return false
}

// SetEditableUntil gets a reference to the given time.Time and assigns it to the EditableUntil field.
func (o *SelectedDeliveryWindow) SetEditableUntil(v time.Time) {
	o.EditableUntil = &v
}

// GetEndDate returns the EndDate field value
func (o *SelectedDeliveryWindow) GetEndDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *SelectedDeliveryWindow) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *SelectedDeliveryWindow) SetEndDate(v time.Time) {
	o.EndDate = v
}

// GetStartDate returns the StartDate field value
func (o *SelectedDeliveryWindow) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *SelectedDeliveryWindow) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *SelectedDeliveryWindow) SetStartDate(v time.Time) {
	o.StartDate = v
}

func (o SelectedDeliveryWindow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["availabilityType"] = o.AvailabilityType
	toSerialize["deliveryWindowOptionId"] = o.DeliveryWindowOptionId
	if !IsNil(o.EditableUntil) {
		toSerialize["editableUntil"] = o.EditableUntil
	}
	toSerialize["endDate"] = o.EndDate
	toSerialize["startDate"] = o.StartDate
	return toSerialize, nil
}

type NullableSelectedDeliveryWindow struct {
	value *SelectedDeliveryWindow
	isSet bool
}

func (v NullableSelectedDeliveryWindow) Get() *SelectedDeliveryWindow {
	return v.value
}

func (v *NullableSelectedDeliveryWindow) Set(val *SelectedDeliveryWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectedDeliveryWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectedDeliveryWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectedDeliveryWindow(val *SelectedDeliveryWindow) *NullableSelectedDeliveryWindow {
	return &NullableSelectedDeliveryWindow{value: val, isSet: true}
}

func (v NullableSelectedDeliveryWindow) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSelectedDeliveryWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
