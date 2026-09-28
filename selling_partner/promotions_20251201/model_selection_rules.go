package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the SelectionRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectionRules{}

// SelectionRules Rules for catalog-based selections.
type SelectionRules struct {
	// Items to exclude from the catalog selection (maximum 100 items).
	ExcludedItems []Item `json:"excludedItems,omitempty"`
}

// NewSelectionRules instantiates a new SelectionRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectionRules() *SelectionRules {
	this := SelectionRules{}
	return &this
}

// NewSelectionRulesWithDefaults instantiates a new SelectionRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectionRulesWithDefaults() *SelectionRules {
	this := SelectionRules{}
	return &this
}

// GetExcludedItems returns the ExcludedItems field value if set, zero value otherwise.
func (o *SelectionRules) GetExcludedItems() []Item {
	if o == nil || IsNil(o.ExcludedItems) {
		var ret []Item
		return ret
	}
	return o.ExcludedItems
}

// GetExcludedItemsOk returns a tuple with the ExcludedItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectionRules) GetExcludedItemsOk() ([]Item, bool) {
	if o == nil || IsNil(o.ExcludedItems) {
		return nil, false
	}
	return o.ExcludedItems, true
}

// HasExcludedItems returns a boolean if a field has been set.
func (o *SelectionRules) HasExcludedItems() bool {
	if o != nil && !IsNil(o.ExcludedItems) {
		return true
	}

	return false
}

// SetExcludedItems gets a reference to the given []Item and assigns it to the ExcludedItems field.
func (o *SelectionRules) SetExcludedItems(v []Item) {
	o.ExcludedItems = v
}

func (o SelectionRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExcludedItems) {
		toSerialize["excludedItems"] = o.ExcludedItems
	}
	return toSerialize, nil
}

type NullableSelectionRules struct {
	value *SelectionRules
	isSet bool
}

func (v NullableSelectionRules) Get() *SelectionRules {
	return v.value
}

func (v *NullableSelectionRules) Set(val *SelectionRules) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectionRules) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectionRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectionRules(val *SelectionRules) *NullableSelectionRules {
	return &NullableSelectionRules{value: val, isSet: true}
}

func (v NullableSelectionRules) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSelectionRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
