package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CurvePoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurvePoint{}

// CurvePoint A single point on a curve.
type CurvePoint struct {
	// If this point is the point with the focus circle.
	IsFocus *bool `json:"isFocus,omitempty"`
	// x-axis value.
	X []CurvePointFixedValue `json:"x,omitempty"`
	// y-axis value of multiple KPI types.
	Y []CurvePointRangedValue `json:"y,omitempty"`
}

// NewCurvePoint instantiates a new CurvePoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurvePoint() *CurvePoint {
	this := CurvePoint{}
	return &this
}

// NewCurvePointWithDefaults instantiates a new CurvePoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurvePointWithDefaults() *CurvePoint {
	this := CurvePoint{}
	return &this
}

// GetIsFocus returns the IsFocus field value if set, zero value otherwise.
func (o *CurvePoint) GetIsFocus() bool {
	if o == nil || IsNil(o.IsFocus) {
		var ret bool
		return ret
	}
	return *o.IsFocus
}

// GetIsFocusOk returns a tuple with the IsFocus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurvePoint) GetIsFocusOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFocus) {
		return nil, false
	}
	return o.IsFocus, true
}

// HasIsFocus returns a boolean if a field has been set.
func (o *CurvePoint) HasIsFocus() bool {
	if o != nil && !IsNil(o.IsFocus) {
		return true
	}

	return false
}

// SetIsFocus gets a reference to the given bool and assigns it to the IsFocus field.
func (o *CurvePoint) SetIsFocus(v bool) {
	o.IsFocus = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *CurvePoint) GetX() []CurvePointFixedValue {
	if o == nil || IsNil(o.X) {
		var ret []CurvePointFixedValue
		return ret
	}
	return o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurvePoint) GetXOk() ([]CurvePointFixedValue, bool) {
	if o == nil || IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *CurvePoint) HasX() bool {
	if o != nil && !IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given []CurvePointFixedValue and assigns it to the X field.
func (o *CurvePoint) SetX(v []CurvePointFixedValue) {
	o.X = v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *CurvePoint) GetY() []CurvePointRangedValue {
	if o == nil || IsNil(o.Y) {
		var ret []CurvePointRangedValue
		return ret
	}
	return o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurvePoint) GetYOk() ([]CurvePointRangedValue, bool) {
	if o == nil || IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *CurvePoint) HasY() bool {
	if o != nil && !IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given []CurvePointRangedValue and assigns it to the Y field.
func (o *CurvePoint) SetY(v []CurvePointRangedValue) {
	o.Y = v
}

func (o CurvePoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsFocus) {
		toSerialize["isFocus"] = o.IsFocus
	}
	if !IsNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !IsNil(o.Y) {
		toSerialize["y"] = o.Y
	}
	return toSerialize, nil
}

type NullableCurvePoint struct {
	value *CurvePoint
	isSet bool
}

func (v NullableCurvePoint) Get() *CurvePoint {
	return v.value
}

func (v *NullableCurvePoint) Set(val *CurvePoint) {
	v.value = val
	v.isSet = true
}

func (v NullableCurvePoint) IsSet() bool {
	return v.isSet
}

func (v *NullableCurvePoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurvePoint(val *CurvePoint) *NullableCurvePoint {
	return &NullableCurvePoint{value: val, isSet: true}
}

func (v NullableCurvePoint) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCurvePoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
