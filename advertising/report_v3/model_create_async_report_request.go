package report_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateAsyncReportRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAsyncReportRequest{}

// CreateAsyncReportRequest struct for CreateAsyncReportRequest
type CreateAsyncReportRequest struct {
	// YYYY-MM-DD format. The maximum lookback window supported depends on the selection of reportTypeId. Most report types support `95 days` as lookback window.
	EndDate       string                   `json:"endDate"`
	Configuration AsyncReportConfiguration `json:"configuration"`
	// The name of the report.
	Name *string `json:"name,omitempty"`
	// YYYY-MM-DD format. The maximum lookback window supported depends on the selection of reportTypeId. Most report types support `95 days` as lookback window.
	StartDate string `json:"startDate"`
}

type _CreateAsyncReportRequest CreateAsyncReportRequest

// NewCreateAsyncReportRequest instantiates a new CreateAsyncReportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAsyncReportRequest(endDate string, configuration AsyncReportConfiguration, startDate string) *CreateAsyncReportRequest {
	this := CreateAsyncReportRequest{}
	this.EndDate = endDate
	this.Configuration = configuration
	this.StartDate = startDate
	return &this
}

// NewCreateAsyncReportRequestWithDefaults instantiates a new CreateAsyncReportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAsyncReportRequestWithDefaults() *CreateAsyncReportRequest {
	this := CreateAsyncReportRequest{}
	return &this
}

// GetEndDate returns the EndDate field value
func (o *CreateAsyncReportRequest) GetEndDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *CreateAsyncReportRequest) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *CreateAsyncReportRequest) SetEndDate(v string) {
	o.EndDate = v
}

// GetConfiguration returns the Configuration field value
func (o *CreateAsyncReportRequest) GetConfiguration() AsyncReportConfiguration {
	if o == nil {
		var ret AsyncReportConfiguration
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *CreateAsyncReportRequest) GetConfigurationOk() (*AsyncReportConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// SetConfiguration sets field value
func (o *CreateAsyncReportRequest) SetConfiguration(v AsyncReportConfiguration) {
	o.Configuration = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateAsyncReportRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAsyncReportRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateAsyncReportRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateAsyncReportRequest) SetName(v string) {
	o.Name = &v
}

// GetStartDate returns the StartDate field value
func (o *CreateAsyncReportRequest) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *CreateAsyncReportRequest) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *CreateAsyncReportRequest) SetStartDate(v string) {
	o.StartDate = v
}

func (o CreateAsyncReportRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["endDate"] = o.EndDate
	toSerialize["configuration"] = o.Configuration
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["startDate"] = o.StartDate
	return toSerialize, nil
}

type NullableCreateAsyncReportRequest struct {
	value *CreateAsyncReportRequest
	isSet bool
}

func (v NullableCreateAsyncReportRequest) Get() *CreateAsyncReportRequest {
	return v.value
}

func (v *NullableCreateAsyncReportRequest) Set(val *CreateAsyncReportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAsyncReportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAsyncReportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAsyncReportRequest(val *CreateAsyncReportRequest) *NullableCreateAsyncReportRequest {
	return &NullableCreateAsyncReportRequest{value: val, isSet: true}
}

func (v NullableCreateAsyncReportRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateAsyncReportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
