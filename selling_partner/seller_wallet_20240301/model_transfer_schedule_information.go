package seller_wallet_20240301

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the TransferScheduleInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferScheduleInformation{}

// TransferScheduleInformation Mandatory information for initiating a schedule transfer.
type TransferScheduleInformation struct {
	// The start date of the scheduled transfer.
	ScheduleStartDate *time.Time `json:"scheduleStartDate,omitempty"`
	// The end date of the scheduled transfer.
	ScheduleEndDate    *time.Time            `json:"scheduleEndDate,omitempty"`
	ScheduleExpression *ScheduleExpression   `json:"scheduleExpression,omitempty"`
	ScheduleType       *ScheduleTransferType `json:"scheduleType,omitempty"`
}

// NewTransferScheduleInformation instantiates a new TransferScheduleInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferScheduleInformation() *TransferScheduleInformation {
	this := TransferScheduleInformation{}
	return &this
}

// NewTransferScheduleInformationWithDefaults instantiates a new TransferScheduleInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferScheduleInformationWithDefaults() *TransferScheduleInformation {
	this := TransferScheduleInformation{}
	return &this
}

// GetScheduleStartDate returns the ScheduleStartDate field value if set, zero value otherwise.
func (o *TransferScheduleInformation) GetScheduleStartDate() time.Time {
	if o == nil || IsNil(o.ScheduleStartDate) {
		var ret time.Time
		return ret
	}
	return *o.ScheduleStartDate
}

// GetScheduleStartDateOk returns a tuple with the ScheduleStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferScheduleInformation) GetScheduleStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduleStartDate) {
		return nil, false
	}
	return o.ScheduleStartDate, true
}

// HasScheduleStartDate returns a boolean if a field has been set.
func (o *TransferScheduleInformation) HasScheduleStartDate() bool {
	if o != nil && !IsNil(o.ScheduleStartDate) {
		return true
	}

	return false
}

// SetScheduleStartDate gets a reference to the given time.Time and assigns it to the ScheduleStartDate field.
func (o *TransferScheduleInformation) SetScheduleStartDate(v time.Time) {
	o.ScheduleStartDate = &v
}

// GetScheduleEndDate returns the ScheduleEndDate field value if set, zero value otherwise.
func (o *TransferScheduleInformation) GetScheduleEndDate() time.Time {
	if o == nil || IsNil(o.ScheduleEndDate) {
		var ret time.Time
		return ret
	}
	return *o.ScheduleEndDate
}

// GetScheduleEndDateOk returns a tuple with the ScheduleEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferScheduleInformation) GetScheduleEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduleEndDate) {
		return nil, false
	}
	return o.ScheduleEndDate, true
}

// HasScheduleEndDate returns a boolean if a field has been set.
func (o *TransferScheduleInformation) HasScheduleEndDate() bool {
	if o != nil && !IsNil(o.ScheduleEndDate) {
		return true
	}

	return false
}

// SetScheduleEndDate gets a reference to the given time.Time and assigns it to the ScheduleEndDate field.
func (o *TransferScheduleInformation) SetScheduleEndDate(v time.Time) {
	o.ScheduleEndDate = &v
}

// GetScheduleExpression returns the ScheduleExpression field value if set, zero value otherwise.
func (o *TransferScheduleInformation) GetScheduleExpression() ScheduleExpression {
	if o == nil || IsNil(o.ScheduleExpression) {
		var ret ScheduleExpression
		return ret
	}
	return *o.ScheduleExpression
}

// GetScheduleExpressionOk returns a tuple with the ScheduleExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferScheduleInformation) GetScheduleExpressionOk() (*ScheduleExpression, bool) {
	if o == nil || IsNil(o.ScheduleExpression) {
		return nil, false
	}
	return o.ScheduleExpression, true
}

// HasScheduleExpression returns a boolean if a field has been set.
func (o *TransferScheduleInformation) HasScheduleExpression() bool {
	if o != nil && !IsNil(o.ScheduleExpression) {
		return true
	}

	return false
}

// SetScheduleExpression gets a reference to the given ScheduleExpression and assigns it to the ScheduleExpression field.
func (o *TransferScheduleInformation) SetScheduleExpression(v ScheduleExpression) {
	o.ScheduleExpression = &v
}

// GetScheduleType returns the ScheduleType field value if set, zero value otherwise.
func (o *TransferScheduleInformation) GetScheduleType() ScheduleTransferType {
	if o == nil || IsNil(o.ScheduleType) {
		var ret ScheduleTransferType
		return ret
	}
	return *o.ScheduleType
}

// GetScheduleTypeOk returns a tuple with the ScheduleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferScheduleInformation) GetScheduleTypeOk() (*ScheduleTransferType, bool) {
	if o == nil || IsNil(o.ScheduleType) {
		return nil, false
	}
	return o.ScheduleType, true
}

// HasScheduleType returns a boolean if a field has been set.
func (o *TransferScheduleInformation) HasScheduleType() bool {
	if o != nil && !IsNil(o.ScheduleType) {
		return true
	}

	return false
}

// SetScheduleType gets a reference to the given ScheduleTransferType and assigns it to the ScheduleType field.
func (o *TransferScheduleInformation) SetScheduleType(v ScheduleTransferType) {
	o.ScheduleType = &v
}

func (o TransferScheduleInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScheduleStartDate) {
		toSerialize["scheduleStartDate"] = o.ScheduleStartDate
	}
	if !IsNil(o.ScheduleEndDate) {
		toSerialize["scheduleEndDate"] = o.ScheduleEndDate
	}
	if !IsNil(o.ScheduleExpression) {
		toSerialize["scheduleExpression"] = o.ScheduleExpression
	}
	if !IsNil(o.ScheduleType) {
		toSerialize["scheduleType"] = o.ScheduleType
	}
	return toSerialize, nil
}

type NullableTransferScheduleInformation struct {
	value *TransferScheduleInformation
	isSet bool
}

func (v NullableTransferScheduleInformation) Get() *TransferScheduleInformation {
	return v.value
}

func (v *NullableTransferScheduleInformation) Set(val *TransferScheduleInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferScheduleInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferScheduleInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferScheduleInformation(val *TransferScheduleInformation) *NullableTransferScheduleInformation {
	return &NullableTransferScheduleInformation{value: val, isSet: true}
}

func (v NullableTransferScheduleInformation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTransferScheduleInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
