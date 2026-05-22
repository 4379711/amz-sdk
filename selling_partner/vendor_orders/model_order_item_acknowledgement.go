package vendor_orders

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the OrderItemAcknowledgement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderItemAcknowledgement{}

// OrderItemAcknowledgement Represents the acknowledgement details for an individual order item, including the acknowledgement code, acknowledged quantity, scheduled ship and delivery dates, and rejection reason (if applicable).
type OrderItemAcknowledgement struct {
	// This indicates the acknowledgement code.
	AcknowledgementCode  string       `json:"acknowledgementCode"`
	AcknowledgedQuantity ItemQuantity `json:"acknowledgedQuantity"`
	// Estimated ship date per line item. Must be in ISO-8601 date/time format.
	ScheduledShipDate *time.Time `json:"scheduledShipDate,omitempty"`
	// Estimated delivery date per line item. Must be in ISO-8601 date/time format.
	ScheduledDeliveryDate *time.Time `json:"scheduledDeliveryDate,omitempty"`
	// Indicates the reason for rejection.
	RejectionReason *string `json:"rejectionReason,omitempty"`
}

type _OrderItemAcknowledgement OrderItemAcknowledgement

// NewOrderItemAcknowledgement instantiates a new OrderItemAcknowledgement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderItemAcknowledgement(acknowledgementCode string, acknowledgedQuantity ItemQuantity) *OrderItemAcknowledgement {
	this := OrderItemAcknowledgement{}
	this.AcknowledgementCode = acknowledgementCode
	this.AcknowledgedQuantity = acknowledgedQuantity
	return &this
}

// NewOrderItemAcknowledgementWithDefaults instantiates a new OrderItemAcknowledgement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderItemAcknowledgementWithDefaults() *OrderItemAcknowledgement {
	this := OrderItemAcknowledgement{}
	return &this
}

// GetAcknowledgementCode returns the AcknowledgementCode field value
func (o *OrderItemAcknowledgement) GetAcknowledgementCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AcknowledgementCode
}

// GetAcknowledgementCodeOk returns a tuple with the AcknowledgementCode field value
// and a boolean to check if the value has been set.
func (o *OrderItemAcknowledgement) GetAcknowledgementCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcknowledgementCode, true
}

// SetAcknowledgementCode sets field value
func (o *OrderItemAcknowledgement) SetAcknowledgementCode(v string) {
	o.AcknowledgementCode = v
}

// GetAcknowledgedQuantity returns the AcknowledgedQuantity field value
func (o *OrderItemAcknowledgement) GetAcknowledgedQuantity() ItemQuantity {
	if o == nil {
		var ret ItemQuantity
		return ret
	}

	return o.AcknowledgedQuantity
}

// GetAcknowledgedQuantityOk returns a tuple with the AcknowledgedQuantity field value
// and a boolean to check if the value has been set.
func (o *OrderItemAcknowledgement) GetAcknowledgedQuantityOk() (*ItemQuantity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcknowledgedQuantity, true
}

// SetAcknowledgedQuantity sets field value
func (o *OrderItemAcknowledgement) SetAcknowledgedQuantity(v ItemQuantity) {
	o.AcknowledgedQuantity = v
}

// GetScheduledShipDate returns the ScheduledShipDate field value if set, zero value otherwise.
func (o *OrderItemAcknowledgement) GetScheduledShipDate() time.Time {
	if o == nil || IsNil(o.ScheduledShipDate) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledShipDate
}

// GetScheduledShipDateOk returns a tuple with the ScheduledShipDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemAcknowledgement) GetScheduledShipDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduledShipDate) {
		return nil, false
	}
	return o.ScheduledShipDate, true
}

// HasScheduledShipDate returns a boolean if a field has been set.
func (o *OrderItemAcknowledgement) HasScheduledShipDate() bool {
	if o != nil && !IsNil(o.ScheduledShipDate) {
		return true
	}

	return false
}

// SetScheduledShipDate gets a reference to the given time.Time and assigns it to the ScheduledShipDate field.
func (o *OrderItemAcknowledgement) SetScheduledShipDate(v time.Time) {
	o.ScheduledShipDate = &v
}

// GetScheduledDeliveryDate returns the ScheduledDeliveryDate field value if set, zero value otherwise.
func (o *OrderItemAcknowledgement) GetScheduledDeliveryDate() time.Time {
	if o == nil || IsNil(o.ScheduledDeliveryDate) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledDeliveryDate
}

// GetScheduledDeliveryDateOk returns a tuple with the ScheduledDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemAcknowledgement) GetScheduledDeliveryDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduledDeliveryDate) {
		return nil, false
	}
	return o.ScheduledDeliveryDate, true
}

// HasScheduledDeliveryDate returns a boolean if a field has been set.
func (o *OrderItemAcknowledgement) HasScheduledDeliveryDate() bool {
	if o != nil && !IsNil(o.ScheduledDeliveryDate) {
		return true
	}

	return false
}

// SetScheduledDeliveryDate gets a reference to the given time.Time and assigns it to the ScheduledDeliveryDate field.
func (o *OrderItemAcknowledgement) SetScheduledDeliveryDate(v time.Time) {
	o.ScheduledDeliveryDate = &v
}

// GetRejectionReason returns the RejectionReason field value if set, zero value otherwise.
func (o *OrderItemAcknowledgement) GetRejectionReason() string {
	if o == nil || IsNil(o.RejectionReason) {
		var ret string
		return ret
	}
	return *o.RejectionReason
}

// GetRejectionReasonOk returns a tuple with the RejectionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemAcknowledgement) GetRejectionReasonOk() (*string, bool) {
	if o == nil || IsNil(o.RejectionReason) {
		return nil, false
	}
	return o.RejectionReason, true
}

// HasRejectionReason returns a boolean if a field has been set.
func (o *OrderItemAcknowledgement) HasRejectionReason() bool {
	if o != nil && !IsNil(o.RejectionReason) {
		return true
	}

	return false
}

// SetRejectionReason gets a reference to the given string and assigns it to the RejectionReason field.
func (o *OrderItemAcknowledgement) SetRejectionReason(v string) {
	o.RejectionReason = &v
}

func (o OrderItemAcknowledgement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["acknowledgementCode"] = o.AcknowledgementCode
	toSerialize["acknowledgedQuantity"] = o.AcknowledgedQuantity
	if !IsNil(o.ScheduledShipDate) {
		toSerialize["scheduledShipDate"] = o.ScheduledShipDate
	}
	if !IsNil(o.ScheduledDeliveryDate) {
		toSerialize["scheduledDeliveryDate"] = o.ScheduledDeliveryDate
	}
	if !IsNil(o.RejectionReason) {
		toSerialize["rejectionReason"] = o.RejectionReason
	}
	return toSerialize, nil
}

type NullableOrderItemAcknowledgement struct {
	value *OrderItemAcknowledgement
	isSet bool
}

func (v NullableOrderItemAcknowledgement) Get() *OrderItemAcknowledgement {
	return v.value
}

func (v *NullableOrderItemAcknowledgement) Set(val *OrderItemAcknowledgement) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderItemAcknowledgement) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderItemAcknowledgement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderItemAcknowledgement(val *OrderItemAcknowledgement) *NullableOrderItemAcknowledgement {
	return &NullableOrderItemAcknowledgement{value: val, isSet: true}
}

func (v NullableOrderItemAcknowledgement) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderItemAcknowledgement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
