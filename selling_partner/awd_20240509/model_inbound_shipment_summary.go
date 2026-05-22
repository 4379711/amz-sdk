package awd_20240509

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the InboundShipmentSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InboundShipmentSummary{}

// InboundShipmentSummary Summary for an AWD inbound shipment containing the shipment ID, which can be used to retrieve the actual shipment.
type InboundShipmentSummary struct {
	// Timestamp when the shipment was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Optional client-provided reference ID that can be used to correlate this shipment with client resources. For example, to map this shipment to an internal bookkeeping order record.
	ExternalReferenceId *string `json:"externalReferenceId,omitempty"`
	// The AWD inbound order ID that this inbound shipment belongs to.
	OrderId string `json:"orderId"`
	// A unique shipment ID.
	ShipmentId     string                `json:"shipmentId"`
	ShipmentStatus InboundShipmentStatus `json:"shipmentStatus"`
	// Timestamp when the shipment was updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type _InboundShipmentSummary InboundShipmentSummary

// NewInboundShipmentSummary instantiates a new InboundShipmentSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundShipmentSummary(orderId string, shipmentId string, shipmentStatus InboundShipmentStatus) *InboundShipmentSummary {
	this := InboundShipmentSummary{}
	this.OrderId = orderId
	this.ShipmentId = shipmentId
	this.ShipmentStatus = shipmentStatus
	return &this
}

// NewInboundShipmentSummaryWithDefaults instantiates a new InboundShipmentSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundShipmentSummaryWithDefaults() *InboundShipmentSummary {
	this := InboundShipmentSummary{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InboundShipmentSummary) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundShipmentSummary) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InboundShipmentSummary) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InboundShipmentSummary) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExternalReferenceId returns the ExternalReferenceId field value if set, zero value otherwise.
func (o *InboundShipmentSummary) GetExternalReferenceId() string {
	if o == nil || IsNil(o.ExternalReferenceId) {
		var ret string
		return ret
	}
	return *o.ExternalReferenceId
}

// GetExternalReferenceIdOk returns a tuple with the ExternalReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundShipmentSummary) GetExternalReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalReferenceId) {
		return nil, false
	}
	return o.ExternalReferenceId, true
}

// HasExternalReferenceId returns a boolean if a field has been set.
func (o *InboundShipmentSummary) HasExternalReferenceId() bool {
	if o != nil && !IsNil(o.ExternalReferenceId) {
		return true
	}

	return false
}

// SetExternalReferenceId gets a reference to the given string and assigns it to the ExternalReferenceId field.
func (o *InboundShipmentSummary) SetExternalReferenceId(v string) {
	o.ExternalReferenceId = &v
}

// GetOrderId returns the OrderId field value
func (o *InboundShipmentSummary) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *InboundShipmentSummary) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *InboundShipmentSummary) SetOrderId(v string) {
	o.OrderId = v
}

// GetShipmentId returns the ShipmentId field value
func (o *InboundShipmentSummary) GetShipmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value
// and a boolean to check if the value has been set.
func (o *InboundShipmentSummary) GetShipmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentId, true
}

// SetShipmentId sets field value
func (o *InboundShipmentSummary) SetShipmentId(v string) {
	o.ShipmentId = v
}

// GetShipmentStatus returns the ShipmentStatus field value
func (o *InboundShipmentSummary) GetShipmentStatus() InboundShipmentStatus {
	if o == nil {
		var ret InboundShipmentStatus
		return ret
	}

	return o.ShipmentStatus
}

// GetShipmentStatusOk returns a tuple with the ShipmentStatus field value
// and a boolean to check if the value has been set.
func (o *InboundShipmentSummary) GetShipmentStatusOk() (*InboundShipmentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentStatus, true
}

// SetShipmentStatus sets field value
func (o *InboundShipmentSummary) SetShipmentStatus(v InboundShipmentStatus) {
	o.ShipmentStatus = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InboundShipmentSummary) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundShipmentSummary) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InboundShipmentSummary) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *InboundShipmentSummary) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o InboundShipmentSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.ExternalReferenceId) {
		toSerialize["externalReferenceId"] = o.ExternalReferenceId
	}
	toSerialize["orderId"] = o.OrderId
	toSerialize["shipmentId"] = o.ShipmentId
	toSerialize["shipmentStatus"] = o.ShipmentStatus
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableInboundShipmentSummary struct {
	value *InboundShipmentSummary
	isSet bool
}

func (v NullableInboundShipmentSummary) Get() *InboundShipmentSummary {
	return v.value
}

func (v *NullableInboundShipmentSummary) Set(val *InboundShipmentSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundShipmentSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundShipmentSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundShipmentSummary(val *InboundShipmentSummary) *NullableInboundShipmentSummary {
	return &NullableInboundShipmentSummary{value: val, isSet: true}
}

func (v NullableInboundShipmentSummary) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInboundShipmentSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
