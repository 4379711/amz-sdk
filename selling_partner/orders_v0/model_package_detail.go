package orders_v0

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the PackageDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageDetail{}

// PackageDetail Properties of packages
type PackageDetail struct {
	// A seller-supplied identifier that uniquely identifies a package within the scope of an order. Only positive numeric values are supported.
	PackageReferenceId string `json:"packageReferenceId"`
	// Identifies the carrier that will deliver the package. This field is required for all marketplaces. For more information, refer to the [`CarrierCode` announcement](https://developer-docs.amazon.com/sp-api/changelog/carriercode-value-required-in-shipment-confirmations-for-br-mx-ca-sg-au-in-jp-marketplaces).
	CarrierCode string `json:"carrierCode"`
	// Carrier Name that will deliver the package. Required when `carrierCode` is \"Others\"
	CarrierName *string `json:"carrierName,omitempty"`
	// Ship method to be used for shipping the order.
	ShippingMethod *string `json:"shippingMethod,omitempty"`
	// The tracking number used to obtain tracking and delivery information.
	TrackingNumber string `json:"trackingNumber"`
	// The shipping date for the package. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date/time format.
	ShipDate time.Time `json:"shipDate"`
	// The unique identifier for the supply source.
	ShipFromSupplySourceId *string `json:"shipFromSupplySourceId,omitempty"`
	// A list of order items.
	OrderItems []ConfirmShipmentOrderItem `json:"orderItems"`
}

type _PackageDetail PackageDetail

// NewPackageDetail instantiates a new PackageDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageDetail(packageReferenceId string, carrierCode string, trackingNumber string, shipDate time.Time, orderItems []ConfirmShipmentOrderItem) *PackageDetail {
	this := PackageDetail{}
	this.PackageReferenceId = packageReferenceId
	this.CarrierCode = carrierCode
	this.TrackingNumber = trackingNumber
	this.ShipDate = shipDate
	this.OrderItems = orderItems
	return &this
}

// NewPackageDetailWithDefaults instantiates a new PackageDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageDetailWithDefaults() *PackageDetail {
	this := PackageDetail{}
	return &this
}

// GetPackageReferenceId returns the PackageReferenceId field value
func (o *PackageDetail) GetPackageReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageReferenceId
}

// GetPackageReferenceIdOk returns a tuple with the PackageReferenceId field value
// and a boolean to check if the value has been set.
func (o *PackageDetail) GetPackageReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageReferenceId, true
}

// SetPackageReferenceId sets field value
func (o *PackageDetail) SetPackageReferenceId(v string) {
	o.PackageReferenceId = v
}

// GetCarrierCode returns the CarrierCode field value
func (o *PackageDetail) GetCarrierCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierCode
}

// GetCarrierCodeOk returns a tuple with the CarrierCode field value
// and a boolean to check if the value has been set.
func (o *PackageDetail) GetCarrierCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierCode, true
}

// SetCarrierCode sets field value
func (o *PackageDetail) SetCarrierCode(v string) {
	o.CarrierCode = v
}

// GetCarrierName returns the CarrierName field value if set, zero value otherwise.
func (o *PackageDetail) GetCarrierName() string {
	if o == nil || IsNil(o.CarrierName) {
		var ret string
		return ret
	}
	return *o.CarrierName
}

// GetCarrierNameOk returns a tuple with the CarrierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDetail) GetCarrierNameOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierName) {
		return nil, false
	}
	return o.CarrierName, true
}

// HasCarrierName returns a boolean if a field has been set.
func (o *PackageDetail) HasCarrierName() bool {
	if o != nil && !IsNil(o.CarrierName) {
		return true
	}

	return false
}

// SetCarrierName gets a reference to the given string and assigns it to the CarrierName field.
func (o *PackageDetail) SetCarrierName(v string) {
	o.CarrierName = &v
}

// GetShippingMethod returns the ShippingMethod field value if set, zero value otherwise.
func (o *PackageDetail) GetShippingMethod() string {
	if o == nil || IsNil(o.ShippingMethod) {
		var ret string
		return ret
	}
	return *o.ShippingMethod
}

