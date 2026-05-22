package fulfillment_inbound_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the AmazonPrepFeesDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmazonPrepFeesDetails{}

// AmazonPrepFeesDetails The fees for Amazon to prep goods for shipment.
type AmazonPrepFeesDetails struct {
	PrepInstruction *PrepInstruction `json:"PrepInstruction,omitempty"`
	FeePerUnit      *Amount          `json:"FeePerUnit,omitempty"`
}

// NewAmazonPrepFeesDetails instantiates a new AmazonPrepFeesDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonPrepFeesDetails() *AmazonPrepFeesDetails {
	this := AmazonPrepFeesDetails{}
	return &this
}

// NewAmazonPrepFeesDetailsWithDefaults instantiates a new AmazonPrepFeesDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonPrepFeesDetailsWithDefaults() *AmazonPrepFeesDetails {
	this := AmazonPrepFeesDetails{}
	return &this
}

// GetPrepInstruction returns the PrepInstruction field value if set, zero value otherwise.
func (o *AmazonPrepFeesDetails) GetPrepInstruction() PrepInstruction {
	if o == nil || IsNil(o.PrepInstruction) {
		var ret PrepInstruction
		return ret
	}
	return *o.PrepInstruction
}

// GetPrepInstructionOk returns a tuple with the PrepInstruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonPrepFeesDetails) GetPrepInstructionOk() (*PrepInstruction, bool) {
	if o == nil || IsNil(o.PrepInstruction) {
		return nil, false
	}
	return o.PrepInstruction, true
}

// HasPrepInstruction returns a boolean if a field has been set.
func (o *AmazonPrepFeesDetails) HasPrepInstruction() bool {
	if o != nil && !IsNil(o.PrepInstruction) {
		return true
	}

	return false
}

// SetPrepInstruction gets a reference to the given PrepInstruction and assigns it to the PrepInstruction field.
func (o *AmazonPrepFeesDetails) SetPrepInstruction(v PrepInstruction) {
	o.PrepInstruction = &v
}

// GetFeePerUnit returns the FeePerUnit field value if set, zero value otherwise.
func (o *AmazonPrepFeesDetails) GetFeePerUnit() Amount {
	if o == nil || IsNil(o.FeePerUnit) {
		var ret Amount
		return ret
	}
	return *o.FeePerUnit
}

// GetFeePerUnitOk returns a tuple with the FeePerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonPrepFeesDetails) GetFeePerUnitOk() (*Amount, bool) {
	if o == nil || IsNil(o.FeePerUnit) {
		return nil, false
	}
	return o.FeePerUnit, true
}

// HasFeePerUnit returns a boolean if a field has been set.
func (o *AmazonPrepFeesDetails) HasFeePerUnit() bool {
	if o != nil && !IsNil(o.FeePerUnit) {
		return true
	}

	return false
}

// SetFeePerUnit gets a reference to the given Amount and assigns it to the FeePerUnit field.
func (o *AmazonPrepFeesDetails) SetFeePerUnit(v Amount) {
	o.FeePerUnit = &v
}

func (o AmazonPrepFeesDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrepInstruction) {
		toSerialize["PrepInstruction"] = o.PrepInstruction
	}
	if !IsNil(o.FeePerUnit) {
		toSerialize["FeePerUnit"] = o.FeePerUnit
	}
	return toSerialize, nil
}

type NullableAmazonPrepFeesDetails struct {
	value *AmazonPrepFeesDetails
	isSet bool
}

func (v NullableAmazonPrepFeesDetails) Get() *AmazonPrepFeesDetails {
	return v.value
}

func (v *NullableAmazonPrepFeesDetails) Set(val *AmazonPrepFeesDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonPrepFeesDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonPrepFeesDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonPrepFeesDetails(val *AmazonPrepFeesDetails) *NullableAmazonPrepFeesDetails {
	return &NullableAmazonPrepFeesDetails{value: val, isSet: true}
}

func (v NullableAmazonPrepFeesDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAmazonPrepFeesDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
