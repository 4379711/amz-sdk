package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderBuyerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderBuyerInfo{}

// OrderBuyerInfo Buyer information for an order.
type OrderBuyerInfo struct {
	// An Amazon-defined order identifier, in 3-7-7 format.
	AmazonOrderId string `json:"AmazonOrderId"`
	// The anonymized email address of the buyer.
	BuyerEmail *string `json:"BuyerEmail,omitempty"`
	// The buyer name or the recipient name.
	BuyerName *string `json:"BuyerName,omitempty"`
	// The county of the buyer.  **Note**: This attribute is only available in the Brazil marketplace.
	BuyerCounty  *string       `json:"BuyerCounty,omitempty"`
	BuyerTaxInfo *BuyerTaxInfo `json:"BuyerTaxInfo,omitempty"`
	// The purchase order (PO) number entered by the buyer at checkout. Only returned for orders where the buyer entered a PO number at checkout.
	PurchaseOrderNumber *string `json:"PurchaseOrderNumber,omitempty"`
}

type _OrderBuyerInfo OrderBuyerInfo

// NewOrderBuyerInfo instantiates a new OrderBuyerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderBuyerInfo(amazonOrderId string) *OrderBuyerInfo {
	this := OrderBuyerInfo{}
	this.AmazonOrderId = amazonOrderId
	return &this
}

// NewOrderBuyerInfoWithDefaults instantiates a new OrderBuyerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderBuyerInfoWithDefaults() *OrderBuyerInfo {
	this := OrderBuyerInfo{}
	return &this
}

// GetAmazonOrderId returns the AmazonOrderId field value
func (o *OrderBuyerInfo) GetAmazonOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmazonOrderId
}

// GetAmazonOrderIdOk returns a tuple with the AmazonOrderId field value
// and a boolean to check if the value has been set.
func (o *OrderBuyerInfo) GetAmazonOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmazonOrderId, true
}

// SetAmazonOrderId sets field value
func (o *OrderBuyerInfo) SetAmazonOrderId(v string) {
	o.AmazonOrderId = v
}

// GetBuyerEmail returns the BuyerEmail field value if set, zero value otherwise.
func (o *OrderBuyerInfo) GetBuyerEmail() string {
	if o == nil || IsNil(o.BuyerEmail) {
		var ret string
		return ret
	}
	return *o.BuyerEmail
}

// GetBuyerEmailOk returns a tuple with the BuyerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBuyerInfo) GetBuyerEmailOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerEmail) {
		return nil, false
	}
	return o.BuyerEmail, true
}

// HasBuyerEmail returns a boolean if a field has been set.
func (o *OrderBuyerInfo) HasBuyerEmail() bool {
	if o != nil && !IsNil(o.BuyerEmail) {
		return true
	}

	return false
}

// SetBuyerEmail gets a reference to the given string and assigns it to the BuyerEmail field.
func (o *OrderBuyerInfo) SetBuyerEmail(v string) {
	o.BuyerEmail = &v
}

// GetBuyerName returns the BuyerName field value if set, zero value otherwise.
func (o *OrderBuyerInfo) GetBuyerName() string {
	if o == nil || IsNil(o.BuyerName) {
		var ret string
		return ret
	}
	return *o.BuyerName
}

// GetBuyerNameOk returns a tuple with the BuyerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBuyerInfo) GetBuyerNameOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerName) {
		return nil, false
	}
	return o.BuyerName, true
}

// HasBuyerName returns a boolean if a field has been set.
func (o *OrderBuyerInfo) HasBuyerName() bool {
	if o != nil && !IsNil(o.BuyerName) {
		return true
	}

	return false
}

// SetBuyerName gets a reference to the given string and assigns it to the BuyerName field.
func (o *OrderBuyerInfo) SetBuyerName(v string) {
	o.BuyerName = &v
}

// GetBuyerCounty returns the BuyerCounty field value if set, zero value otherwise.
func (o *OrderBuyerInfo) GetBuyerCounty() string {
	if o == nil || IsNil(o.BuyerCounty) {
		var ret string
		return ret
	}
	return *o.BuyerCounty
}

