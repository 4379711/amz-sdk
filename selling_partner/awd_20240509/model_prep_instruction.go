package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the PrepInstruction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrepInstruction{}

// PrepInstruction Information pertaining to the preparation of inbound products.
type PrepInstruction struct {
	PrepOwner *PrepOwner `json:"prepOwner,omitempty"`
	// The type of preparation to be done. For more information about preparing items, refer to [Prep guidance](https://sellercentral.amazon.com/help/hub/reference/external/GF4G7547KSLDX2KC) on Seller Central.
	PrepType *string `json:"prepType,omitempty"`
}

// NewPrepInstruction instantiates a new PrepInstruction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrepInstruction() *PrepInstruction {
	this := PrepInstruction{}
	return &this
}

// NewPrepInstructionWithDefaults instantiates a new PrepInstruction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrepInstructionWithDefaults() *PrepInstruction {
	this := PrepInstruction{}
	return &this
}

// GetPrepOwner returns the PrepOwner field value if set, zero value otherwise.
func (o *PrepInstruction) GetPrepOwner() PrepOwner {
	if o == nil || IsNil(o.PrepOwner) {
		var ret PrepOwner
		return ret
	}
	return *o.PrepOwner
}

// GetPrepOwnerOk returns a tuple with the PrepOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrepInstruction) GetPrepOwnerOk() (*PrepOwner, bool) {
	if o == nil || IsNil(o.PrepOwner) {
		return nil, false
	}
	return o.PrepOwner, true
}

// HasPrepOwner returns a boolean if a field has been set.
func (o *PrepInstruction) HasPrepOwner() bool {
	if o != nil && !IsNil(o.PrepOwner) {
		return true
	}

	return false
}

// SetPrepOwner gets a reference to the given PrepOwner and assigns it to the PrepOwner field.
func (o *PrepInstruction) SetPrepOwner(v PrepOwner) {
	o.PrepOwner = &v
}

// GetPrepType returns the PrepType field value if set, zero value otherwise.
func (o *PrepInstruction) GetPrepType() string {
	if o == nil || IsNil(o.PrepType) {
		var ret string
		return ret
	}
	return *o.PrepType
}

// GetPrepTypeOk returns a tuple with the PrepType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrepInstruction) GetPrepTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PrepType) {
		return nil, false
	}
	return o.PrepType, true
}

// HasPrepType returns a boolean if a field has been set.
func (o *PrepInstruction) HasPrepType() bool {
	if o != nil && !IsNil(o.PrepType) {
		return true
	}

	return false
}

// SetPrepType gets a reference to the given string and assigns it to the PrepType field.
func (o *PrepInstruction) SetPrepType(v string) {
	o.PrepType = &v
}

func (o PrepInstruction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrepOwner) {
		toSerialize["prepOwner"] = o.PrepOwner
	}
	if !IsNil(o.PrepType) {
		toSerialize["prepType"] = o.PrepType
	}
	return toSerialize, nil
}

type NullablePrepInstruction struct {
	value *PrepInstruction
	isSet bool
}

func (v NullablePrepInstruction) Get() *PrepInstruction {
	return v.value
}

func (v *NullablePrepInstruction) Set(val *PrepInstruction) {
	v.value = val
	v.isSet = true
}

func (v NullablePrepInstruction) IsSet() bool {
	return v.isSet
}

func (v *NullablePrepInstruction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrepInstruction(val *PrepInstruction) *NullablePrepInstruction {
	return &NullablePrepInstruction{value: val, isSet: true}
}

func (v NullablePrepInstruction) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePrepInstruction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
