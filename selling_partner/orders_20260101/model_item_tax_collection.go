package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemTaxCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemTaxCollection{}

// ItemTaxCollection Tax collection information for an order item.
type ItemTaxCollection struct {
	// The tax collection model applied to the item.  **Possible values**: - `MARKETPLACE_FACILITATOR` (Tax is withheld and remitted to the taxing authority by Amazon on behalf of the seller)
	Model *string `json:"model,omitempty"`
	// The party responsible for withholding the taxes and remitting them to the taxing authority.
	ResponsibleParty *string `json:"responsibleParty,omitempty"`
}

// NewItemTaxCollection instantiates a new ItemTaxCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemTaxCollection() *ItemTaxCollection {
	this := ItemTaxCollection{}
	return &this
}

// NewItemTaxCollectionWithDefaults instantiates a new ItemTaxCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemTaxCollectionWithDefaults() *ItemTaxCollection {
	this := ItemTaxCollection{}
	return &this
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ItemTaxCollection) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemTaxCollection) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ItemTaxCollection) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ItemTaxCollection) SetModel(v string) {
	o.Model = &v
}

// GetResponsibleParty returns the ResponsibleParty field value if set, zero value otherwise.
func (o *ItemTaxCollection) GetResponsibleParty() string {
	if o == nil || IsNil(o.ResponsibleParty) {
		var ret string
		return ret
	}
	return *o.ResponsibleParty
}

// GetResponsiblePartyOk returns a tuple with the ResponsibleParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemTaxCollection) GetResponsiblePartyOk() (*string, bool) {
	if o == nil || IsNil(o.ResponsibleParty) {
		return nil, false
	}
	return o.ResponsibleParty, true
}

// HasResponsibleParty returns a boolean if a field has been set.
func (o *ItemTaxCollection) HasResponsibleParty() bool {
	if o != nil && !IsNil(o.ResponsibleParty) {
		return true
	}

	return false
}

// SetResponsibleParty gets a reference to the given string and assigns it to the ResponsibleParty field.
func (o *ItemTaxCollection) SetResponsibleParty(v string) {
	o.ResponsibleParty = &v
}

func (o ItemTaxCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.ResponsibleParty) {
		toSerialize["responsibleParty"] = o.ResponsibleParty
	}
	return toSerialize, nil
}

type NullableItemTaxCollection struct {
	value *ItemTaxCollection
	isSet bool
}

func (v NullableItemTaxCollection) Get() *ItemTaxCollection {
	return v.value
}

func (v *NullableItemTaxCollection) Set(val *ItemTaxCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableItemTaxCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableItemTaxCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemTaxCollection(val *ItemTaxCollection) *NullableItemTaxCollection {
	return &NullableItemTaxCollection{value: val, isSet: true}
}

func (v NullableItemTaxCollection) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemTaxCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
