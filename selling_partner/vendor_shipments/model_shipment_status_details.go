package vendor_shipments

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the ShipmentStatusDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentStatusDetails{}

// ShipmentStatusDetails Shipment Status details.
type ShipmentStatusDetails struct {
	// Current status of the shipment on whether it is picked up or scheduled.
	ShipmentStatus *string `json:"shipmentStatus,omitempty"`
	// Date and time on last status update received for the shipment
	ShipmentStatusDate *time.Time `json:"shipmentStatusDate,omitempty"`
}

// NewShipmentStatusDetails instantiates a new ShipmentStatusDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentStatusDetails() *ShipmentStatusDetails {
	this := ShipmentStatusDetails{}
	return &this
}

// NewShipmentStatusDetailsWithDefaults instantiates a new ShipmentStatusDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentStatusDetailsWithDefaults() *ShipmentStatusDetails {
	this := ShipmentStatusDetails{}
	return &this
}

// GetShipmentStatus returns the ShipmentStatus field value if set, zero value otherwise.
func (o *ShipmentStatusDetails) GetShipmentStatus() string {
	if o == nil || IsNil(o.ShipmentStatus) {
		var ret string
		return ret
	}
	return *o.ShipmentStatus
}

// GetShipmentStatusOk returns a tuple with the ShipmentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentStatusDetails) GetShipmentStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ShipmentStatus) {
		return nil, false
	}
	return o.ShipmentStatus, true
}

// HasShipmentStatus returns a boolean if a field has been set.
func (o *ShipmentStatusDetails) HasShipmentStatus() bool {
	if o != nil && !IsNil(o.ShipmentStatus) {
		return true
	}

	return false
}

// SetShipmentStatus gets a reference to the given string and assigns it to the ShipmentStatus field.
func (o *ShipmentStatusDetails) SetShipmentStatus(v string) {
	o.ShipmentStatus = &v
}

// GetShipmentStatusDate returns the ShipmentStatusDate field value if set, zero value otherwise.
func (o *ShipmentStatusDetails) GetShipmentStatusDate() time.Time {
	if o == nil || IsNil(o.ShipmentStatusDate) {
		var ret time.Time
		return ret
	}
	return *o.ShipmentStatusDate
}

// GetShipmentStatusDateOk returns a tuple with the ShipmentStatusDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentStatusDetails) GetShipmentStatusDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ShipmentStatusDate) {
		return nil, false
	}
	return o.ShipmentStatusDate, true
}

// HasShipmentStatusDate returns a boolean if a field has been set.
func (o *ShipmentStatusDetails) HasShipmentStatusDate() bool {
	if o != nil && !IsNil(o.ShipmentStatusDate) {
		return true
	}

	return false
}

// SetShipmentStatusDate gets a reference to the given time.Time and assigns it to the ShipmentStatusDate field.
func (o *ShipmentStatusDetails) SetShipmentStatusDate(v time.Time) {
	o.ShipmentStatusDate = &v
}

func (o ShipmentStatusDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShipmentStatus) {
		toSerialize["shipmentStatus"] = o.ShipmentStatus
	}
	if !IsNil(o.ShipmentStatusDate) {
		toSerialize["shipmentStatusDate"] = o.ShipmentStatusDate
	}
	return toSerialize, nil
}

type NullableShipmentStatusDetails struct {
	value *ShipmentStatusDetails
	isSet bool
}

func (v NullableShipmentStatusDetails) Get() *ShipmentStatusDetails {
	return v.value
}

func (v *NullableShipmentStatusDetails) Set(val *ShipmentStatusDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentStatusDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentStatusDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentStatusDetails(val *ShipmentStatusDetails) *NullableShipmentStatusDetails {
	return &NullableShipmentStatusDetails{value: val, isSet: true}
}

func (v NullableShipmentStatusDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShipmentStatusDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
