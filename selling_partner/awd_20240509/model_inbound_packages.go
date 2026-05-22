package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the InboundPackages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InboundPackages{}

// InboundPackages Represents the packages to inbound.
type InboundPackages struct {
	// List of packages to be inbounded.
	PackagesToInbound []DistributionPackageQuantity `json:"packagesToInbound"`
}

type _InboundPackages InboundPackages

// NewInboundPackages instantiates a new InboundPackages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundPackages(packagesToInbound []DistributionPackageQuantity) *InboundPackages {
	this := InboundPackages{}
	this.PackagesToInbound = packagesToInbound
	return &this
}

// NewInboundPackagesWithDefaults instantiates a new InboundPackages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundPackagesWithDefaults() *InboundPackages {
	this := InboundPackages{}
	return &this
}

// GetPackagesToInbound returns the PackagesToInbound field value
func (o *InboundPackages) GetPackagesToInbound() []DistributionPackageQuantity {
	if o == nil {
		var ret []DistributionPackageQuantity
		return ret
	}

	return o.PackagesToInbound
}

// GetPackagesToInboundOk returns a tuple with the PackagesToInbound field value
// and a boolean to check if the value has been set.
func (o *InboundPackages) GetPackagesToInboundOk() ([]DistributionPackageQuantity, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackagesToInbound, true
}

// SetPackagesToInbound sets field value
func (o *InboundPackages) SetPackagesToInbound(v []DistributionPackageQuantity) {
	o.PackagesToInbound = v
}

func (o InboundPackages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["packagesToInbound"] = o.PackagesToInbound
	return toSerialize, nil
}

type NullableInboundPackages struct {
	value *InboundPackages
	isSet bool
}

func (v NullableInboundPackages) Get() *InboundPackages {
	return v.value
}

func (v *NullableInboundPackages) Set(val *InboundPackages) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundPackages) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundPackages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundPackages(val *InboundPackages) *NullableInboundPackages {
	return &NullableInboundPackages{value: val, isSet: true}
}

func (v NullableInboundPackages) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInboundPackages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
