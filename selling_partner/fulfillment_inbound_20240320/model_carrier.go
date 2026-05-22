package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the Carrier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Carrier{}

// Carrier The carrier for the inbound shipment.
type Carrier struct {
	// The carrier code. For example, USPS or DHLEX.
	AlphaCode *string `json:"alphaCode,omitempty"`
	// The name of the carrier.
	Name *string `json:"name,omitempty"`
}

// NewCarrier instantiates a new Carrier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCarrier() *Carrier {
	this := Carrier{}
	return &this
}

// NewCarrierWithDefaults instantiates a new Carrier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCarrierWithDefaults() *Carrier {
	this := Carrier{}
	return &this
}

// GetAlphaCode returns the AlphaCode field value if set, zero value otherwise.
func (o *Carrier) GetAlphaCode() string {
	if o == nil || IsNil(o.AlphaCode) {
		var ret string
		return ret
	}
	return *o.AlphaCode
}

// GetAlphaCodeOk returns a tuple with the AlphaCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Carrier) GetAlphaCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AlphaCode) {
		return nil, false
	}
	return o.AlphaCode, true
}

// HasAlphaCode returns a boolean if a field has been set.
func (o *Carrier) HasAlphaCode() bool {
	if o != nil && !IsNil(o.AlphaCode) {
		return true
	}

	return false
}

// SetAlphaCode gets a reference to the given string and assigns it to the AlphaCode field.
func (o *Carrier) SetAlphaCode(v string) {
	o.AlphaCode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Carrier) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Carrier) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Carrier) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Carrier) SetName(v string) {
	o.Name = &v
}

func (o Carrier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlphaCode) {
		toSerialize["alphaCode"] = o.AlphaCode
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCarrier struct {
	value *Carrier
	isSet bool
}

func (v NullableCarrier) Get() *Carrier {
	return v.value
}

func (v *NullableCarrier) Set(val *Carrier) {
	v.value = val
	v.isSet = true
}

func (v NullableCarrier) IsSet() bool {
	return v.isSet
}

func (v *NullableCarrier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarrier(val *Carrier) *NullableCarrier {
	return &NullableCarrier{value: val, isSet: true}
}

func (v NullableCarrier) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCarrier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
