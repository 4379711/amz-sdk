package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the ScheduledDeliveryInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduledDeliveryInfo{}

// ScheduledDeliveryInfo Delivery information for a scheduled delivery. This is only available in the JP marketplace.
type ScheduledDeliveryInfo struct {
	// The time zone of the destination address for the fulfillment order preview. Must be an IANA time zone name. Example: Asia/Tokyo.
	DeliveryTimeZone string `json:"deliveryTimeZone"`
	// An array of delivery windows.
	DeliveryWindows []DeliveryWindow `json:"deliveryWindows"`
}

type _ScheduledDeliveryInfo ScheduledDeliveryInfo

// NewScheduledDeliveryInfo instantiates a new ScheduledDeliveryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledDeliveryInfo(deliveryTimeZone string, deliveryWindows []DeliveryWindow) *ScheduledDeliveryInfo {
	this := ScheduledDeliveryInfo{}
	this.DeliveryTimeZone = deliveryTimeZone
	this.DeliveryWindows = deliveryWindows
	return &this
}

// NewScheduledDeliveryInfoWithDefaults instantiates a new ScheduledDeliveryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledDeliveryInfoWithDefaults() *ScheduledDeliveryInfo {
	this := ScheduledDeliveryInfo{}
	return &this
}

// GetDeliveryTimeZone returns the DeliveryTimeZone field value
func (o *ScheduledDeliveryInfo) GetDeliveryTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliveryTimeZone
}

// GetDeliveryTimeZoneOk returns a tuple with the DeliveryTimeZone field value
// and a boolean to check if the value has been set.
func (o *ScheduledDeliveryInfo) GetDeliveryTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryTimeZone, true
}

// SetDeliveryTimeZone sets field value
func (o *ScheduledDeliveryInfo) SetDeliveryTimeZone(v string) {
	o.DeliveryTimeZone = v
}

// GetDeliveryWindows returns the DeliveryWindows field value
func (o *ScheduledDeliveryInfo) GetDeliveryWindows() []DeliveryWindow {
	if o == nil {
		var ret []DeliveryWindow
		return ret
	}

	return o.DeliveryWindows
}

// GetDeliveryWindowsOk returns a tuple with the DeliveryWindows field value
// and a boolean to check if the value has been set.
func (o *ScheduledDeliveryInfo) GetDeliveryWindowsOk() ([]DeliveryWindow, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryWindows, true
}

// SetDeliveryWindows sets field value
func (o *ScheduledDeliveryInfo) SetDeliveryWindows(v []DeliveryWindow) {
	o.DeliveryWindows = v
}

func (o ScheduledDeliveryInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deliveryTimeZone"] = o.DeliveryTimeZone
	toSerialize["deliveryWindows"] = o.DeliveryWindows
	return toSerialize, nil
}

type NullableScheduledDeliveryInfo struct {
	value *ScheduledDeliveryInfo
	isSet bool
}

func (v NullableScheduledDeliveryInfo) Get() *ScheduledDeliveryInfo {
	return v.value
}

func (v *NullableScheduledDeliveryInfo) Set(val *ScheduledDeliveryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledDeliveryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledDeliveryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledDeliveryInfo(val *ScheduledDeliveryInfo) *NullableScheduledDeliveryInfo {
	return &NullableScheduledDeliveryInfo{value: val, isSet: true}
}

func (v NullableScheduledDeliveryInfo) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableScheduledDeliveryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
