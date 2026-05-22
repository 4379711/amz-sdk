package reports_20210630

import (
	"github.com/bytedance/sonic"
)

// checks if the ReportScheduleList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportScheduleList{}

// ReportScheduleList A list of report schedules.
type ReportScheduleList struct {
	// Detailed information about a report schedule.
	ReportSchedules []ReportSchedule `json:"reportSchedules"`
}

type _ReportScheduleList ReportScheduleList

// NewReportScheduleList instantiates a new ReportScheduleList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportScheduleList(reportSchedules []ReportSchedule) *ReportScheduleList {
	this := ReportScheduleList{}
	this.ReportSchedules = reportSchedules
	return &this
}

// NewReportScheduleListWithDefaults instantiates a new ReportScheduleList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportScheduleListWithDefaults() *ReportScheduleList {
	this := ReportScheduleList{}
	return &this
}

// GetReportSchedules returns the ReportSchedules field value
func (o *ReportScheduleList) GetReportSchedules() []ReportSchedule {
	if o == nil {
		var ret []ReportSchedule
		return ret
	}

	return o.ReportSchedules
}

// GetReportSchedulesOk returns a tuple with the ReportSchedules field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleList) GetReportSchedulesOk() ([]ReportSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReportSchedules, true
}

// SetReportSchedules sets field value
func (o *ReportScheduleList) SetReportSchedules(v []ReportSchedule) {
	o.ReportSchedules = v
}

func (o ReportScheduleList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reportSchedules"] = o.ReportSchedules
	return toSerialize, nil
}

type NullableReportScheduleList struct {
	value *ReportScheduleList
	isSet bool
}

func (v NullableReportScheduleList) Get() *ReportScheduleList {
	return v.value
}

func (v *NullableReportScheduleList) Set(val *ReportScheduleList) {
	v.value = val
	v.isSet = true
}

func (v NullableReportScheduleList) IsSet() bool {
	return v.isSet
}

func (v *NullableReportScheduleList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportScheduleList(val *ReportScheduleList) *NullableReportScheduleList {
	return &NullableReportScheduleList{value: val, isSet: true}
}

func (v NullableReportScheduleList) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableReportScheduleList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
