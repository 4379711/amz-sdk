package reports_20210630

import (
	"github.com/bytedance/sonic"
)

// checks if the GetReportsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReportsResponse{}

// GetReportsResponse The response for the `getReports` operation.
type GetReportsResponse struct {
	// A list of reports.
	Reports []Report `json:"reports"`
	// Returned when the number of results exceeds `pageSize`. To get the next page of results, call `getReports` with this token as the only parameter.
	NextToken *string `json:"nextToken,omitempty"`
}

type _GetReportsResponse GetReportsResponse

// NewGetReportsResponse instantiates a new GetReportsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReportsResponse(reports []Report) *GetReportsResponse {
	this := GetReportsResponse{}
	this.Reports = reports
	return &this
}

// NewGetReportsResponseWithDefaults instantiates a new GetReportsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReportsResponseWithDefaults() *GetReportsResponse {
	this := GetReportsResponse{}
	return &this
}

// GetReports returns the Reports field value
func (o *GetReportsResponse) GetReports() []Report {
	if o == nil {
		var ret []Report
		return ret
	}

	return o.Reports
}

// GetReportsOk returns a tuple with the Reports field value
// and a boolean to check if the value has been set.
func (o *GetReportsResponse) GetReportsOk() ([]Report, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reports, true
}

// SetReports sets field value
func (o *GetReportsResponse) SetReports(v []Report) {
	o.Reports = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *GetReportsResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *GetReportsResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *GetReportsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o GetReportsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reports"] = o.Reports
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableGetReportsResponse struct {
	value *GetReportsResponse
	isSet bool
}

func (v NullableGetReportsResponse) Get() *GetReportsResponse {
	return v.value
}

func (v *NullableGetReportsResponse) Set(val *GetReportsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReportsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReportsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReportsResponse(val *GetReportsResponse) *NullableGetReportsResponse {
	return &NullableGetReportsResponse{value: val, isSet: true}
}

func (v NullableGetReportsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetReportsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
