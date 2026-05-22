package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the Dates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dates{}

// Dates Specifies the date that the seller expects their shipment will be shipped.
type Dates struct {
	ReadyToShipWindow *Window `json:"readyToShipWindow,omitempty"`
}

// NewDates instantiates a new Dates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDates() *Dates {
	this := Dates{}
	return &this
}

// NewDatesWithDefaults instantiates a new Dates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatesWithDefaults() *Dates {
	this := Dates{}
	return &this
}

// GetReadyToShipWindow returns the ReadyToShipWindow field value if set, zero value otherwise.
func (o *Dates) GetReadyToShipWindow() Window {
	if o == nil || IsNil(o.ReadyToShipWindow) {
		var ret Window
		return ret
	}
	return *o.ReadyToShipWindow
}

// GetReadyToShipWindowOk returns a tuple with the ReadyToShipWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dates) GetReadyToShipWindowOk() (*Window, bool) {
	if o == nil || IsNil(o.ReadyToShipWindow) {
		return nil, false
	}
	return o.ReadyToShipWindow, true
}

// HasReadyToShipWindow returns a boolean if a field has been set.
func (o *Dates) HasReadyToShipWindow() bool {
	if o != nil && !IsNil(o.ReadyToShipWindow) {
		return true
	}

	return false
}

// SetReadyToShipWindow gets a reference to the given Window and assigns it to the ReadyToShipWindow field.
func (o *Dates) SetReadyToShipWindow(v Window) {
	o.ReadyToShipWindow = &v
}

func (o Dates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReadyToShipWindow) {
		toSerialize["readyToShipWindow"] = o.ReadyToShipWindow
	}
	return toSerialize, nil
}

type NullableDates struct {
	value *Dates
	isSet bool
}

func (v NullableDates) Get() *Dates {
	return v.value
}

func (v *NullableDates) Set(val *Dates) {
	v.value = val
	v.isSet = true
}

func (v NullableDates) IsSet() bool {
	return v.isSet
}

func (v *NullableDates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDates(val *Dates) *NullableDates {
	return &NullableDates{value: val, isSet: true}
}

func (v NullableDates) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
