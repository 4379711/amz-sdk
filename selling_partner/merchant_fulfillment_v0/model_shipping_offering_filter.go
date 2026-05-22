package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ShippingOfferingFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingOfferingFilter{}

// ShippingOfferingFilter Filter for use when requesting eligible shipping services.
type ShippingOfferingFilter struct {
	// When true, include a packing slip with the label.
	IncludePackingSlipWithLabel *bool `json:"IncludePackingSlipWithLabel,omitempty"`
	// When true, include complex shipping options.
	IncludeComplexShippingOptions *bool                     `json:"IncludeComplexShippingOptions,omitempty"`
	CarrierWillPickUp             *CarrierWillPickUpOption  `json:"CarrierWillPickUp,omitempty"`
	DeliveryExperience            *DeliveryExperienceOption `json:"DeliveryExperience,omitempty"`
}

// NewShippingOfferingFilter instantiates a new ShippingOfferingFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingOfferingFilter() *ShippingOfferingFilter {
	this := ShippingOfferingFilter{}
	return &this
}

// NewShippingOfferingFilterWithDefaults instantiates a new ShippingOfferingFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingOfferingFilterWithDefaults() *ShippingOfferingFilter {
	this := ShippingOfferingFilter{}
	return &this
}

// GetIncludePackingSlipWithLabel returns the IncludePackingSlipWithLabel field value if set, zero value otherwise.
func (o *ShippingOfferingFilter) GetIncludePackingSlipWithLabel() bool {
	if o == nil || IsNil(o.IncludePackingSlipWithLabel) {
		var ret bool
		return ret
	}
	return *o.IncludePackingSlipWithLabel
}

// GetIncludePackingSlipWithLabelOk returns a tuple with the IncludePackingSlipWithLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOfferingFilter) GetIncludePackingSlipWithLabelOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludePackingSlipWithLabel) {
		return nil, false
	}
	return o.IncludePackingSlipWithLabel, true
}

// HasIncludePackingSlipWithLabel returns a boolean if a field has been set.
func (o *ShippingOfferingFilter) HasIncludePackingSlipWithLabel() bool {
	if o != nil && !IsNil(o.IncludePackingSlipWithLabel) {
		return true
	}

	return false
}

// SetIncludePackingSlipWithLabel gets a reference to the given bool and assigns it to the IncludePackingSlipWithLabel field.
func (o *ShippingOfferingFilter) SetIncludePackingSlipWithLabel(v bool) {
	o.IncludePackingSlipWithLabel = &v
}

// GetIncludeComplexShippingOptions returns the IncludeComplexShippingOptions field value if set, zero value otherwise.
func (o *ShippingOfferingFilter) GetIncludeComplexShippingOptions() bool {
	if o == nil || IsNil(o.IncludeComplexShippingOptions) {
		var ret bool
		return ret
	}
	return *o.IncludeComplexShippingOptions
}

// GetIncludeComplexShippingOptionsOk returns a tuple with the IncludeComplexShippingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOfferingFilter) GetIncludeComplexShippingOptionsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeComplexShippingOptions) {
		return nil, false
	}
	return o.IncludeComplexShippingOptions, true
}

// HasIncludeComplexShippingOptions returns a boolean if a field has been set.
func (o *ShippingOfferingFilter) HasIncludeComplexShippingOptions() bool {
	if o != nil && !IsNil(o.IncludeComplexShippingOptions) {
		return true
	}

	return false
}

// SetIncludeComplexShippingOptions gets a reference to the given bool and assigns it to the IncludeComplexShippingOptions field.
func (o *ShippingOfferingFilter) SetIncludeComplexShippingOptions(v bool) {
	o.IncludeComplexShippingOptions = &v
}

// GetCarrierWillPickUp returns the CarrierWillPickUp field value if set, zero value otherwise.
func (o *ShippingOfferingFilter) GetCarrierWillPickUp() CarrierWillPickUpOption {
	if o == nil || IsNil(o.CarrierWillPickUp) {
		var ret CarrierWillPickUpOption
		return ret
	}
	return *o.CarrierWillPickUp
}

// GetCarrierWillPickUpOk returns a tuple with the CarrierWillPickUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOfferingFilter) GetCarrierWillPickUpOk() (*CarrierWillPickUpOption, bool) {
	if o == nil || IsNil(o.CarrierWillPickUp) {
		return nil, false
	}
	return o.CarrierWillPickUp, true
}

// HasCarrierWillPickUp returns a boolean if a field has been set.
func (o *ShippingOfferingFilter) HasCarrierWillPickUp() bool {
	if o != nil && !IsNil(o.CarrierWillPickUp) {
		return true
	}

	return false
}

// SetCarrierWillPickUp gets a reference to the given CarrierWillPickUpOption and assigns it to the CarrierWillPickUp field.
func (o *ShippingOfferingFilter) SetCarrierWillPickUp(v CarrierWillPickUpOption) {
	o.CarrierWillPickUp = &v
}

// GetDeliveryExperience returns the DeliveryExperience field value if set, zero value otherwise.
func (o *ShippingOfferingFilter) GetDeliveryExperience() DeliveryExperienceOption {
	if o == nil || IsNil(o.DeliveryExperience) {
		var ret DeliveryExperienceOption
		return ret
	}
	return *o.DeliveryExperience
}

// GetDeliveryExperienceOk returns a tuple with the DeliveryExperience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOfferingFilter) GetDeliveryExperienceOk() (*DeliveryExperienceOption, bool) {
	if o == nil || IsNil(o.DeliveryExperience) {
		return nil, false
	}
	return o.DeliveryExperience, true
}

// HasDeliveryExperience returns a boolean if a field has been set.
func (o *ShippingOfferingFilter) HasDeliveryExperience() bool {
	if o != nil && !IsNil(o.DeliveryExperience) {
		return true
	}

	return false
}

// SetDeliveryExperience gets a reference to the given DeliveryExperienceOption and assigns it to the DeliveryExperience field.
func (o *ShippingOfferingFilter) SetDeliveryExperience(v DeliveryExperienceOption) {
	o.DeliveryExperience = &v
}

func (o ShippingOfferingFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncludePackingSlipWithLabel) {
		toSerialize["IncludePackingSlipWithLabel"] = o.IncludePackingSlipWithLabel
	}
	if !IsNil(o.IncludeComplexShippingOptions) {
		toSerialize["IncludeComplexShippingOptions"] = o.IncludeComplexShippingOptions
	}
	if !IsNil(o.CarrierWillPickUp) {
		toSerialize["CarrierWillPickUp"] = o.CarrierWillPickUp
	}
	if !IsNil(o.DeliveryExperience) {
		toSerialize["DeliveryExperience"] = o.DeliveryExperience
	}
	return toSerialize, nil
}

type NullableShippingOfferingFilter struct {
	value *ShippingOfferingFilter
	isSet bool
}

func (v NullableShippingOfferingFilter) Get() *ShippingOfferingFilter {
	return v.value
}

func (v *NullableShippingOfferingFilter) Set(val *ShippingOfferingFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingOfferingFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingOfferingFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingOfferingFilter(val *ShippingOfferingFilter) *NullableShippingOfferingFilter {
	return &NullableShippingOfferingFilter{value: val, isSet: true}
}

func (v NullableShippingOfferingFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShippingOfferingFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
