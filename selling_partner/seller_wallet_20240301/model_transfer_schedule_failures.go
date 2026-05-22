package seller_wallet_20240301

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the TransferScheduleFailures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferScheduleFailures{}

// TransferScheduleFailures The time of and reason for the transfer schedule failure.
type TransferScheduleFailures struct {
	// The transfer schedule failure date.
	TransferScheduleFailureDate time.Time `json:"transferScheduleFailureDate"`
	// The reason listed for the failure of the transfer schedule.
	TransferScheduleFailureReason string `json:"transferScheduleFailureReason"`
}

type _TransferScheduleFailures TransferScheduleFailures

// NewTransferScheduleFailures instantiates a new TransferScheduleFailures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferScheduleFailures(transferScheduleFailureDate time.Time, transferScheduleFailureReason string) *TransferScheduleFailures {
	this := TransferScheduleFailures{}
	this.TransferScheduleFailureDate = transferScheduleFailureDate
	this.TransferScheduleFailureReason = transferScheduleFailureReason
	return &this
}

// NewTransferScheduleFailuresWithDefaults instantiates a new TransferScheduleFailures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferScheduleFailuresWithDefaults() *TransferScheduleFailures {
	this := TransferScheduleFailures{}
	return &this
}

// GetTransferScheduleFailureDate returns the TransferScheduleFailureDate field value
func (o *TransferScheduleFailures) GetTransferScheduleFailureDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TransferScheduleFailureDate
}

// GetTransferScheduleFailureDateOk returns a tuple with the TransferScheduleFailureDate field value
// and a boolean to check if the value has been set.
func (o *TransferScheduleFailures) GetTransferScheduleFailureDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransferScheduleFailureDate, true
}

// SetTransferScheduleFailureDate sets field value
func (o *TransferScheduleFailures) SetTransferScheduleFailureDate(v time.Time) {
	o.TransferScheduleFailureDate = v
}

// GetTransferScheduleFailureReason returns the TransferScheduleFailureReason field value
func (o *TransferScheduleFailures) GetTransferScheduleFailureReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferScheduleFailureReason
}

// GetTransferScheduleFailureReasonOk returns a tuple with the TransferScheduleFailureReason field value
// and a boolean to check if the value has been set.
func (o *TransferScheduleFailures) GetTransferScheduleFailureReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransferScheduleFailureReason, true
}

// SetTransferScheduleFailureReason sets field value
func (o *TransferScheduleFailures) SetTransferScheduleFailureReason(v string) {
	o.TransferScheduleFailureReason = v
}

func (o TransferScheduleFailures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transferScheduleFailureDate"] = o.TransferScheduleFailureDate
	toSerialize["transferScheduleFailureReason"] = o.TransferScheduleFailureReason
	return toSerialize, nil
}

type NullableTransferScheduleFailures struct {
	value *TransferScheduleFailures
	isSet bool
}

func (v NullableTransferScheduleFailures) Get() *TransferScheduleFailures {
	return v.value
}

func (v *NullableTransferScheduleFailures) Set(val *TransferScheduleFailures) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferScheduleFailures) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferScheduleFailures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferScheduleFailures(val *TransferScheduleFailures) *NullableTransferScheduleFailures {
	return &NullableTransferScheduleFailures{value: val, isSet: true}
}

func (v NullableTransferScheduleFailures) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTransferScheduleFailures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
