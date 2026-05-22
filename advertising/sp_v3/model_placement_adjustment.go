package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PlacementAdjustment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlacementAdjustment{}

// PlacementAdjustment Specifies bid adjustments based on the placement location. Use `PLACEMENT_TOP` for the top of the search page. Use `PLACEMENT_REST_OF_SEARCH` for the rest of the search page. Use `PLACEMENT_PRODUCT_PAGE` for a product page.
type PlacementAdjustment struct {
	Predicate  *string `json:"predicate,omitempty"`
	Percentage *int32  `json:"percentage,omitempty"`
}

// NewPlacementAdjustment instantiates a new PlacementAdjustment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlacementAdjustment() *PlacementAdjustment {
	this := PlacementAdjustment{}
	return &this
}

// NewPlacementAdjustmentWithDefaults instantiates a new PlacementAdjustment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlacementAdjustmentWithDefaults() *PlacementAdjustment {
	this := PlacementAdjustment{}
	return &this
}

// GetPredicate returns the Predicate field value if set, zero value otherwise.
func (o *PlacementAdjustment) GetPredicate() string {
	if o == nil || IsNil(o.Predicate) {
		var ret string
		return ret
	}
	return *o.Predicate
}

// GetPredicateOk returns a tuple with the Predicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlacementAdjustment) GetPredicateOk() (*string, bool) {
	if o == nil || IsNil(o.Predicate) {
		return nil, false
	}
	return o.Predicate, true
}

// HasPredicate returns a boolean if a field has been set.
func (o *PlacementAdjustment) HasPredicate() bool {
	if o != nil && !IsNil(o.Predicate) {
		return true
	}

	return false
}

// SetPredicate gets a reference to the given string and assigns it to the Predicate field.
func (o *PlacementAdjustment) SetPredicate(v string) {
	o.Predicate = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *PlacementAdjustment) GetPercentage() int32 {
	if o == nil || IsNil(o.Percentage) {
		var ret int32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlacementAdjustment) GetPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *PlacementAdjustment) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given int32 and assigns it to the Percentage field.
func (o *PlacementAdjustment) SetPercentage(v int32) {
	o.Percentage = &v
}

func (o PlacementAdjustment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Predicate) {
		toSerialize["predicate"] = o.Predicate
	}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return toSerialize, nil
}

type NullablePlacementAdjustment struct {
	value *PlacementAdjustment
	isSet bool
}

func (v NullablePlacementAdjustment) Get() *PlacementAdjustment {
	return v.value
}

func (v *NullablePlacementAdjustment) Set(val *PlacementAdjustment) {
	v.value = val
	v.isSet = true
}

func (v NullablePlacementAdjustment) IsSet() bool {
	return v.isSet
}

func (v *NullablePlacementAdjustment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlacementAdjustment(val *PlacementAdjustment) *NullablePlacementAdjustment {
	return &NullablePlacementAdjustment{value: val, isSet: true}
}

func (v NullablePlacementAdjustment) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePlacementAdjustment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
