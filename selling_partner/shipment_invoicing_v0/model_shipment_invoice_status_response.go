package shipment_invoicing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ShipmentInvoiceStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentInvoiceStatusResponse{}

// ShipmentInvoiceStatusResponse The shipment invoice status response.
type ShipmentInvoiceStatusResponse struct {
	Shipments *ShipmentInvoiceStatusInfo `json:"Shipments,omitempty"`
}

// NewShipmentInvoiceStatusResponse instantiates a new ShipmentInvoiceStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentInvoiceStatusResponse() *ShipmentInvoiceStatusResponse {
	this := ShipmentInvoiceStatusResponse{}
	return &this
}

// NewShipmentInvoiceStatusResponseWithDefaults instantiates a new ShipmentInvoiceStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentInvoiceStatusResponseWithDefaults() *ShipmentInvoiceStatusResponse {
	this := ShipmentInvoiceStatusResponse{}
	return &this
}

// GetShipments returns the Shipments field value if set, zero value otherwise.
func (o *ShipmentInvoiceStatusResponse) GetShipments() ShipmentInvoiceStatusInfo {
	if o == nil || IsNil(o.Shipments) {
		var ret ShipmentInvoiceStatusInfo
		return ret
	}
	return *o.Shipments
}

// GetShipmentsOk returns a tuple with the Shipments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInvoiceStatusResponse) GetShipmentsOk() (*ShipmentInvoiceStatusInfo, bool) {
	if o == nil || IsNil(o.Shipments) {
		return nil, false
	}
	return o.Shipments, true
}

// HasShipments returns a boolean if a field has been set.
func (o *ShipmentInvoiceStatusResponse) HasShipments() bool {
	if o != nil && !IsNil(o.Shipments) {
		return true
	}

	return false
}

// SetShipments gets a reference to the given ShipmentInvoiceStatusInfo and assigns it to the Shipments field.
func (o *ShipmentInvoiceStatusResponse) SetShipments(v ShipmentInvoiceStatusInfo) {
	o.Shipments = &v
}

func (o ShipmentInvoiceStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Shipments) {
		toSerialize["Shipments"] = o.Shipments
	}
	return toSerialize, nil
}

type NullableShipmentInvoiceStatusResponse struct {
	value *ShipmentInvoiceStatusResponse
	isSet bool
}

func (v NullableShipmentInvoiceStatusResponse) Get() *ShipmentInvoiceStatusResponse {
	return v.value
}

func (v *NullableShipmentInvoiceStatusResponse) Set(val *ShipmentInvoiceStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentInvoiceStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentInvoiceStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentInvoiceStatusResponse(val *ShipmentInvoiceStatusResponse) *NullableShipmentInvoiceStatusResponse {
	return &NullableShipmentInvoiceStatusResponse{value: val, isSet: true}
}

func (v NullableShipmentInvoiceStatusResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShipmentInvoiceStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
