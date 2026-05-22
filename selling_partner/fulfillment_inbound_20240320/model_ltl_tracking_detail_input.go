package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the LtlTrackingDetailInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LtlTrackingDetailInput{}

// LtlTrackingDetailInput Contains input information to update Less-Than-Truckload (LTL) tracking information.
type LtlTrackingDetailInput struct {
	// The number of the carrier shipment acknowledgement document.
	BillOfLadingNumber *string `json:"billOfLadingNumber,omitempty"`
	// Number associated with the freight bill.
	FreightBillNumber []string `json:"freightBillNumber"`
}

type _LtlTrackingDetailInput LtlTrackingDetailInput

// NewLtlTrackingDetailInput instantiates a new LtlTrackingDetailInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLtlTrackingDetailInput(freightBillNumber []string) *LtlTrackingDetailInput {
	this := LtlTrackingDetailInput{}
	this.FreightBillNumber = freightBillNumber
	return &this
}

// NewLtlTrackingDetailInputWithDefaults instantiates a new LtlTrackingDetailInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLtlTrackingDetailInputWithDefaults() *LtlTrackingDetailInput {
	this := LtlTrackingDetailInput{}
	return &this
}

// GetBillOfLadingNumber returns the BillOfLadingNumber field value if set, zero value otherwise.
func (o *LtlTrackingDetailInput) GetBillOfLadingNumber() string {
	if o == nil || IsNil(o.BillOfLadingNumber) {
		var ret string
		return ret
	}
	return *o.BillOfLadingNumber
}

// GetBillOfLadingNumberOk returns a tuple with the BillOfLadingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LtlTrackingDetailInput) GetBillOfLadingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BillOfLadingNumber) {
		return nil, false
	}
	return o.BillOfLadingNumber, true
}

// HasBillOfLadingNumber returns a boolean if a field has been set.
func (o *LtlTrackingDetailInput) HasBillOfLadingNumber() bool {
	if o != nil && !IsNil(o.BillOfLadingNumber) {
		return true
	}

	return false
}

// SetBillOfLadingNumber gets a reference to the given string and assigns it to the BillOfLadingNumber field.
func (o *LtlTrackingDetailInput) SetBillOfLadingNumber(v string) {
	o.BillOfLadingNumber = &v
}

// GetFreightBillNumber returns the FreightBillNumber field value
func (o *LtlTrackingDetailInput) GetFreightBillNumber() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FreightBillNumber
}

// GetFreightBillNumberOk returns a tuple with the FreightBillNumber field value
// and a boolean to check if the value has been set.
func (o *LtlTrackingDetailInput) GetFreightBillNumberOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FreightBillNumber, true
}

// SetFreightBillNumber sets field value
func (o *LtlTrackingDetailInput) SetFreightBillNumber(v []string) {
	o.FreightBillNumber = v
}

func (o LtlTrackingDetailInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillOfLadingNumber) {
		toSerialize["billOfLadingNumber"] = o.BillOfLadingNumber
	}
	toSerialize["freightBillNumber"] = o.FreightBillNumber
	return toSerialize, nil
}

type NullableLtlTrackingDetailInput struct {
	value *LtlTrackingDetailInput
	isSet bool
}

func (v NullableLtlTrackingDetailInput) Get() *LtlTrackingDetailInput {
	return v.value
}

func (v *NullableLtlTrackingDetailInput) Set(val *LtlTrackingDetailInput) {
	v.value = val
	v.isSet = true
}

func (v NullableLtlTrackingDetailInput) IsSet() bool {
	return v.isSet
}

func (v *NullableLtlTrackingDetailInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLtlTrackingDetailInput(val *LtlTrackingDetailInput) *NullableLtlTrackingDetailInput {
	return &NullableLtlTrackingDetailInput{value: val, isSet: true}
}

func (v NullableLtlTrackingDetailInput) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLtlTrackingDetailInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