// GetShippingMethodOk returns a tuple with the ShippingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDetail) GetShippingMethodOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingMethod) {
		return nil, false
	}
	return o.ShippingMethod, true
}

// HasShippingMethod returns a boolean if a field has been set.
func (o *PackageDetail) HasShippingMethod() bool {
	if o != nil && !IsNil(o.ShippingMethod) {
		return true
	}

	return false
}

// SetShippingMethod gets a reference to the given string and assigns it to the ShippingMethod field.
func (o *PackageDetail) SetShippingMethod(v string) {
	o.ShippingMethod = &v
}

// GetTrackingNumber returns the TrackingNumber field value
func (o *PackageDetail) GetTrackingNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value
// and a boolean to check if the value has been set.
func (o *PackageDetail) GetTrackingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrackingNumber, true
}

// SetTrackingNumber sets field value
func (o *PackageDetail) SetTrackingNumber(v string) {
	o.TrackingNumber = v
}

// GetShipDate returns the ShipDate field value
func (o *PackageDetail) GetShipDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ShipDate
}

// GetShipDateOk returns a tuple with the ShipDate field value
// and a boolean to check if the value has been set.
func (o *PackageDetail) GetShipDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipDate, true
}

// SetShipDate sets field value
func (o *PackageDetail) SetShipDate(v time.Time) {
	o.ShipDate = v
}

// GetShipFromSupplySourceId returns the ShipFromSupplySourceId field value if set, zero value otherwise.
func (o *PackageDetail) GetShipFromSupplySourceId() string {
	if o == nil || IsNil(o.ShipFromSupplySourceId) {
		var ret string
		return ret
	}
	return *o.ShipFromSupplySourceId
}

// GetShipFromSupplySourceIdOk returns a tuple with the ShipFromSupplySourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDetail) GetShipFromSupplySourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShipFromSupplySourceId) {
		return nil, false
	}
	return o.ShipFromSupplySourceId, true
}

// HasShipFromSupplySourceId returns a boolean if a field has been set.
func (o *PackageDetail) HasShipFromSupplySourceId() bool {
	if o != nil && !IsNil(o.ShipFromSupplySourceId) {
		return true
	}

	return false
}

// SetShipFromSupplySourceId gets a reference to the given string and assigns it to the ShipFromSupplySourceId field.
func (o *PackageDetail) SetShipFromSupplySourceId(v string) {
	o.ShipFromSupplySourceId = &v
}

// GetOrderItems returns the OrderItems field value
func (o *PackageDetail) GetOrderItems() []ConfirmShipmentOrderItem {
	if o == nil {
		var ret []ConfirmShipmentOrderItem
		return ret
	}

	return o.OrderItems
}

// GetOrderItemsOk returns a tuple with the OrderItems field value
// and a boolean to check if the value has been set.
func (o *PackageDetail) GetOrderItemsOk() ([]ConfirmShipmentOrderItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderItems, true
}

// SetOrderItems sets field value
func (o *PackageDetail) SetOrderItems(v []ConfirmShipmentOrderItem) {
	o.OrderItems = v
}

func (o PackageDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["packageReferenceId"] = o.PackageReferenceId
	toSerialize["carrierCode"] = o.CarrierCode
	if !IsNil(o.CarrierName) {
		toSerialize["carrierName"] = o.CarrierName
	}
	if !IsNil(o.ShippingMethod) {
		toSerialize["shippingMethod"] = o.ShippingMethod
	}
	toSerialize["trackingNumber"] = o.TrackingNumber
	toSerialize["shipDate"] = o.ShipDate
	if !IsNil(o.ShipFromSupplySourceId) {
		toSerialize["shipFromSupplySourceId"] = o.ShipFromSupplySourceId
	}
	toSerialize["orderItems"] = o.OrderItems
	return toSerialize, nil
}

type NullablePackageDetail struct {
	value *PackageDetail
	isSet bool
}

func (v NullablePackageDetail) Get() *PackageDetail {
	return v.value
}

func (v *NullablePackageDetail) Set(val *PackageDetail) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageDetail) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageDetail(val *PackageDetail) *NullablePackageDetail {
	return &NullablePackageDetail{value: val, isSet: true}
}

func (v NullablePackageDetail) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackageDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
