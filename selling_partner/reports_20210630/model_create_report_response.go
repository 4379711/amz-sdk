package reports_20210630

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateReportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateReportResponse{}

// CreateReportResponse The response schema.
type CreateReportResponse struct {
	// The identifier for the report. This identifier is unique only in combination with a seller ID.
	ReportId string `json:"reportId"`
}

type _CreateReportResponse CreateReportResponse

// NewCreateReportResponse instantiates a new CreateReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReportResponse(reportId string) *CreateReportResponse {
	this := CreateReportResponse{}
	this.ReportId = reportId
	return &this
}

// NewCreateReportResponseWithDefaults instantiates a new CreateReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReportResponseWithDefaults() *CreateReportResponse {
	this := CreateReportResponse{}
	return &this
}

// GetReportId returns the ReportId field value
func (o *CreateReportResponse) GetReportId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value
// and a boolean to check if the value has been set.
func (o *CreateReportResponse) GetReportIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportId, true
}

// SetReportId sets field value
func (o *CreateReportResponse) SetReportId(v string) {
	o.ReportId = v
}

func (o CreateReportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reportId"] = o.ReportId
	return toSerialize, nil
}

type NullableCreateReportResponse struct {
	value *CreateReportResponse
	isSet bool
}

func (v NullableCreateReportResponse) Get() *CreateReportResponse {
	return v.value
}

func (v *NullableCreateReportResponse) Set(val *CreateReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReportResponse(val *CreateReportResponse) *NullableCreateReportResponse {
	return &NullableCreateReportResponse{value: val, isSet: true}
}

func (v NullableCreateReportResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
