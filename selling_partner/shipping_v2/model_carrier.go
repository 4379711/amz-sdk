package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the Carrier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Carrier{}

// Carrier Carrier Related Info
type Carrier struct {
	// The carrier identifier for the offering, provided by the carrier.
	Id string `json:"id"`
	// The carrier name for the offering.
	Name string `json:"name"`
}

type _Carrier Carrier

// NewCarrier instantiates a new Carrier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCarrier(id string, name string) *Carrier {
	this := Carrier{}
	this.Id = id
	this.Name = name
	return &this
}

// NewCarrierWithDefaults instantiates a new Carrier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCarrierWithDefaults() *Carrier {
	this := Carrier{}
	return &this
}

// GetId returns the Id field value
func (o *Carrier) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Carrier) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Carrier) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Carrier) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Carrier) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Carrier) SetName(v string) {
	o.Name = v
}

func (o Carrier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
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
