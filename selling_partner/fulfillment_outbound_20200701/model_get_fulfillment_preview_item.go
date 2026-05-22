package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the GetFulfillmentPreviewItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFulfillmentPreviewItem{}

// GetFulfillmentPreviewItem Item information for a fulfillment order preview.
type GetFulfillmentPreviewItem struct {
	// The seller SKU of the item.
	SellerSku string `json:"sellerSku"`
	// The item quantity.
	Quantity             int32  `json:"quantity"`
	PerUnitDeclaredValue *Money `json:"perUnitDeclaredValue,omitempty"`
	// A fulfillment order item identifier that the seller creates to track items in the fulfillment preview.
	SellerFulfillmentOrderItemId string `json:"sellerFulfillmentOrderItemId"`
}

type _GetFulfillmentPreviewItem GetFulfillmentPreviewItem

// NewGetFulfillmentPreviewItem instantiates a new GetFulfillmentPreviewItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFulfillmentPreviewItem(sellerSku string, quantity int32, sellerFulfillmentOrderItemId string) *GetFulfillmentPreviewItem {
	this := GetFulfillmentPreviewItem{}
	this.SellerSku = sellerSku
	this.Quantity = quantity
	this.SellerFulfillmentOrderItemId = sellerFulfillmentOrderItemId
	return &this
}

// NewGetFulfillmentPreviewItemWithDefaults instantiates a new GetFulfillmentPreviewItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFulfillmentPreviewItemWithDefaults() *GetFulfillmentPreviewItem {
	this := GetFulfillmentPreviewItem{}
	return &this
}

// GetSellerSku returns the SellerSku field value
func (o *GetFulfillmentPreviewItem) GetSellerSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerSku
}

// GetSellerSkuOk returns a tuple with the SellerSku field value
// and a boolean to check if the value has been set.
func (o *GetFulfillmentPreviewItem) GetSellerSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerSku, true
}

// SetSellerSku sets field value
func (o *GetFulfillmentPreviewItem) SetSellerSku(v string) {
	o.SellerSku = v
}

// GetQuantity returns the Quantity field value
func (o *GetFulfillmentPreviewItem) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *GetFulfillmentPreviewItem) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *GetFulfillmentPreviewItem) SetQuantity(v int32) {
	o.Quantity = v
}

// GetPerUnitDeclaredValue returns the PerUnitDeclaredValue field value if set, zero value otherwise.
func (o *GetFulfillmentPreviewItem) GetPerUnitDeclaredValue() Money {
	if o == nil || IsNil(o.PerUnitDeclaredValue) {
		var ret Money
		return ret
	}
	return *o.PerUnitDeclaredValue
}

// GetPerUnitDeclaredValueOk returns a tuple with the PerUnitDeclaredValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFulfillmentPreviewItem) GetPerUnitDeclaredValueOk() (*Money, bool) {
	if o == nil || IsNil(o.PerUnitDeclaredValue) {
		return nil, false
	}
	return o.PerUnitDeclaredValue, true
}

// HasPerUnitDeclaredValue returns a boolean if a field has been set.
func (o *GetFulfillmentPreviewItem) HasPerUnitDeclaredValue() bool {
	if o != nil && !IsNil(o.PerUnitDeclaredValue) {
		return true
	}

	return false
}

// SetPerUnitDeclaredValue gets a reference to the given Money and assigns it to the PerUnitDeclaredValue field.
func (o *GetFulfillmentPreviewItem) SetPerUnitDeclaredValue(v Money) {
	o.PerUnitDeclaredValue = &v
}

// GetSellerFulfillmentOrderItemId returns the SellerFulfillmentOrderItemId field value
func (o *GetFulfillmentPreviewItem) GetSellerFulfillmentOrderItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerFulfillmentOrderItemId
}

// GetSellerFulfillmentOrderItemIdOk returns a tuple with the SellerFulfillmentOrderItemId field value
// and a boolean to check if the value has been set.
func (o *GetFulfillmentPreviewItem) GetSellerFulfillmentOrderItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerFulfillmentOrderItemId, true
}

// SetSellerFulfillmentOrderItemId sets field value
func (o *GetFulfillmentPreviewItem) SetSellerFulfillmentOrderItemId(v string) {
	o.SellerFulfillmentOrderItemId = v
}

func (o GetFulfillmentPreviewItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sellerSku"] = o.SellerSku
	toSerialize["quantity"] = o.Quantity
	if !IsNil(o.PerUnitDeclaredValue) {
		toSerialize["perUnitDeclaredValue"] = o.PerUnitDeclaredValue
	}
	toSerialize["sellerFulfillmentOrderItemId"] = o.SellerFulfillmentOrderItemId
	return toSerialize, nil
}

type NullableGetFulfillmentPreviewItem struct {
	value *GetFulfillmentPreviewItem
	isSet bool
}

func (v NullableGetFulfillmentPreviewItem) Get() *GetFulfillmentPreviewItem {
	return v.value
}

func (v *NullableGetFulfillmentPreviewItem) Set(val *GetFulfillmentPreviewItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFulfillmentPreviewItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFulfillmentPreviewItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFulfillmentPreviewItem(val *GetFulfillmentPreviewItem) *NullableGetFulfillmentPreviewItem {
	return &NullableGetFulfillmentPreviewItem{value: val, isSet: true}
}

func (v NullableGetFulfillmentPreviewItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetFulfillmentPreviewItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
