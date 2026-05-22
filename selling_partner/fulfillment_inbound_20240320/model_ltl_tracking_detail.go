package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the LtlTrackingDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LtlTrackingDetail{}

// LtlTrackingDetail Contains information related to Less-Than-Truckload (LTL) shipment tracking.
type LtlTrackingDetail struct {
	// The number of the carrier shipment acknowledgement document.
	BillOfLadingNumber *string `json:"billOfLadingNumber,omitempty"`
	// The number associated with the freight bill.
	FreightBillNumber []string `json:"freightBillNumber,omitempty"`
}

// NewLtlTrackingDetail instantiates a new LtlTrackingDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLtlTrackingDetail() *LtlTrackingDetail {
	this := LtlTrackingDetail{}
	return &this
}

// NewLtlTrackingDetailWithDefaults instantiates a new LtlTrackingDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLtlTrackingDetailWithDefaults() *LtlTrackingDetail {
	this := LtlTrackingDetail{}
	return &this
}

// GetBillOfLadingNumber returns the BillOfLadingNumber field value if set, zero value otherwise.
func (o *LtlTrackingDetail) GetBillOfLadingNumber() string {
	if o == nil || IsNil(o.BillOfLadingNumber) {
		var ret string
		return ret
	}
	return *o.BillOfLadingNumber
}

// GetBillOfLadingNumberOk returns a tuple with the BillOfLadingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LtlTrackingDetail) GetBillOfLadingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BillOfLadingNumber) {
		return nil, false
	}
	return o.BillOfLadingNumber, true
}

// HasBillOfLadingNumber returns a boolean if a field has been set.
func (o *LtlTrackingDetail) HasBillOfLadingNumber() bool {
	if o != nil && !IsNil(o.BillOfLadingNumber) {
		return true
	}

	return false
}

// SetBillOfLadingNumber gets a reference to the given string and assigns it to the BillOfLadingNumber field.
func (o *LtlTrackingDetail) SetBillOfLadingNumber(v string) {
	o.BillOfLadingNumber = &v
}

// GetFreightBillNumber returns the FreightBillNumber field value if set, zero value otherwise.
func (o *LtlTrackingDetail) GetFreightBillNumber() []string {
	if o == nil || IsNil(o.FreightBillNumber) {
		var ret []string
		return ret
	}
	return o.FreightBillNumber
}

// GetFreightBillNumberOk returns a tuple with the FreightBillNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LtlTrackingDetail) GetFreightBillNumberOk() ([]string, bool) {
	if o == nil || IsNil(o.FreightBillNumber) {
		return nil, false
	}
	return o.FreightBillNumber, true
}

// HasFreightBillNumber returns a boolean if a field has been set.
func (o *LtlTrackingDetail) HasFreightBillNumber() bool {
	if o != nil && !IsNil(o.FreightBillNumber) {
		return true
	}

	return false
}

// SetFreightBillNumber gets a reference to the given []string and assigns it to the FreightBillNumber field.
func (o *LtlTrackingDetail) SetFreightBillNumber(v []string) {
	o.FreightBillNumber = v
}

func (o LtlTrackingDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillOfLadingNumber) {
		toSerialize["billOfLadingNumber"] = o.BillOfLadingNumber
	}
	if !IsNil(o.FreightBillNumber) {
		toSerialize["freightBillNumber"] = o.FreightBillNumber
	}
	return toSerialize, nil
}

type NullableLtlTrackingDetail struct {
	value *LtlTrackingDetail
	isSet bool
}

func (v NullableLtlTrackingDetail) Get() *LtlTrackingDetail {
	return v.value
}

func (v *NullableLtlTrackingDetail) Set(val *LtlTrackingDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableLtlTrackingDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableLtlTrackingDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLtlTrackingDetail(val *LtlTrackingDetail) *NullableLtlTrackingDetail {
	return &NullableLtlTrackingDetail{value: val, isSet: true}
}

func (v NullableLtlTrackingDetail) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLtlTrackingDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
