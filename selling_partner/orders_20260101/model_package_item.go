package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the PackageItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageItem{}

// PackageItem Individual order item contained within a shipping package.
type PackageItem struct {
	// Unique identifier of the order item included in this package.
	OrderItemId string `json:"orderItemId"`
	// Number of units of this item included in the package shipment.
	Quantity int32 `json:"quantity"`
	// The transparency codes associated with this item for product authentication.
	TransparencyCodes []string `json:"transparencyCodes,omitempty"`
}

// NewPackageItem instantiates a new PackageItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageItem(orderItemId string, quantity int32) *PackageItem {
	this := PackageItem{}
	this.OrderItemId = orderItemId
	this.Quantity = quantity
	return &this
}

// NewPackageItemWithDefaults instantiates a new PackageItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageItemWithDefaults() *PackageItem {
	this := PackageItem{}
	return &this
}

// GetOrderItemId returns the OrderItemId field value
func (o *PackageItem) GetOrderItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderItemId
}

// GetOrderItemIdOk returns a tuple with the OrderItemId field value
// and a boolean to check if the value has been set.
func (o *PackageItem) GetOrderItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderItemId, true
}

// SetOrderItemId sets field value
func (o *PackageItem) SetOrderItemId(v string) {
	o.OrderItemId = v
}

// GetQuantity returns the Quantity field value
func (o *PackageItem) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *PackageItem) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *PackageItem) SetQuantity(v int32) {
	o.Quantity = v
}

// GetTransparencyCodes returns the TransparencyCodes field value if set, zero value otherwise.
func (o *PackageItem) GetTransparencyCodes() []string {
	if o == nil || IsNil(o.TransparencyCodes) {
		var ret []string
		return ret
	}
	return o.TransparencyCodes
}

// GetTransparencyCodesOk returns a tuple with the TransparencyCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageItem) GetTransparencyCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.TransparencyCodes) {
		return nil, false
	}
	return o.TransparencyCodes, true
}

// HasTransparencyCodes returns a boolean if a field has been set.
func (o *PackageItem) HasTransparencyCodes() bool {
	if o != nil && !IsNil(o.TransparencyCodes) {
		return true
	}

	return false
}

// SetTransparencyCodes gets a reference to the given []string and assigns it to the TransparencyCodes field.
func (o *PackageItem) SetTransparencyCodes(v []string) {
	o.TransparencyCodes = v
}

func (o PackageItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orderItemId"] = o.OrderItemId
	toSerialize["quantity"] = o.Quantity
	if !IsNil(o.TransparencyCodes) {
		toSerialize["transparencyCodes"] = o.TransparencyCodes
	}
	return toSerialize, nil
}

type NullablePackageItem struct {
	value *PackageItem
	isSet bool
}

func (v NullablePackageItem) Get() *PackageItem {
	return v.value
}

func (v *NullablePackageItem) Set(val *PackageItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageItem(val *PackageItem) *NullablePackageItem {
	return &NullablePackageItem{value: val, isSet: true}
}

func (v NullablePackageItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackageItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
