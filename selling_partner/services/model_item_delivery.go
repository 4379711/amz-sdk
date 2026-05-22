package services

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the ItemDelivery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemDelivery{}

// ItemDelivery Delivery information for the item.
type ItemDelivery struct {
	// The date and time of the latest Estimated Delivery Date (EDD) of all the items with an EDD. In ISO 8601 format.
	EstimatedDeliveryDate *time.Time           `json:"estimatedDeliveryDate,omitempty"`
	ItemDeliveryPromise   *ItemDeliveryPromise `json:"itemDeliveryPromise,omitempty"`
}

// NewItemDelivery instantiates a new ItemDelivery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemDelivery() *ItemDelivery {
	this := ItemDelivery{}
	return &this
}

// NewItemDeliveryWithDefaults instantiates a new ItemDelivery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemDeliveryWithDefaults() *ItemDelivery {
	this := ItemDelivery{}
	return &this
}

// GetEstimatedDeliveryDate returns the EstimatedDeliveryDate field value if set, zero value otherwise.
func (o *ItemDelivery) GetEstimatedDeliveryDate() time.Time {
	if o == nil || IsNil(o.EstimatedDeliveryDate) {
		var ret time.Time
		return ret
	}
	return *o.EstimatedDeliveryDate
}

// GetEstimatedDeliveryDateOk returns a tuple with the EstimatedDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDelivery) GetEstimatedDeliveryDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EstimatedDeliveryDate) {
		return nil, false
	}
	return o.EstimatedDeliveryDate, true
}

// HasEstimatedDeliveryDate returns a boolean if a field has been set.
func (o *ItemDelivery) HasEstimatedDeliveryDate() bool {
	if o != nil && !IsNil(o.EstimatedDeliveryDate) {
		return true
	}

	return false
}

// SetEstimatedDeliveryDate gets a reference to the given time.Time and assigns it to the EstimatedDeliveryDate field.
func (o *ItemDelivery) SetEstimatedDeliveryDate(v time.Time) {
	o.EstimatedDeliveryDate = &v
}

// GetItemDeliveryPromise returns the ItemDeliveryPromise field value if set, zero value otherwise.
func (o *ItemDelivery) GetItemDeliveryPromise() ItemDeliveryPromise {
	if o == nil || IsNil(o.ItemDeliveryPromise) {
		var ret ItemDeliveryPromise
		return ret
	}
	return *o.ItemDeliveryPromise
}

// GetItemDeliveryPromiseOk returns a tuple with the ItemDeliveryPromise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDelivery) GetItemDeliveryPromiseOk() (*ItemDeliveryPromise, bool) {
	if o == nil || IsNil(o.ItemDeliveryPromise) {
		return nil, false
	}
	return o.ItemDeliveryPromise, true
}

// HasItemDeliveryPromise returns a boolean if a field has been set.
func (o *ItemDelivery) HasItemDeliveryPromise() bool {
	if o != nil && !IsNil(o.ItemDeliveryPromise) {
		return true
	}

	return false
}

// SetItemDeliveryPromise gets a reference to the given ItemDeliveryPromise and assigns it to the ItemDeliveryPromise field.
func (o *ItemDelivery) SetItemDeliveryPromise(v ItemDeliveryPromise) {
	o.ItemDeliveryPromise = &v
}

func (o ItemDelivery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EstimatedDeliveryDate) {
		toSerialize["estimatedDeliveryDate"] = o.EstimatedDeliveryDate
	}
	if !IsNil(o.ItemDeliveryPromise) {
		toSerialize["itemDeliveryPromise"] = o.ItemDeliveryPromise
	}
	return toSerialize, nil
}

type NullableItemDelivery struct {
	value *ItemDelivery
	isSet bool
}

func (v NullableItemDelivery) Get() *ItemDelivery {
	return v.value
}

func (v *NullableItemDelivery) Set(val *ItemDelivery) {
	v.value = val
	v.isSet = true
}

func (v NullableItemDelivery) IsSet() bool {
	return v.isSet
}

func (v *NullableItemDelivery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemDelivery(val *ItemDelivery) *NullableItemDelivery {
	return &NullableItemDelivery{value: val, isSet: true}
}

func (v NullableItemDelivery) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemDelivery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
