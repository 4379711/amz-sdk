package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the Item type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Item{}

// Item An item in the Amazon catalog for promotion selection.
type Item struct {
	// The Amazon Standard Identification Number (ASIN) is the unique identifier for an item in the Amazon catalog.
	Asin *string `json:"asin,omitempty"`
	// The Stock Keeping Unit (SKU).
	Sku     *string      `json:"sku,omitempty"`
	Benefit *ItemBenefit `json:"benefit,omitempty"`
	Budget  *Budget      `json:"budget,omitempty"`
}

// NewItem instantiates a new Item object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItem() *Item {
	this := Item{}
	return &this
}

// NewItemWithDefaults instantiates a new Item object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemWithDefaults() *Item {
	this := Item{}
	return &this
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *Item) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *Item) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *Item) SetAsin(v string) {
	o.Asin = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *Item) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *Item) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *Item) SetSku(v string) {
	o.Sku = &v
}

// GetBenefit returns the Benefit field value if set, zero value otherwise.
func (o *Item) GetBenefit() ItemBenefit {
	if o == nil || IsNil(o.Benefit) {
		var ret ItemBenefit
		return ret
	}
	return *o.Benefit
}

// GetBenefitOk returns a tuple with the Benefit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetBenefitOk() (*ItemBenefit, bool) {
	if o == nil || IsNil(o.Benefit) {
		return nil, false
	}
	return o.Benefit, true
}

// HasBenefit returns a boolean if a field has been set.
func (o *Item) HasBenefit() bool {
	if o != nil && !IsNil(o.Benefit) {
		return true
	}

	return false
}

// SetBenefit gets a reference to the given ItemBenefit and assigns it to the Benefit field.
func (o *Item) SetBenefit(v ItemBenefit) {
	o.Benefit = &v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *Item) GetBudget() Budget {
	if o == nil || IsNil(o.Budget) {
		var ret Budget
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetBudgetOk() (*Budget, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *Item) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given Budget and assigns it to the Budget field.
func (o *Item) SetBudget(v Budget) {
	o.Budget = &v
}

func (o Item) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.Benefit) {
		toSerialize["benefit"] = o.Benefit
	}
	if !IsNil(o.Budget) {
		toSerialize["budget"] = o.Budget
	}
	return toSerialize, nil
}

type NullableItem struct {
	value *Item
	isSet bool
}

func (v NullableItem) Get() *Item {
	return v.value
}

func (v *NullableItem) Set(val *Item) {
	v.value = val
	v.isSet = true
}

func (v NullableItem) IsSet() bool {
	return v.isSet
}

func (v *NullableItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItem(val *Item) *NullableItem {
	return &NullableItem{value: val, isSet: true}
}

func (v NullableItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
