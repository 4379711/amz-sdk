package report_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the AsyncReportConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsyncReportConfiguration{}

// AsyncReportConfiguration struct for AsyncReportConfiguration
type AsyncReportConfiguration struct {
	AdProduct AsyncReportAdProduct `json:"adProduct"`
	// The list of columns to be used for report. The availability of columns depends on the selection of reportTypeId. This list cannot be null or empty.
	Columns []string `json:"columns"`
	// The identifier of the Report Type to be generated.
	ReportTypeId string `json:"reportTypeId"`
	// The report file format.
	Format string `json:"format"`
	// This field determines the aggregation level of the report data and also makes additional fields available for selection. This field cannot be null or empty.
	GroupBy []string `json:"groupBy"`
	// The list of filters supported by a report type. The availability of filters fields depends on the selection of reportTypeId.
	Filters []AsyncReportFilter `json:"filters,omitempty"`
	// The aggregation level of report data. If the timeUnit is set to `SUMMARY`, the report data is aggregated at the time period specified. The availability of time unit breakdowns depends on the selection of reportTypeId.
	TimeUnit string `json:"timeUnit"`
}

type _AsyncReportConfiguration AsyncReportConfiguration

// NewAsyncReportConfiguration instantiates a new AsyncReportConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsyncReportConfiguration(adProduct AsyncReportAdProduct, columns []string, reportTypeId string, format string, groupBy []string, timeUnit string) *AsyncReportConfiguration {
	this := AsyncReportConfiguration{}
	this.AdProduct = adProduct
	this.Columns = columns
	this.ReportTypeId = reportTypeId
	this.Format = format
	this.GroupBy = groupBy
	this.TimeUnit = timeUnit
	return &this
}

// NewAsyncReportConfigurationWithDefaults instantiates a new AsyncReportConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsyncReportConfigurationWithDefaults() *AsyncReportConfiguration {
	this := AsyncReportConfiguration{}
	return &this
}

// GetAdProduct returns the AdProduct field value
func (o *AsyncReportConfiguration) GetAdProduct() AsyncReportAdProduct {
	if o == nil {
		var ret AsyncReportAdProduct
		return ret
	}

	return o.AdProduct
}

// GetAdProductOk returns a tuple with the AdProduct field value
// and a boolean to check if the value has been set.
func (o *AsyncReportConfiguration) GetAdProductOk() (*AsyncReportAdProduct, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdProduct, true
}

// SetAdProduct sets field value
func (o *AsyncReportConfiguration) SetAdProduct(v AsyncReportAdProduct) {
	o.AdProduct = v
}

// GetColumns returns the Columns field value
func (o *AsyncReportConfiguration) GetColumns() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *AsyncReportConfiguration) GetColumnsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Columns, true
}

// SetColumns sets field value
func (o *AsyncReportConfiguration) SetColumns(v []string) {
	o.Columns = v
}

// GetReportTypeId returns the ReportTypeId field value
func (o *AsyncReportConfiguration) GetReportTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportTypeId
}

// GetReportTypeIdOk returns a tuple with the ReportTypeId field value
// and a boolean to check if the value has been set.
func (o *AsyncReportConfiguration) GetReportTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportTypeId, true
}

// SetReportTypeId sets field value
func (o *AsyncReportConfiguration) SetReportTypeId(v string) {
	o.ReportTypeId = v
}

// GetFormat returns the Format field value
func (o *AsyncReportConfiguration) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *AsyncReportConfiguration) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *AsyncReportConfiguration) SetFormat(v string) {
	o.Format = v
}

// GetGroupBy returns the GroupBy field value
func (o *AsyncReportConfiguration) GetGroupBy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value
// and a boolean to check if the value has been set.
func (o *AsyncReportConfiguration) GetGroupByOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupBy, true
}

// SetGroupBy sets field value
func (o *AsyncReportConfiguration) SetGroupBy(v []string) {
	o.GroupBy = v
}

// GetFilters returns the Filters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AsyncReportConfiguration) GetFilters() []AsyncReportFilter {
	if o == nil {
		var ret []AsyncReportFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AsyncReportConfiguration) GetFiltersOk() ([]AsyncReportFilter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *AsyncReportConfiguration) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []AsyncReportFilter and assigns it to the Filters field.
func (o *AsyncReportConfiguration) SetFilters(v []AsyncReportFilter) {
	o.Filters = v
}

// GetTimeUnit returns the TimeUnit field value
func (o *AsyncReportConfiguration) GetTimeUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeUnit
}

// GetTimeUnitOk returns a tuple with the TimeUnit field value
// and a boolean to check if the value has been set.
func (o *AsyncReportConfiguration) GetTimeUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeUnit, true
}

// SetTimeUnit sets field value
func (o *AsyncReportConfiguration) SetTimeUnit(v string) {
	o.TimeUnit = v
}

func (o AsyncReportConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adProduct"] = o.AdProduct
	toSerialize["columns"] = o.Columns
	toSerialize["reportTypeId"] = o.ReportTypeId
	toSerialize["format"] = o.Format
	toSerialize["groupBy"] = o.GroupBy
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	toSerialize["timeUnit"] = o.TimeUnit
	return toSerialize, nil
}

type NullableAsyncReportConfiguration struct {
	value *AsyncReportConfiguration
	isSet bool
}

func (v NullableAsyncReportConfiguration) Get() *AsyncReportConfiguration {
	return v.value
}

func (v *NullableAsyncReportConfiguration) Set(val *AsyncReportConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableAsyncReportConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableAsyncReportConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsyncReportConfiguration(val *AsyncReportConfiguration) *NullableAsyncReportConfiguration {
	return &NullableAsyncReportConfiguration{value: val, isSet: true}
}

func (v NullableAsyncReportConfiguration) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAsyncReportConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
