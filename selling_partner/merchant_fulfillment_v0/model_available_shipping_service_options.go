package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the AvailableShippingServiceOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailableShippingServiceOptions{}

// AvailableShippingServiceOptions The available shipping service options.
type AvailableShippingServiceOptions struct {
	// List of available carrier pickup options.
	AvailableCarrierWillPickUpOptions []AvailableCarrierWillPickUpOption `json:"AvailableCarrierWillPickUpOptions"`
	// List of available delivery experience options.
	AvailableDeliveryExperienceOptions []AvailableDeliveryExperienceOption `json:"AvailableDeliveryExperienceOptions"`
}

type _AvailableShippingServiceOptions AvailableShippingServiceOptions

// NewAvailableShippingServiceOptions instantiates a new AvailableShippingServiceOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableShippingServiceOptions(availableCarrierWillPickUpOptions []AvailableCarrierWillPickUpOption, availableDeliveryExperienceOptions []AvailableDeliveryExperienceOption) *AvailableShippingServiceOptions {
	this := AvailableShippingServiceOptions{}
	this.AvailableCarrierWillPickUpOptions = availableCarrierWillPickUpOptions
	this.AvailableDeliveryExperienceOptions = availableDeliveryExperienceOptions
	return &this
}

// NewAvailableShippingServiceOptionsWithDefaults instantiates a new AvailableShippingServiceOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableShippingServiceOptionsWithDefaults() *AvailableShippingServiceOptions {
	this := AvailableShippingServiceOptions{}
	return &this
}

// GetAvailableCarrierWillPickUpOptions returns the AvailableCarrierWillPickUpOptions field value
func (o *AvailableShippingServiceOptions) GetAvailableCarrierWillPickUpOptions() []AvailableCarrierWillPickUpOption {
	if o == nil {
		var ret []AvailableCarrierWillPickUpOption
		return ret
	}

	return o.AvailableCarrierWillPickUpOptions
}

// GetAvailableCarrierWillPickUpOptionsOk returns a tuple with the AvailableCarrierWillPickUpOptions field value
// and a boolean to check if the value has been set.
func (o *AvailableShippingServiceOptions) GetAvailableCarrierWillPickUpOptionsOk() ([]AvailableCarrierWillPickUpOption, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableCarrierWillPickUpOptions, true
}

// SetAvailableCarrierWillPickUpOptions sets field value
func (o *AvailableShippingServiceOptions) SetAvailableCarrierWillPickUpOptions(v []AvailableCarrierWillPickUpOption) {
	o.AvailableCarrierWillPickUpOptions = v
}

// GetAvailableDeliveryExperienceOptions returns the AvailableDeliveryExperienceOptions field value
func (o *AvailableShippingServiceOptions) GetAvailableDeliveryExperienceOptions() []AvailableDeliveryExperienceOption {
	if o == nil {
		var ret []AvailableDeliveryExperienceOption
		return ret
	}

	return o.AvailableDeliveryExperienceOptions
}

// GetAvailableDeliveryExperienceOptionsOk returns a tuple with the AvailableDeliveryExperienceOptions field value
// and a boolean to check if the value has been set.
func (o *AvailableShippingServiceOptions) GetAvailableDeliveryExperienceOptionsOk() ([]AvailableDeliveryExperienceOption, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableDeliveryExperienceOptions, true
}

// SetAvailableDeliveryExperienceOptions sets field value
func (o *AvailableShippingServiceOptions) SetAvailableDeliveryExperienceOptions(v []AvailableDeliveryExperienceOption) {
	o.AvailableDeliveryExperienceOptions = v
}

func (o AvailableShippingServiceOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AvailableCarrierWillPickUpOptions"] = o.AvailableCarrierWillPickUpOptions
	toSerialize["AvailableDeliveryExperienceOptions"] = o.AvailableDeliveryExperienceOptions
	return toSerialize, nil
}

type NullableAvailableShippingServiceOptions struct {
	value *AvailableShippingServiceOptions
	isSet bool
}

func (v NullableAvailableShippingServiceOptions) Get() *AvailableShippingServiceOptions {
	return v.value
}

func (v *NullableAvailableShippingServiceOptions) Set(val *AvailableShippingServiceOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableShippingServiceOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableShippingServiceOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableShippingServiceOptions(val *AvailableShippingServiceOptions) *NullableAvailableShippingServiceOptions {
	return &NullableAvailableShippingServiceOptions{value: val, isSet: true}
}

func (v NullableAvailableShippingServiceOptions) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAvailableShippingServiceOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
