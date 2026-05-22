package vendor_direct_fulfillment_shipping_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the ShippingLabelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingLabelRequest{}

// ShippingLabelRequest Represents the request payload for creating a shipping label, containing the purchase order number, selling party, ship from party, and a list of containers or packages in the shipment.
type ShippingLabelRequest struct {
	// Purchase order number of the order for which to create a shipping label.
	PurchaseOrderNumber string              `json:"purchaseOrderNumber" validate:"regexp=^[a-zA-Z0-9]+$"`
	SellingParty        PartyIdentification `json:"sellingParty"`
	ShipFromParty       PartyIdentification `json:"shipFromParty"`
	// A list of the packages in this shipment.
	Containers []Container `json:"containers,omitempty"`
}

type _ShippingLabelRequest ShippingLabelRequest

// NewShippingLabelRequest instantiates a new ShippingLabelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingLabelRequest(purchaseOrderNumber string, sellingParty PartyIdentification, shipFromParty PartyIdentification) *ShippingLabelRequest {
	this := ShippingLabelRequest{}
	this.PurchaseOrderNumber = purchaseOrderNumber
	this.SellingParty = sellingParty
	this.ShipFromParty = shipFromParty
	return &this
}

// NewShippingLabelRequestWithDefaults instantiates a new ShippingLabelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingLabelRequestWithDefaults() *ShippingLabelRequest {
	this := ShippingLabelRequest{}
	return &this
}

// GetPurchaseOrderNumber returns the PurchaseOrderNumber field value
func (o *ShippingLabelRequest) GetPurchaseOrderNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PurchaseOrderNumber
}

// GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field value
// and a boolean to check if the value has been set.
func (o *ShippingLabelRequest) GetPurchaseOrderNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseOrderNumber, true
}

// SetPurchaseOrderNumber sets field value
func (o *ShippingLabelRequest) SetPurchaseOrderNumber(v string) {
	o.PurchaseOrderNumber = v
}

// GetSellingParty returns the SellingParty field value
func (o *ShippingLabelRequest) GetSellingParty() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.SellingParty
}

// GetSellingPartyOk returns a tuple with the SellingParty field value
// and a boolean to check if the value has been set.
func (o *ShippingLabelRequest) GetSellingPartyOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellingParty, true
}

// SetSellingParty sets field value
func (o *ShippingLabelRequest) SetSellingParty(v PartyIdentification) {
	o.SellingParty = v
}

// GetShipFromParty returns the ShipFromParty field value
func (o *ShippingLabelRequest) GetShipFromParty() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.ShipFromParty
}

// GetShipFromPartyOk returns a tuple with the ShipFromParty field value
// and a boolean to check if the value has been set.
func (o *ShippingLabelRequest) GetShipFromPartyOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipFromParty, true
}

// SetShipFromParty sets field value
func (o *ShippingLabelRequest) SetShipFromParty(v PartyIdentification) {
	o.ShipFromParty = v
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *ShippingLabelRequest) GetContainers() []Container {
	if o == nil || IsNil(o.Containers) {
		var ret []Container
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingLabelRequest) GetContainersOk() ([]Container, bool) {
	if o == nil || IsNil(o.Containers) {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *ShippingLabelRequest) HasContainers() bool {
	if o != nil && !IsNil(o.Containers) {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []Container and assigns it to the Containers field.
func (o *ShippingLabelRequest) SetContainers(v []Container) {
	o.Containers = v
}

func (o ShippingLabelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["purchaseOrderNumber"] = o.PurchaseOrderNumber
	toSerialize["sellingParty"] = o.SellingParty
	toSerialize["shipFromParty"] = o.ShipFromParty
	if !IsNil(o.Containers) {
		toSerialize["containers"] = o.Containers
	}
	return toSerialize, nil
}

type NullableShippingLabelRequest struct {
	value *ShippingLabelRequest
	isSet bool
}

func (v NullableShippingLabelRequest) Get() *ShippingLabelRequest {
	return v.value
}

func (v *NullableShippingLabelRequest) Set(val *ShippingLabelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingLabelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingLabelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingLabelRequest(val *ShippingLabelRequest) *NullableShippingLabelRequest {
	return &NullableShippingLabelRequest{value: val, isSet: true}
}

func (v NullableShippingLabelRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShippingLabelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