// GetBuyerCountyOk returns a tuple with the BuyerCounty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBuyerInfo) GetBuyerCountyOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerCounty) {
		return nil, false
	}
	return o.BuyerCounty, true
}

// HasBuyerCounty returns a boolean if a field has been set.
func (o *OrderBuyerInfo) HasBuyerCounty() bool {
	if o != nil && !IsNil(o.BuyerCounty) {
		return true
	}

	return false
}

// SetBuyerCounty gets a reference to the given string and assigns it to the BuyerCounty field.
func (o *OrderBuyerInfo) SetBuyerCounty(v string) {
	o.BuyerCounty = &v
}

// GetBuyerTaxInfo returns the BuyerTaxInfo field value if set, zero value otherwise.
func (o *OrderBuyerInfo) GetBuyerTaxInfo() BuyerTaxInfo {
	if o == nil || IsNil(o.BuyerTaxInfo) {
		var ret BuyerTaxInfo
		return ret
	}
	return *o.BuyerTaxInfo
}

// GetBuyerTaxInfoOk returns a tuple with the BuyerTaxInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBuyerInfo) GetBuyerTaxInfoOk() (*BuyerTaxInfo, bool) {
	if o == nil || IsNil(o.BuyerTaxInfo) {
		return nil, false
	}
	return o.BuyerTaxInfo, true
}

// HasBuyerTaxInfo returns a boolean if a field has been set.
func (o *OrderBuyerInfo) HasBuyerTaxInfo() bool {
	if o != nil && !IsNil(o.BuyerTaxInfo) {
		return true
	}

	return false
}

// SetBuyerTaxInfo gets a reference to the given BuyerTaxInfo and assigns it to the BuyerTaxInfo field.
func (o *OrderBuyerInfo) SetBuyerTaxInfo(v BuyerTaxInfo) {
	o.BuyerTaxInfo = &v
}

// GetPurchaseOrderNumber returns the PurchaseOrderNumber field value if set, zero value otherwise.
func (o *OrderBuyerInfo) GetPurchaseOrderNumber() string {
	if o == nil || IsNil(o.PurchaseOrderNumber) {
		var ret string
		return ret
	}
	return *o.PurchaseOrderNumber
}

// GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBuyerInfo) GetPurchaseOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseOrderNumber) {
		return nil, false
	}
	return o.PurchaseOrderNumber, true
}

// HasPurchaseOrderNumber returns a boolean if a field has been set.
func (o *OrderBuyerInfo) HasPurchaseOrderNumber() bool {
	if o != nil && !IsNil(o.PurchaseOrderNumber) {
		return true
	}

	return false
}

// SetPurchaseOrderNumber gets a reference to the given string and assigns it to the PurchaseOrderNumber field.
func (o *OrderBuyerInfo) SetPurchaseOrderNumber(v string) {
	o.PurchaseOrderNumber = &v
}

func (o OrderBuyerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AmazonOrderId"] = o.AmazonOrderId
	if !IsNil(o.BuyerEmail) {
		toSerialize["BuyerEmail"] = o.BuyerEmail
	}
	if !IsNil(o.BuyerName) {
		toSerialize["BuyerName"] = o.BuyerName
	}
	if !IsNil(o.BuyerCounty) {
		toSerialize["BuyerCounty"] = o.BuyerCounty
	}
	if !IsNil(o.BuyerTaxInfo) {
		toSerialize["BuyerTaxInfo"] = o.BuyerTaxInfo
	}
	if !IsNil(o.PurchaseOrderNumber) {
		toSerialize["PurchaseOrderNumber"] = o.PurchaseOrderNumber
	}
	return toSerialize, nil
}

type NullableOrderBuyerInfo struct {
	value *OrderBuyerInfo
	isSet bool
}

func (v NullableOrderBuyerInfo) Get() *OrderBuyerInfo {
	return v.value
}

func (v *NullableOrderBuyerInfo) Set(val *OrderBuyerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderBuyerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderBuyerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderBuyerInfo(val *OrderBuyerInfo) *NullableOrderBuyerInfo {
	return &NullableOrderBuyerInfo{value: val, isSet: true}
}

func (v NullableOrderBuyerInfo) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderBuyerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
