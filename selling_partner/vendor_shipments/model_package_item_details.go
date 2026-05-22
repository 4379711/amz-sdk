package vendor_shipments

import (
	"github.com/bytedance/sonic"
)

// checks if the PackageItemDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageItemDetails{}

// PackageItemDetails Item details for be provided for every item in shipment at either the item or carton or pallet level, whichever is appropriate.
type PackageItemDetails struct {
	// The purchase order number for the shipment being confirmed. If the items in this shipment belong to multiple purchase order numbers that are in particular carton or pallet within the shipment, then provide the purchaseOrderNumber at the appropriate carton or pallet level. Formatting Notes: 8-character alpha-numeric code.
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
	// The batch or lot number associates an item with information the manufacturer considers relevant for traceability of the trade item to which the Element String is applied. The data may refer to the trade item itself or to items contained. This field is mandatory for all perishable items.
	LotNumber *string `json:"lotNumber,omitempty"`
	Expiry    *Expiry `json:"expiry,omitempty"`
}

// NewPackageItemDetails instantiates a new PackageItemDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageItemDetails() *PackageItemDetails {
	this := PackageItemDetails{}
	return &this
}

// NewPackageItemDetailsWithDefaults instantiates a new PackageItemDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageItemDetailsWithDefaults() *PackageItemDetails {
	this := PackageItemDetails{}
	return &this
}

// GetPurchaseOrderNumber returns the PurchaseOrderNumber field value if set, zero value otherwise.
func (o *PackageItemDetails) GetPurchaseOrderNumber() string {
	if o == nil || IsNil(o.PurchaseOrderNumber) {
		var ret string
		return ret
	}
	return *o.PurchaseOrderNumber
}

// GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageItemDetails) GetPurchaseOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseOrderNumber) {
		return nil, false
	}
	return o.PurchaseOrderNumber, true
}

// HasPurchaseOrderNumber returns a boolean if a field has been set.
func (o *PackageItemDetails) HasPurchaseOrderNumber() bool {
	if o != nil && !IsNil(o.PurchaseOrderNumber) {
		return true
	}

	return false
}

// SetPurchaseOrderNumber gets a reference to the given string and assigns it to the PurchaseOrderNumber field.
func (o *PackageItemDetails) SetPurchaseOrderNumber(v string) {
	o.PurchaseOrderNumber = &v
}

// GetLotNumber returns the LotNumber field value if set, zero value otherwise.
func (o *PackageItemDetails) GetLotNumber() string {
	if o == nil || IsNil(o.LotNumber) {
		var ret string
		return ret
	}
	return *o.LotNumber
}

// GetLotNumberOk returns a tuple with the LotNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageItemDetails) GetLotNumberOk() (*string, bool) {
	if o == nil || IsNil(o.LotNumber) {
		return nil, false
	}
	return o.LotNumber, true
}

// HasLotNumber returns a boolean if a field has been set.
func (o *PackageItemDetails) HasLotNumber() bool {
	if o != nil && !IsNil(o.LotNumber) {
		return true
	}

	return false
}

// SetLotNumber gets a reference to the given string and assigns it to the LotNumber field.
func (o *PackageItemDetails) SetLotNumber(v string) {
	o.LotNumber = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *PackageItemDetails) GetExpiry() Expiry {
	if o == nil || IsNil(o.Expiry) {
		var ret Expiry
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageItemDetails) GetExpiryOk() (*Expiry, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *PackageItemDetails) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given Expiry and assigns it to the Expiry field.
func (o *PackageItemDetails) SetExpiry(v Expiry) {
	o.Expiry = &v
}

func (o PackageItemDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PurchaseOrderNumber) {
		toSerialize["purchaseOrderNumber"] = o.PurchaseOrderNumber
	}
	if !IsNil(o.LotNumber) {
		toSerialize["lotNumber"] = o.LotNumber
	}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	return toSerialize, nil
}

type NullablePackageItemDetails struct {
	value *PackageItemDetails
	isSet bool
}

func (v NullablePackageItemDetails) Get() *PackageItemDetails {
	return v.value
}

func (v *NullablePackageItemDetails) Set(val *PackageItemDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageItemDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageItemDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageItemDetails(val *PackageItemDetails) *NullablePackageItemDetails {
	return &NullablePackageItemDetails{value: val, isSet: true}
}

func (v NullablePackageItemDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackageItemDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
