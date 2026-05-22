package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the DestinationDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DestinationDetails{}

// DestinationDetails Destination details of an inbound order based on the assigned region and DC for the order.
type DestinationDetails struct {
	DestinationAddress *Address `json:"destinationAddress,omitempty"`
	// Assigned region where the order will be shipped. This can differ from what was passed as preference. AWD currently supports following region IDs: [us-west, us-east, us-southcentral, us-southeast]
	DestinationRegion *string `json:"destinationRegion,omitempty"`
	// Unique ID of the confirmed shipment being shipped to the assigned destination. This will be available only after an inbound order is confirmed and can be used to track the shipment.
	ShipmentId *string `json:"shipmentId,omitempty"`
}

// NewDestinationDetails instantiates a new DestinationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationDetails() *DestinationDetails {
	this := DestinationDetails{}
	return &this
}

// NewDestinationDetailsWithDefaults instantiates a new DestinationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationDetailsWithDefaults() *DestinationDetails {
	this := DestinationDetails{}
	return &this
}

// GetDestinationAddress returns the DestinationAddress field value if set, zero value otherwise.
func (o *DestinationDetails) GetDestinationAddress() Address {
	if o == nil || IsNil(o.DestinationAddress) {
		var ret Address
		return ret
	}
	return *o.DestinationAddress
}

// GetDestinationAddressOk returns a tuple with the DestinationAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationDetails) GetDestinationAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.DestinationAddress) {
		return nil, false
	}
	return o.DestinationAddress, true
}

// HasDestinationAddress returns a boolean if a field has been set.
func (o *DestinationDetails) HasDestinationAddress() bool {
	if o != nil && !IsNil(o.DestinationAddress) {
		return true
	}

	return false
}

// SetDestinationAddress gets a reference to the given Address and assigns it to the DestinationAddress field.
func (o *DestinationDetails) SetDestinationAddress(v Address) {
	o.DestinationAddress = &v
}

// GetDestinationRegion returns the DestinationRegion field value if set, zero value otherwise.
func (o *DestinationDetails) GetDestinationRegion() string {
	if o == nil || IsNil(o.DestinationRegion) {
		var ret string
		return ret
	}
	return *o.DestinationRegion
}

// GetDestinationRegionOk returns a tuple with the DestinationRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationDetails) GetDestinationRegionOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationRegion) {
		return nil, false
	}
	return o.DestinationRegion, true
}

// HasDestinationRegion returns a boolean if a field has been set.
func (o *DestinationDetails) HasDestinationRegion() bool {
	if o != nil && !IsNil(o.DestinationRegion) {
		return true
	}

	return false
}

// SetDestinationRegion gets a reference to the given string and assigns it to the DestinationRegion field.
func (o *DestinationDetails) SetDestinationRegion(v string) {
	o.DestinationRegion = &v
}

// GetShipmentId returns the ShipmentId field value if set, zero value otherwise.
func (o *DestinationDetails) GetShipmentId() string {
	if o == nil || IsNil(o.ShipmentId) {
		var ret string
		return ret
	}
	return *o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationDetails) GetShipmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShipmentId) {
		return nil, false
	}
	return o.ShipmentId, true
}

// HasShipmentId returns a boolean if a field has been set.
func (o *DestinationDetails) HasShipmentId() bool {
	if o != nil && !IsNil(o.ShipmentId) {
		return true
	}

	return false
}

// SetShipmentId gets a reference to the given string and assigns it to the ShipmentId field.
func (o *DestinationDetails) SetShipmentId(v string) {
	o.ShipmentId = &v
}

func (o DestinationDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DestinationAddress) {
		toSerialize["destinationAddress"] = o.DestinationAddress
	}
	if !IsNil(o.DestinationRegion) {
		toSerialize["destinationRegion"] = o.DestinationRegion
	}
	if !IsNil(o.ShipmentId) {
		toSerialize["shipmentId"] = o.ShipmentId
	}
	return toSerialize, nil
}

type NullableDestinationDetails struct {
	value *DestinationDetails
	isSet bool
}

func (v NullableDestinationDetails) Get() *DestinationDetails {
	return v.value
}

func (v *NullableDestinationDetails) Set(val *DestinationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationDetails(val *DestinationDetails) *NullableDestinationDetails {
	return &NullableDestinationDetails{value: val, isSet: true}
}

func (v NullableDestinationDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDestinationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
