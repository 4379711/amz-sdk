package vendor_direct_fulfillment_orders_20211228

import (
	"github.com/bytedance/sonic"
)

// checks if the ShipmentDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentDetails{}

// ShipmentDetails Shipment details required for the shipment.
type ShipmentDetails struct {
	// When true, this is a priority shipment.
	IsPriorityShipment bool `json:"isPriorityShipment"`
	// When true, this order is part of a scheduled delivery program.
	IsScheduledDeliveryShipment *bool `json:"isScheduledDeliveryShipment,omitempty"`
	// When true, a packing slip is required to be sent to the customer.
	IsPslipRequired bool `json:"isPslipRequired"`
	// When true, the order contain a gift. Include the gift message and gift wrap information.
	IsGift *bool `json:"isGift,omitempty"`
	// Ship method to be used for shipping the order. Amazon defines ship method codes indicating the shipping carrier and shipment service level. To see the full list of ship methods in use, including both the code and the friendly name, search the 'Help' section on Vendor Central for 'ship methods'.
	ShipMethod    string        `json:"shipMethod"`
	ShipmentDates ShipmentDates `json:"shipmentDates"`
	// Message to customer for order status.
	MessageToCustomer string `json:"messageToCustomer"`
}

type _ShipmentDetails ShipmentDetails

// NewShipmentDetails instantiates a new ShipmentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentDetails(isPriorityShipment bool, isPslipRequired bool, shipMethod string, shipmentDates ShipmentDates, messageToCustomer string) *ShipmentDetails {
	this := ShipmentDetails{}
	this.IsPriorityShipment = isPriorityShipment
	this.IsPslipRequired = isPslipRequired
	this.ShipMethod = shipMethod
	this.ShipmentDates = shipmentDates
	this.MessageToCustomer = messageToCustomer
	return &this
}

// NewShipmentDetailsWithDefaults instantiates a new ShipmentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentDetailsWithDefaults() *ShipmentDetails {
	this := ShipmentDetails{}
	return &this
}

// GetIsPriorityShipment returns the IsPriorityShipment field value
func (o *ShipmentDetails) GetIsPriorityShipment() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPriorityShipment
}

// GetIsPriorityShipmentOk returns a tuple with the IsPriorityShipment field value
// and a boolean to check if the value has been set.
func (o *ShipmentDetails) GetIsPriorityShipmentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPriorityShipment, true
}

// SetIsPriorityShipment sets field value
func (o *ShipmentDetails) SetIsPriorityShipment(v bool) {
	o.IsPriorityShipment = v
}

// GetIsScheduledDeliveryShipment returns the IsScheduledDeliveryShipment field value if set, zero value otherwise.
func (o *ShipmentDetails) GetIsScheduledDeliveryShipment() bool {
	if o == nil || IsNil(o.IsScheduledDeliveryShipment) {
		var ret bool
		return ret
	}
	return *o.IsScheduledDeliveryShipment
}

// GetIsScheduledDeliveryShipmentOk returns a tuple with the IsScheduledDeliveryShipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDetails) GetIsScheduledDeliveryShipmentOk() (*bool, bool) {
	if o == nil || IsNil(o.IsScheduledDeliveryShipment) {
		return nil, false
	}
	return o.IsScheduledDeliveryShipment, true
}

// HasIsScheduledDeliveryShipment returns a boolean if a field has been set.
func (o *ShipmentDetails) HasIsScheduledDeliveryShipment() bool {
	if o != nil && !IsNil(o.IsScheduledDeliveryShipment) {
		return true
	}

	return false
}

// SetIsScheduledDeliveryShipment gets a reference to the given bool and assigns it to the IsScheduledDeliveryShipment field.
func (o *ShipmentDetails) SetIsScheduledDeliveryShipment(v bool) {
	o.IsScheduledDeliveryShipment = &v
}

