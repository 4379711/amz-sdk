package fba_inventory

import (
	"github.com/bytedance/sonic"
)

// checks if the ResearchingQuantityEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResearchingQuantityEntry{}

// ResearchingQuantityEntry The misplaced or warehouse damaged inventory that is actively being confirmed at our fulfillment centers.
type ResearchingQuantityEntry struct {
	// The duration of the research.
	Name string `json:"name"`
	// The number of units.
	Quantity int32 `json:"quantity"`
}

type _ResearchingQuantityEntry ResearchingQuantityEntry

// NewResearchingQuantityEntry instantiates a new ResearchingQuantityEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResearchingQuantityEntry(name string, quantity int32) *ResearchingQuantityEntry {
	this := ResearchingQuantityEntry{}
	this.Name = name
	this.Quantity = quantity
	return &this
}

// NewResearchingQuantityEntryWithDefaults instantiates a new ResearchingQuantityEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResearchingQuantityEntryWithDefaults() *ResearchingQuantityEntry {
	this := ResearchingQuantityEntry{}
	return &this
}

// GetName returns the Name field value
func (o *ResearchingQuantityEntry) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResearchingQuantityEntry) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResearchingQuantityEntry) SetName(v string) {
	o.Name = v
}

// GetQuantity returns the Quantity field value
func (o *ResearchingQuantityEntry) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *ResearchingQuantityEntry) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *ResearchingQuantityEntry) SetQuantity(v int32) {
	o.Quantity = v
}

func (o ResearchingQuantityEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["quantity"] = o.Quantity
	return toSerialize, nil
}

type NullableResearchingQuantityEntry struct {
	value *ResearchingQuantityEntry
	isSet bool
}

func (v NullableResearchingQuantityEntry) Get() *ResearchingQuantityEntry {
	return v.value
}

func (v *NullableResearchingQuantityEntry) Set(val *ResearchingQuantityEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableResearchingQuantityEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableResearchingQuantityEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResearchingQuantityEntry(val *ResearchingQuantityEntry) *NullableResearchingQuantityEntry {
	return &NullableResearchingQuantityEntry{value: val, isSet: true}
}

func (v NullableResearchingQuantityEntry) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableResearchingQuantityEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
