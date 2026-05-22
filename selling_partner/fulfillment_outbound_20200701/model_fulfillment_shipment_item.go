package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the FulfillmentShipmentItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FulfillmentShipmentItem{}

// FulfillmentShipmentItem Item information for a shipment in a fulfillment order.
type FulfillmentShipmentItem struct {
	// The seller SKU of the item.
	SellerSku string `json:"sellerSku"`
	// The fulfillment order item identifier that the seller created and submitted with a call to the `createFulfillmentOrder` operation.
	SellerFulfillmentOrderItemId string `json:"sellerFulfillmentOrderItemId"`
	// The item quantity.
	Quantity int32 `json:"quantity"`
	// An identifier for the package that contains the item quantity.
	PackageNumber *int32 `json:"packageNumber,omitempty"`
	// The serial number of the shipped item.
	SerialNumber *string `json:"serialNumber,omitempty"`
	// String list
	ManufacturerLotCodes []string `json:"manufacturerLotCodes,omitempty"`
}

type _FulfillmentShipmentItem FulfillmentShipmentItem

// NewFulfillmentShipmentItem instantiates a new FulfillmentShipmentItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillmentShipmentItem(sellerSku string, sellerFulfillmentOrderItemId string, quantity int32) *FulfillmentShipmentItem {
	this := FulfillmentShipmentItem{}
	this.SellerSku = sellerSku
	this.SellerFulfillmentOrderItemId = sellerFulfillmentOrderItemId
	this.Quantity = quantity
	return &this
}

// NewFulfillmentShipmentItemWithDefaults instantiates a new FulfillmentShipmentItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillmentShipmentItemWithDefaults() *FulfillmentShipmentItem {
	this := FulfillmentShipmentItem{}
	return &this
}

// GetSellerSku returns the SellerSku field value
func (o *FulfillmentShipmentItem) GetSellerSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerSku
}

// GetSellerSkuOk returns a tuple with the SellerSku field value
// and a boolean to check if the value has been set.
func (o *FulfillmentShipmentItem) GetSellerSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerSku, true
}

// SetSellerSku sets field value
func (o *FulfillmentShipmentItem) SetSellerSku(v string) {
	o.SellerSku = v
}

// GetSellerFulfillmentOrderItemId returns the SellerFulfillmentOrderItemId field value
func (o *FulfillmentShipmentItem) GetSellerFulfillmentOrderItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerFulfillmentOrderItemId
}

// GetSellerFulfillmentOrderItemIdOk returns a tuple with the SellerFulfillmentOrderItemId field value
// and a boolean to check if the value has been set.
func (o *FulfillmentShipmentItem) GetSellerFulfillmentOrderItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerFulfillmentOrderItemId, true
}

// SetSellerFulfillmentOrderItemId sets field value
func (o *FulfillmentShipmentItem) SetSellerFulfillmentOrderItemId(v string) {
	o.SellerFulfillmentOrderItemId = v
}

// GetQuantity returns the Quantity field value
func (o *FulfillmentShipmentItem) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *FulfillmentShipmentItem) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *FulfillmentShipmentItem) SetQuantity(v int32) {
	o.Quantity = v
}

// GetPackageNumber returns the PackageNumber field value if set, zero value otherwise.
func (o *FulfillmentShipmentItem) GetPackageNumber() int32 {
	if o == nil || IsNil(o.PackageNumber) {
		var ret int32
		return ret
	}
	return *o.PackageNumber
}

// GetPackageNumberOk returns a tuple with the PackageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentShipmentItem) GetPackageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PackageNumber) {
		return nil, false
	}
	return o.PackageNumber, true
}

// HasPackageNumber returns a boolean if a field has been set.
func (o *FulfillmentShipmentItem) HasPackageNumber() bool {
	if o != nil && !IsNil(o.PackageNumber) {
		return true
	}

	return false
}

// SetPackageNumber gets a reference to the given int32 and assigns it to the PackageNumber field.
func (o *FulfillmentShipmentItem) SetPackageNumber(v int32) {
	o.PackageNumber = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *FulfillmentShipmentItem) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentShipmentItem) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *FulfillmentShipmentItem) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *FulfillmentShipmentItem) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetManufacturerLotCodes returns the ManufacturerLotCodes field value if set, zero value otherwise.
func (o *FulfillmentShipmentItem) GetManufacturerLotCodes() []string {
	if o == nil || IsNil(o.ManufacturerLotCodes) {
		var ret []string
		return ret
	}
	return o.ManufacturerLotCodes
}

// GetManufacturerLotCodesOk returns a tuple with the ManufacturerLotCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentShipmentItem) GetManufacturerLotCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.ManufacturerLotCodes) {
		return nil, false
	}
	return o.ManufacturerLotCodes, true
}

// HasManufacturerLotCodes returns a boolean if a field has been set.
func (o *FulfillmentShipmentItem) HasManufacturerLotCodes() bool {
	if o != nil && !IsNil(o.ManufacturerLotCodes) {
		return true
	}

	return false
}

// SetManufacturerLotCodes gets a reference to the given []string and assigns it to the ManufacturerLotCodes field.
func (o *FulfillmentShipmentItem) SetManufacturerLotCodes(v []string) {
	o.ManufacturerLotCodes = v
}

func (o FulfillmentShipmentItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sellerSku"] = o.SellerSku
	toSerialize["sellerFulfillmentOrderItemId"] = o.SellerFulfillmentOrderItemId
	toSerialize["quantity"] = o.Quantity
	if !IsNil(o.PackageNumber) {
		toSerialize["packageNumber"] = o.PackageNumber
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if !IsNil(o.ManufacturerLotCodes) {
		toSerialize["manufacturerLotCodes"] = o.ManufacturerLotCodes
	}
	return toSerialize, nil
}

type NullableFulfillmentShipmentItem struct {
	value *FulfillmentShipmentItem
	isSet bool
}

func (v NullableFulfillmentShipmentItem) Get() *FulfillmentShipmentItem {
	return v.value
}

func (v *NullableFulfillmentShipmentItem) Set(val *FulfillmentShipmentItem) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentShipmentItem) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentShipmentItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentShipmentItem(val *FulfillmentShipmentItem) *NullableFulfillmentShipmentItem {
	return &NullableFulfillmentShipmentItem{value: val, isSet: true}
}

func (v NullableFulfillmentShipmentItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFulfillmentShipmentItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
