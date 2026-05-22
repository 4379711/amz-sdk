package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the DeliveryWindowOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryWindowOption{}

// DeliveryWindowOption Contains information pertaining to a delivery window option.
type DeliveryWindowOption struct {
	// Identifies type of Delivery Window Availability. Values: `AVAILABLE`, `CONGESTED`
	AvailabilityType string `json:"availabilityType"`
	// Identifier of a delivery window option. A delivery window option represent one option for when a shipment is expected to be delivered.
	DeliveryWindowOptionId string `json:"deliveryWindowOptionId" validate:"regexp=^[a-zA-Z0-9-]*$"`
	// The time at which this delivery window option ends. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime format with pattern `yyyy-MM-ddTHH:mmZ`.
	EndDate time.Time `json:"endDate"`
	// The time at which this delivery window option starts. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime format with pattern `yyyy-MM-ddTHH:mmZ`.
	StartDate time.Time `json:"startDate"`
	// The time at which this window delivery option is no longer valid. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime format with pattern `yyyy-MM-ddTHH:mmZ`.
	ValidUntil time.Time `json:"validUntil"`
}

type _DeliveryWindowOption DeliveryWindowOption

// NewDeliveryWindowOption instantiates a new DeliveryWindowOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryWindowOption(availabilityType string, deliveryWindowOptionId string, endDate time.Time, startDate time.Time, validUntil time.Time) *DeliveryWindowOption {
	this := DeliveryWindowOption{}
	this.AvailabilityType = availabilityType
	this.DeliveryWindowOptionId = deliveryWindowOptionId
	this.EndDate = endDate
	this.StartDate = startDate
	this.ValidUntil = validUntil
	return &this
}

// NewDeliveryWindowOptionWithDefaults instantiates a new DeliveryWindowOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryWindowOptionWithDefaults() *DeliveryWindowOption {
	this := DeliveryWindowOption{}
	return &this
}

// GetAvailabilityType returns the AvailabilityType field value
func (o *DeliveryWindowOption) GetAvailabilityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvailabilityType
}

// GetAvailabilityTypeOk returns a tuple with the AvailabilityType field value
// and a boolean to check if the value has been set.
func (o *DeliveryWindowOption) GetAvailabilityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailabilityType, true
}

// SetAvailabilityType sets field value
func (o *DeliveryWindowOption) SetAvailabilityType(v string) {
	o.AvailabilityType = v
}

// GetDeliveryWindowOptionId returns the DeliveryWindowOptionId field value
func (o *DeliveryWindowOption) GetDeliveryWindowOptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliveryWindowOptionId
}

// GetDeliveryWindowOptionIdOk returns a tuple with the DeliveryWindowOptionId field value
// and a boolean to check if the value has been set.
func (o *DeliveryWindowOption) GetDeliveryWindowOptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryWindowOptionId, true
}

// SetDeliveryWindowOptionId sets field value
func (o *DeliveryWindowOption) SetDeliveryWindowOptionId(v string) {
	o.DeliveryWindowOptionId = v
}

// GetEndDate returns the EndDate field value
func (o *DeliveryWindowOption) GetEndDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *DeliveryWindowOption) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *DeliveryWindowOption) SetEndDate(v time.Time) {
	o.EndDate = v
}

// GetStartDate returns the StartDate field value
func (o *DeliveryWindowOption) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *DeliveryWindowOption) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *DeliveryWindowOption) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetValidUntil returns the ValidUntil field value
func (o *DeliveryWindowOption) GetValidUntil() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value
// and a boolean to check if the value has been set.
func (o *DeliveryWindowOption) GetValidUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidUntil, true
}

// SetValidUntil sets field value
func (o *DeliveryWindowOption) SetValidUntil(v time.Time) {
	o.ValidUntil = v
}

func (o DeliveryWindowOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["availabilityType"] = o.AvailabilityType
	toSerialize["deliveryWindowOptionId"] = o.DeliveryWindowOptionId
	toSerialize["endDate"] = o.EndDate
	toSerialize["startDate"] = o.StartDate
	toSerialize["validUntil"] = o.ValidUntil
	return toSerialize, nil
}

type NullableDeliveryWindowOption struct {
	value *DeliveryWindowOption
	isSet bool
}

func (v NullableDeliveryWindowOption) Get() *DeliveryWindowOption {
	return v.value
}

func (v *NullableDeliveryWindowOption) Set(val *DeliveryWindowOption) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryWindowOption) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryWindowOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryWindowOption(val *DeliveryWindowOption) *NullableDeliveryWindowOption {
	return &NullableDeliveryWindowOption{value: val, isSet: true}
}

func (v NullableDeliveryWindowOption) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeliveryWindowOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
