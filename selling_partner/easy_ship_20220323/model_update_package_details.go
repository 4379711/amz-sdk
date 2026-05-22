package easy_ship_20220323

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdatePackageDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePackageDetails{}

// UpdatePackageDetails Request to update the time slot of a package.
type UpdatePackageDetails struct {
	ScheduledPackageId ScheduledPackageId `json:"scheduledPackageId"`
	PackageTimeSlot    TimeSlot           `json:"packageTimeSlot"`
}

type _UpdatePackageDetails UpdatePackageDetails

// NewUpdatePackageDetails instantiates a new UpdatePackageDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePackageDetails(scheduledPackageId ScheduledPackageId, packageTimeSlot TimeSlot) *UpdatePackageDetails {
	this := UpdatePackageDetails{}
	this.ScheduledPackageId = scheduledPackageId
	this.PackageTimeSlot = packageTimeSlot
	return &this
}

// NewUpdatePackageDetailsWithDefaults instantiates a new UpdatePackageDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePackageDetailsWithDefaults() *UpdatePackageDetails {
	this := UpdatePackageDetails{}
	return &this
}

// GetScheduledPackageId returns the ScheduledPackageId field value
func (o *UpdatePackageDetails) GetScheduledPackageId() ScheduledPackageId {
	if o == nil {
		var ret ScheduledPackageId
		return ret
	}

	return o.ScheduledPackageId
}

// GetScheduledPackageIdOk returns a tuple with the ScheduledPackageId field value
// and a boolean to check if the value has been set.
func (o *UpdatePackageDetails) GetScheduledPackageIdOk() (*ScheduledPackageId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledPackageId, true
}

// SetScheduledPackageId sets field value
func (o *UpdatePackageDetails) SetScheduledPackageId(v ScheduledPackageId) {
	o.ScheduledPackageId = v
}

// GetPackageTimeSlot returns the PackageTimeSlot field value
func (o *UpdatePackageDetails) GetPackageTimeSlot() TimeSlot {
	if o == nil {
		var ret TimeSlot
		return ret
	}

	return o.PackageTimeSlot
}

// GetPackageTimeSlotOk returns a tuple with the PackageTimeSlot field value
// and a boolean to check if the value has been set.
func (o *UpdatePackageDetails) GetPackageTimeSlotOk() (*TimeSlot, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageTimeSlot, true
}

// SetPackageTimeSlot sets field value
func (o *UpdatePackageDetails) SetPackageTimeSlot(v TimeSlot) {
	o.PackageTimeSlot = v
}

func (o UpdatePackageDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["scheduledPackageId"] = o.ScheduledPackageId
	toSerialize["packageTimeSlot"] = o.PackageTimeSlot
	return toSerialize, nil
}

type NullableUpdatePackageDetails struct {
	value *UpdatePackageDetails
	isSet bool
}

func (v NullableUpdatePackageDetails) Get() *UpdatePackageDetails {
	return v.value
}

func (v *NullableUpdatePackageDetails) Set(val *UpdatePackageDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePackageDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePackageDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePackageDetails(val *UpdatePackageDetails) *NullableUpdatePackageDetails {
	return &NullableUpdatePackageDetails{value: val, isSet: true}
}

func (v NullableUpdatePackageDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdatePackageDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
