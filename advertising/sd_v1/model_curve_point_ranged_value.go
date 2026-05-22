package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CurvePointRangedValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurvePointRangedValue{}

// CurvePointRangedValue A ranged value.
type CurvePointRangedValue struct {
	// KPI label.
	Label *string              `json:"label,omitempty"`
	Value *ForecastRangeDouble `json:"value,omitempty"`
}

// NewCurvePointRangedValue instantiates a new CurvePointRangedValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurvePointRangedValue() *CurvePointRangedValue {
	this := CurvePointRangedValue{}
	return &this
}

// NewCurvePointRangedValueWithDefaults instantiates a new CurvePointRangedValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurvePointRangedValueWithDefaults() *CurvePointRangedValue {
	this := CurvePointRangedValue{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CurvePointRangedValue) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurvePointRangedValue) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CurvePointRangedValue) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CurvePointRangedValue) SetLabel(v string) {
	o.Label = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CurvePointRangedValue) GetValue() ForecastRangeDouble {
	if o == nil || IsNil(o.Value) {
		var ret ForecastRangeDouble
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurvePointRangedValue) GetValueOk() (*ForecastRangeDouble, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CurvePointRangedValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given ForecastRangeDouble and assigns it to the Value field.
func (o *CurvePointRangedValue) SetValue(v ForecastRangeDouble) {
	o.Value = &v
}

func (o CurvePointRangedValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCurvePointRangedValue struct {
	value *CurvePointRangedValue
	isSet bool
}

func (v NullableCurvePointRangedValue) Get() *CurvePointRangedValue {
	return v.value
}

func (v *NullableCurvePointRangedValue) Set(val *CurvePointRangedValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCurvePointRangedValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCurvePointRangedValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurvePointRangedValue(val *CurvePointRangedValue) *NullableCurvePointRangedValue {
	return &NullableCurvePointRangedValue{value: val, isSet: true}
}

func (v NullableCurvePointRangedValue) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCurvePointRangedValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
