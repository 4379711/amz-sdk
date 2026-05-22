package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the DropOffLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DropOffLocation{}

// DropOffLocation The preferred location to leave packages at the destination address.
type DropOffLocation struct {
	// Specifies the preferred location to leave the package at the destination address.
	Type string `json:"type"`
	// Additional information about the drop-off location that can vary depending on the type of drop-off location specified in the `type` field. If the `type` is set to `FALLBACK_NEIGHBOR_DELIVERY`, the `attributes` object should include the exact keys `neighborName` and `houseNumber` to provide the name and house number of the designated neighbor.
	Attributes *map[string]string `json:"attributes,omitempty"`
}

type _DropOffLocation DropOffLocation

// NewDropOffLocation instantiates a new DropOffLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDropOffLocation(type_ string) *DropOffLocation {
	this := DropOffLocation{}
	this.Type = type_
	return &this
}

// NewDropOffLocationWithDefaults instantiates a new DropOffLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDropOffLocationWithDefaults() *DropOffLocation {
	this := DropOffLocation{}
	return &this
}

// GetType returns the Type field value
func (o *DropOffLocation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DropOffLocation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DropOffLocation) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DropOffLocation) GetAttributes() map[string]string {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropOffLocation) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DropOffLocation) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *DropOffLocation) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

func (o DropOffLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableDropOffLocation struct {
	value *DropOffLocation
	isSet bool
}

func (v NullableDropOffLocation) Get() *DropOffLocation {
	return v.value
}

func (v *NullableDropOffLocation) Set(val *DropOffLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableDropOffLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableDropOffLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDropOffLocation(val *DropOffLocation) *NullableDropOffLocation {
	return &NullableDropOffLocation{value: val, isSet: true}
}

func (v NullableDropOffLocation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDropOffLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
