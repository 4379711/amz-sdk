package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ShippingConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingConfiguration{}

// ShippingConfiguration The shipping configurations supported for the packing option. Available modes are ground small parcel, freight less-than-truckload (LTL), freight full-truckload (FTL) palletized, freight FTL non-palletized, ocean less-than-container-load (LCL), ocean full-container load (FCL), air small parcel, and air small parcel express.
type ShippingConfiguration struct {
	// Mode of shipment transportation that this option will provide.  Possible values: `GROUND_SMALL_PARCEL`, `FREIGHT_LTL`, `FREIGHT_FTL_PALLET`, `FREIGHT_FTL_NONPALLET`, `OCEAN_LCL`, `OCEAN_FCL`, `AIR_SMALL_PARCEL`, `AIR_SMALL_PARCEL_EXPRESS`.
	ShippingMode *string `json:"shippingMode,omitempty"`
	// Shipping program for the option. Possible values: `AMAZON_PARTNERED_CARRIER`, `USE_YOUR_OWN_CARRIER`.
	ShippingSolution *string `json:"shippingSolution,omitempty"`
}

// NewShippingConfiguration instantiates a new ShippingConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingConfiguration() *ShippingConfiguration {
	this := ShippingConfiguration{}
	return &this
}

// NewShippingConfigurationWithDefaults instantiates a new ShippingConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingConfigurationWithDefaults() *ShippingConfiguration {
	this := ShippingConfiguration{}
	return &this
}

// GetShippingMode returns the ShippingMode field value if set, zero value otherwise.
func (o *ShippingConfiguration) GetShippingMode() string {
	if o == nil || IsNil(o.ShippingMode) {
		var ret string
		return ret
	}
	return *o.ShippingMode
}

// GetShippingModeOk returns a tuple with the ShippingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingConfiguration) GetShippingModeOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingMode) {
		return nil, false
	}
	return o.ShippingMode, true
}

// HasShippingMode returns a boolean if a field has been set.
func (o *ShippingConfiguration) HasShippingMode() bool {
	if o != nil && !IsNil(o.ShippingMode) {
		return true
	}

	return false
}

// SetShippingMode gets a reference to the given string and assigns it to the ShippingMode field.
func (o *ShippingConfiguration) SetShippingMode(v string) {
	o.ShippingMode = &v
}

// GetShippingSolution returns the ShippingSolution field value if set, zero value otherwise.
func (o *ShippingConfiguration) GetShippingSolution() string {
	if o == nil || IsNil(o.ShippingSolution) {
		var ret string
		return ret
	}
	return *o.ShippingSolution
}

// GetShippingSolutionOk returns a tuple with the ShippingSolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingConfiguration) GetShippingSolutionOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingSolution) {
		return nil, false
	}
	return o.ShippingSolution, true
}

// HasShippingSolution returns a boolean if a field has been set.
func (o *ShippingConfiguration) HasShippingSolution() bool {
	if o != nil && !IsNil(o.ShippingSolution) {
		return true
	}

	return false
}

// SetShippingSolution gets a reference to the given string and assigns it to the ShippingSolution field.
func (o *ShippingConfiguration) SetShippingSolution(v string) {
	o.ShippingSolution = &v
}

func (o ShippingConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShippingMode) {
		toSerialize["shippingMode"] = o.ShippingMode
	}
	if !IsNil(o.ShippingSolution) {
		toSerialize["shippingSolution"] = o.ShippingSolution
	}
	return toSerialize, nil
}

type NullableShippingConfiguration struct {
	value *ShippingConfiguration
	isSet bool
}

func (v NullableShippingConfiguration) Get() *ShippingConfiguration {
	return v.value
}

func (v *NullableShippingConfiguration) Set(val *ShippingConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingConfiguration(val *ShippingConfiguration) *NullableShippingConfiguration {
	return &NullableShippingConfiguration{value: val, isSet: true}
}

func (v NullableShippingConfiguration) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShippingConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
