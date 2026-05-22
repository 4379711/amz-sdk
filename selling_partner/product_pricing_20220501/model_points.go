package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the Points type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Points{}

// Points The number of Amazon Points that are offered with the purchase of an item and the monetary value of these points.
type Points struct {
	// The number of Amazon Points.
	PointsNumber        *int32     `json:"pointsNumber,omitempty"`
	PointsMonetaryValue *MoneyType `json:"pointsMonetaryValue,omitempty"`
}

// NewPoints instantiates a new Points object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoints() *Points {
	this := Points{}
	return &this
}

// NewPointsWithDefaults instantiates a new Points object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointsWithDefaults() *Points {
	this := Points{}
	return &this
}

// GetPointsNumber returns the PointsNumber field value if set, zero value otherwise.
func (o *Points) GetPointsNumber() int32 {
	if o == nil || IsNil(o.PointsNumber) {
		var ret int32
		return ret
	}
	return *o.PointsNumber
}

// GetPointsNumberOk returns a tuple with the PointsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Points) GetPointsNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PointsNumber) {
		return nil, false
	}
	return o.PointsNumber, true
}

// HasPointsNumber returns a boolean if a field has been set.
func (o *Points) HasPointsNumber() bool {
	if o != nil && !IsNil(o.PointsNumber) {
		return true
	}

	return false
}

// SetPointsNumber gets a reference to the given int32 and assigns it to the PointsNumber field.
func (o *Points) SetPointsNumber(v int32) {
	o.PointsNumber = &v
}

// GetPointsMonetaryValue returns the PointsMonetaryValue field value if set, zero value otherwise.
func (o *Points) GetPointsMonetaryValue() MoneyType {
	if o == nil || IsNil(o.PointsMonetaryValue) {
		var ret MoneyType
		return ret
	}
	return *o.PointsMonetaryValue
}

// GetPointsMonetaryValueOk returns a tuple with the PointsMonetaryValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Points) GetPointsMonetaryValueOk() (*MoneyType, bool) {
	if o == nil || IsNil(o.PointsMonetaryValue) {
		return nil, false
	}
	return o.PointsMonetaryValue, true
}

// HasPointsMonetaryValue returns a boolean if a field has been set.
func (o *Points) HasPointsMonetaryValue() bool {
	if o != nil && !IsNil(o.PointsMonetaryValue) {
		return true
	}

	return false
}

// SetPointsMonetaryValue gets a reference to the given MoneyType and assigns it to the PointsMonetaryValue field.
func (o *Points) SetPointsMonetaryValue(v MoneyType) {
	o.PointsMonetaryValue = &v
}

func (o Points) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PointsNumber) {
		toSerialize["pointsNumber"] = o.PointsNumber
	}
	if !IsNil(o.PointsMonetaryValue) {
		toSerialize["pointsMonetaryValue"] = o.PointsMonetaryValue
	}
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
