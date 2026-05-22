package shipment_invoicing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ShipmentInvoiceStatusInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentInvoiceStatusInfo{}

// ShipmentInvoiceStatusInfo The shipment invoice status information.
type ShipmentInvoiceStatusInfo struct {
	// The Amazon-defined shipment identifier.
	AmazonShipmentId *string                `json:"AmazonShipmentId,omitempty"`
	InvoiceStatus    *ShipmentInvoiceStatus `json:"InvoiceStatus,omitempty"`
}

// NewShipmentInvoiceStatusInfo instantiates a new ShipmentInvoiceStatusInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentInvoiceStatusInfo() *ShipmentInvoiceStatusInfo {
	this := ShipmentInvoiceStatusInfo{}
	return &this
}

// NewShipmentInvoiceStatusInfoWithDefaults instantiates a new ShipmentInvoiceStatusInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentInvoiceStatusInfoWithDefaults() *ShipmentInvoiceStatusInfo {
	this := ShipmentInvoiceStatusInfo{}
	return &this
}

// GetAmazonShipmentId returns the AmazonShipmentId field value if set, zero value otherwise.
func (o *ShipmentInvoiceStatusInfo) GetAmazonShipmentId() string {
	if o == nil || IsNil(o.AmazonShipmentId) {
		var ret string
		return ret
	}
	return *o.AmazonShipmentId
}

// GetAmazonShipmentIdOk returns a tuple with the AmazonShipmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInvoiceStatusInfo) GetAmazonShipmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.AmazonShipmentId) {
		return nil, false
	}
	return o.AmazonShipmentId, true
}

// HasAmazonShipmentId returns a boolean if a field has been set.
func (o *ShipmentInvoiceStatusInfo) HasAmazonShipmentId() bool {
	if o != nil && !IsNil(o.AmazonShipmentId) {
		return true
	}

	return false
}

// SetAmazonShipmentId gets a reference to the given string and assigns it to the AmazonShipmentId field.
func (o *ShipmentInvoiceStatusInfo) SetAmazonShipmentId(v string) {
	o.AmazonShipmentId = &v
}

// GetInvoiceStatus returns the InvoiceStatus field value if set, zero value otherwise.
func (o *ShipmentInvoiceStatusInfo) GetInvoiceStatus() ShipmentInvoiceStatus {
	if o == nil || IsNil(o.InvoiceStatus) {
		var ret ShipmentInvoiceStatus
		return ret
	}
	return *o.InvoiceStatus
}

// GetInvoiceStatusOk returns a tuple with the InvoiceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInvoiceStatusInfo) GetInvoiceStatusOk() (*ShipmentInvoiceStatus, bool) {
	if o == nil || IsNil(o.InvoiceStatus) {
		return nil, false
	}
	return o.InvoiceStatus, true
}

// HasInvoiceStatus returns a boolean if a field has been set.
func (o *ShipmentInvoiceStatusInfo) HasInvoiceStatus() bool {
	if o != nil && !IsNil(o.InvoiceStatus) {
		return true
	}

	return false
}

// SetInvoiceStatus gets a reference to the given ShipmentInvoiceStatus and assigns it to the InvoiceStatus field.
func (o *ShipmentInvoiceStatusInfo) SetInvoiceStatus(v ShipmentInvoiceStatus) {
	o.InvoiceStatus = &v
}

func (o ShipmentInvoiceStatusInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AmazonShipmentId) {
		toSerialize["AmazonShipmentId"] = o.AmazonShipmentId
	}
	if !IsNil(o.InvoiceStatus) {
		toSerialize["InvoiceStatus"] = o.InvoiceStatus
	}
	return toSerialize, nil
}

type NullableShipmentInvoiceStatusInfo struct {
	value *ShipmentInvoiceStatusInfo
	isSet bool
}

func (v NullableShipmentInvoiceStatusInfo) Get() *ShipmentInvoiceStatusInfo {
	return v.value
}

func (v *NullableShipmentInvoiceStatusInfo) Set(val *ShipmentInvoiceStatusInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentInvoiceStatusInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentInvoiceStatusInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentInvoiceStatusInfo(val *ShipmentInvoiceStatusInfo) *NullableShipmentInvoiceStatusInfo {
	return &NullableShipmentInvoiceStatusInfo{value: val, isSet: true}
}

func (v NullableShipmentInvoiceStatusInfo) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShipmentInvoiceStatusInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
