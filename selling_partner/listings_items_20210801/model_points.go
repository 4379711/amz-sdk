package listings_items_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the Points type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Points{}

// Points The number of Amazon Points offered with the purchase of an item, and their monetary value. Note that the `Points` element is only returned in Japan (JP).
type Points struct {
	PointsNumber int32 `json:"pointsNumber"`
}

type _Points Points

// NewPoints instantiates a new Points object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoints(pointsNumber int32) *Points {
	this := Points{}
	this.PointsNumber = pointsNumber
	return &this
}

// NewPointsWithDefaults instantiates a new Points object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointsWithDefaults() *Points {
	this := Points{}
	return &this
}

// GetPointsNumber returns the PointsNumber field value
func (o *Points) GetPointsNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PointsNumber
}

// GetPointsNumberOk returns a tuple with the PointsNumber field value
// and a boolean to check if the value has been set.
func (o *Points) GetPointsNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PointsNumber, true
}

// SetPointsNumber sets field value
func (o *Points) SetPointsNumber(v int32) {
	o.PointsNumber = v
}

func (o Points) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pointsNumber"] = o.PointsNumber
	return toSerialize, nil
}

type NullablePoints struct {
	value *Points
	isSet bool
}

func (v NullablePoints) Get() *Points {
	return v.value
}

func (v *NullablePoints) Set(val *Points) {
	v.value = val
	v.isSet = true
}

func (v NullablePoints) IsSet() bool {
	return v.isSet
}

func (v *NullablePoints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoints(val *Points) *NullablePoints {
	return &NullablePoints{value: val, isSet: true}
}

func (v NullablePoints) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePoints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
