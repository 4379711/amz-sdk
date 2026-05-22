package fulfillment_inbound_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the PrepDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrepDetails{}

// PrepDetails Preparation instructions and who is responsible for the preparation.
type PrepDetails struct {
	PrepInstruction PrepInstruction `json:"PrepInstruction"`
	PrepOwner       PrepOwner       `json:"PrepOwner"`
}

type _PrepDetails PrepDetails

// NewPrepDetails instantiates a new PrepDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrepDetails(prepInstruction PrepInstruction, prepOwner PrepOwner) *PrepDetails {
	this := PrepDetails{}
	this.PrepInstruction = prepInstruction
	this.PrepOwner = prepOwner
	return &this
}

// NewPrepDetailsWithDefaults instantiates a new PrepDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrepDetailsWithDefaults() *PrepDetails {
	this := PrepDetails{}
	return &this
}

// GetPrepInstruction returns the PrepInstruction field value
func (o *PrepDetails) GetPrepInstruction() PrepInstruction {
	if o == nil {
		var ret PrepInstruction
		return ret
	}

	return o.PrepInstruction
}

// GetPrepInstructionOk returns a tuple with the PrepInstruction field value
// and a boolean to check if the value has been set.
func (o *PrepDetails) GetPrepInstructionOk() (*PrepInstruction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrepInstruction, true
}

// SetPrepInstruction sets field value
func (o *PrepDetails) SetPrepInstruction(v PrepInstruction) {
	o.PrepInstruction = v
}

// GetPrepOwner returns the PrepOwner field value
func (o *PrepDetails) GetPrepOwner() PrepOwner {
	if o == nil {
		var ret PrepOwner
		return ret
	}

	return o.PrepOwner
}

// GetPrepOwnerOk returns a tuple with the PrepOwner field value
// and a boolean to check if the value has been set.
func (o *PrepDetails) GetPrepOwnerOk() (*PrepOwner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrepOwner, true
}

// SetPrepOwner sets field value
func (o *PrepDetails) SetPrepOwner(v PrepOwner) {
	o.PrepOwner = v
}

func (o PrepDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["PrepInstruction"] = o.PrepInstruction
	toSerialize["PrepOwner"] = o.PrepOwner
	return toSerialize, nil
}

type NullablePrepDetails struct {
	value *PrepDetails
	isSet bool
}

func (v NullablePrepDetails) Get() *PrepDetails {
	return v.value
}

func (v *NullablePrepDetails) Set(val *PrepDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePrepDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePrepDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrepDetails(val *PrepDetails) *NullablePrepDetails {
	return &NullablePrepDetails{value: val, isSet: true}
}

func (v NullablePrepDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePrepDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
