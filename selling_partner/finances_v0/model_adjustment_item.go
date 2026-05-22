package finances_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the AdjustmentItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdjustmentItem{}

// AdjustmentItem An item in an adjustment to the seller's account.
type AdjustmentItem struct {
	// Represents the number of units in the seller's inventory when the AdustmentType is FBAInventoryReimbursement.
	Quantity      *string   `json:"Quantity,omitempty"`
	PerUnitAmount *Currency `json:"PerUnitAmount,omitempty"`
	TotalAmount   *Currency `json:"TotalAmount,omitempty"`
	// The seller SKU of the item. The seller SKU is qualified by the seller's seller ID, which is included with every call to the Selling Partner API.
	SellerSKU *string `json:"SellerSKU,omitempty"`
	// A unique identifier assigned to products stored in and fulfilled from a fulfillment center.
	FnSKU *string `json:"FnSKU,omitempty"`
	// A short description of the item.
	ProductDescription *string `json:"ProductDescription,omitempty"`
	// The Amazon Standard Identification Number (ASIN) of the item.
	ASIN *string `json:"ASIN,omitempty"`
	// The transaction number that is related to the adjustment.
	TransactionNumber *string `json:"TransactionNumber,omitempty"`
}

// NewAdjustmentItem instantiates a new AdjustmentItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdjustmentItem() *AdjustmentItem {
	this := AdjustmentItem{}
	return &this
}

// NewAdjustmentItemWithDefaults instantiates a new AdjustmentItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdjustmentItemWithDefaults() *AdjustmentItem {
	this := AdjustmentItem{}
	return &this
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *AdjustmentItem) GetQuantity() string {
	if o == nil || IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustmentItem) GetQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *AdjustmentItem) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *AdjustmentItem) SetQuantity(v string) {
	o.Quantity = &v
}

// GetPerUnitAmount returns the PerUnitAmount field value if set, zero value otherwise.
func (o *AdjustmentItem) GetPerUnitAmount() Currency {
	if o == nil || IsNil(o.PerUnitAmount) {
		var ret Currency
		return ret
	}
	return *o.PerUnitAmount
}

// GetPerUnitAmountOk returns a tuple with the PerUnitAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustmentItem) GetPerUnitAmountOk() (*Currency, bool) {
	if o == nil || IsNil(o.PerUnitAmount) {
		return nil, false
	}
	return o.PerUnitAmount, true
}

// HasPerUnitAmount returns a boolean if a field has been set.
func (o *AdjustmentItem) HasPerUnitAmount() bool {
	if o != nil && !IsNil(o.PerUnitAmount) {
		return true
	}

	return false
}

// SetPerUnitAmount gets a reference to the given Currency and assigns it to the PerUnitAmount field.
func (o *AdjustmentItem) SetPerUnitAmount(v Currency) {
	o.PerUnitAmount = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *AdjustmentItem) GetTotalAmount() Currency {
	if o == nil || IsNil(o.TotalAmount) {
		var ret Currency
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustmentItem) GetTotalAmountOk() (*Currency, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *AdjustmentItem) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given Currency and assigns it to the TotalAmount field.
func (o *AdjustmentItem) SetTotalAmount(v Currency) {
	o.TotalAmount = &v
}

// GetSellerSKU returns the SellerSKU field value if set, zero value otherwise.
func (o *AdjustmentItem) GetSellerSKU() string {
	if o == nil || IsNil(o.SellerSKU) {
		var ret string
		return ret
	}
	return *o.SellerSKU
}

// GetSellerSKUOk returns a tuple with the SellerSKU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustmentItem) GetSellerSKUOk() (*string, bool) {
	if o == nil || IsNil(o.SellerSKU) {
		return nil, false
	}
	return o.SellerSKU, true
}

// HasSellerSKU returns a boolean if a field has been set.
func (o *AdjustmentItem) HasSellerSKU() bool {
	if o != nil && !IsNil(o.SellerSKU) {
		return true
	}

	return false
}

// SetSellerSKU gets a reference to the given string and assigns it to the SellerSKU field.
func (o *AdjustmentItem) SetSellerSKU(v string) {
	o.SellerSKU = &v
}

// GetFnSKU returns the FnSKU field value if set, zero value otherwise.
func (o *AdjustmentItem) GetFnSKU() string {
	if o == nil || IsNil(o.FnSKU) {
		var ret string
		return ret
	}
	return *o.FnSKU
}

