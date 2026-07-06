package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the PointsGranted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PointsGranted{}

// PointsGranted Information about Amazon Points awarded with an item purchase.
type PointsGranted struct {
	// Total number of Amazon Points granted to the customer's account for this item purchase.
	PointsNumber        *int32 `json:"pointsNumber,omitempty"`
	PointsMonetaryValue *Money `json:"pointsMonetaryValue,omitempty"`
}

// NewPointsGranted instantiates a new PointsGranted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPointsGranted() *PointsGranted {
	this := PointsGranted{}
	return &this
}

// NewPointsGrantedWithDefaults instantiates a new PointsGranted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointsGrantedWithDefaults() *PointsGranted {
	this := PointsGranted{}
	return &this
}

// GetPointsNumber returns the PointsNumber field value if set, zero value otherwise.
func (o *PointsGranted) GetPointsNumber() int32 {
	if o == nil || IsNil(o.PointsNumber) {
		var ret int32
		return ret
	}
	return *o.PointsNumber
}

// GetPointsNumberOk returns a tuple with the PointsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointsGranted) GetPointsNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PointsNumber) {
		return nil, false
	}
	return o.PointsNumber, true
}

// HasPointsNumber returns a boolean if a field has been set.
func (o *PointsGranted) HasPointsNumber() bool {
	if o != nil && !IsNil(o.PointsNumber) {
		return true
	}

	return false
}

// SetPointsNumber gets a reference to the given int32 and assigns it to the PointsNumber field.
func (o *PointsGranted) SetPointsNumber(v int32) {
	o.PointsNumber = &v
}

// GetPointsMonetaryValue returns the PointsMonetaryValue field value if set, zero value otherwise.
func (o *PointsGranted) GetPointsMonetaryValue() Money {
	if o == nil || IsNil(o.PointsMonetaryValue) {
		var ret Money
		return ret
	}
	return *o.PointsMonetaryValue
}

// GetPointsMonetaryValueOk returns a tuple with the PointsMonetaryValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointsGranted) GetPointsMonetaryValueOk() (*Money, bool) {
	if o == nil || IsNil(o.PointsMonetaryValue) {
		return nil, false
	}
	return o.PointsMonetaryValue, true
}

// HasPointsMonetaryValue returns a boolean if a field has been set.
func (o *PointsGranted) HasPointsMonetaryValue() bool {
	if o != nil && !IsNil(o.PointsMonetaryValue) {
		return true
	}

	return false
}

// SetPointsMonetaryValue gets a reference to the given Money and assigns it to the PointsMonetaryValue field.
func (o *PointsGranted) SetPointsMonetaryValue(v Money) {
	o.PointsMonetaryValue = &v
}

func (o PointsGranted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PointsNumber) {
		toSerialize["pointsNumber"] = o.PointsNumber
	}
	if !IsNil(o.PointsMonetaryValue) {
		toSerialize["pointsMonetaryValue"] = o.PointsMonetaryValue
	}
	return toSerialize, nil
}

type NullablePointsGranted struct {
	value *PointsGranted
	isSet bool
}

func (v NullablePointsGranted) Get() *PointsGranted {
	return v.value
}

func (v *NullablePointsGranted) Set(val *PointsGranted) {
	v.value = val
	v.isSet = true
}

func (v NullablePointsGranted) IsSet() bool {
	return v.isSet
}

func (v *NullablePointsGranted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePointsGranted(val *PointsGranted) *NullablePointsGranted {
	return &NullablePointsGranted{value: val, isSet: true}
}

func (v NullablePointsGranted) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePointsGranted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
