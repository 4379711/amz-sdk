package reports_20210630

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateReportScheduleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateReportScheduleResponse{}

// CreateReportScheduleResponse Response schema.
type CreateReportScheduleResponse struct {
	// The identifier for the report schedule. This identifier is unique only in combination with a seller ID.
	ReportScheduleId string `json:"reportScheduleId"`
}

type _CreateReportScheduleResponse CreateReportScheduleResponse

// NewCreateReportScheduleResponse instantiates a new CreateReportScheduleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReportScheduleResponse(reportScheduleId string) *CreateReportScheduleResponse {
	this := CreateReportScheduleResponse{}
	this.ReportScheduleId = reportScheduleId
	return &this
}

// NewCreateReportScheduleResponseWithDefaults instantiates a new CreateReportScheduleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReportScheduleResponseWithDefaults() *CreateReportScheduleResponse {
	this := CreateReportScheduleResponse{}
	return &this
}

// GetReportScheduleId returns the ReportScheduleId field value
func (o *CreateReportScheduleResponse) GetReportScheduleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportScheduleId
}

// GetReportScheduleIdOk returns a tuple with the ReportScheduleId field value
// and a boolean to check if the value has been set.
func (o *CreateReportScheduleResponse) GetReportScheduleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportScheduleId, true
}

// SetReportScheduleId sets field value
func (o *CreateReportScheduleResponse) SetReportScheduleId(v string) {
	o.ReportScheduleId = v
}

func (o CreateReportScheduleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reportScheduleId"] = o.ReportScheduleId
	return toSerialize, nil
}

type NullableCreateReportScheduleResponse struct {
	value *CreateReportScheduleResponse
	isSet bool
}

func (v NullableCreateReportScheduleResponse) Get() *CreateReportScheduleResponse {
	return v.value
}

func (v *NullableCreateReportScheduleResponse) Set(val *CreateReportScheduleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReportScheduleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReportScheduleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReportScheduleResponse(val *CreateReportScheduleResponse) *NullableCreateReportScheduleResponse {
	return &NullableCreateReportScheduleResponse{value: val, isSet: true}
}

func (v NullableCreateReportScheduleResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateReportScheduleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
