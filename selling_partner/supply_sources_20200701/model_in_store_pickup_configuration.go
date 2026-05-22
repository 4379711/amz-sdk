package supply_sources_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the InStorePickupConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InStorePickupConfiguration{}

// InStorePickupConfiguration The in-store pickup configuration of a supply source.
type InStorePickupConfiguration struct {
	// When true, in-store pickup is supported by the supply source (default: `isSupported` value in `PickupChannel`).
	IsSupported          *bool                 `json:"isSupported,omitempty"`
	ParkingConfiguration *ParkingConfiguration `json:"parkingConfiguration,omitempty"`
}

// NewInStorePickupConfiguration instantiates a new InStorePickupConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInStorePickupConfiguration() *InStorePickupConfiguration {
	this := InStorePickupConfiguration{}
	return &this
}

// NewInStorePickupConfigurationWithDefaults instantiates a new InStorePickupConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInStorePickupConfigurationWithDefaults() *InStorePickupConfiguration {
	this := InStorePickupConfiguration{}
	return &this
}

// GetIsSupported returns the IsSupported field value if set, zero value otherwise.
func (o *InStorePickupConfiguration) GetIsSupported() bool {
	if o == nil || IsNil(o.IsSupported) {
		var ret bool
		return ret
	}
	return *o.IsSupported
}

// GetIsSupportedOk returns a tuple with the IsSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStorePickupConfiguration) GetIsSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSupported) {
		return nil, false
	}
	return o.IsSupported, true
}

// HasIsSupported returns a boolean if a field has been set.
func (o *InStorePickupConfiguration) HasIsSupported() bool {
	if o != nil && !IsNil(o.IsSupported) {
		return true
	}

	return false
}

// SetIsSupported gets a reference to the given bool and assigns it to the IsSupported field.
func (o *InStorePickupConfiguration) SetIsSupported(v bool) {
	o.IsSupported = &v
}

// GetParkingConfiguration returns the ParkingConfiguration field value if set, zero value otherwise.
func (o *InStorePickupConfiguration) GetParkingConfiguration() ParkingConfiguration {
	if o == nil || IsNil(o.ParkingConfiguration) {
		var ret ParkingConfiguration
		return ret
	}
	return *o.ParkingConfiguration
}

// GetParkingConfigurationOk returns a tuple with the ParkingConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStorePickupConfiguration) GetParkingConfigurationOk() (*ParkingConfiguration, bool) {
	if o == nil || IsNil(o.ParkingConfiguration) {
		return nil, false
	}
	return o.ParkingConfiguration, true
}

// HasParkingConfiguration returns a boolean if a field has been set.
func (o *InStorePickupConfiguration) HasParkingConfiguration() bool {
	if o != nil && !IsNil(o.ParkingConfiguration) {
		return true
	}

	return false
}

// SetParkingConfiguration gets a reference to the given ParkingConfiguration and assigns it to the ParkingConfiguration field.
func (o *InStorePickupConfiguration) SetParkingConfiguration(v ParkingConfiguration) {
	o.ParkingConfiguration = &v
}

func (o InStorePickupConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsSupported) {
		toSerialize["isSupported"] = o.IsSupported
	}
	if !IsNil(o.ParkingConfiguration) {
		toSerialize["parkingConfiguration"] = o.ParkingConfiguration
	}
	return toSerialize, nil
}

type NullableInStorePickupConfiguration struct {
	value *InStorePickupConfiguration
	isSet bool
}

func (v NullableInStorePickupConfiguration) Get() *InStorePickupConfiguration {
	return v.value
}

func (v *NullableInStorePickupConfiguration) Set(val *InStorePickupConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableInStorePickupConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableInStorePickupConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInStorePickupConfiguration(val *InStorePickupConfiguration) *NullableInStorePickupConfiguration {
	return &NullableInStorePickupConfiguration{value: val, isSet: true}
}

func (v NullableInStorePickupConfiguration) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInStorePickupConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
