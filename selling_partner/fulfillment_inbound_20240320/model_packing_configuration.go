package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the PackingConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackingConfiguration{}

// PackingConfiguration A way to configure this packing option. Some box content information sources might not be allowed. Non-standard minimum and maximum box weights might be enforced.
type PackingConfiguration struct {
	// The box content information sources that are allowed.
	BoxPackingMethods []BoxContentInformationSource `json:"boxPackingMethods,omitempty"`
	BoxRequirements   *BoxRequirements              `json:"boxRequirements,omitempty"`
	// A list of supported shipping requirements for this packing configuration.
	ShippingRequirements []ShippingRequirements `json:"shippingRequirements,omitempty"`
}

// NewPackingConfiguration instantiates a new PackingConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackingConfiguration() *PackingConfiguration {
	this := PackingConfiguration{}
	return &this
}

// NewPackingConfigurationWithDefaults instantiates a new PackingConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackingConfigurationWithDefaults() *PackingConfiguration {
	this := PackingConfiguration{}
	return &this
}

// GetBoxPackingMethods returns the BoxPackingMethods field value if set, zero value otherwise.
func (o *PackingConfiguration) GetBoxPackingMethods() []BoxContentInformationSource {
	if o == nil || IsNil(o.BoxPackingMethods) {
		var ret []BoxContentInformationSource
		return ret
	}
	return o.BoxPackingMethods
}

// GetBoxPackingMethodsOk returns a tuple with the BoxPackingMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackingConfiguration) GetBoxPackingMethodsOk() ([]BoxContentInformationSource, bool) {
	if o == nil || IsNil(o.BoxPackingMethods) {
		return nil, false
	}
	return o.BoxPackingMethods, true
}

// HasBoxPackingMethods returns a boolean if a field has been set.
func (o *PackingConfiguration) HasBoxPackingMethods() bool {
	if o != nil && !IsNil(o.BoxPackingMethods) {
		return true
	}

	return false
}

// SetBoxPackingMethods gets a reference to the given []BoxContentInformationSource and assigns it to the BoxPackingMethods field.
func (o *PackingConfiguration) SetBoxPackingMethods(v []BoxContentInformationSource) {
	o.BoxPackingMethods = v
}

// GetBoxRequirements returns the BoxRequirements field value if set, zero value otherwise.
func (o *PackingConfiguration) GetBoxRequirements() BoxRequirements {
	if o == nil || IsNil(o.BoxRequirements) {
		var ret BoxRequirements
		return ret
	}
	return *o.BoxRequirements
}

// GetBoxRequirementsOk returns a tuple with the BoxRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackingConfiguration) GetBoxRequirementsOk() (*BoxRequirements, bool) {
	if o == nil || IsNil(o.BoxRequirements) {
		return nil, false
	}
	return o.BoxRequirements, true
}

// HasBoxRequirements returns a boolean if a field has been set.
func (o *PackingConfiguration) HasBoxRequirements() bool {
	if o != nil && !IsNil(o.BoxRequirements) {
		return true
	}

	return false
}

// SetBoxRequirements gets a reference to the given BoxRequirements and assigns it to the BoxRequirements field.
func (o *PackingConfiguration) SetBoxRequirements(v BoxRequirements) {
	o.BoxRequirements = &v
}

// GetShippingRequirements returns the ShippingRequirements field value if set, zero value otherwise.
func (o *PackingConfiguration) GetShippingRequirements() []ShippingRequirements {
	if o == nil || IsNil(o.ShippingRequirements) {
		var ret []ShippingRequirements
		return ret
	}
	return o.ShippingRequirements
}

// GetShippingRequirementsOk returns a tuple with the ShippingRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackingConfiguration) GetShippingRequirementsOk() ([]ShippingRequirements, bool) {
	if o == nil || IsNil(o.ShippingRequirements) {
		return nil, false
	}
	return o.ShippingRequirements, true
}

// HasShippingRequirements returns a boolean if a field has been set.
func (o *PackingConfiguration) HasShippingRequirements() bool {
	if o != nil && !IsNil(o.ShippingRequirements) {
		return true
	}

	return false
}

// SetShippingRequirements gets a reference to the given []ShippingRequirements and assigns it to the ShippingRequirements field.
func (o *PackingConfiguration) SetShippingRequirements(v []ShippingRequirements) {
	o.ShippingRequirements = v
}

func (o PackingConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BoxPackingMethods) {
		toSerialize["boxPackingMethods"] = o.BoxPackingMethods
	}
	if !IsNil(o.BoxRequirements) {
		toSerialize["boxRequirements"] = o.BoxRequirements
	}
	if !IsNil(o.ShippingRequirements) {
		toSerialize["shippingRequirements"] = o.ShippingRequirements
	}
	return toSerialize, nil
}

type NullablePackingConfiguration struct {
	value *PackingConfiguration
	isSet bool
}

func (v NullablePackingConfiguration) Get() *PackingConfiguration {
	return v.value
}

func (v *NullablePackingConfiguration) Set(val *PackingConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullablePackingConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullablePackingConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackingConfiguration(val *PackingConfiguration) *NullablePackingConfiguration {
	return &NullablePackingConfiguration{value: val, isSet: true}
}

func (v NullablePackingConfiguration) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackingConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
