package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ViolatingAsinEvidence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViolatingAsinEvidence{}

// ViolatingAsinEvidence struct for ViolatingAsinEvidence
type ViolatingAsinEvidence struct {
	// ASIN which has the ad policy violation.
	Asin *string `json:"asin,omitempty"`
}

// NewViolatingAsinEvidence instantiates a new ViolatingAsinEvidence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViolatingAsinEvidence() *ViolatingAsinEvidence {
	this := ViolatingAsinEvidence{}
	return &this
}

// NewViolatingAsinEvidenceWithDefaults instantiates a new ViolatingAsinEvidence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViolatingAsinEvidenceWithDefaults() *ViolatingAsinEvidence {
	this := ViolatingAsinEvidence{}
	return &this
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *ViolatingAsinEvidence) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolatingAsinEvidence) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *ViolatingAsinEvidence) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *ViolatingAsinEvidence) SetAsin(v string) {
	o.Asin = &v
}

func (o ViolatingAsinEvidence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	return toSerialize, nil
}

type NullableViolatingAsinEvidence struct {
	value *ViolatingAsinEvidence
	isSet bool
}

func (v NullableViolatingAsinEvidence) Get() *ViolatingAsinEvidence {
	return v.value
}

func (v *NullableViolatingAsinEvidence) Set(val *ViolatingAsinEvidence) {
	v.value = val
	v.isSet = true
}

func (v NullableViolatingAsinEvidence) IsSet() bool {
	return v.isSet
}

func (v *NullableViolatingAsinEvidence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViolatingAsinEvidence(val *ViolatingAsinEvidence) *NullableViolatingAsinEvidence {
	return &NullableViolatingAsinEvidence{value: val, isSet: true}
}

func (v NullableViolatingAsinEvidence) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableViolatingAsinEvidence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
