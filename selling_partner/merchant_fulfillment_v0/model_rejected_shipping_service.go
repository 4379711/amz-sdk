package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the RejectedShippingService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RejectedShippingService{}

// RejectedShippingService Information about a rejected shipping service
type RejectedShippingService struct {
	// The rejected shipping carrier name. For example, USPS.
	CarrierName string `json:"CarrierName"`
	// The rejected shipping service localized name. For example, FedEx Standard Overnight.
	ShippingServiceName string `json:"ShippingServiceName"`
	// An Amazon-defined shipping service identifier.
	ShippingServiceId string `json:"ShippingServiceId"`
	// A reason code meant to be consumed programatically. For example, `CARRIER_CANNOT_SHIP_TO_POBOX`.
	RejectionReasonCode string `json:"RejectionReasonCode"`
	// A localized human readable description of the rejected reason.
	RejectionReasonMessage *string `json:"RejectionReasonMessage,omitempty"`
}

type _RejectedShippingService RejectedShippingService

// NewRejectedShippingService instantiates a new RejectedShippingService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRejectedShippingService(carrierName string, shippingServiceName string, shippingServiceId string, rejectionReasonCode string) *RejectedShippingService {
	this := RejectedShippingService{}
	this.CarrierName = carrierName
	this.ShippingServiceName = shippingServiceName
	this.ShippingServiceId = shippingServiceId
	this.RejectionReasonCode = rejectionReasonCode
	return &this
}

// NewRejectedShippingServiceWithDefaults instantiates a new RejectedShippingService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRejectedShippingServiceWithDefaults() *RejectedShippingService {
	this := RejectedShippingService{}
	return &this
}

// GetCarrierName returns the CarrierName field value
func (o *RejectedShippingService) GetCarrierName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierName
}

// GetCarrierNameOk returns a tuple with the CarrierName field value
// and a boolean to check if the value has been set.
func (o *RejectedShippingService) GetCarrierNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierName, true
}

// SetCarrierName sets field value
func (o *RejectedShippingService) SetCarrierName(v string) {
	o.CarrierName = v
}

// GetShippingServiceName returns the ShippingServiceName field value
func (o *RejectedShippingService) GetShippingServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShippingServiceName
}

// GetShippingServiceNameOk returns a tuple with the ShippingServiceName field value
// and a boolean to check if the value has been set.
func (o *RejectedShippingService) GetShippingServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingServiceName, true
}

// SetShippingServiceName sets field value
func (o *RejectedShippingService) SetShippingServiceName(v string) {
	o.ShippingServiceName = v
}

// GetShippingServiceId returns the ShippingServiceId field value
func (o *RejectedShippingService) GetShippingServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShippingServiceId
}

// GetShippingServiceIdOk returns a tuple with the ShippingServiceId field value
// and a boolean to check if the value has been set.
func (o *RejectedShippingService) GetShippingServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingServiceId, true
}

// SetShippingServiceId sets field value
func (o *RejectedShippingService) SetShippingServiceId(v string) {
	o.ShippingServiceId = v
}

// GetRejectionReasonCode returns the RejectionReasonCode field value
func (o *RejectedShippingService) GetRejectionReasonCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RejectionReasonCode
}

// GetRejectionReasonCodeOk returns a tuple with the RejectionReasonCode field value
// and a boolean to check if the value has been set.
func (o *RejectedShippingService) GetRejectionReasonCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RejectionReasonCode, true
}

// SetRejectionReasonCode sets field value
func (o *RejectedShippingService) SetRejectionReasonCode(v string) {
	o.RejectionReasonCode = v
}

// GetRejectionReasonMessage returns the RejectionReasonMessage field value if set, zero value otherwise.
func (o *RejectedShippingService) GetRejectionReasonMessage() string {
	if o == nil || IsNil(o.RejectionReasonMessage) {
		var ret string
		return ret
	}
	return *o.RejectionReasonMessage
}

// GetRejectionReasonMessageOk returns a tuple with the RejectionReasonMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectedShippingService) GetRejectionReasonMessageOk() (*string, bool) {
	if o == nil || IsNil(o.RejectionReasonMessage) {
		return nil, false
	}
	return o.RejectionReasonMessage, true
}

// HasRejectionReasonMessage returns a boolean if a field has been set.
func (o *RejectedShippingService) HasRejectionReasonMessage() bool {
	if o != nil && !IsNil(o.RejectionReasonMessage) {
		return true
	}

	return false
}

// SetRejectionReasonMessage gets a reference to the given string and assigns it to the RejectionReasonMessage field.
func (o *RejectedShippingService) SetRejectionReasonMessage(v string) {
	o.RejectionReasonMessage = &v
}

func (o RejectedShippingService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CarrierName"] = o.CarrierName
	toSerialize["ShippingServiceName"] = o.ShippingServiceName
	toSerialize["ShippingServiceId"] = o.ShippingServiceId
	toSerialize["RejectionReasonCode"] = o.RejectionReasonCode
	if !IsNil(o.RejectionReasonMessage) {
		toSerialize["RejectionReasonMessage"] = o.RejectionReasonMessage
	}
	return toSerialize, nil
}

type NullableRejectedShippingService struct {
	value *RejectedShippingService
	isSet bool
}

func (v NullableRejectedShippingService) Get() *RejectedShippingService {
	return v.value
}

func (v *NullableRejectedShippingService) Set(val *RejectedShippingService) {
	v.value = val
	v.isSet = true
}

func (v NullableRejectedShippingService) IsSet() bool {
	return v.isSet
}

func (v *NullableRejectedShippingService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRejectedShippingService(val *RejectedShippingService) *NullableRejectedShippingService {
	return &NullableRejectedShippingService{value: val, isSet: true}
}

func (v NullableRejectedShippingService) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRejectedShippingService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
