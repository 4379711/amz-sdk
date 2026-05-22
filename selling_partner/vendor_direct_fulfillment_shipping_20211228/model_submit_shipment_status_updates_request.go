package vendor_direct_fulfillment_shipping_20211228

import (
	"github.com/bytedance/sonic"
)

// checks if the SubmitShipmentStatusUpdatesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitShipmentStatusUpdatesRequest{}

// SubmitShipmentStatusUpdatesRequest The request schema for the `submitShipmentStatusUpdates` operation.
type SubmitShipmentStatusUpdatesRequest struct {
	// Contains a list of one or more `ShipmentStatusUpdate` objects. Each `ShipmentStatusUpdate` object represents an update to the status of a specific shipment.
	ShipmentStatusUpdates []ShipmentStatusUpdate `json:"shipmentStatusUpdates,omitempty"`
}

// NewSubmitShipmentStatusUpdatesRequest instantiates a new SubmitShipmentStatusUpdatesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitShipmentStatusUpdatesRequest() *SubmitShipmentStatusUpdatesRequest {
	this := SubmitShipmentStatusUpdatesRequest{}
	return &this
}

// NewSubmitShipmentStatusUpdatesRequestWithDefaults instantiates a new SubmitShipmentStatusUpdatesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitShipmentStatusUpdatesRequestWithDefaults() *SubmitShipmentStatusUpdatesRequest {
	this := SubmitShipmentStatusUpdatesRequest{}
	return &this
}

// GetShipmentStatusUpdates returns the ShipmentStatusUpdates field value if set, zero value otherwise.
func (o *SubmitShipmentStatusUpdatesRequest) GetShipmentStatusUpdates() []ShipmentStatusUpdate {
	if o == nil || IsNil(o.ShipmentStatusUpdates) {
		var ret []ShipmentStatusUpdate
		return ret
	}
	return o.ShipmentStatusUpdates
}

// GetShipmentStatusUpdatesOk returns a tuple with the ShipmentStatusUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitShipmentStatusUpdatesRequest) GetShipmentStatusUpdatesOk() ([]ShipmentStatusUpdate, bool) {
	if o == nil || IsNil(o.ShipmentStatusUpdates) {
		return nil, false
	}
	return o.ShipmentStatusUpdates, true
}

// HasShipmentStatusUpdates returns a boolean if a field has been set.
func (o *SubmitShipmentStatusUpdatesRequest) HasShipmentStatusUpdates() bool {
	if o != nil && !IsNil(o.ShipmentStatusUpdates) {
		return true
	}

	return false
}

// SetShipmentStatusUpdates gets a reference to the given []ShipmentStatusUpdate and assigns it to the ShipmentStatusUpdates field.
func (o *SubmitShipmentStatusUpdatesRequest) SetShipmentStatusUpdates(v []ShipmentStatusUpdate) {
	o.ShipmentStatusUpdates = v
}

func (o SubmitShipmentStatusUpdatesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShipmentStatusUpdates) {
		toSerialize["shipmentStatusUpdates"] = o.ShipmentStatusUpdates
	}
	return toSerialize, nil
}

type NullableSubmitShipmentStatusUpdatesRequest struct {
	value *SubmitShipmentStatusUpdatesRequest
	isSet bool
}

func (v NullableSubmitShipmentStatusUpdatesRequest) Get() *SubmitShipmentStatusUpdatesRequest {
	return v.value
}

func (v *NullableSubmitShipmentStatusUpdatesRequest) Set(val *SubmitShipmentStatusUpdatesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitShipmentStatusUpdatesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitShipmentStatusUpdatesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitShipmentStatusUpdatesRequest(val *SubmitShipmentStatusUpdatesRequest) *NullableSubmitShipmentStatusUpdatesRequest {
	return &NullableSubmitShipmentStatusUpdatesRequest{value: val, isSet: true}
}

func (v NullableSubmitShipmentStatusUpdatesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubmitShipmentStatusUpdatesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
