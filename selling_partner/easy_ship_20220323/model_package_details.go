package easy_ship_20220323

import (
	"github.com/bytedance/sonic"
)

// checks if the PackageDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageDetails{}

// PackageDetails Package details. Includes `packageItems`, `packageTimeSlot`, and `packageIdentifier`.
type PackageDetails struct {
	// A list of items contained in the package.
	PackageItems    []Item   `json:"packageItems,omitempty"`
	PackageTimeSlot TimeSlot `json:"packageTimeSlot"`
	// Optional seller-created identifier that is printed on the shipping label to help the seller identify the package.
	PackageIdentifier *string `json:"packageIdentifier,omitempty"`
}

type _PackageDetails PackageDetails

// NewPackageDetails instantiates a new PackageDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageDetails(packageTimeSlot TimeSlot) *PackageDetails {
	this := PackageDetails{}
	this.PackageTimeSlot = packageTimeSlot
	return &this
}

// NewPackageDetailsWithDefaults instantiates a new PackageDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageDetailsWithDefaults() *PackageDetails {
	this := PackageDetails{}
	return &this
}

// GetPackageItems returns the PackageItems field value if set, zero value otherwise.
func (o *PackageDetails) GetPackageItems() []Item {
	if o == nil || IsNil(o.PackageItems) {
		var ret []Item
		return ret
	}
	return o.PackageItems
}

// GetPackageItemsOk returns a tuple with the PackageItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDetails) GetPackageItemsOk() ([]Item, bool) {
	if o == nil || IsNil(o.PackageItems) {
		return nil, false
	}
	return o.PackageItems, true
}

// HasPackageItems returns a boolean if a field has been set.
func (o *PackageDetails) HasPackageItems() bool {
	if o != nil && !IsNil(o.PackageItems) {
		return true
	}

	return false
}

// SetPackageItems gets a reference to the given []Item and assigns it to the PackageItems field.
func (o *PackageDetails) SetPackageItems(v []Item) {
	o.PackageItems = v
}

// GetPackageTimeSlot returns the PackageTimeSlot field value
func (o *PackageDetails) GetPackageTimeSlot() TimeSlot {
	if o == nil {
		var ret TimeSlot
		return ret
	}

	return o.PackageTimeSlot
}

// GetPackageTimeSlotOk returns a tuple with the PackageTimeSlot field value
// and a boolean to check if the value has been set.
func (o *PackageDetails) GetPackageTimeSlotOk() (*TimeSlot, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageTimeSlot, true
}

// SetPackageTimeSlot sets field value
func (o *PackageDetails) SetPackageTimeSlot(v TimeSlot) {
	o.PackageTimeSlot = v
}

// GetPackageIdentifier returns the PackageIdentifier field value if set, zero value otherwise.
func (o *PackageDetails) GetPackageIdentifier() string {
	if o == nil || IsNil(o.PackageIdentifier) {
		var ret string
		return ret
	}
	return *o.PackageIdentifier
}

// GetPackageIdentifierOk returns a tuple with the PackageIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDetails) GetPackageIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.PackageIdentifier) {
		return nil, false
	}
	return o.PackageIdentifier, true
}

// HasPackageIdentifier returns a boolean if a field has been set.
func (o *PackageDetails) HasPackageIdentifier() bool {
	if o != nil && !IsNil(o.PackageIdentifier) {
		return true
	}

	return false
}

// SetPackageIdentifier gets a reference to the given string and assigns it to the PackageIdentifier field.
func (o *PackageDetails) SetPackageIdentifier(v string) {
	o.PackageIdentifier = &v
}

func (o PackageDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PackageItems) {
		toSerialize["packageItems"] = o.PackageItems
	}
	toSerialize["packageTimeSlot"] = o.PackageTimeSlot
	if !IsNil(o.PackageIdentifier) {
		toSerialize["packageIdentifier"] = o.PackageIdentifier
	}
	return toSerialize, nil
}

type NullablePackageDetails struct {
	value *PackageDetails
	isSet bool
}

func (v NullablePackageDetails) Get() *PackageDetails {
	return v.value
}

func (v *NullablePackageDetails) Set(val *PackageDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageDetails(val *PackageDetails) *NullablePackageDetails {
	return &NullablePackageDetails{value: val, isSet: true}
}

func (v NullablePackageDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackageDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
