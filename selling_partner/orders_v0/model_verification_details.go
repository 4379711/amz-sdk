package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the VerificationDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerificationDetails{}

// VerificationDetails Additional information related to the verification of a regulated order.
type VerificationDetails struct {
	PrescriptionDetail *PrescriptionDetail `json:"prescriptionDetail,omitempty"`
}

// NewVerificationDetails instantiates a new VerificationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationDetails() *VerificationDetails {
	this := VerificationDetails{}
	return &this
}

// NewVerificationDetailsWithDefaults instantiates a new VerificationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationDetailsWithDefaults() *VerificationDetails {
	this := VerificationDetails{}
	return &this
}

// GetPrescriptionDetail returns the PrescriptionDetail field value if set, zero value otherwise.
func (o *VerificationDetails) GetPrescriptionDetail() PrescriptionDetail {
	if o == nil || IsNil(o.PrescriptionDetail) {
		var ret PrescriptionDetail
		return ret
	}
	return *o.PrescriptionDetail
}

// GetPrescriptionDetailOk returns a tuple with the PrescriptionDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationDetails) GetPrescriptionDetailOk() (*PrescriptionDetail, bool) {
	if o == nil || IsNil(o.PrescriptionDetail) {
		return nil, false
	}
	return o.PrescriptionDetail, true
}

// HasPrescriptionDetail returns a boolean if a field has been set.
func (o *VerificationDetails) HasPrescriptionDetail() bool {
	if o != nil && !IsNil(o.PrescriptionDetail) {
		return true
	}

	return false
}

// SetPrescriptionDetail gets a reference to the given PrescriptionDetail and assigns it to the PrescriptionDetail field.
func (o *VerificationDetails) SetPrescriptionDetail(v PrescriptionDetail) {
	o.PrescriptionDetail = &v
}

func (o VerificationDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrescriptionDetail) {
		toSerialize["prescriptionDetail"] = o.PrescriptionDetail
	}
	return toSerialize, nil
}

type NullableVerificationDetails struct {
	value *VerificationDetails
	isSet bool
}

func (v NullableVerificationDetails) Get() *VerificationDetails {
	return v.value
}

func (v *NullableVerificationDetails) Set(val *VerificationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationDetails(val *VerificationDetails) *NullableVerificationDetails {
	return &NullableVerificationDetails{value: val, isSet: true}
}

func (v NullableVerificationDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableVerificationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
