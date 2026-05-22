package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ShipmentSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentSummary{}

// ShipmentSummary Summary information about a shipment.
type ShipmentSummary struct {
	// Identifier of a shipment. A shipment contains the boxes and units being inbounded.
	ShipmentId string `json:"shipmentId" validate:"regexp=^[a-zA-Z0-9-]*$"`
	// The status of a shipment. The state of the shipment will typically start as `UNCONFIRMED`, then transition to `WORKING` after a placement option has been confirmed, and then to `READY_TO_SHIP` once labels are generated.  Possible values: `ABANDONED`, `CANCELLED`, `CHECKED_IN`, `CLOSED`, `DELETED`, `DELIVERED`, `IN_TRANSIT`, `MIXED`, `READY_TO_SHIP`, `RECEIVING`, `SHIPPED`, `UNCONFIRMED`, `WORKING`
	Status string `json:"status"`
}

type _ShipmentSummary ShipmentSummary

// NewShipmentSummary instantiates a new ShipmentSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentSummary(shipmentId string, status string) *ShipmentSummary {
	this := ShipmentSummary{}
	this.ShipmentId = shipmentId
	this.Status = status
	return &this
}

// NewShipmentSummaryWithDefaults instantiates a new ShipmentSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentSummaryWithDefaults() *ShipmentSummary {
	this := ShipmentSummary{}
	return &this
}

// GetShipmentId returns the ShipmentId field value
func (o *ShipmentSummary) GetShipmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value
// and a boolean to check if the value has been set.
func (o *ShipmentSummary) GetShipmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentId, true
}

// SetShipmentId sets field value
func (o *ShipmentSummary) SetShipmentId(v string) {
	o.ShipmentId = v
}

// GetStatus returns the Status field value
func (o *ShipmentSummary) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ShipmentSummary) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ShipmentSummary) SetStatus(v string) {
	o.Status = v
}

func (o ShipmentSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipmentId"] = o.ShipmentId
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableShipmentSummary struct {
	value *ShipmentSummary
	isSet bool
}

func (v NullableShipmentSummary) Get() *ShipmentSummary {
	return v.value
}

func (v *NullableShipmentSummary) Set(val *ShipmentSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentSummary(val *ShipmentSummary) *NullableShipmentSummary {
	return &NullableShipmentSummary{value: val, isSet: true}
}

func (v NullableShipmentSummary) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShipmentSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
