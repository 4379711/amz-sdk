package fulfillment_outbound_20200701

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the DeliveryWindow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryWindow{}

// DeliveryWindow The time range within which a Scheduled Delivery fulfillment order should be delivered. This is only available in the JP marketplace.
type DeliveryWindow struct {
	// Date timestamp
	StartDate time.Time `json:"startDate"`
	// Date timestamp
	EndDate time.Time `json:"endDate"`
}

type _DeliveryWindow DeliveryWindow

// NewDeliveryWindow instantiates a new DeliveryWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryWindow(startDate time.Time, endDate time.Time) *DeliveryWindow {
	this := DeliveryWindow{}
	this.StartDate = startDate
	this.EndDate = endDate
	return &this
}

// NewDeliveryWindowWithDefaults instantiates a new DeliveryWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryWindowWithDefaults() *DeliveryWindow {
	this := DeliveryWindow{}
	return &this
}

// GetStartDate returns the StartDate field value
func (o *DeliveryWindow) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *DeliveryWindow) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *DeliveryWindow) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *DeliveryWindow) GetEndDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *DeliveryWindow) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *DeliveryWindow) SetEndDate(v time.Time) {
	o.EndDate = v
}

func (o DeliveryWindow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startDate"] = o.StartDate
	toSerialize["endDate"] = o.EndDate
	return toSerialize, nil
}

type NullableDeliveryWindow struct {
	value *DeliveryWindow
	isSet bool
}

func (v NullableDeliveryWindow) Get() *DeliveryWindow {
	return v.value
}

func (v *NullableDeliveryWindow) Set(val *DeliveryWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryWindow(val *DeliveryWindow) *NullableDeliveryWindow {
	return &NullableDeliveryWindow{value: val, isSet: true}
}

func (v NullableDeliveryWindow) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeliveryWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
