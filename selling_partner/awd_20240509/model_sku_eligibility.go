package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the SkuEligibility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuEligibility{}

// SkuEligibility Represents eligibility of one SKU.
type SkuEligibility struct {
	// If not eligible, these are list of error codes and descriptions.
	IneligibilityReasons []SkuIneligibilityReason    `json:"ineligibilityReasons,omitempty"`
	PackageQuantity      DistributionPackageQuantity `json:"packageQuantity"`
	Status               InboundEligibilityStatus    `json:"status"`
}

type _SkuEligibility SkuEligibility

// NewSkuEligibility instantiates a new SkuEligibility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuEligibility(packageQuantity DistributionPackageQuantity, status InboundEligibilityStatus) *SkuEligibility {
	this := SkuEligibility{}
	this.PackageQuantity = packageQuantity
	this.Status = status
	return &this
}

// NewSkuEligibilityWithDefaults instantiates a new SkuEligibility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuEligibilityWithDefaults() *SkuEligibility {
	this := SkuEligibility{}
	return &this
}

// GetIneligibilityReasons returns the IneligibilityReasons field value if set, zero value otherwise.
func (o *SkuEligibility) GetIneligibilityReasons() []SkuIneligibilityReason {
	if o == nil || IsNil(o.IneligibilityReasons) {
		var ret []SkuIneligibilityReason
		return ret
	}
	return o.IneligibilityReasons
}

// GetIneligibilityReasonsOk returns a tuple with the IneligibilityReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuEligibility) GetIneligibilityReasonsOk() ([]SkuIneligibilityReason, bool) {
	if o == nil || IsNil(o.IneligibilityReasons) {
		return nil, false
	}
	return o.IneligibilityReasons, true
}

// HasIneligibilityReasons returns a boolean if a field has been set.
func (o *SkuEligibility) HasIneligibilityReasons() bool {
	if o != nil && !IsNil(o.IneligibilityReasons) {
		return true
	}

	return false
}

// SetIneligibilityReasons gets a reference to the given []SkuIneligibilityReason and assigns it to the IneligibilityReasons field.
func (o *SkuEligibility) SetIneligibilityReasons(v []SkuIneligibilityReason) {
	o.IneligibilityReasons = v
}

// GetPackageQuantity returns the PackageQuantity field value
func (o *SkuEligibility) GetPackageQuantity() DistributionPackageQuantity {
	if o == nil {
		var ret DistributionPackageQuantity
		return ret
	}

	return o.PackageQuantity
}

// GetPackageQuantityOk returns a tuple with the PackageQuantity field value
// and a boolean to check if the value has been set.
func (o *SkuEligibility) GetPackageQuantityOk() (*DistributionPackageQuantity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageQuantity, true
}

// SetPackageQuantity sets field value
func (o *SkuEligibility) SetPackageQuantity(v DistributionPackageQuantity) {
	o.PackageQuantity = v
}

// GetStatus returns the Status field value
func (o *SkuEligibility) GetStatus() InboundEligibilityStatus {
	if o == nil {
		var ret InboundEligibilityStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SkuEligibility) GetStatusOk() (*InboundEligibilityStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SkuEligibility) SetStatus(v InboundEligibilityStatus) {
	o.Status = v
}

func (o SkuEligibility) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IneligibilityReasons) {
		toSerialize["ineligibilityReasons"] = o.IneligibilityReasons
	}
	toSerialize["packageQuantity"] = o.PackageQuantity
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableSkuEligibility struct {
	value *SkuEligibility
	isSet bool
}

func (v NullableSkuEligibility) Get() *SkuEligibility {
	return v.value
}

func (v *NullableSkuEligibility) Set(val *SkuEligibility) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuEligibility) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuEligibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuEligibility(val *SkuEligibility) *NullableSkuEligibility {
	return &NullableSkuEligibility{value: val, isSet: true}
}

func (v NullableSkuEligibility) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSkuEligibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
