package report_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the AsyncReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsyncReport{}

// AsyncReport struct for AsyncReport
type AsyncReport struct {
	// The identifier of the requested report.
	ReportId string `json:"reportId"`
	// The end date for the reporting period in YYYY-mm-dd format.
	EndDate       string                   `json:"endDate"`
	Configuration AsyncReportConfiguration `json:"configuration"`
	// The date at which the download URL for the generated report expires. urlExpires at this time defaults to 3600 seconds but may vary in the future.
	UrlExpiresAt NullableString `json:"urlExpiresAt,omitempty"`
	// URL of the generated report.
	Url NullableString `json:"url,omitempty"`
	// The date at which the report was created in ISO 8601 date time format.
	CreatedAt string `json:"createdAt"`
	// The size of the report file, in bytes.
	FileSize NullableFloat32 `json:"fileSize,omitempty"`
	// Present for failed reports only. The reason why a report failed to generate.
	FailureReason NullableString `json:"failureReason,omitempty"`
	// Optional. The name of the generated report.
	Name NullableString `json:"name,omitempty"`
	// The date at which the report was generated in ISO 8601 date time format.
	GeneratedAt NullableString `json:"generatedAt,omitempty"`
	// The start date for the reporting period in YYYY-mm-dd format.
	StartDate string `json:"startDate"`
	// The build status of the report.   - `PENDING` - Report is created and awaiting processing.   - `PROCESSING` - Report is processing. Please wait.   - `COMPLETED` - Report has completed.  Check the `url` for the output file.   - `FAILED` - Report generation failed.  Check the `failureReason` for details.
	Status string `json:"status"`
	// The date at which the report was last updated in ISO 8601 date time format.
	UpdatedAt string `json:"updatedAt"`
}

type _AsyncReport AsyncReport

// NewAsyncReport instantiates a new AsyncReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsyncReport(reportId string, endDate string, configuration AsyncReportConfiguration, createdAt string, startDate string, status string, updatedAt string) *AsyncReport {
	this := AsyncReport{}
	this.ReportId = reportId
	this.EndDate = endDate
	this.Configuration = configuration
	this.CreatedAt = createdAt
	this.StartDate = startDate
	this.Status = status
	this.UpdatedAt = updatedAt
	return &this
}

// NewAsyncReportWithDefaults instantiates a new AsyncReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsyncReportWithDefaults() *AsyncReport {
	this := AsyncReport{}
	return &this
}

// GetReportId returns the ReportId field value
func (o *AsyncReport) GetReportId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value
// and a boolean to check if the value has been set.
func (o *AsyncReport) GetReportIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportId, true
}

// SetReportId sets field value
func (o *AsyncReport) SetReportId(v string) {
	o.ReportId = v
}

// GetEndDate returns the EndDate field value
func (o *AsyncReport) GetEndDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *AsyncReport) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *AsyncReport) SetEndDate(v string) {
	o.EndDate = v
}

// GetConfiguration returns the Configuration field value
func (o *AsyncReport) GetConfiguration() AsyncReportConfiguration {
	if o == nil {
		var ret AsyncReportConfiguration
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *AsyncReport) GetConfigurationOk() (*AsyncReportConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// SetConfiguration sets field value
func (o *AsyncReport) SetConfiguration(v AsyncReportConfiguration) {
	o.Configuration = v
}

// GetUrlExpiresAt returns the UrlExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AsyncReport) GetUrlExpiresAt() string {
	if o == nil || IsNil(o.UrlExpiresAt.Get()) {
		var ret string
		return ret
	}
	return *o.UrlExpiresAt.Get()
}

// GetUrlExpiresAtOk returns a tuple with the UrlExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AsyncReport) GetUrlExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UrlExpiresAt.Get(), o.UrlExpiresAt.IsSet()
}

// HasUrlExpiresAt returns a boolean if a field has been set.
func (o *AsyncReport) HasUrlExpiresAt() bool {
	if o != nil && o.UrlExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetUrlExpiresAt gets a reference to the given NullableString and assigns it to the UrlExpiresAt field.
func (o *AsyncReport) SetUrlExpiresAt(v string) {
	o.UrlExpiresAt.Set(&v)
}

// SetUrlExpiresAtNil sets the value for UrlExpiresAt to be an explicit nil
func (o *AsyncReport) SetUrlExpiresAtNil() {
	o.UrlExpiresAt.Set(nil)
}

// UnsetUrlExpiresAt ensures that no value is present for UrlExpiresAt, not even an explicit nil
func (o *AsyncReport) UnsetUrlExpiresAt() {
	o.UrlExpiresAt.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AsyncReport) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AsyncReport) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *AsyncReport) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *AsyncReport) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *AsyncReport) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *AsyncReport) UnsetUrl() {
	o.Url.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *AsyncReport) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AsyncReport) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AsyncReport) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AsyncReport) GetFileSize() float32 {
	if o == nil || IsNil(o.FileSize.Get()) {
		var ret float32
		return ret
	}
	return *o.FileSize.Get()
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AsyncReport) GetFileSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileSize.Get(), o.FileSize.IsSet()
}