// GetIsPslipRequired returns the IsPslipRequired field value
func (o *ShipmentDetails) GetIsPslipRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPslipRequired
}

// GetIsPslipRequiredOk returns a tuple with the IsPslipRequired field value
// and a boolean to check if the value has been set.
func (o *ShipmentDetails) GetIsPslipRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPslipRequired, true
}

// SetIsPslipRequired sets field value
func (o *ShipmentDetails) SetIsPslipRequired(v bool) {
	o.IsPslipRequired = v
}

// GetIsGift returns the IsGift field value if set, zero value otherwise.
func (o *ShipmentDetails) GetIsGift() bool {
	if o == nil || IsNil(o.IsGift) {
		var ret bool
		return ret
	}
	return *o.IsGift
}

// GetIsGiftOk returns a tuple with the IsGift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDetails) GetIsGiftOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGift) {
		return nil, false
	}
	return o.IsGift, true
}

// HasIsGift returns a boolean if a field has been set.
func (o *ShipmentDetails) HasIsGift() bool {
	if o != nil && !IsNil(o.IsGift) {
		return true
	}

	return false
}

// SetIsGift gets a reference to the given bool and assigns it to the IsGift field.
func (o *ShipmentDetails) SetIsGift(v bool) {
	o.IsGift = &v
}

// GetShipMethod returns the ShipMethod field value
func (o *ShipmentDetails) GetShipMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipMethod
}

// GetShipMethodOk returns a tuple with the ShipMethod field value
// and a boolean to check if the value has been set.
func (o *ShipmentDetails) GetShipMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipMethod, true
}

// SetShipMethod sets field value
func (o *ShipmentDetails) SetShipMethod(v string) {
	o.ShipMethod = v
}

// GetShipmentDates returns the ShipmentDates field value
func (o *ShipmentDetails) GetShipmentDates() ShipmentDates {
	if o == nil {
		var ret ShipmentDates
		return ret
	}

	return o.ShipmentDates
}

// GetShipmentDatesOk returns a tuple with the ShipmentDates field value
// and a boolean to check if the value has been set.
func (o *ShipmentDetails) GetShipmentDatesOk() (*ShipmentDates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentDates, true
}

// SetShipmentDates sets field value
func (o *ShipmentDetails) SetShipmentDates(v ShipmentDates) {
	o.ShipmentDates = v
}

// GetMessageToCustomer returns the MessageToCustomer field value
func (o *ShipmentDetails) GetMessageToCustomer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageToCustomer
}

// GetMessageToCustomerOk returns a tuple with the MessageToCustomer field value
// and a boolean to check if the value has been set.
func (o *ShipmentDetails) GetMessageToCustomerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageToCustomer, true
}

// SetMessageToCustomer sets field value
func (o *ShipmentDetails) SetMessageToCustomer(v string) {
	o.MessageToCustomer = v
}

func (o ShipmentDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isPriorityShipment"] = o.IsPriorityShipment
	if !IsNil(o.IsScheduledDeliveryShipment) {
		toSerialize["isScheduledDeliveryShipment"] = o.IsScheduledDeliveryShipment
	}
	toSerialize["isPslipRequired"] = o.IsPslipRequired
	if !IsNil(o.IsGift) {
		toSerialize["isGift"] = o.IsGift
	}
	toSerialize["shipMethod"] = o.ShipMethod
	toSerialize["shipmentDates"] = o.ShipmentDates
	toSerialize["messageToCustomer"] = o.MessageToCustomer
	return toSerialize, nil
}

type NullableShipmentDetails struct {
	value *ShipmentDetails
	isSet bool
}

func (v NullableShipmentDetails) Get() *ShipmentDetails {
	return v.value
}

func (v *NullableShipmentDetails) Set(val *ShipmentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentDetails(val *ShipmentDetails) *NullableShipmentDetails {
	return &NullableShipmentDetails{value: val, isSet: true}
}

func (v NullableShipmentDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShipmentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
