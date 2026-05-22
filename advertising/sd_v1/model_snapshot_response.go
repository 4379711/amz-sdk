package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SnapshotResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotResponse{}

// SnapshotResponse struct for SnapshotResponse
type SnapshotResponse struct {
	// The identifier of the snapshot that was requested.
	SnapshotId *string `json:"snapshotId,omitempty"`
	// The record type of the snapshot file.
	RecordType *string `json:"recordType,omitempty"`
	// The status of the generation of the snapshot.
	Status *string `json:"status,omitempty"`
	// Optional description of the status.
	StatusDetails *string `json:"statusDetails,omitempty"`
	// The URI for the snapshot. It's only available if status is SUCCESS.
	Location *string `json:"location,omitempty"`
	// The size of the snapshot file in bytes. It's only available if status is SUCCESS.
	FileSize *float32 `json:"fileSize,omitempty"`
	// The epoch time for expiration of the snapshot file and each snapshot file will be expired in 30 mins after generated. It's only available if status is SUCCESS.
	Expiration *float32 `json:"expiration,omitempty"`
}

// NewSnapshotResponse instantiates a new SnapshotResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotResponse() *SnapshotResponse {
	this := SnapshotResponse{}
	return &this
}

// NewSnapshotResponseWithDefaults instantiates a new SnapshotResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotResponseWithDefaults() *SnapshotResponse {
	this := SnapshotResponse{}
	return &this
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *SnapshotResponse) GetSnapshotId() string {
	if o == nil || IsNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotResponse) GetSnapshotIdOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotId) {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *SnapshotResponse) HasSnapshotId() bool {
	if o != nil && !IsNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *SnapshotResponse) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *SnapshotResponse) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotResponse) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *SnapshotResponse) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *SnapshotResponse) SetRecordType(v string) {
	o.RecordType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SnapshotResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SnapshotResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SnapshotResponse) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDetails returns the StatusDetails field value if set, zero value otherwise.
func (o *SnapshotResponse) GetStatusDetails() string {
	if o == nil || IsNil(o.StatusDetails) {
		var ret string
		return ret
	}
	return *o.StatusDetails
}

// GetStatusDetailsOk returns a tuple with the StatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotResponse) GetStatusDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.StatusDetails) {
		return nil, false
	}
	return o.StatusDetails, true
}

// HasStatusDetails returns a boolean if a field has been set.
func (o *SnapshotResponse) HasStatusDetails() bool {
	if o != nil && !IsNil(o.StatusDetails) {
		return true
	}

	return false
}

// SetStatusDetails gets a reference to the given string and assigns it to the StatusDetails field.
func (o *SnapshotResponse) SetStatusDetails(v string) {
	o.StatusDetails = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SnapshotResponse) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotResponse) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SnapshotResponse) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *SnapshotResponse) SetLocation(v string) {
	o.Location = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *SnapshotResponse) GetFileSize() float32 {
	if o == nil || IsNil(o.FileSize) {
		var ret float32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotResponse) GetFileSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *SnapshotResponse) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given float32 and assigns it to the FileSize field.
func (o *SnapshotResponse) SetFileSize(v float32) {
	o.FileSize = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *SnapshotResponse) GetExpiration() float32 {
	if o == nil || IsNil(o.Expiration) {
		var ret float32
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotResponse) GetExpirationOk() (*float32, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *SnapshotResponse) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given float32 and assigns it to the Expiration field.
func (o *SnapshotResponse) SetExpiration(v float32) {
	o.Expiration = &v
}

func (o SnapshotResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SnapshotId) {
		toSerialize["snapshotId"] = o.SnapshotId
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

type NullableSnapshotResponse struct {
	value *SnapshotResponse
	isSet bool
}

func (v NullableSnapshotResponse) Get() *SnapshotResponse {
	return v.value
}

func (v *NullableSnapshotResponse) Set(val *SnapshotResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotResponse(val *SnapshotResponse) *NullableSnapshotResponse {
	return &NullableSnapshotResponse{value: val, isSet: true}
}

func (v NullableSnapshotResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSnapshotResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
