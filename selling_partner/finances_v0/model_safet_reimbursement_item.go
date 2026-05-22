package finances_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the SAFETReimbursementItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAFETReimbursementItem{}

// SAFETReimbursementItem An item from a SAFE-T claim reimbursement.
type SAFETReimbursementItem struct {
	// A list of charge information on the seller's account.
	ItemChargeList []ChargeComponent `json:"itemChargeList,omitempty"`
	// The description of the item as shown on the product detail page on the retail website.
	ProductDescription *string `json:"productDescription,omitempty"`
	// The number of units of the item being reimbursed.
	Quantity *string `json:"quantity,omitempty"`
}

// NewSAFETReimbursementItem instantiates a new SAFETReimbursementItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAFETReimbursementItem() *SAFETReimbursementItem {
	this := SAFETReimbursementItem{}
	return &this
}

// NewSAFETReimbursementItemWithDefaults instantiates a new SAFETReimbursementItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAFETReimbursementItemWithDefaults() *SAFETReimbursementItem {
	this := SAFETReimbursementItem{}
	return &this
}

// GetItemChargeList returns the ItemChargeList field value if set, zero value otherwise.
func (o *SAFETReimbursementItem) GetItemChargeList() []ChargeComponent {
	if o == nil || IsNil(o.ItemChargeList) {
		var ret []ChargeComponent
		return ret
	}
	return o.ItemChargeList
}

// GetItemChargeListOk returns a tuple with the ItemChargeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAFETReimbursementItem) GetItemChargeListOk() ([]ChargeComponent, bool) {
	if o == nil || IsNil(o.ItemChargeList) {
		return nil, false
	}
	return o.ItemChargeList, true
}

// HasItemChargeList returns a boolean if a field has been set.
func (o *SAFETReimbursementItem) HasItemChargeList() bool {
	if o != nil && !IsNil(o.ItemChargeList) {
		return true
	}

	return false
}

// SetItemChargeList gets a reference to the given []ChargeComponent and assigns it to the ItemChargeList field.
func (o *SAFETReimbursementItem) SetItemChargeList(v []ChargeComponent) {
	o.ItemChargeList = v
}

// GetProductDescription returns the ProductDescription field value if set, zero value otherwise.
func (o *SAFETReimbursementItem) GetProductDescription() string {
	if o == nil || IsNil(o.ProductDescription) {
		var ret string
		return ret
	}
	return *o.ProductDescription
}

// GetProductDescriptionOk returns a tuple with the ProductDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAFETReimbursementItem) GetProductDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductDescription) {
		return nil, false
	}
	return o.ProductDescription, true
}

// HasProductDescription returns a boolean if a field has been set.
func (o *SAFETReimbursementItem) HasProductDescription() bool {
	if o != nil && !IsNil(o.ProductDescription) {
		return true
	}

	return false
}

// SetProductDescription gets a reference to the given string and assigns it to the ProductDescription field.
func (o *SAFETReimbursementItem) SetProductDescription(v string) {
	o.ProductDescription = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *SAFETReimbursementItem) GetQuantity() string {
	if o == nil || IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAFETReimbursementItem) GetQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *SAFETReimbursementItem) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *SAFETReimbursementItem) SetQuantity(v string) {
	o.Quantity = &v
}

func (o SAFETReimbursementItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemChargeList) {
		toSerialize["itemChargeList"] = o.ItemChargeList
	}
	if !IsNil(o.ProductDescription) {
		toSerialize["productDescription"] = o.ProductDescription
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableSAFETReimbursementItem struct {
	value *SAFETReimbursementItem
	isSet bool
}

func (v NullableSAFETReimbursementItem) Get() *SAFETReimbursementItem {
	return v.value
}

func (v *NullableSAFETReimbursementItem) Set(val *SAFETReimbursementItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSAFETReimbursementItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSAFETReimbursementItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAFETReimbursementItem(val *SAFETReimbursementItem) *NullableSAFETReimbursementItem {
	return &NullableSAFETReimbursementItem{value: val, isSet: true}
}

func (v NullableSAFETReimbursementItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSAFETReimbursementItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
