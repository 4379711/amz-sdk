package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the PackageStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageStatus{}

// PackageStatus Current status and detailed tracking information for a shipping package throughout the delivery process.
type PackageStatus struct {
	// Primary status classification of the package in the shipping workflow.
	Status string `json:"status"`
	// Granular status information providing specific details about the package's current location and handling stage.   **Possible values**: - `PENDING_SCHEDULE` (Package awaiting pickup scheduling) - `PENDING_PICK_UP` (Package ready for carrier collection from seller) - `PENDING_DROP_OFF` (Package awaiting seller delivery to carrier) - `LABEL_CANCELLED` (Shipping label canceled by seller) - `PICKED_UP` (Package collected by carrier from seller location) - `DROPPED_OFF` (Package delivered to carrier by seller) - `AT_ORIGIN_FC` (Package at originating fulfillment center) - `AT_DESTINATION_FC` (Package at destination fulfillment center) - `DELIVERED` (Package successfully delivered to recipient) - `REJECTED_BY_BUYER` (Package refused by intended recipient) - `UNDELIVERABLE` (Package cannot be delivered due to address or access issues) - `RETURNING_TO_SELLER` (Package in transit back to seller) - `RETURNED_TO_SELLER` (Package successfully returned to seller) - `LOST` (Package location unknown or confirmed lost) - `OUT_FOR_DELIVERY` (Package on delivery vehicle for final delivery) - `DAMAGED` (Package damaged during transit)
	DetailedStatus *string `json:"detailedStatus,omitempty"`
}

// NewPackageStatus instantiates a new PackageStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageStatus(status string) *PackageStatus {
	this := PackageStatus{}
	this.Status = status
	return &this
}

// NewPackageStatusWithDefaults instantiates a new PackageStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageStatusWithDefaults() *PackageStatus {
	this := PackageStatus{}
	return &this
}

// GetStatus returns the Status field value
func (o *PackageStatus) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PackageStatus) SetStatus(v string) {
	o.Status = v
}

// GetDetailedStatus returns the DetailedStatus field value if set, zero value otherwise.
func (o *PackageStatus) GetDetailedStatus() string {
	if o == nil || IsNil(o.DetailedStatus) {
		var ret string
		return ret
	}
	return *o.DetailedStatus
}

// GetDetailedStatusOk returns a tuple with the DetailedStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetDetailedStatusOk() (*string, bool) {
	if o == nil || IsNil(o.DetailedStatus) {
		return nil, false
	}
	return o.DetailedStatus, true
}

// HasDetailedStatus returns a boolean if a field has been set.
func (o *PackageStatus) HasDetailedStatus() bool {
	if o != nil && !IsNil(o.DetailedStatus) {
		return true
	}

	return false
}

// SetDetailedStatus gets a reference to the given string and assigns it to the DetailedStatus field.
func (o *PackageStatus) SetDetailedStatus(v string) {
	o.DetailedStatus = &v
}

func (o PackageStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if !IsNil(o.DetailedStatus) {
		toSerialize["detailedStatus"] = o.DetailedStatus
	}
	return toSerialize, nil
}

type NullablePackageStatus struct {
	value *PackageStatus
	isSet bool
}

func (v NullablePackageStatus) Get() *PackageStatus {
	return v.value
}

func (v *NullablePackageStatus) Set(val *PackageStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageStatus(val *PackageStatus) *NullablePackageStatus {
	return &NullablePackageStatus{value: val, isSet: true}
}

func (v NullablePackageStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackageStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
