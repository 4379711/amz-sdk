package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BidPlacementAdjustment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BidPlacementAdjustment{}

// BidPlacementAdjustment Specifies bid adjustments based on the placement location. Use `PLACEMENT_TOP` for the top of the search page. Use `PLACEMENT_REST_OF_SEARCH` for the rest of the search page. Use `PLACEMENT_PRODUCT_PAGE` for a product page.
type BidPlacementAdjustment struct {
	Predicate  *string `json:"predicate,omitempty"`
	Percentage *int32  `json:"percentage,omitempty"`
}

// NewBidPlacementAdjustment instantiates a new BidPlacementAdjustment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBidPlacementAdjustment() *BidPlacementAdjustment {
	this := BidPlacementAdjustment{}
	return &this
}

// NewBidPlacementAdjustmentWithDefaults instantiates a new BidPlacementAdjustment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBidPlacementAdjustmentWithDefaults() *BidPlacementAdjustment {
	this := BidPlacementAdjustment{}
	return &this
}

// GetPredicate returns the Predicate field value if set, zero value otherwise.
func (o *BidPlacementAdjustment) GetPredicate() string {
	if o == nil || IsNil(o.Predicate) {
		var ret string
		return ret
	}
	return *o.Predicate
}

// GetPredicateOk returns a tuple with the Predicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidPlacementAdjustment) GetPredicateOk() (*string, bool) {
	if o == nil || IsNil(o.Predicate) {
		return nil, false
	}
	return o.Predicate, true
}

// HasPredicate returns a boolean if a field has been set.
func (o *BidPlacementAdjustment) HasPredicate() bool {
	if o != nil && !IsNil(o.Predicate) {
		return true
	}

	return false
}

// SetPredicate gets a reference to the given string and assigns it to the Predicate field.
func (o *BidPlacementAdjustment) SetPredicate(v string) {
	o.Predicate = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *BidPlacementAdjustment) GetPercentage() int32 {
	if o == nil || IsNil(o.Percentage) {
		var ret int32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidPlacementAdjustment) GetPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *BidPlacementAdjustment) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given int32 and assigns it to the Percentage field.
func (o *BidPlacementAdjustment) SetPercentage(v int32) {
	o.Percentage = &v
}

func (o BidPlacementAdjustment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Predicate) {
		toSerialize["predicate"] = o.Predicate
	}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return toSerialize, nil
}

type NullableBidPlacementAdjustment struct {
	value *BidPlacementAdjustment
	isSet bool
}

func (v NullableBidPlacementAdjustment) Get() *BidPlacementAdjustment {
	return v.value
}

func (v *NullableBidPlacementAdjustment) Set(val *BidPlacementAdjustment) {
	v.value = val
	v.isSet = true
}

func (v NullableBidPlacementAdjustment) IsSet() bool {
	return v.isSet
}

func (v *NullableBidPlacementAdjustment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBidPlacementAdjustment(val *BidPlacementAdjustment) *NullableBidPlacementAdjustment {
	return &NullableBidPlacementAdjustment{value: val, isSet: true}
}

func (v NullableBidPlacementAdjustment) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBidPlacementAdjustment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
