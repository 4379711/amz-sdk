package services

import (
	"github.com/bytedance/sonic"
)

// checks if the Technician type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Technician{}

// Technician A technician who is assigned to perform the service job in part or in full.
type Technician struct {
	// The technician identifier.
	TechnicianId *string `json:"technicianId,omitempty"`
	// The name of the technician.
	Name *string `json:"name,omitempty"`
}

// NewTechnician instantiates a new Technician object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechnician() *Technician {
	this := Technician{}
	return &this
}

// NewTechnicianWithDefaults instantiates a new Technician object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechnicianWithDefaults() *Technician {
	this := Technician{}
	return &this
}

// GetTechnicianId returns the TechnicianId field value if set, zero value otherwise.
func (o *Technician) GetTechnicianId() string {
	if o == nil || IsNil(o.TechnicianId) {
		var ret string
		return ret
	}
	return *o.TechnicianId
}

// GetTechnicianIdOk returns a tuple with the TechnicianId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Technician) GetTechnicianIdOk() (*string, bool) {
	if o == nil || IsNil(o.TechnicianId) {
		return nil, false
	}
	return o.TechnicianId, true
}

// HasTechnicianId returns a boolean if a field has been set.
func (o *Technician) HasTechnicianId() bool {
	if o != nil && !IsNil(o.TechnicianId) {
		return true
	}

	return false
}

// SetTechnicianId gets a reference to the given string and assigns it to the TechnicianId field.
func (o *Technician) SetTechnicianId(v string) {
	o.TechnicianId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Technician) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Technician) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Technician) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Technician) SetName(v string) {
	o.Name = &v
}

func (o Technician) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TechnicianId) {
		toSerialize["technicianId"] = o.TechnicianId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableTechnician struct {
	value *Technician
	isSet bool
}

func (v NullableTechnician) Get() *Technician {
	return v.value
}

func (v *NullableTechnician) Set(val *Technician) {
	v.value = val
	v.isSet = true
}

func (v NullableTechnician) IsSet() bool {
	return v.isSet
}

func (v *NullableTechnician) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechnician(val *Technician) *NullableTechnician {
	return &NullableTechnician{value: val, isSet: true}
}

func (v NullableTechnician) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTechnician) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
