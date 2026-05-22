package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemInput{}

// ItemInput Defines an item's input parameters.
type ItemInput struct {
	// The expiration date of the MSKU. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime format with pattern `YYYY-MM-DD`. Items with the same MSKU but different expiration dates cannot go into the same box.
	Expiration *string    `json:"expiration,omitempty" validate:"regexp=^([0-9]{4})-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$"`
	LabelOwner LabelOwner `json:"labelOwner"`
	// The manufacturing lot code.
	ManufacturingLotCode *string `json:"manufacturingLotCode,omitempty"`
	// The merchant SKU, a merchant-supplied identifier of a specific SKU.
	Msku      string    `json:"msku"`
	PrepOwner PrepOwner `json:"prepOwner"`
	// The number of units of the specified MSKU that will be shipped.
	Quantity int32 `json:"quantity"`
}

type _ItemInput ItemInput

// NewItemInput instantiates a new ItemInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemInput(labelOwner LabelOwner, msku string, prepOwner PrepOwner, quantity int32) *ItemInput {
	this := ItemInput{}
	this.LabelOwner = labelOwner
	this.Msku = msku
	this.PrepOwner = prepOwner
	this.Quantity = quantity
	return &this
}

// NewItemInputWithDefaults instantiates a new ItemInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemInputWithDefaults() *ItemInput {
	this := ItemInput{}
	return &this
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *ItemInput) GetExpiration() string {
	if o == nil || IsNil(o.Expiration) {
		var ret string
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemInput) GetExpirationOk() (*string, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *ItemInput) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given string and assigns it to the Expiration field.
func (o *ItemInput) SetExpiration(v string) {
	o.Expiration = &v
}

// GetLabelOwner returns the LabelOwner field value
func (o *ItemInput) GetLabelOwner() LabelOwner {
	if o == nil {
		var ret LabelOwner
		return ret
	}

	return o.LabelOwner
}

// GetLabelOwnerOk returns a tuple with the LabelOwner field value
// and a boolean to check if the value has been set.
func (o *ItemInput) GetLabelOwnerOk() (*LabelOwner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelOwner, true
}

// SetLabelOwner sets field value
func (o *ItemInput) SetLabelOwner(v LabelOwner) {
	o.LabelOwner = v
}

// GetManufacturingLotCode returns the ManufacturingLotCode field value if set, zero value otherwise.
func (o *ItemInput) GetManufacturingLotCode() string {
	if o == nil || IsNil(o.ManufacturingLotCode) {
		var ret string
		return ret
	}
	return *o.ManufacturingLotCode
}

// GetManufacturingLotCodeOk returns a tuple with the ManufacturingLotCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemInput) GetManufacturingLotCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ManufacturingLotCode) {
		return nil, false
	}
	return o.ManufacturingLotCode, true
}

// HasManufacturingLotCode returns a boolean if a field has been set.
func (o *ItemInput) HasManufacturingLotCode() bool {
	if o != nil && !IsNil(o.ManufacturingLotCode) {
		return true
	}

	return false
}

// SetManufacturingLotCode gets a reference to the given string and assigns it to the ManufacturingLotCode field.
func (o *ItemInput) SetManufacturingLotCode(v string) {
	o.ManufacturingLotCode = &v
}

// GetMsku returns the Msku field value
func (o *ItemInput) GetMsku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msku
}

// GetMskuOk returns a tuple with the Msku field value
// and a boolean to check if the value has been set.
func (o *ItemInput) GetMskuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msku, true
}

// SetMsku sets field value
func (o *ItemInput) SetMsku(v string) {
	o.Msku = v
}

// GetPrepOwner returns the PrepOwner field value
func (o *ItemInput) GetPrepOwner() PrepOwner {
	if o == nil {
		var ret PrepOwner
		return ret
	}

	return o.PrepOwner
}

// GetPrepOwnerOk returns a tuple with the PrepOwner field value
// and a boolean to check if the value has been set.
func (o *ItemInput) GetPrepOwnerOk() (*PrepOwner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrepOwner, true
}

// SetPrepOwner sets field value
func (o *ItemInput) SetPrepOwner(v PrepOwner) {
	o.PrepOwner = v
}

// GetQuantity returns the Quantity field value
func (o *ItemInput) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *ItemInput) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *ItemInput) SetQuantity(v int32) {
	o.Quantity = v
}

func (o ItemInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	toSerialize["labelOwner"] = o.LabelOwner
	if !IsNil(o.ManufacturingLotCode) {
		toSerialize["manufacturingLotCode"] = o.ManufacturingLotCode
	}
	toSerialize["msku"] = o.Msku
	toSerialize["prepOwner"] = o.PrepOwner
	toSerialize["quantity"] = o.Quantity
	return toSerialize, nil
}

type NullableItemInput struct {
	value *ItemInput
	isSet bool
}

func (v NullableItemInput) Get() *ItemInput {
	return v.value
}

func (v *NullableItemInput) Set(val *ItemInput) {
	v.value = val
	v.isSet = true
}

func (v NullableItemInput) IsSet() bool {
	return v.isSet
}

func (v *NullableItemInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemInput(val *ItemInput) *NullableItemInput {
	return &NullableItemInput{value: val, isSet: true}
}

func (v NullableItemInput) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
