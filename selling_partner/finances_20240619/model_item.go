package finances_20240619

import (
	"github.com/bytedance/sonic"
)

// checks if the Item type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Item{}

// Item Additional information about the items in Transaction.
type Item struct {
	// Description of items in the transaction
	Description *string `json:"description,omitempty"`
	// Related Business identifiers of the item in Transaction.
	RelatedIdentifiers []ItemRelatedIdentifier `json:"relatedIdentifiers,omitempty"`
	TotalAmount        *Currency               `json:"totalAmount,omitempty"`
	// List of breakdowns which will provide the details on how the total amount is calculated for the financial transaction.
	Breakdowns []Breakdown `json:"breakdowns,omitempty"`
	// List of additional Information about the item.
	Contexts []Context `json:"contexts,omitempty"`
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

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Item) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Item) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Item) SetDescription(v string) {
	o.Description = &v
}

// GetRelatedIdentifiers returns the RelatedIdentifiers field value if set, zero value otherwise.
func (o *Item) GetRelatedIdentifiers() []ItemRelatedIdentifier {
	if o == nil || IsNil(o.RelatedIdentifiers) {
		var ret []ItemRelatedIdentifier
		return ret
	}
	return o.RelatedIdentifiers
}

// GetRelatedIdentifiersOk returns a tuple with the RelatedIdentifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetRelatedIdentifiersOk() ([]ItemRelatedIdentifier, bool) {
	if o == nil || IsNil(o.RelatedIdentifiers) {
		return nil, false
	}
	return o.RelatedIdentifiers, true
}

// HasRelatedIdentifiers returns a boolean if a field has been set.
func (o *Item) HasRelatedIdentifiers() bool {
	if o != nil && !IsNil(o.RelatedIdentifiers) {
		return true
	}

	return false
}

// SetRelatedIdentifiers gets a reference to the given []ItemRelatedIdentifier and assigns it to the RelatedIdentifiers field.
func (o *Item) SetRelatedIdentifiers(v []ItemRelatedIdentifier) {
	o.RelatedIdentifiers = v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *Item) GetTotalAmount() Currency {
	if o == nil || IsNil(o.TotalAmount) {
		var ret Currency
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetTotalAmountOk() (*Currency, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *Item) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given Currency and assigns it to the TotalAmount field.
func (o *Item) SetTotalAmount(v Currency) {
	o.TotalAmount = &v
}

// GetBreakdowns returns the Breakdowns field value if set, zero value otherwise.
func (o *Item) GetBreakdowns() []Breakdown {
	if o == nil || IsNil(o.Breakdowns) {
		var ret []Breakdown
		return ret
	}
	return o.Breakdowns
}

// GetBreakdownsOk returns a tuple with the Breakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetBreakdownsOk() ([]Breakdown, bool) {
	if o == nil || IsNil(o.Breakdowns) {
		return nil, false
	}
	return o.Breakdowns, true
}

// HasBreakdowns returns a boolean if a field has been set.
func (o *Item) HasBreakdowns() bool {
	if o != nil && !IsNil(o.Breakdowns) {
		return true
	}

	return false
}

// SetBreakdowns gets a reference to the given []Breakdown and assigns it to the Breakdowns field.
func (o *Item) SetBreakdowns(v []Breakdown) {
	o.Breakdowns = v
}

// GetContexts returns the Contexts field value if set, zero value otherwise.
func (o *Item) GetContexts() []Context {
	if o == nil || IsNil(o.Contexts) {
		var ret []Context
		return ret
	}
	return o.Contexts
}

// GetContextsOk returns a tuple with the Contexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetContextsOk() ([]Context, bool) {
	if o == nil || IsNil(o.Contexts) {
		return nil, false
	}
	return o.Contexts, true
}

// HasContexts returns a boolean if a field has been set.
func (o *Item) HasContexts() bool {
	if o != nil && !IsNil(o.Contexts) {
		return true
	}

	return false
}

// SetContexts gets a reference to the given []Context and assigns it to the Contexts field.
func (o *Item) SetContexts(v []Context) {
	o.Contexts = v
}

func (o Item) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.RelatedIdentifiers) {
		toSerialize["relatedIdentifiers"] = o.RelatedIdentifiers
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if !IsNil(o.Breakdowns) {
		toSerialize["breakdowns"] = o.Breakdowns
	}
	if !IsNil(o.Contexts) {
		toSerialize["contexts"] = o.Contexts
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
