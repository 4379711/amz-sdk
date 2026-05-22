package easy_ship_20220323

import (
	"github.com/bytedance/sonic"
)

// checks if the Package type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Package{}

// Package This object contains all the details of the scheduled Easy Ship package.
type Package struct {
	ScheduledPackageId ScheduledPackageId `json:"scheduledPackageId"`
	PackageDimensions  Dimensions         `json:"packageDimensions"`
	PackageWeight      Weight             `json:"packageWeight"`
	// A list of items contained in the package.
	PackageItems    []Item   `json:"packageItems,omitempty"`
	PackageTimeSlot TimeSlot `json:"packageTimeSlot"`
	// Optional seller-created identifier that is printed on the shipping label to help the seller identify the package.
	PackageIdentifier *string          `json:"packageIdentifier,omitempty"`
	Invoice           *InvoiceData     `json:"invoice,omitempty"`
	PackageStatus     *PackageStatus   `json:"packageStatus,omitempty"`
	TrackingDetails   *TrackingDetails `json:"trackingDetails,omitempty"`
}

type _Package Package

// NewPackage instantiates a new Package object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackage(scheduledPackageId ScheduledPackageId, packageDimensions Dimensions, packageWeight Weight, packageTimeSlot TimeSlot) *Package {
	this := Package{}
	this.ScheduledPackageId = scheduledPackageId
	this.PackageDimensions = packageDimensions
	this.PackageWeight = packageWeight
	this.PackageTimeSlot = packageTimeSlot
	return &this
}

// NewPackageWithDefaults instantiates a new Package object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageWithDefaults() *Package {
	this := Package{}
	return &this
}

// GetScheduledPackageId returns the ScheduledPackageId field value
func (o *Package) GetScheduledPackageId() ScheduledPackageId {
	if o == nil {
		var ret ScheduledPackageId
		return ret
	}

	return o.ScheduledPackageId
}

// GetScheduledPackageIdOk returns a tuple with the ScheduledPackageId field value
// and a boolean to check if the value has been set.
func (o *Package) GetScheduledPackageIdOk() (*ScheduledPackageId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledPackageId, true
}

// SetScheduledPackageId sets field value
func (o *Package) SetScheduledPackageId(v ScheduledPackageId) {
	o.ScheduledPackageId = v
}

// GetPackageDimensions returns the PackageDimensions field value
func (o *Package) GetPackageDimensions() Dimensions {
	if o == nil {
		var ret Dimensions
		return ret
	}

	return o.PackageDimensions
}

// GetPackageDimensionsOk returns a tuple with the PackageDimensions field value
// and a boolean to check if the value has been set.
func (o *Package) GetPackageDimensionsOk() (*Dimensions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageDimensions, true
}

// SetPackageDimensions sets field value
func (o *Package) SetPackageDimensions(v Dimensions) {
	o.PackageDimensions = v
}

// GetPackageWeight returns the PackageWeight field value
func (o *Package) GetPackageWeight() Weight {
	if o == nil {
		var ret Weight
		return ret
	}

	return o.PackageWeight
}

// GetPackageWeightOk returns a tuple with the PackageWeight field value
// and a boolean to check if the value has been set.
func (o *Package) GetPackageWeightOk() (*Weight, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageWeight, true
}

// SetPackageWeight sets field value
func (o *Package) SetPackageWeight(v Weight) {
	o.PackageWeight = v
}

// GetPackageItems returns the PackageItems field value if set, zero value otherwise.
func (o *Package) GetPackageItems() []Item {
	if o == nil || IsNil(o.PackageItems) {
		var ret []Item
		return ret
	}
	return o.PackageItems
}

// GetPackageItemsOk returns a tuple with the PackageItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetPackageItemsOk() ([]Item, bool) {
	if o == nil || IsNil(o.PackageItems) {
		return nil, false
	}
	return o.PackageItems, true
}

// HasPackageItems returns a boolean if a field has been set.
func (o *Package) HasPackageItems() bool {
	if o != nil && !IsNil(o.PackageItems) {
		return true
	}

	return false
}

// SetPackageItems gets a reference to the given []Item and assigns it to the PackageItems field.
func (o *Package) SetPackageItems(v []Item) {
	o.PackageItems = v
}

// GetPackageTimeSlot returns the PackageTimeSlot field value
func (o *Package) GetPackageTimeSlot() TimeSlot {
	if o == nil {
		var ret TimeSlot
		return ret
	}

	return o.PackageTimeSlot
}