// HasFileSize returns a boolean if a field has been set.
func (o *AsyncReport) HasFileSize() bool {
	if o != nil && o.FileSize.IsSet() {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given NullableFloat32 and assigns it to the FileSize field.
func (o *AsyncReport) SetFileSize(v float32) {
	o.FileSize.Set(&v)
}

// SetFileSizeNil sets the value for FileSize to be an explicit nil
func (o *AsyncReport) SetFileSizeNil() {
	o.FileSize.Set(nil)
}

// UnsetFileSize ensures that no value is present for FileSize, not even an explicit nil
func (o *AsyncReport) UnsetFileSize() {
	o.FileSize.Unset()
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AsyncReport) GetFailureReason() string {
	if o == nil || IsNil(o.FailureReason.Get()) {
		var ret string
		return ret
	}
	return *o.FailureReason.Get()
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AsyncReport) GetFailureReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureReason.Get(), o.FailureReason.IsSet()
}

// HasFailureReason returns a boolean if a field has been set.
func (o *AsyncReport) HasFailureReason() bool {
	if o != nil && o.FailureReason.IsSet() {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given NullableString and assigns it to the FailureReason field.
func (o *AsyncReport) SetFailureReason(v string) {
	o.FailureReason.Set(&v)
}

// SetFailureReasonNil sets the value for FailureReason to be an explicit nil
func (o *AsyncReport) SetFailureReasonNil() {
	o.FailureReason.Set(nil)
}

// UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
func (o *AsyncReport) UnsetFailureReason() {
	o.FailureReason.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AsyncReport) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AsyncReport) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AsyncReport) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AsyncReport) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *AsyncReport) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AsyncReport) UnsetName() {
	o.Name.Unset()
}

// GetGeneratedAt returns the GeneratedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AsyncReport) GetGeneratedAt() string {
	if o == nil || IsNil(o.GeneratedAt.Get()) {
		var ret string
		return ret
	}
	return *o.GeneratedAt.Get()
}

// GetGeneratedAtOk returns a tuple with the GeneratedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AsyncReport) GetGeneratedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GeneratedAt.Get(), o.GeneratedAt.IsSet()
}

// HasGeneratedAt returns a boolean if a field has been set.
func (o *AsyncReport) HasGeneratedAt() bool {
	if o != nil && o.GeneratedAt.IsSet() {
		return true
	}

	return false
}

// SetGeneratedAt gets a reference to the given NullableString and assigns it to the GeneratedAt field.
func (o *AsyncReport) SetGeneratedAt(v string) {
	o.GeneratedAt.Set(&v)
}

// SetGeneratedAtNil sets the value for GeneratedAt to be an explicit nil
func (o *AsyncReport) SetGeneratedAtNil() {
	o.GeneratedAt.Set(nil)
}

// UnsetGeneratedAt ensures that no value is present for GeneratedAt, not even an explicit nil
func (o *AsyncReport) UnsetGeneratedAt() {
	o.GeneratedAt.Unset()
}

// GetStartDate returns the StartDate field value
func (o *AsyncReport) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *AsyncReport) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *AsyncReport) SetStartDate(v string) {
	o.StartDate = v
}

// GetStatus returns the Status field value
func (o *AsyncReport) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AsyncReport) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AsyncReport) SetStatus(v string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *AsyncReport) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AsyncReport) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *AsyncReport) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o AsyncReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reportId"] = o.ReportId
	toSerialize["endDate"] = o.EndDate
	toSerialize["configuration"] = o.Configuration
	if o.UrlExpiresAt.IsSet() {
		toSerialize["urlExpiresAt"] = o.UrlExpiresAt.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	toSerialize["createdAt"] = o.CreatedAt
	if o.FileSize.IsSet() {
		toSerialize["fileSize"] = o.FileSize.Get()
	}
	if o.FailureReason.IsSet() {
		toSerialize["failureReason"] = o.FailureReason.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.GeneratedAt.IsSet() {
		toSerialize["generatedAt"] = o.GeneratedAt.Get()
	}
	toSerialize["startDate"] = o.StartDate
	toSerialize["status"] = o.Status
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableAsyncReport struct {
	value *AsyncReport
	isSet bool
}

func (v NullableAsyncReport) Get() *AsyncReport {
	return v.value
}

func (v *NullableAsyncReport) Set(val *AsyncReport) {
	v.value = val
	v.isSet = true
}

func (v NullableAsyncReport) IsSet() bool {
	return v.isSet
}

func (v *NullableAsyncReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsyncReport(val *AsyncReport) *NullableAsyncReport {
	return &NullableAsyncReport{value: val, isSet: true}
}

func (v NullableAsyncReport) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAsyncReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
