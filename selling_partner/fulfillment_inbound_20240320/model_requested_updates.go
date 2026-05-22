package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the RequestedUpdates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestedUpdates{}

// RequestedUpdates Objects that were included in the update request.
type RequestedUpdates struct {
	// A list of boxes that will be present in the shipment after the update.
	Boxes []BoxUpdateInput `json:"boxes,omitempty"`
	// A list of all items that will be present in the shipment after the update.
	Items []ItemInput `json:"items,omitempty"`
}

// NewRequestedUpdates instantiates a new RequestedUpdates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestedUpdates() *RequestedUpdates {
	this := RequestedUpdates{}
	return &this
}

// NewRequestedUpdatesWithDefaults instantiates a new RequestedUpdates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestedUpdatesWithDefaults() *RequestedUpdates {
	this := RequestedUpdates{}
	return &this
}

// GetBoxes returns the Boxes field value if set, zero value otherwise.
func (o *RequestedUpdates) GetBoxes() []BoxUpdateInput {
	if o == nil || IsNil(o.Boxes) {
		var ret []BoxUpdateInput
		return ret
	}
	return o.Boxes
}

// GetBoxesOk returns a tuple with the Boxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedUpdates) GetBoxesOk() ([]BoxUpdateInput, bool) {
	if o == nil || IsNil(o.Boxes) {
		return nil, false
	}
	return o.Boxes, true
}

// HasBoxes returns a boolean if a field has been set.
func (o *RequestedUpdates) HasBoxes() bool {
	if o != nil && !IsNil(o.Boxes) {
		return true
	}

	return false
}

// SetBoxes gets a reference to the given []BoxUpdateInput and assigns it to the Boxes field.
func (o *RequestedUpdates) SetBoxes(v []BoxUpdateInput) {
	o.Boxes = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *RequestedUpdates) GetItems() []ItemInput {
	if o == nil || IsNil(o.Items) {
		var ret []ItemInput
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedUpdates) GetItemsOk() ([]ItemInput, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *RequestedUpdates) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ItemInput and assigns it to the Items field.
func (o *RequestedUpdates) SetItems(v []ItemInput) {
	o.Items = v
}

func (o RequestedUpdates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Boxes) {
		toSerialize["boxes"] = o.Boxes
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableRequestedUpdates struct {
	value *RequestedUpdates
	isSet bool
}

func (v NullableRequestedUpdates) Get() *RequestedUpdates {
	return v.value
}

func (v *NullableRequestedUpdates) Set(val *RequestedUpdates) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestedUpdates) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestedUpdates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestedUpdates(val *RequestedUpdates) *NullableRequestedUpdates {
	return &NullableRequestedUpdates{value: val, isSet: true}
}

func (v NullableRequestedUpdates) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRequestedUpdates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
