package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderFulfillment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderFulfillment{}

// OrderFulfillment Information about how the order is being processed, packed, and shipped to the customer.
type OrderFulfillment struct {
	FulfillmentStatus FulfillmentStatus `json:"fulfillmentStatus"`
	// Specifies whether Amazon or the merchant is responsible for fulfilling this order.  **Possible values**: `AMAZON`, `MERCHANT`.
	FulfilledBy *string `json:"fulfilledBy,omitempty"`
	// The category of the shipping speed option selected by the customer at checkout.  **Possible values**: `EXPEDITED`, `FREE_ECONOMY`, `NEXT_DAY`, `PRIORITY`, `SAME_DAY`, `SECOND_DAY`, `SCHEDULED`, `STANDARD`.
	FulfillmentServiceLevel *string        `json:"fulfillmentServiceLevel,omitempty"`
	ShipByWindow            *DateTimeRange `json:"shipByWindow,omitempty"`
	DeliverByWindow         *DateTimeRange `json:"deliverByWindow,omitempty"`
}

// NewOrderFulfillment instantiates a new OrderFulfillment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderFulfillment(fulfillmentStatus FulfillmentStatus) *OrderFulfillment {
	this := OrderFulfillment{}
	this.FulfillmentStatus = fulfillmentStatus
	return &this
}

// NewOrderFulfillmentWithDefaults instantiates a new OrderFulfillment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderFulfillmentWithDefaults() *OrderFulfillment {
	this := OrderFulfillment{}
	return &this
}

// GetFulfillmentStatus returns the FulfillmentStatus field value
func (o *OrderFulfillment) GetFulfillmentStatus() FulfillmentStatus {
	if o == nil {
		var ret FulfillmentStatus
		return ret
	}

	return o.FulfillmentStatus
}

// GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field value
// and a boolean to check if the value has been set.
func (o *OrderFulfillment) GetFulfillmentStatusOk() (*FulfillmentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulfillmentStatus, true
}

// SetFulfillmentStatus sets field value
func (o *OrderFulfillment) SetFulfillmentStatus(v FulfillmentStatus) {
	o.FulfillmentStatus = v
}

// GetFulfilledBy returns the FulfilledBy field value if set, zero value otherwise.
func (o *OrderFulfillment) GetFulfilledBy() string {
	if o == nil || IsNil(o.FulfilledBy) {
		var ret string
		return ret
	}
	return *o.FulfilledBy
}

// GetFulfilledByOk returns a tuple with the FulfilledBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFulfillment) GetFulfilledByOk() (*string, bool) {
	if o == nil || IsNil(o.FulfilledBy) {
		return nil, false
	}
	return o.FulfilledBy, true
}

// HasFulfilledBy returns a boolean if a field has been set.
func (o *OrderFulfillment) HasFulfilledBy() bool {
	if o != nil && !IsNil(o.FulfilledBy) {
		return true
	}

	return false
}

// SetFulfilledBy gets a reference to the given string and assigns it to the FulfilledBy field.
func (o *OrderFulfillment) SetFulfilledBy(v string) {
	o.FulfilledBy = &v
}

// GetFulfillmentServiceLevel returns the FulfillmentServiceLevel field value if set, zero value otherwise.
func (o *OrderFulfillment) GetFulfillmentServiceLevel() string {
	if o == nil || IsNil(o.FulfillmentServiceLevel) {
		var ret string
		return ret
	}
	return *o.FulfillmentServiceLevel
}

// GetFulfillmentServiceLevelOk returns a tuple with the FulfillmentServiceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFulfillment) GetFulfillmentServiceLevelOk() (*string, bool) {
	if o == nil || IsNil(o.FulfillmentServiceLevel) {
		return nil, false
	}
	return o.FulfillmentServiceLevel, true
}

// HasFulfillmentServiceLevel returns a boolean if a field has been set.
func (o *OrderFulfillment) HasFulfillmentServiceLevel() bool {
	if o != nil && !IsNil(o.FulfillmentServiceLevel) {
		return true
	}

	return false
}

// SetFulfillmentServiceLevel gets a reference to the given string and assigns it to the FulfillmentServiceLevel field.
func (o *OrderFulfillment) SetFulfillmentServiceLevel(v string) {
	o.FulfillmentServiceLevel = &v
}

// GetShipByWindow returns the ShipByWindow field value if set, zero value otherwise.
func (o *OrderFulfillment) GetShipByWindow() DateTimeRange {
	if o == nil || IsNil(o.ShipByWindow) {
		var ret DateTimeRange
		return ret
	}
	return *o.ShipByWindow
}

// GetShipByWindowOk returns a tuple with the ShipByWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFulfillment) GetShipByWindowOk() (*DateTimeRange, bool) {
	if o == nil || IsNil(o.ShipByWindow) {
		return nil, false
	}
	return o.ShipByWindow, true
}

// HasShipByWindow returns a boolean if a field has been set.
func (o *OrderFulfillment) HasShipByWindow() bool {
	if o != nil && !IsNil(o.ShipByWindow) {
		return true
	}

	return false
}

// SetShipByWindow gets a reference to the given DateTimeRange and assigns it to the ShipByWindow field.
func (o *OrderFulfillment) SetShipByWindow(v DateTimeRange) {
	o.ShipByWindow = &v
}

// GetDeliverByWindow returns the DeliverByWindow field value if set, zero value otherwise.
func (o *OrderFulfillment) GetDeliverByWindow() DateTimeRange {
	if o == nil || IsNil(o.DeliverByWindow) {
		var ret DateTimeRange
		return ret
	}
	return *o.DeliverByWindow
}

// GetDeliverByWindowOk returns a tuple with the DeliverByWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFulfillment) GetDeliverByWindowOk() (*DateTimeRange, bool) {
	if o == nil || IsNil(o.DeliverByWindow) {
		return nil, false
	}
	return o.DeliverByWindow, true
}

// HasDeliverByWindow returns a boolean if a field has been set.
func (o *OrderFulfillment) HasDeliverByWindow() bool {
	if o != nil && !IsNil(o.DeliverByWindow) {
		return true
	}

	return false
}

// SetDeliverByWindow gets a reference to the given DateTimeRange and assigns it to the DeliverByWindow field.
func (o *OrderFulfillment) SetDeliverByWindow(v DateTimeRange) {
	o.DeliverByWindow = &v
}

func (o OrderFulfillment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fulfillmentStatus"] = o.FulfillmentStatus
	if !IsNil(o.FulfilledBy) {
		toSerialize["fulfilledBy"] = o.FulfilledBy
	}
	if !IsNil(o.FulfillmentServiceLevel) {
		toSerialize["fulfillmentServiceLevel"] = o.FulfillmentServiceLevel
	}
	if !IsNil(o.ShipByWindow) {
		toSerialize["shipByWindow"] = o.ShipByWindow
	}
	if !IsNil(o.DeliverByWindow) {
		toSerialize["deliverByWindow"] = o.DeliverByWindow
	}
	return toSerialize, nil
}

type NullableOrderFulfillment struct {
	value *OrderFulfillment
	isSet bool
}

func (v NullableOrderFulfillment) Get() *OrderFulfillment {
	return v.value
}

func (v *NullableOrderFulfillment) Set(val *OrderFulfillment) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderFulfillment) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderFulfillment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderFulfillment(val *OrderFulfillment) *NullableOrderFulfillment {
	return &NullableOrderFulfillment{value: val, isSet: true}
}

func (v NullableOrderFulfillment) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderFulfillment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
