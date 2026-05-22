package report_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the DeleteAsyncReportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteAsyncReportResponse{}

// DeleteAsyncReportResponse struct for DeleteAsyncReportResponse
type DeleteAsyncReportResponse struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// The identifier of the report.
	ReportId *string `json:"reportId,omitempty"`
	// A human-readable description of the response.
	Detail *string `json:"detail,omitempty"`
}

// NewDeleteAsyncReportResponse instantiates a new DeleteAsyncReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAsyncReportResponse() *DeleteAsyncReportResponse {
	this := DeleteAsyncReportResponse{}
	return &this
}

// NewDeleteAsyncReportResponseWithDefaults instantiates a new DeleteAsyncReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteAsyncReportResponseWithDefaults() *DeleteAsyncReportResponse {
	this := DeleteAsyncReportResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *DeleteAsyncReportResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAsyncReportResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *DeleteAsyncReportResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *DeleteAsyncReportResponse) SetCode(v string) {
	o.Code = &v
}

// GetReportId returns the ReportId field value if set, zero value otherwise.
func (o *DeleteAsyncReportResponse) GetReportId() string {
	if o == nil || IsNil(o.ReportId) {
		var ret string
		return ret
	}
	return *o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAsyncReportResponse) GetReportIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReportId) {
		return nil, false
	}
	return o.ReportId, true
}

// HasReportId returns a boolean if a field has been set.
func (o *DeleteAsyncReportResponse) HasReportId() bool {
	if o != nil && !IsNil(o.ReportId) {
		return true
	}

	return false
}

// SetReportId gets a reference to the given string and assigns it to the ReportId field.
func (o *DeleteAsyncReportResponse) SetReportId(v string) {
	o.ReportId = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *DeleteAsyncReportResponse) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAsyncReportResponse) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *DeleteAsyncReportResponse) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *DeleteAsyncReportResponse) SetDetail(v string) {
	o.Detail = &v
}

func (o DeleteAsyncReportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.ReportId) {
		toSerialize["reportId"] = o.ReportId
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

type NullableDeleteAsyncReportResponse struct {
	value *DeleteAsyncReportResponse
	isSet bool
}

func (v NullableDeleteAsyncReportResponse) Get() *DeleteAsyncReportResponse {
	return v.value
}

func (v *NullableDeleteAsyncReportResponse) Set(val *DeleteAsyncReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAsyncReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAsyncReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAsyncReportResponse(val *DeleteAsyncReportResponse) *NullableDeleteAsyncReportResponse {
	return &NullableDeleteAsyncReportResponse{value: val, isSet: true}
}

func (v NullableDeleteAsyncReportResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeleteAsyncReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
