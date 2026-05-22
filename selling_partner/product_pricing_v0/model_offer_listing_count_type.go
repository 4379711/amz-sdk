package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the OfferListingCountType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferListingCountType{}

// OfferListingCountType The number of offer listings with the specified condition.
type OfferListingCountType struct {
	// The number of offer listings.
	Count int32 `json:"Count"`
	// The condition of the item.
	Condition string `json:"condition"`
}

type _OfferListingCountType OfferListingCountType

// NewOfferListingCountType instantiates a new OfferListingCountType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferListingCountType(count int32, condition string) *OfferListingCountType {
	this := OfferListingCountType{}
	this.Count = count
	this.Condition = condition
	return &this
}

// NewOfferListingCountTypeWithDefaults instantiates a new OfferListingCountType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferListingCountTypeWithDefaults() *OfferListingCountType {
	this := OfferListingCountType{}
	return &this
}

// GetCount returns the Count field value
func (o *OfferListingCountType) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *OfferListingCountType) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *OfferListingCountType) SetCount(v int32) {
	o.Count = v
}

// GetCondition returns the Condition field value
func (o *OfferListingCountType) GetCondition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *OfferListingCountType) GetConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *OfferListingCountType) SetCondition(v string) {
	o.Condition = v
}

func (o OfferListingCountType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Count"] = o.Count
	toSerialize["condition"] = o.Condition
	return toSerialize, nil
}

type NullableOfferListingCountType struct {
	value *OfferListingCountType
	isSet bool
}

func (v NullableOfferListingCountType) Get() *OfferListingCountType {
	return v.value
}

func (v *NullableOfferListingCountType) Set(val *OfferListingCountType) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferListingCountType) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferListingCountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferListingCountType(val *OfferListingCountType) *NullableOfferListingCountType {
	return &NullableOfferListingCountType{value: val, isSet: true}
}

func (v NullableOfferListingCountType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOfferListingCountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
