package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ShippingServiceOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingServiceOptions{}

// ShippingServiceOptions Extra services provided by a carrier.
type ShippingServiceOptions struct {
	DeliveryExperience DeliveryExperienceType `json:"DeliveryExperience"`
	DeclaredValue      *CurrencyAmount        `json:"DeclaredValue,omitempty"`
	// When true, the carrier will pick up the package. Note: Scheduled carrier pickup is available only using Dynamex (US), DPD (UK), and Royal Mail (UK).
	CarrierWillPickUp       bool                     `json:"CarrierWillPickUp"`
	CarrierWillPickUpOption *CarrierWillPickUpOption `json:"CarrierWillPickUpOption,omitempty"`
	LabelFormat             *LabelFormat             `json:"LabelFormat,omitempty"`
}

type _ShippingServiceOptions ShippingServiceOptions

// NewShippingServiceOptions instantiates a new ShippingServiceOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingServiceOptions(deliveryExperience DeliveryExperienceType, carrierWillPickUp bool) *ShippingServiceOptions {
	this := ShippingServiceOptions{}
	this.DeliveryExperience = deliveryExperience
	this.CarrierWillPickUp = carrierWillPickUp
	return &this
}

// NewShippingServiceOptionsWithDefaults instantiates a new ShippingServiceOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingServiceOptionsWithDefaults() *ShippingServiceOptions {
	this := ShippingServiceOptions{}
	return &this
}

// GetDeliveryExperience returns the DeliveryExperience field value
func (o *ShippingServiceOptions) GetDeliveryExperience() DeliveryExperienceType {
	if o == nil {
		var ret DeliveryExperienceType
		return ret
	}

	return o.DeliveryExperience
}

// GetDeliveryExperienceOk returns a tuple with the DeliveryExperience field value
// and a boolean to check if the value has been set.
func (o *ShippingServiceOptions) GetDeliveryExperienceOk() (*DeliveryExperienceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryExperience, true
}

// SetDeliveryExperience sets field value
func (o *ShippingServiceOptions) SetDeliveryExperience(v DeliveryExperienceType) {
	o.DeliveryExperience = v
}

// GetDeclaredValue returns the DeclaredValue field value if set, zero value otherwise.
func (o *ShippingServiceOptions) GetDeclaredValue() CurrencyAmount {
	if o == nil || IsNil(o.DeclaredValue) {
		var ret CurrencyAmount
		return ret
	}
	return *o.DeclaredValue
}

// GetDeclaredValueOk returns a tuple with the DeclaredValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingServiceOptions) GetDeclaredValueOk() (*CurrencyAmount, bool) {
	if o == nil || IsNil(o.DeclaredValue) {
		return nil, false
	}
	return o.DeclaredValue, true
}

// HasDeclaredValue returns a boolean if a field has been set.
func (o *ShippingServiceOptions) HasDeclaredValue() bool {
	if o != nil && !IsNil(o.DeclaredValue) {
		return true
	}

	return false
}

// SetDeclaredValue gets a reference to the given CurrencyAmount and assigns it to the DeclaredValue field.
func (o *ShippingServiceOptions) SetDeclaredValue(v CurrencyAmount) {
	o.DeclaredValue = &v
}

// GetCarrierWillPickUp returns the CarrierWillPickUp field value
func (o *ShippingServiceOptions) GetCarrierWillPickUp() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CarrierWillPickUp
}

// GetCarrierWillPickUpOk returns a tuple with the CarrierWillPickUp field value
// and a boolean to check if the value has been set.
func (o *ShippingServiceOptions) GetCarrierWillPickUpOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierWillPickUp, true
}

// SetCarrierWillPickUp sets field value
func (o *ShippingServiceOptions) SetCarrierWillPickUp(v bool) {
	o.CarrierWillPickUp = v
}

// GetCarrierWillPickUpOption returns the CarrierWillPickUpOption field value if set, zero value otherwise.
func (o *ShippingServiceOptions) GetCarrierWillPickUpOption() CarrierWillPickUpOption {
	if o == nil || IsNil(o.CarrierWillPickUpOption) {
		var ret CarrierWillPickUpOption
		return ret
	}
	return *o.CarrierWillPickUpOption
}

// GetCarrierWillPickUpOptionOk returns a tuple with the CarrierWillPickUpOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingServiceOptions) GetCarrierWillPickUpOptionOk() (*CarrierWillPickUpOption, bool) {
	if o == nil || IsNil(o.CarrierWillPickUpOption) {
		return nil, false
	}
	return o.CarrierWillPickUpOption, true
}

// HasCarrierWillPickUpOption returns a boolean if a field has been set.
func (o *ShippingServiceOptions) HasCarrierWillPickUpOption() bool {
	if o != nil && !IsNil(o.CarrierWillPickUpOption) {
		return true
	}

	return false
}

// SetCarrierWillPickUpOption gets a reference to the given CarrierWillPickUpOption and assigns it to the CarrierWillPickUpOption field.
func (o *ShippingServiceOptions) SetCarrierWillPickUpOption(v CarrierWillPickUpOption) {
	o.CarrierWillPickUpOption = &v
}

// GetLabelFormat returns the LabelFormat field value if set, zero value otherwise.
func (o *ShippingServiceOptions) GetLabelFormat() LabelFormat {
	if o == nil || IsNil(o.LabelFormat) {
		var ret LabelFormat
		return ret
	}
	return *o.LabelFormat
}

// GetLabelFormatOk returns a tuple with the LabelFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingServiceOptions) GetLabelFormatOk() (*LabelFormat, bool) {
	if o == nil || IsNil(o.LabelFormat) {
		return nil, false
	}
	return o.LabelFormat, true
}

// HasLabelFormat returns a boolean if a field has been set.
func (o *ShippingServiceOptions) HasLabelFormat() bool {
	if o != nil && !IsNil(o.LabelFormat) {
		return true
	}

	return false
}

// SetLabelFormat gets a reference to the given LabelFormat and assigns it to the LabelFormat field.
func (o *ShippingServiceOptions) SetLabelFormat(v LabelFormat) {
	o.LabelFormat = &v
}

func (o ShippingServiceOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DeliveryExperience"] = o.DeliveryExperience
	if !IsNil(o.DeclaredValue) {
		toSerialize["DeclaredValue"] = o.DeclaredValue
	}
	toSerialize["CarrierWillPickUp"] = o.CarrierWillPickUp
	if !IsNil(o.CarrierWillPickUpOption) {
		toSerialize["CarrierWillPickUpOption"] = o.CarrierWillPickUpOption
	}
	if !IsNil(o.LabelFormat) {
		toSerialize["LabelFormat"] = o.LabelFormat
	}
	return toSerialize, nil
}

type NullableShippingServiceOptions struct {
	value *ShippingServiceOptions
	isSet bool
}

func (v NullableShippingServiceOptions) Get() *ShippingServiceOptions {
	return v.value
}

func (v *NullableShippingServiceOptions) Set(val *ShippingServiceOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingServiceOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingServiceOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingServiceOptions(val *ShippingServiceOptions) *NullableShippingServiceOptions {
	return &NullableShippingServiceOptions{value: val, isSet: true}
}

func (v NullableShippingServiceOptions) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShippingServiceOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
