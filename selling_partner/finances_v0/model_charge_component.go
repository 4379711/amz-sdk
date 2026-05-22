package finances_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ChargeComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeComponent{}

// ChargeComponent A charge on the seller's account.  Possible values:  * Principal - The selling price of the order item, equal to the selling price of the item multiplied by the quantity ordered.  * Tax - The tax collected by the seller on the Principal.  * MarketplaceFacilitatorTax-Principal - The tax withheld on the Principal.  * MarketplaceFacilitatorTax-Shipping - The tax withheld on the ShippingCharge.  * MarketplaceFacilitatorTax-Giftwrap - The tax withheld on the Giftwrap charge.  * MarketplaceFacilitatorTax-Other - The tax withheld on other miscellaneous charges.  * Discount - The promotional discount for an order item.  * TaxDiscount - The tax amount deducted for promotional rebates.  * CODItemCharge - The COD charge for an order item.  * CODItemTaxCharge - The tax collected by the seller on a CODItemCharge.  * CODOrderCharge - The COD charge for an order.  * CODOrderTaxCharge - The tax collected by the seller on a CODOrderCharge.  * CODShippingCharge - Shipping charges for a COD order.  * CODShippingTaxCharge - The tax collected by the seller on a CODShippingCharge.  * ShippingCharge - The shipping charge.  * ShippingTax - The tax collected by the seller on a ShippingCharge.  * Goodwill - The amount given to a buyer as a gesture of goodwill or to compensate for pain and suffering in the buying experience.  * Giftwrap - The gift wrap charge.  * GiftwrapTax - The tax collected by the seller on a Giftwrap charge.  * RestockingFee - The charge applied to the buyer when returning a product in certain categories.  * ReturnShipping - The amount given to the buyer to compensate for shipping the item back in the event we are at fault.  * PointsFee - The value of Amazon Points deducted from the refund if the buyer does not have enough Amazon Points to cover the deduction.  * GenericDeduction - A generic bad debt deduction.  * FreeReplacementReturnShipping - The compensation for return shipping when a buyer receives the wrong item, requests a free replacement, and returns the incorrect item.  * PaymentMethodFee - The fee collected for certain payment methods in certain marketplaces.  * ExportCharge - The export duty that is charged when an item is shipped to an international destination as part of the Amazon Global program.  * SAFE-TReimbursement - The SAFE-T claim amount for the item.  * TCS-CGST - Tax Collected at Source (TCS) for Central Goods and Services Tax (CGST).  * TCS-SGST - Tax Collected at Source for State Goods and Services Tax (SGST).  * TCS-IGST - Tax Collected at Source for Integrated Goods and Services Tax (IGST).  * TCS-UTGST - Tax Collected at Source for Union Territories Goods and Services Tax (UTGST).
type ChargeComponent struct {
	// The type of charge.
	ChargeType   *string   `json:"ChargeType,omitempty"`
	ChargeAmount *Currency `json:"ChargeAmount,omitempty"`
}

// NewChargeComponent instantiates a new ChargeComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeComponent() *ChargeComponent {
	this := ChargeComponent{}
	return &this
}

// NewChargeComponentWithDefaults instantiates a new ChargeComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeComponentWithDefaults() *ChargeComponent {
	this := ChargeComponent{}
	return &this
}

// GetChargeType returns the ChargeType field value if set, zero value otherwise.
func (o *ChargeComponent) GetChargeType() string {
	if o == nil || IsNil(o.ChargeType) {
		var ret string
		return ret
	}
	return *o.ChargeType
}

// GetChargeTypeOk returns a tuple with the ChargeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeComponent) GetChargeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ChargeType) {
		return nil, false
	}
	return o.ChargeType, true
}

// HasChargeType returns a boolean if a field has been set.
func (o *ChargeComponent) HasChargeType() bool {
	if o != nil && !IsNil(o.ChargeType) {
		return true
	}

	return false
}

// SetChargeType gets a reference to the given string and assigns it to the ChargeType field.
func (o *ChargeComponent) SetChargeType(v string) {
	o.ChargeType = &v
}

// GetChargeAmount returns the ChargeAmount field value if set, zero value otherwise.
func (o *ChargeComponent) GetChargeAmount() Currency {
	if o == nil || IsNil(o.ChargeAmount) {
		var ret Currency
		return ret
	}
	return *o.ChargeAmount
}

// GetChargeAmountOk returns a tuple with the ChargeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeComponent) GetChargeAmountOk() (*Currency, bool) {
	if o == nil || IsNil(o.ChargeAmount) {
		return nil, false
	}
	return o.ChargeAmount, true
}

// HasChargeAmount returns a boolean if a field has been set.
func (o *ChargeComponent) HasChargeAmount() bool {
	if o != nil && !IsNil(o.ChargeAmount) {
		return true
	}

	return false
}

// SetChargeAmount gets a reference to the given Currency and assigns it to the ChargeAmount field.
func (o *ChargeComponent) SetChargeAmount(v Currency) {
	o.ChargeAmount = &v
}

func (o ChargeComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChargeType) {
		toSerialize["ChargeType"] = o.ChargeType
	}
	if !IsNil(o.ChargeAmount) {
		toSerialize["ChargeAmount"] = o.ChargeAmount
	}
	return toSerialize, nil
}

type NullableChargeComponent struct {
	value *ChargeComponent
	isSet bool
}

func (v NullableChargeComponent) Get() *ChargeComponent {
	return v.value
}

func (v *NullableChargeComponent) Set(val *ChargeComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeComponent(val *ChargeComponent) *NullableChargeComponent {
	return &NullableChargeComponent{value: val, isSet: true}
}

func (v NullableChargeComponent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableChargeComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
