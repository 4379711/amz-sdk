package vendor_direct_fulfillment_shipping_20211228

import (
	"github.com/bytedance/sonic"
)

// checks if the Package type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Package{}

// Package The package that is associated with the container.
type Package struct {
	// The tracking number on the label of shipment package, that you can fetch from the `shippingLabels` response. You can also scan the bar code on the shipping label to get the tracking number.
	PackageTrackingNumber string `json:"packageTrackingNumber"`
}

type _Package Package

// NewPackage instantiates a new Package object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackage(packageTrackingNumber string) *Package {
	this := Package{}
	this.PackageTrackingNumber = packageTrackingNumber
	return &this
}

// NewPackageWithDefaults instantiates a new Package object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageWithDefaults() *Package {
	this := Package{}
	return &this
}

// GetPackageTrackingNumber returns the PackageTrackingNumber field value
func (o *Package) GetPackageTrackingNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageTrackingNumber
}

// GetPackageTrackingNumberOk returns a tuple with the PackageTrackingNumber field value
// and a boolean to check if the value has been set.
func (o *Package) GetPackageTrackingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageTrackingNumber, true
}

// SetPackageTrackingNumber sets field value
func (o *Package) SetPackageTrackingNumber(v string) {
	o.PackageTrackingNumber = v
}

func (o Package) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["packageTrackingNumber"] = o.PackageTrackingNumber
	return toSerialize, nil
}

type NullablePackage struct {
	value *Package
	isSet bool
}

func (v NullablePackage) Get() *Package {
	return v.value
}

func (v *NullablePackage) Set(val *Package) {
	v.value = val
	v.isSet = true
}

func (v NullablePackage) IsSet() bool {
	return v.isSet
}

func (v *NullablePackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackage(val *Package) *NullablePackage {
	return &NullablePackage{value: val, isSet: true}
}

func (v NullablePackage) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
