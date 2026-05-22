package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the ReportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportResponse{}

// ReportResponse struct for ReportResponse
type ReportResponse struct {
	// The identifier of the report.
	ReportId *string `json:"reportId,omitempty"`
	// The type of report requested.
	RecordType *string `json:"recordType,omitempty"`
	// The build status of the report.
	Status *string `json:"status,omitempty"`
	// A human-readable description of the current status.
	StatusDetails *string `json:"statusDetails,omitempty"`
	// The URI location of the report.
	Location *string `json:"location,omitempty"`
	// The size of the report file, in bytes.
	FileSize *int64 `json:"fileSize,omitempty"`
	// Epoch date of the expiration of the URI in the `location` property.
	Expiration *int64 `json:"expiration,omitempty"`
}

// NewReportResponse instantiates a new ReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportResponse() *ReportResponse {
	this := ReportResponse{}
	return &this
}

// NewReportResponseWithDefaults instantiates a new ReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportResponseWithDefaults() *ReportResponse {
	this := ReportResponse{}
	return &this
}

// GetReportId returns the ReportId field value if set, zero value otherwise.
func (o *ReportResponse) GetReportId() string {
	if o == nil || IsNil(o.ReportId) {
		var ret string
		return ret
	}
	return *o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResponse) GetReportIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReportId) {
		return nil, false
	}
	return o.ReportId, true
}

// HasReportId returns a boolean if a field has been set.
func (o *ReportResponse) HasReportId() bool {
	if o != nil && !IsNil(o.ReportId) {
		return true
	}

	return false
}

// SetReportId gets a reference to the given string and assigns it to the ReportId field.
func (o *ReportResponse) SetReportId(v string) {
	o.ReportId = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *ReportResponse) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResponse) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *ReportResponse) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *ReportResponse) SetRecordType(v string) {
	o.RecordType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ReportResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ReportResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ReportResponse) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDetails returns the StatusDetails field value if set, zero value otherwise.
func (o *ReportResponse) GetStatusDetails() string {
	if o == nil || IsNil(o.StatusDetails) {
		var ret string
		return ret
	}
	return *o.StatusDetails
}

// GetStatusDetailsOk returns a tuple with the StatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResponse) GetStatusDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.StatusDetails) {
		return nil, false
	}
	return o.StatusDetails, true
}

// HasStatusDetails returns a boolean if a field has been set.
func (o *ReportResponse) HasStatusDetails() bool {
	if o != nil && !IsNil(o.StatusDetails) {
		return true
	}

	return false
}

// SetStatusDetails gets a reference to the given string and assigns it to the StatusDetails field.
func (o *ReportResponse) SetStatusDetails(v string) {
	o.StatusDetails = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ReportResponse) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResponse) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ReportResponse) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *ReportResponse) SetLocation(v string) {
	o.Location = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *ReportResponse) GetFileSize() int64 {
	if o == nil || IsNil(o.FileSize) {
		var ret int64
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResponse) GetFileSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *ReportResponse) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int64 and assigns it to the FileSize field.
func (o *ReportResponse) SetFileSize(v int64) {
	o.FileSize = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *ReportResponse) GetExpiration() int64 {
	if o == nil || IsNil(o.Expiration) {
		var ret int64
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResponse) GetExpirationOk() (*int64, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *ReportResponse) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given int64 and assigns it to the Expiration field.
func (o *ReportResponse) SetExpiration(v int64) {
	o.Expiration = &v
}

func (o ReportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReportId) {
		toSerialize["reportId"] = o.ReportId
	}
	if !IsNil(o.RecordType) {
		toSerialize["recordType"] = o.RecordType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusDetails) {
		toSerialize["statusDetails"] = o.StatusDetails
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.FileSize) {
		toSerialize["fileSize"] = o.FileSize
	}
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	return toSerialize, nil
}

type NullableReportResponse struct {
	value *ReportResponse
	isSet bool
}

func (v NullableReportResponse) Get() *ReportResponse {
	return v.value
}

func (v *NullableReportResponse) Set(val *ReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportResponse(val *ReportResponse) *NullableReportResponse {
	return &NullableReportResponse{value: val, isSet: true}
}

func (v NullableReportResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
