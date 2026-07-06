package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemPointsCost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemPointsCost{}

// ItemPointsCost Information about Amazon Points granted with the purchase of an item, including both quantity and monetary equivalent value.
type ItemPointsCost struct {
	PointsGranted *PointsGranted `json:"pointsGranted,omitempty"`
}

// NewItemPointsCost instantiates a new ItemPointsCost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemPointsCost() *ItemPointsCost {
	this := ItemPointsCost{}
	return &this
}

// NewItemPointsCostWithDefaults instantiates a new ItemPointsCost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemPointsCostWithDefaults() *ItemPointsCost {
	this := ItemPointsCost{}
	return &this
}

// GetPointsGranted returns the PointsGranted field value if set, zero value otherwise.
func (o *ItemPointsCost) GetPointsGranted() PointsGranted {
	if o == nil || IsNil(o.PointsGranted) {
		var ret PointsGranted
		return ret
	}
	return *o.PointsGranted
}

// GetPointsGrantedOk returns a tuple with the PointsGranted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemPointsCost) GetPointsGrantedOk() (*PointsGranted, bool) {
	if o == nil || IsNil(o.PointsGranted) {
		return nil, false
	}
	return o.PointsGranted, true
}

// HasPointsGranted returns a boolean if a field has been set.
func (o *ItemPointsCost) HasPointsGranted() bool {
	if o != nil && !IsNil(o.PointsGranted) {
		return true
	}

	return false
}

// SetPointsGranted gets a reference to the given PointsGranted and assigns it to the PointsGranted field.
func (o *ItemPointsCost) SetPointsGranted(v PointsGranted) {
	o.PointsGranted = &v
}

func (o ItemPointsCost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PointsGranted) {
		toSerialize["pointsGranted"] = o.PointsGranted
	}
	return toSerialize, nil
}

type NullableItemPointsCost struct {
	value *ItemPointsCost
	isSet bool
}

func (v NullableItemPointsCost) Get() *ItemPointsCost {
	return v.value
}

func (v *NullableItemPointsCost) Set(val *ItemPointsCost) {
	v.value = val
	v.isSet = true
}

func (v NullableItemPointsCost) IsSet() bool {
	return v.isSet
}

func (v *NullableItemPointsCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemPointsCost(val *ItemPointsCost) *NullableItemPointsCost {
	return &NullableItemPointsCost{value: val, isSet: true}
}

func (v NullableItemPointsCost) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemPointsCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