// GetPackageTimeSlotOk returns a tuple with the PackageTimeSlot field value
// and a boolean to check if the value has been set.
func (o *Package) GetPackageTimeSlotOk() (*TimeSlot, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageTimeSlot, true
}

// SetPackageTimeSlot sets field value
func (o *Package) SetPackageTimeSlot(v TimeSlot) {
	o.PackageTimeSlot = v
}

// GetPackageIdentifier returns the PackageIdentifier field value if set, zero value otherwise.
func (o *Package) GetPackageIdentifier() string {
	if o == nil || IsNil(o.PackageIdentifier) {
		var ret string
		return ret
	}
	return *o.PackageIdentifier
}

// GetPackageIdentifierOk returns a tuple with the PackageIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetPackageIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.PackageIdentifier) {
		return nil, false
	}
	return o.PackageIdentifier, true
}

// HasPackageIdentifier returns a boolean if a field has been set.
func (o *Package) HasPackageIdentifier() bool {
	if o != nil && !IsNil(o.PackageIdentifier) {
		return true
	}

	return false
}

// SetPackageIdentifier gets a reference to the given string and assigns it to the PackageIdentifier field.
func (o *Package) SetPackageIdentifier(v string) {
	o.PackageIdentifier = &v
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *Package) GetInvoice() InvoiceData {
	if o == nil || IsNil(o.Invoice) {
		var ret InvoiceData
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetInvoiceOk() (*InvoiceData, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *Package) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given InvoiceData and assigns it to the Invoice field.
func (o *Package) SetInvoice(v InvoiceData) {
	o.Invoice = &v
}

// GetPackageStatus returns the PackageStatus field value if set, zero value otherwise.
func (o *Package) GetPackageStatus() PackageStatus {
	if o == nil || IsNil(o.PackageStatus) {
		var ret PackageStatus
		return ret
	}
	return *o.PackageStatus
}

// GetPackageStatusOk returns a tuple with the PackageStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetPackageStatusOk() (*PackageStatus, bool) {
	if o == nil || IsNil(o.PackageStatus) {
		return nil, false
	}
	return o.PackageStatus, true
}

// HasPackageStatus returns a boolean if a field has been set.
func (o *Package) HasPackageStatus() bool {
	if o != nil && !IsNil(o.PackageStatus) {
		return true
	}

	return false
}

// SetPackageStatus gets a reference to the given PackageStatus and assigns it to the PackageStatus field.
func (o *Package) SetPackageStatus(v PackageStatus) {
	o.PackageStatus = &v
}

// GetTrackingDetails returns the TrackingDetails field value if set, zero value otherwise.
func (o *Package) GetTrackingDetails() TrackingDetails {
	if o == nil || IsNil(o.TrackingDetails) {
		var ret TrackingDetails
		return ret
	}
	return *o.TrackingDetails
}

// GetTrackingDetailsOk returns a tuple with the TrackingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetTrackingDetailsOk() (*TrackingDetails, bool) {
	if o == nil || IsNil(o.TrackingDetails) {
		return nil, false
	}
	return o.TrackingDetails, true
}

// HasTrackingDetails returns a boolean if a field has been set.
func (o *Package) HasTrackingDetails() bool {
	if o != nil && !IsNil(o.TrackingDetails) {
		return true
	}

	return false
}

// SetTrackingDetails gets a reference to the given TrackingDetails and assigns it to the TrackingDetails field.
func (o *Package) SetTrackingDetails(v TrackingDetails) {
	o.TrackingDetails = &v
}

func (o Package) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["scheduledPackageId"] = o.ScheduledPackageId
	toSerialize["packageDimensions"] = o.PackageDimensions
	toSerialize["packageWeight"] = o.PackageWeight
	if !IsNil(o.PackageItems) {
		toSerialize["packageItems"] = o.PackageItems
	}
	toSerialize["packageTimeSlot"] = o.PackageTimeSlot
	if !IsNil(o.PackageIdentifier) {
		toSerialize["packageIdentifier"] = o.PackageIdentifier
	}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	if !IsNil(o.PackageStatus) {
		toSerialize["packageStatus"] = o.PackageStatus
	}
	if !IsNil(o.TrackingDetails) {
		toSerialize["trackingDetails"] = o.TrackingDetails
	}
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
