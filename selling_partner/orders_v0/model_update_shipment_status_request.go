package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateShipmentStatusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateShipmentStatusRequest{}

// UpdateShipmentStatusRequest The request body for the `updateShipmentStatus` operation.
type UpdateShipmentStatusRequest struct {
	// The unobfuscated marketplace identifier.
	MarketplaceId  string         `json:"marketplaceId"`
	ShipmentStatus ShipmentStatus `json:"shipmentStatus"`
	// For partial shipment status updates, the list of order items and quantities to be updated.
	OrderItems []OrderItemsInner `json:"orderItems,omitempty"`
}

type _UpdateShipmentStatusRequest UpdateShipmentStatusRequest

// NewUpdateShipmentStatusRequest instantiates a new UpdateShipmentStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateShipmentStatusRequest(marketplaceId string, shipmentStatus ShipmentStatus) *UpdateShipmentStatusRequest {
	this := UpdateShipmentStatusRequest{}
	this.MarketplaceId = marketplaceId
	this.ShipmentStatus = shipmentStatus
	return &this
}

// NewUpdateShipmentStatusRequestWithDefaults instantiates a new UpdateShipmentStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateShipmentStatusRequestWithDefaults() *UpdateShipmentStatusRequest {
	this := UpdateShipmentStatusRequest{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *UpdateShipmentStatusRequest) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *UpdateShipmentStatusRequest) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *UpdateShipmentStatusRequest) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetShipmentStatus returns the ShipmentStatus field value
func (o *UpdateShipmentStatusRequest) GetShipmentStatus() ShipmentStatus {
	if o == nil {
		var ret ShipmentStatus
		return ret
	}

	return o.ShipmentStatus
}

// GetShipmentStatusOk returns a tuple with the ShipmentStatus field value
// and a boolean to check if the value has been set.
func (o *UpdateShipmentStatusRequest) GetShipmentStatusOk() (*ShipmentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentStatus, true
}

// SetShipmentStatus sets field value
func (o *UpdateShipmentStatusRequest) SetShipmentStatus(v ShipmentStatus) {
	o.ShipmentStatus = v
}

// GetOrderItems returns the OrderItems field value if set, zero value otherwise.
func (o *UpdateShipmentStatusRequest) GetOrderItems() []OrderItemsInner {
	if o == nil || IsNil(o.OrderItems) {
		var ret []OrderItemsInner
		return ret
	}
	return o.OrderItems
}

// GetOrderItemsOk returns a tuple with the OrderItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateShipmentStatusRequest) GetOrderItemsOk() ([]OrderItemsInner, bool) {
	if o == nil || IsNil(o.OrderItems) {
		return nil, false
	}
	return o.OrderItems, true
}

// HasOrderItems returns a boolean if a field has been set.
func (o *UpdateShipmentStatusRequest) HasOrderItems() bool {
	if o != nil && !IsNil(o.OrderItems) {
		return true
	}

	return false
}

// SetOrderItems gets a reference to the given []OrderItemsInner and assigns it to the OrderItems field.
func (o *UpdateShipmentStatusRequest) SetOrderItems(v []OrderItemsInner) {
	o.OrderItems = v
}

func (o UpdateShipmentStatusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["shipmentStatus"] = o.ShipmentStatus
	if !IsNil(o.OrderItems) {
		toSerialize["orderItems"] = o.OrderItems
	}
	return toSerialize, nil
}

type NullableUpdateShipmentStatusRequest struct {
	value *UpdateShipmentStatusRequest
	isSet bool
}

func (v NullableUpdateShipmentStatusRequest) Get() *UpdateShipmentStatusRequest {
	return v.value
}

func (v *NullableUpdateShipmentStatusRequest) Set(val *UpdateShipmentStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateShipmentStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateShipmentStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateShipmentStatusRequest(val *UpdateShipmentStatusRequest) *NullableUpdateShipmentStatusRequest {
	return &NullableUpdateShipmentStatusRequest{value: val, isSet: true}
}

func (v NullableUpdateShipmentStatusRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateShipmentStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