// GetFnSKUOk returns a tuple with the FnSKU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustmentItem) GetFnSKUOk() (*string, bool) {
	if o == nil || IsNil(o.FnSKU) {
		return nil, false
	}
	return o.FnSKU, true
}

// HasFnSKU returns a boolean if a field has been set.
func (o *AdjustmentItem) HasFnSKU() bool {
	if o != nil && !IsNil(o.FnSKU) {
		return true
	}

	return false
}

// SetFnSKU gets a reference to the given string and assigns it to the FnSKU field.
func (o *AdjustmentItem) SetFnSKU(v string) {
	o.FnSKU = &v
}

// GetProductDescription returns the ProductDescription field value if set, zero value otherwise.
func (o *AdjustmentItem) GetProductDescription() string {
	if o == nil || IsNil(o.ProductDescription) {
		var ret string
		return ret
	}
	return *o.ProductDescription
}

// GetProductDescriptionOk returns a tuple with the ProductDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustmentItem) GetProductDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductDescription) {
		return nil, false
	}
	return o.ProductDescription, true
}

// HasProductDescription returns a boolean if a field has been set.
func (o *AdjustmentItem) HasProductDescription() bool {
	if o != nil && !IsNil(o.ProductDescription) {
		return true
	}

	return false
}

// SetProductDescription gets a reference to the given string and assigns it to the ProductDescription field.
func (o *AdjustmentItem) SetProductDescription(v string) {
	o.ProductDescription = &v
}

// GetASIN returns the ASIN field value if set, zero value otherwise.
func (o *AdjustmentItem) GetASIN() string {
	if o == nil || IsNil(o.ASIN) {
		var ret string
		return ret
	}
	return *o.ASIN
}

// GetASINOk returns a tuple with the ASIN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustmentItem) GetASINOk() (*string, bool) {
	if o == nil || IsNil(o.ASIN) {
		return nil, false
	}
	return o.ASIN, true
}

// HasASIN returns a boolean if a field has been set.
func (o *AdjustmentItem) HasASIN() bool {
	if o != nil && !IsNil(o.ASIN) {
		return true
	}

	return false
}

// SetASIN gets a reference to the given string and assigns it to the ASIN field.
func (o *AdjustmentItem) SetASIN(v string) {
	o.ASIN = &v
}

// GetTransactionNumber returns the TransactionNumber field value if set, zero value otherwise.
func (o *AdjustmentItem) GetTransactionNumber() string {
	if o == nil || IsNil(o.TransactionNumber) {
		var ret string
		return ret
	}
	return *o.TransactionNumber
}

// GetTransactionNumberOk returns a tuple with the TransactionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustmentItem) GetTransactionNumberOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionNumber) {
		return nil, false
	}
	return o.TransactionNumber, true
}

// HasTransactionNumber returns a boolean if a field has been set.
func (o *AdjustmentItem) HasTransactionNumber() bool {
	if o != nil && !IsNil(o.TransactionNumber) {
		return true
	}

	return false
}

// SetTransactionNumber gets a reference to the given string and assigns it to the TransactionNumber field.
func (o *AdjustmentItem) SetTransactionNumber(v string) {
	o.TransactionNumber = &v
}

func (o AdjustmentItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Quantity) {
		toSerialize["Quantity"] = o.Quantity
	}
	if !IsNil(o.PerUnitAmount) {
		toSerialize["PerUnitAmount"] = o.PerUnitAmount
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["TotalAmount"] = o.TotalAmount
	}
	if !IsNil(o.SellerSKU) {
		toSerialize["SellerSKU"] = o.SellerSKU
	}
	if !IsNil(o.FnSKU) {
		toSerialize["FnSKU"] = o.FnSKU
	}
	if !IsNil(o.ProductDescription) {
		toSerialize["ProductDescription"] = o.ProductDescription
	}
	if !IsNil(o.ASIN) {
		toSerialize["ASIN"] = o.ASIN
	}
	if !IsNil(o.TransactionNumber) {
		toSerialize["TransactionNumber"] = o.TransactionNumber
	}
	return toSerialize, nil
}

type NullableAdjustmentItem struct {
	value *AdjustmentItem
	isSet bool
}

func (v NullableAdjustmentItem) Get() *AdjustmentItem {
	return v.value
}

func (v *NullableAdjustmentItem) Set(val *AdjustmentItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAdjustmentItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAdjustmentItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdjustmentItem(val *AdjustmentItem) *NullableAdjustmentItem {
	return &NullableAdjustmentItem{value: val, isSet: true}
}

func (v NullableAdjustmentItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdjustmentItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
