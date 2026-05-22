package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the IneligibleRate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IneligibleRate{}

// IneligibleRate Detailed information for an ineligible shipping service offering.
type IneligibleRate struct {
	// An identifier for the shipping service.
	ServiceId string `json:"serviceId"`
	// The name of the shipping service.
	ServiceName string `json:"serviceName"`
	// The carrier name for the offering.
	CarrierName string `json:"carrierName"`
	// The carrier identifier for the offering, provided by the carrier.
	CarrierId string `json:"carrierId"`
	// A list of reasons why a shipping service offering is ineligible.
	IneligibilityReasons []IneligibilityReason `json:"ineligibilityReasons"`
}

type _IneligibleRate IneligibleRate

// NewIneligibleRate instantiates a new IneligibleRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIneligibleRate(serviceId string, serviceName string, carrierName string, carrierId string, ineligibilityReasons []IneligibilityReason) *IneligibleRate {
	this := IneligibleRate{}
	this.ServiceId = serviceId
	this.ServiceName = serviceName
	this.CarrierName = carrierName
	this.CarrierId = carrierId
	this.IneligibilityReasons = ineligibilityReasons
	return &this
}

// NewIneligibleRateWithDefaults instantiates a new IneligibleRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIneligibleRateWithDefaults() *IneligibleRate {
	this := IneligibleRate{}
	return &this
}

// GetServiceId returns the ServiceId field value
func (o *IneligibleRate) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *IneligibleRate) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *IneligibleRate) SetServiceId(v string) {
	o.ServiceId = v
}

// GetServiceName returns the ServiceName field value
func (o *IneligibleRate) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *IneligibleRate) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *IneligibleRate) SetServiceName(v string) {
	o.ServiceName = v
}

// GetCarrierName returns the CarrierName field value
func (o *IneligibleRate) GetCarrierName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierName
}

// GetCarrierNameOk returns a tuple with the CarrierName field value
// and a boolean to check if the value has been set.
func (o *IneligibleRate) GetCarrierNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierName, true
}

// SetCarrierName sets field value
func (o *IneligibleRate) SetCarrierName(v string) {
	o.CarrierName = v
}

// GetCarrierId returns the CarrierId field value
func (o *IneligibleRate) GetCarrierId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierId
}

// GetCarrierIdOk returns a tuple with the CarrierId field value
// and a boolean to check if the value has been set.
func (o *IneligibleRate) GetCarrierIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierId, true
}

// SetCarrierId sets field value
func (o *IneligibleRate) SetCarrierId(v string) {
	o.CarrierId = v
}

// GetIneligibilityReasons returns the IneligibilityReasons field value
func (o *IneligibleRate) GetIneligibilityReasons() []IneligibilityReason {
	if o == nil {
		var ret []IneligibilityReason
		return ret
	}

	return o.IneligibilityReasons
}

// GetIneligibilityReasonsOk returns a tuple with the IneligibilityReasons field value
// and a boolean to check if the value has been set.
func (o *IneligibleRate) GetIneligibilityReasonsOk() ([]IneligibilityReason, bool) {
	if o == nil {
		return nil, false
	}
	return o.IneligibilityReasons, true
}

// SetIneligibilityReasons sets field value
func (o *IneligibleRate) SetIneligibilityReasons(v []IneligibilityReason) {
	o.IneligibilityReasons = v
}

func (o IneligibleRate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["serviceName"] = o.ServiceName
	toSerialize["carrierName"] = o.CarrierName
	toSerialize["carrierId"] = o.CarrierId
	toSerialize["ineligibilityReasons"] = o.IneligibilityReasons
	return toSerialize, nil
}

type NullableIneligibleRate struct {
	value *IneligibleRate
	isSet bool
}

func (v NullableIneligibleRate) Get() *IneligibleRate {
	return v.value
}

func (v *NullableIneligibleRate) Set(val *IneligibleRate) {
	v.value = val
	v.isSet = true
}

func (v NullableIneligibleRate) IsSet() bool {
	return v.isSet
}

func (v *NullableIneligibleRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIneligibleRate(val *IneligibleRate) *NullableIneligibleRate {
	return &NullableIneligibleRate{value: val, isSet: true}
}

func (v NullableIneligibleRate) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableIneligibleRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
