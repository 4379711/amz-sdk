package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the Adjustment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Adjustment{}

// Adjustment struct for Adjustment
type Adjustment struct {
	PlacementAdjustment *PlacementAdjustment `json:"placementAdjustment,omitempty"`
}

// NewAdjustment instantiates a new Adjustment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdjustment() *Adjustment {
	this := Adjustment{}
	return &this
}

// NewAdjustmentWithDefaults instantiates a new Adjustment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdjustmentWithDefaults() *Adjustment {
	this := Adjustment{}
	return &this
}

// GetPlacementAdjustment returns the PlacementAdjustment field value if set, zero value otherwise.
func (o *Adjustment) GetPlacementAdjustment() PlacementAdjustment {
	if o == nil || IsNil(o.PlacementAdjustment) {
		var ret PlacementAdjustment
		return ret
	}
	return *o.PlacementAdjustment
}

// GetPlacementAdjustmentOk returns a tuple with the PlacementAdjustment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Adjustment) GetPlacementAdjustmentOk() (*PlacementAdjustment, bool) {
	if o == nil || IsNil(o.PlacementAdjustment) {
		return nil, false
	}
	return o.PlacementAdjustment, true
}

// HasPlacementAdjustment returns a boolean if a field has been set.
func (o *Adjustment) HasPlacementAdjustment() bool {
	if o != nil && !IsNil(o.PlacementAdjustment) {
		return true
	}

	return false
}

// SetPlacementAdjustment gets a reference to the given PlacementAdjustment and assigns it to the PlacementAdjustment field.
func (o *Adjustment) SetPlacementAdjustment(v PlacementAdjustment) {
	o.PlacementAdjustment = &v
}

func (o Adjustment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlacementAdjustment) {
		toSerialize["placementAdjustment"] = o.PlacementAdjustment
	}
	return toSerialize, nil
}

type NullableAdjustment struct {
	value *Adjustment
	isSet bool
}

func (v NullableAdjustment) Get() *Adjustment {
	return v.value
}

func (v *NullableAdjustment) Set(val *Adjustment) {
	v.value = val
	v.isSet = true
}

func (v NullableAdjustment) IsSet() bool {
	return v.isSet
}

func (v *NullableAdjustment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdjustment(val *Adjustment) *NullableAdjustment {
	return &NullableAdjustment{value: val, isSet: true}
}

func (v NullableAdjustment) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdjustment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
