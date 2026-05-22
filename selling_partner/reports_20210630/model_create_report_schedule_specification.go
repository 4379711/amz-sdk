package reports_20210630

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the CreateReportScheduleSpecification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateReportScheduleSpecification{}

// CreateReportScheduleSpecification Information required to create the report schedule.
type CreateReportScheduleSpecification struct {
	// The report type. Refer to [Report Type Values](https://developer-docs.amazon.com/sp-api/docs/report-type-values) for more information.
	ReportType string `json:"reportType"`
	// A list of marketplace identifiers for the report schedule.
	MarketplaceIds []string `json:"marketplaceIds"`
	// Additional information passed to reports. This varies by report type.
	ReportOptions *map[string]string `json:"reportOptions,omitempty"`
	// One of a set of predefined <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> periods that specifies how often a report should be created.
	Period string `json:"period"`
	// The date and time when the schedule will create its next report, in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date time format.
	NextReportCreationTime *time.Time `json:"nextReportCreationTime,omitempty"`
}

type _CreateReportScheduleSpecification CreateReportScheduleSpecification

// NewCreateReportScheduleSpecification instantiates a new CreateReportScheduleSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReportScheduleSpecification(reportType string, marketplaceIds []string, period string) *CreateReportScheduleSpecification {
	this := CreateReportScheduleSpecification{}
	this.ReportType = reportType
	this.MarketplaceIds = marketplaceIds
	this.Period = period
	return &this
}

// NewCreateReportScheduleSpecificationWithDefaults instantiates a new CreateReportScheduleSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReportScheduleSpecificationWithDefaults() *CreateReportScheduleSpecification {
	this := CreateReportScheduleSpecification{}
	return &this
}

// GetReportType returns the ReportType field value
func (o *CreateReportScheduleSpecification) GetReportType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value
// and a boolean to check if the value has been set.
func (o *CreateReportScheduleSpecification) GetReportTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportType, true
}

// SetReportType sets field value
func (o *CreateReportScheduleSpecification) SetReportType(v string) {
	o.ReportType = v
}

// GetMarketplaceIds returns the MarketplaceIds field value
func (o *CreateReportScheduleSpecification) GetMarketplaceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MarketplaceIds
}

// GetMarketplaceIdsOk returns a tuple with the MarketplaceIds field value
// and a boolean to check if the value has been set.
func (o *CreateReportScheduleSpecification) GetMarketplaceIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MarketplaceIds, true
}

// SetMarketplaceIds sets field value
func (o *CreateReportScheduleSpecification) SetMarketplaceIds(v []string) {
	o.MarketplaceIds = v
}

// GetReportOptions returns the ReportOptions field value if set, zero value otherwise.
func (o *CreateReportScheduleSpecification) GetReportOptions() map[string]string {
	if o == nil || IsNil(o.ReportOptions) {
		var ret map[string]string
		return ret
	}
	return *o.ReportOptions
}

// GetReportOptionsOk returns a tuple with the ReportOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateReportScheduleSpecification) GetReportOptionsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ReportOptions) {
		return nil, false
	}
	return o.ReportOptions, true
}

// HasReportOptions returns a boolean if a field has been set.
func (o *CreateReportScheduleSpecification) HasReportOptions() bool {
	if o != nil && !IsNil(o.ReportOptions) {
		return true
	}

	return false
}

// SetReportOptions gets a reference to the given map[string]string and assigns it to the ReportOptions field.
func (o *CreateReportScheduleSpecification) SetReportOptions(v map[string]string) {
	o.ReportOptions = &v
}

// GetPeriod returns the Period field value
func (o *CreateReportScheduleSpecification) GetPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *CreateReportScheduleSpecification) GetPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *CreateReportScheduleSpecification) SetPeriod(v string) {
	o.Period = v
}

// GetNextReportCreationTime returns the NextReportCreationTime field value if set, zero value otherwise.
func (o *CreateReportScheduleSpecification) GetNextReportCreationTime() time.Time {
	if o == nil || IsNil(o.NextReportCreationTime) {
		var ret time.Time
		return ret
	}
	return *o.NextReportCreationTime
}

// GetNextReportCreationTimeOk returns a tuple with the NextReportCreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateReportScheduleSpecification) GetNextReportCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NextReportCreationTime) {
		return nil, false
	}
	return o.NextReportCreationTime, true
}

// HasNextReportCreationTime returns a boolean if a field has been set.
func (o *CreateReportScheduleSpecification) HasNextReportCreationTime() bool {
	if o != nil && !IsNil(o.NextReportCreationTime) {
		return true
	}

	return false
}

// SetNextReportCreationTime gets a reference to the given time.Time and assigns it to the NextReportCreationTime field.
func (o *CreateReportScheduleSpecification) SetNextReportCreationTime(v time.Time) {
	o.NextReportCreationTime = &v
}

func (o CreateReportScheduleSpecification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reportType"] = o.ReportType
	toSerialize["marketplaceIds"] = o.MarketplaceIds
	if !IsNil(o.ReportOptions) {
		toSerialize["reportOptions"] = o.ReportOptions
	}
	toSerialize["period"] = o.Period
	if !IsNil(o.NextReportCreationTime) {
		toSerialize["nextReportCreationTime"] = o.NextReportCreationTime
	}
	return toSerialize, nil
}

type NullableCreateReportScheduleSpecification struct {
	value *CreateReportScheduleSpecification
	isSet bool
}

func (v NullableCreateReportScheduleSpecification) Get() *CreateReportScheduleSpecification {
	return v.value
}

func (v *NullableCreateReportScheduleSpecification) Set(val *CreateReportScheduleSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReportScheduleSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReportScheduleSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReportScheduleSpecification(val *CreateReportScheduleSpecification) *NullableCreateReportScheduleSpecification {
	return &NullableCreateReportScheduleSpecification{value: val, isSet: true}
}

func (v NullableCreateReportScheduleSpecification) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateReportScheduleSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
