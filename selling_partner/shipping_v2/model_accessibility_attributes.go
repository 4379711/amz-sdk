package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the AccessibilityAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessibilityAttributes{}

// AccessibilityAttributes Defines the accessibility details of the access point.
type AccessibilityAttributes struct {
	// The approximate distance of access point from input postalCode's centroid.
	Distance *string `json:"distance,omitempty"`
	// The approximate (static) drive time from input postal code's centroid.
	DriveTime *int32 `json:"driveTime,omitempty"`
}

// NewAccessibilityAttributes instantiates a new AccessibilityAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessibilityAttributes() *AccessibilityAttributes {
	this := AccessibilityAttributes{}
	return &this
}

// NewAccessibilityAttributesWithDefaults instantiates a new AccessibilityAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessibilityAttributesWithDefaults() *AccessibilityAttributes {
	this := AccessibilityAttributes{}
	return &this
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *AccessibilityAttributes) GetDistance() string {
	if o == nil || IsNil(o.Distance) {
		var ret string
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessibilityAttributes) GetDistanceOk() (*string, bool) {
	if o == nil || IsNil(o.Distance) {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *AccessibilityAttributes) HasDistance() bool {
	if o != nil && !IsNil(o.Distance) {
		return true
	}

	return false
}

// SetDistance gets a reference to the given string and assigns it to the Distance field.
func (o *AccessibilityAttributes) SetDistance(v string) {
	o.Distance = &v
}

// GetDriveTime returns the DriveTime field value if set, zero value otherwise.
func (o *AccessibilityAttributes) GetDriveTime() int32 {
	if o == nil || IsNil(o.DriveTime) {
		var ret int32
		return ret
	}
	return *o.DriveTime
}

// GetDriveTimeOk returns a tuple with the DriveTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessibilityAttributes) GetDriveTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.DriveTime) {
		return nil, false
	}
	return o.DriveTime, true
}

// HasDriveTime returns a boolean if a field has been set.
func (o *AccessibilityAttributes) HasDriveTime() bool {
	if o != nil && !IsNil(o.DriveTime) {
		return true
	}

	return false
}

// SetDriveTime gets a reference to the given int32 and assigns it to the DriveTime field.
func (o *AccessibilityAttributes) SetDriveTime(v int32) {
	o.DriveTime = &v
}

func (o AccessibilityAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Distance) {
		toSerialize["distance"] = o.Distance
	}
	if !IsNil(o.DriveTime) {
		toSerialize["driveTime"] = o.DriveTime
	}
	return toSerialize, nil
}

type NullableAccessibilityAttributes struct {
	value *AccessibilityAttributes
	isSet bool
}

func (v NullableAccessibilityAttributes) Get() *AccessibilityAttributes {
	return v.value
}

func (v *NullableAccessibilityAttributes) Set(val *AccessibilityAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessibilityAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessibilityAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessibilityAttributes(val *AccessibilityAttributes) *NullableAccessibilityAttributes {
	return &NullableAccessibilityAttributes{value: val, isSet: true}
}

func (v NullableAccessibilityAttributes) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAccessibilityAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
