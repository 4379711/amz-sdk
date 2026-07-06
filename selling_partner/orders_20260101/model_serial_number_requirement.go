package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the SerialNumberRequirement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SerialNumberRequirement{}

// SerialNumberRequirement Whether serial numbers must be provided for this line item.
type SerialNumberRequirement struct {
	// The requirement type for this request.   **Possible values**: `REQUIRED`
	RequirementType *string `json:"requirementType,omitempty"`
}

// NewSerialNumberRequirement instantiates a new SerialNumberRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSerialNumberRequirement() *SerialNumberRequirement {
	this := SerialNumberRequirement{}
	return &this
}

// NewSerialNumberRequirementWithDefaults instantiates a new SerialNumberRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSerialNumberRequirementWithDefaults() *SerialNumberRequirement {
	this := SerialNumberRequirement{}
	return &this
}

// GetRequirementType returns the RequirementType field value if set, zero value otherwise.
func (o *SerialNumberRequirement) GetRequirementType() string {
	if o == nil || IsNil(o.RequirementType) {
		var ret string
		return ret
	}
	return *o.RequirementType
}

// GetRequirementTypeOk returns a tuple with the RequirementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SerialNumberRequirement) GetRequirementTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RequirementType) {
		return nil, false
	}
	return o.RequirementType, true
}

// HasRequirementType returns a boolean if a field has been set.
func (o *SerialNumberRequirement) HasRequirementType() bool {
	if o != nil && !IsNil(o.RequirementType) {
		return true
	}

	return false
}

// SetRequirementType gets a reference to the given string and assigns it to the RequirementType field.
func (o *SerialNumberRequirement) SetRequirementType(v string) {
	o.RequirementType = &v
}

func (o SerialNumberRequirement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequirementType) {
		toSerialize["requirementType"] = o.RequirementType
	}
	return toSerialize, nil
}

type NullableSerialNumberRequirement struct {
	value *SerialNumberRequirement
	isSet bool
}

func (v NullableSerialNumberRequirement) Get() *SerialNumberRequirement {
	return v.value
}

func (v *NullableSerialNumberRequirement) Set(val *SerialNumberRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableSerialNumberRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableSerialNumberRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSerialNumberRequirement(val *SerialNumberRequirement) *NullableSerialNumberRequirement {
	return &NullableSerialNumberRequirement{value: val, isSet: true}
}

func (v NullableSerialNumberRequirement) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSerialNumberRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
