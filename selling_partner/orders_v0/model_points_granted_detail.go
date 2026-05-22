package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the PointsGrantedDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PointsGrantedDetail{}

// PointsGrantedDetail The number of Amazon Points offered with the purchase of an item, and their monetary value.
type PointsGrantedDetail struct {
	// The number of Amazon Points granted with the purchase of an item.
	PointsNumber        *int32 `json:"PointsNumber,omitempty"`
	PointsMonetaryValue *Money `json:"PointsMonetaryValue,omitempty"`
}

// NewPointsGrantedDetail instantiates a new PointsGrantedDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPointsGrantedDetail() *PointsGrantedDetail {
	this := PointsGrantedDetail{}
	return &this
}

// NewPointsGrantedDetailWithDefaults instantiates a new PointsGrantedDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointsGrantedDetailWithDefaults() *PointsGrantedDetail {
	this := PointsGrantedDetail{}
	return &this
}

// GetPointsNumber returns the PointsNumber field value if set, zero value otherwise.
func (o *PointsGrantedDetail) GetPointsNumber() int32 {
	if o == nil || IsNil(o.PointsNumber) {
		var ret int32
		return ret
	}
	return *o.PointsNumber
}

// GetPointsNumberOk returns a tuple with the PointsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointsGrantedDetail) GetPointsNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PointsNumber) {
		return nil, false
	}
	return o.PointsNumber, true
}

// HasPointsNumber returns a boolean if a field has been set.
func (o *PointsGrantedDetail) HasPointsNumber() bool {
	if o != nil && !IsNil(o.PointsNumber) {
		return true
	}

	return false
}

// SetPointsNumber gets a reference to the given int32 and assigns it to the PointsNumber field.
func (o *PointsGrantedDetail) SetPointsNumber(v int32) {
	o.PointsNumber = &v
}

// GetPointsMonetaryValue returns the PointsMonetaryValue field value if set, zero value otherwise.
func (o *PointsGrantedDetail) GetPointsMonetaryValue() Money {
	if o == nil || IsNil(o.PointsMonetaryValue) {
		var ret Money
		return ret
	}
	return *o.PointsMonetaryValue
}

// GetPointsMonetaryValueOk returns a tuple with the PointsMonetaryValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointsGrantedDetail) GetPointsMonetaryValueOk() (*Money, bool) {
	if o == nil || IsNil(o.PointsMonetaryValue) {
		return nil, false
	}
	return o.PointsMonetaryValue, true
}

// HasPointsMonetaryValue returns a boolean if a field has been set.
func (o *PointsGrantedDetail) HasPointsMonetaryValue() bool {
	if o != nil && !IsNil(o.PointsMonetaryValue) {
		return true
	}

	return false
}

// SetPointsMonetaryValue gets a reference to the given Money and assigns it to the PointsMonetaryValue field.
func (o *PointsGrantedDetail) SetPointsMonetaryValue(v Money) {
	o.PointsMonetaryValue = &v
}

func (o PointsGrantedDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PointsNumber) {
		toSerialize["PointsNumber"] = o.PointsNumber
	}
	if !IsNil(o.PointsMonetaryValue) {
		toSerialize["PointsMonetaryValue"] = o.PointsMonetaryValue
	}
	return toSerialize, nil
}

type NullablePointsGrantedDetail struct {
	value *PointsGrantedDetail
	isSet bool
}

func (v NullablePointsGrantedDetail) Get() *PointsGrantedDetail {
	return v.value
}

func (v *NullablePointsGrantedDetail) Set(val *PointsGrantedDetail) {
	v.value = val
	v.isSet = true
}

func (v NullablePointsGrantedDetail) IsSet() bool {
	return v.isSet
}

func (v *NullablePointsGrantedDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePointsGrantedDetail(val *PointsGrantedDetail) *NullablePointsGrantedDetail {
	return &NullablePointsGrantedDetail{value: val, isSet: true}
}

func (v NullablePointsGrantedDetail) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePointsGrantedDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
