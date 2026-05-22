package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the ReportRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportRequest{}

// ReportRequest struct for ReportRequest
type ReportRequest struct {
	// Date in YYYYMMDD format. The report contains only metrics generated on the specified date. Note that the time zone used for date calculation is the one associated with the profile used to make the request.
	ReportDate *string       `json:"reportDate,omitempty"`
	Tactic     *TacticReport `json:"tactic,omitempty"`
	Segment    *Segment      `json:"segment,omitempty"`
	// A comma-separated list of the metrics to be included in the report.  Each report type supports different metrics. **To understand supported metrics for each report type, see [Report types](/API/docs/en-us/guides/reporting/v2/report-types).**  **Note**: Campaigns with vCPM costType should use view+click based metrics (viewAttributedConversions14d, viewAttributedDetailPageView14d, viewAttributedSales14d, viewAttributedUnitsOrdered14d, viewImpressions).  **Note**: Detail page view metrics (attributedDetailPageView14d, viewAttributedDetailPageView14d) have an SLA of 3 days.  **Tip**: Use new-to-brand (NTB) metrics to calculate how efficient your campaigns are at driving new shoppers:    1. Percentage of NTB orders = attributedOrdersNewToBrand14d / attributedConversions14d   2. Percentage NTB sales = attributedSalesNewToBrand14d / attributedSales14d   3. Percentage NTB units = attributedUnitsOrderedNewToBrand14d / attributedUnitsOrdered14d   4. NTB order rate = attributedOrdersNewToBrand14 / impressions
	Metrics *string `json:"metrics,omitempty"`
}

// NewReportRequest instantiates a new ReportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportRequest() *ReportRequest {
	this := ReportRequest{}
	return &this
}

// NewReportRequestWithDefaults instantiates a new ReportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportRequestWithDefaults() *ReportRequest {
	this := ReportRequest{}
	return &this
}

// GetReportDate returns the ReportDate field value if set, zero value otherwise.
func (o *ReportRequest) GetReportDate() string {
	if o == nil || IsNil(o.ReportDate) {
		var ret string
		return ret
	}
	return *o.ReportDate
}

// GetReportDateOk returns a tuple with the ReportDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportRequest) GetReportDateOk() (*string, bool) {
	if o == nil || IsNil(o.ReportDate) {
		return nil, false
	}
	return o.ReportDate, true
}

// HasReportDate returns a boolean if a field has been set.
func (o *ReportRequest) HasReportDate() bool {
	if o != nil && !IsNil(o.ReportDate) {
		return true
	}

	return false
}

// SetReportDate gets a reference to the given string and assigns it to the ReportDate field.
func (o *ReportRequest) SetReportDate(v string) {
	o.ReportDate = &v
}

// GetTactic returns the Tactic field value if set, zero value otherwise.
func (o *ReportRequest) GetTactic() TacticReport {
	if o == nil || IsNil(o.Tactic) {
		var ret TacticReport
		return ret
	}
	return *o.Tactic
}

// GetTacticOk returns a tuple with the Tactic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportRequest) GetTacticOk() (*TacticReport, bool) {
	if o == nil || IsNil(o.Tactic) {
		return nil, false
	}
	return o.Tactic, true
}

// HasTactic returns a boolean if a field has been set.
func (o *ReportRequest) HasTactic() bool {
	if o != nil && !IsNil(o.Tactic) {
		return true
	}

	return false
}

// SetTactic gets a reference to the given TacticReport and assigns it to the Tactic field.
func (o *ReportRequest) SetTactic(v TacticReport) {
	o.Tactic = &v
}

// GetSegment returns the Segment field value if set, zero value otherwise.
func (o *ReportRequest) GetSegment() Segment {
	if o == nil || IsNil(o.Segment) {
		var ret Segment
		return ret
	}
	return *o.Segment
}

// GetSegmentOk returns a tuple with the Segment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportRequest) GetSegmentOk() (*Segment, bool) {
	if o == nil || IsNil(o.Segment) {
		return nil, false
	}
	return o.Segment, true
}

// HasSegment returns a boolean if a field has been set.
func (o *ReportRequest) HasSegment() bool {
	if o != nil && !IsNil(o.Segment) {
		return true
	}

	return false
}

// SetSegment gets a reference to the given Segment and assigns it to the Segment field.
func (o *ReportRequest) SetSegment(v Segment) {
	o.Segment = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *ReportRequest) GetMetrics() string {
	if o == nil || IsNil(o.Metrics) {
		var ret string
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportRequest) GetMetricsOk() (*string, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *ReportRequest) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given string and assigns it to the Metrics field.
func (o *ReportRequest) SetMetrics(v string) {
	o.Metrics = &v
}

func (o ReportRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReportDate) {
		toSerialize["reportDate"] = o.ReportDate
	}
	if !IsNil(o.Tactic) {
		toSerialize["tactic"] = o.Tactic
	}
	if !IsNil(o.Segment) {
		toSerialize["segment"] = o.Segment
	}
	if !IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}
	return toSerialize, nil
}

type NullableReportRequest struct {
	value *ReportRequest
	isSet bool
}

func (v NullableReportRequest) Get() *ReportRequest {
	return v.value
}

func (v *NullableReportRequest) Set(val *ReportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportRequest(val *ReportRequest) *NullableReportRequest {
	return &NullableReportRequest{value: val, isSet: true}
}

func (v NullableReportRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableReportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
