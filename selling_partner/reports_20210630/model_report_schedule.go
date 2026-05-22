package reports_20210630

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the ReportSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportSchedule{}

// ReportSchedule Detailed information about a report schedule.
type ReportSchedule struct {
	// The identifier for the report schedule. This identifier is unique only in combination with a seller ID.
	ReportScheduleId string `json:"reportScheduleId"`
	// The report type. Refer to [Report Type Values](https://developer-docs.amazon.com/sp-api/docs/report-type-values) for more information.
	ReportType string `json:"reportType"`
	// A list of marketplace identifiers. The report document's contents will contain data for all of the specified marketplaces, unless the report type indicates otherwise.
	MarketplaceIds []string `json:"marketplaceIds,omitempty"`
	// Additional information passed to reports. This varies by report type.
	ReportOptions *map[string]string `json:"reportOptions,omitempty"`
	// An <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> period value that indicates how often a report should be created.
	Period string `json:"period"`
	// The date and time when the schedule will create its next report, in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date time format.
	NextReportCreationTime *time.Time `json:"nextReportCreationTime,omitempty"`
}

type _ReportSchedule ReportSchedule

// NewReportSchedule instantiates a new ReportSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportSchedule(reportScheduleId string, reportType string, period string) *ReportSchedule {
	this := ReportSchedule{}
	this.ReportScheduleId = reportScheduleId
	this.ReportType = reportType
	this.Period = period
	return &this
}

// NewReportScheduleWithDefaults instantiates a new ReportSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportScheduleWithDefaults() *ReportSchedule {
	this := ReportSchedule{}
	return &this
}

// GetReportScheduleId returns the ReportScheduleId field value
func (o *ReportSchedule) GetReportScheduleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportScheduleId
}

// GetReportScheduleIdOk returns a tuple with the ReportScheduleId field value
// and a boolean to check if the value has been set.
func (o *ReportSchedule) GetReportScheduleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportScheduleId, true
}

// SetReportScheduleId sets field value
func (o *ReportSchedule) SetReportScheduleId(v string) {
	o.ReportScheduleId = v
}

// GetReportType returns the ReportType field value
func (o *ReportSchedule) GetReportType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value
// and a boolean to check if the value has been set.
func (o *ReportSchedule) GetReportTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportType, true
}

// SetReportType sets field value
func (o *ReportSchedule) SetReportType(v string) {
	o.ReportType = v
}

// GetMarketplaceIds returns the MarketplaceIds field value if set, zero value otherwise.
func (o *ReportSchedule) GetMarketplaceIds() []string {
	if o == nil || IsNil(o.MarketplaceIds) {
		var ret []string
		return ret
	}
	return o.MarketplaceIds
}

// GetMarketplaceIdsOk returns a tuple with the MarketplaceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportSchedule) GetMarketplaceIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.MarketplaceIds) {
		return nil, false
	}
	return o.MarketplaceIds, true
}

// HasMarketplaceIds returns a boolean if a field has been set.
func (o *ReportSchedule) HasMarketplaceIds() bool {
	if o != nil && !IsNil(o.MarketplaceIds) {
		return true
	}

	return false
}

// SetMarketplaceIds gets a reference to the given []string and assigns it to the MarketplaceIds field.
func (o *ReportSchedule) SetMarketplaceIds(v []string) {
	o.MarketplaceIds = v
}

// GetReportOptions returns the ReportOptions field value if set, zero value otherwise.
func (o *ReportSchedule) GetReportOptions() map[string]string {
	if o == nil || IsNil(o.ReportOptions) {
		var ret map[string]string
		return ret
	}
	return *o.ReportOptions
}

// GetReportOptionsOk returns a tuple with the ReportOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportSchedule) GetReportOptionsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ReportOptions) {
		return nil, false
	}
	return o.ReportOptions, true
}

// HasReportOptions returns a boolean if a field has been set.
func (o *ReportSchedule) HasReportOptions() bool {
	if o != nil && !IsNil(o.ReportOptions) {
		return true
	}

	return false
}

// SetReportOptions gets a reference to the given map[string]string and assigns it to the ReportOptions field.
func (o *ReportSchedule) SetReportOptions(v map[string]string) {
	o.ReportOptions = &v
}

// GetPeriod returns the Period field value
func (o *ReportSchedule) GetPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *ReportSchedule) GetPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *ReportSchedule) SetPeriod(v string) {
	o.Period = v
}

// GetNextReportCreationTime returns the NextReportCreationTime field value if set, zero value otherwise.
func (o *ReportSchedule) GetNextReportCreationTime() time.Time {
	if o == nil || IsNil(o.NextReportCreationTime) {
		var ret time.Time
		return ret
	}
	return *o.NextReportCreationTime
}

// GetNextReportCreationTimeOk returns a tuple with the NextReportCreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportSchedule) GetNextReportCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NextReportCreationTime) {
		return nil, false
	}
	return o.NextReportCreationTime, true
}

// HasNextReportCreationTime returns a boolean if a field has been set.
func (o *ReportSchedule) HasNextReportCreationTime() bool {
	if o != nil && !IsNil(o.NextReportCreationTime) {
		return true
	}

	return false
}

// SetNextReportCreationTime gets a reference to the given time.Time and assigns it to the NextReportCreationTime field.
func (o *ReportSchedule) SetNextReportCreationTime(v time.Time) {
	o.NextReportCreationTime = &v
}

func (o ReportSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reportScheduleId"] = o.ReportScheduleId
	toSerialize["reportType"] = o.ReportType
	if !IsNil(o.MarketplaceIds) {
		toSerialize["marketplaceIds"] = o.MarketplaceIds
	}
	if !IsNil(o.ReportOptions) {
		toSerialize["reportOptions"] = o.ReportOptions
	}
	toSerialize["period"] = o.Period
	if !IsNil(o.NextReportCreationTime) {
		toSerialize["nextReportCreationTime"] = o.NextReportCreationTime
	}
	return toSerialize, nil
}

type NullableReportSchedule struct {
	value *ReportSchedule
	isSet bool
}

func (v NullableReportSchedule) Get() *ReportSchedule {
	return v.value
}

func (v *NullableReportSchedule) Set(val *ReportSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableReportSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableReportSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportSchedule(val *ReportSchedule) *NullableReportSchedule {
	return &NullableReportSchedule{value: val, isSet: true}
}

func (v NullableReportSchedule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableReportSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
