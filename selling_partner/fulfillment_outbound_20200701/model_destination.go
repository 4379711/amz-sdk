package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the Destination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Destination{}

// Destination The destination for the delivery offer.
type Destination struct {
	DeliveryAddress *VariablePrecisionAddress `json:"deliveryAddress,omitempty"`
	// An IP Address.
	IpAddress *string `json:"ipAddress,omitempty"`
}

// NewDestination instantiates a new Destination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestination() *Destination {
	this := Destination{}
	return &this
}

// NewDestinationWithDefaults instantiates a new Destination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationWithDefaults() *Destination {
	this := Destination{}
	return &this
}

// GetDeliveryAddress returns the DeliveryAddress field value if set, zero value otherwise.
func (o *Destination) GetDeliveryAddress() VariablePrecisionAddress {
	if o == nil || IsNil(o.DeliveryAddress) {
		var ret VariablePrecisionAddress
		return ret
	}
	return *o.DeliveryAddress
}

// GetDeliveryAddressOk returns a tuple with the DeliveryAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetDeliveryAddressOk() (*VariablePrecisionAddress, bool) {
	if o == nil || IsNil(o.DeliveryAddress) {
		return nil, false
	}
	return o.DeliveryAddress, true
}

// HasDeliveryAddress returns a boolean if a field has been set.
func (o *Destination) HasDeliveryAddress() bool {
	if o != nil && !IsNil(o.DeliveryAddress) {
		return true
	}

	return false
}

// SetDeliveryAddress gets a reference to the given VariablePrecisionAddress and assigns it to the DeliveryAddress field.
func (o *Destination) SetDeliveryAddress(v VariablePrecisionAddress) {
	o.DeliveryAddress = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *Destination) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *Destination) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *Destination) SetIpAddress(v string) {
	o.IpAddress = &v
}

func (o Destination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeliveryAddress) {
		toSerialize["deliveryAddress"] = o.DeliveryAddress
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	return toSerialize, nil
}

type NullableDestination struct {
	value *Destination
	isSet bool
}

func (v NullableDestination) Get() *Destination {
	return v.value
}

func (v *NullableDestination) Set(val *Destination) {
	v.value = val
	v.isSet = true
}

func (v NullableDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestination(val *Destination) *NullableDestination {
	return &NullableDestination{value: val, isSet: true}
}

func (v NullableDestination) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
