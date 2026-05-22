package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the BrandSafetyRequestStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandSafetyRequestStatus{}

// BrandSafetyRequestStatus struct for BrandSafetyRequestStatus
type BrandSafetyRequestStatus struct {
	// Request ID
	RequestId *string `json:"requestId,omitempty"`
	// Request timestamp
	Timestamp *string `json:"timestamp,omitempty"`
	// The status of the request
	Status *string `json:"status,omitempty"`
	// Details related to the request status
	StatusDetails *string `json:"statusDetails,omitempty"`
}

// NewBrandSafetyRequestStatus instantiates a new BrandSafetyRequestStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandSafetyRequestStatus() *BrandSafetyRequestStatus {
	this := BrandSafetyRequestStatus{}
	return &this
}

// NewBrandSafetyRequestStatusWithDefaults instantiates a new BrandSafetyRequestStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandSafetyRequestStatusWithDefaults() *BrandSafetyRequestStatus {
	this := BrandSafetyRequestStatus{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *BrandSafetyRequestStatus) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyRequestStatus) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *BrandSafetyRequestStatus) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *BrandSafetyRequestStatus) SetRequestId(v string) {
	o.RequestId = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *BrandSafetyRequestStatus) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyRequestStatus) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *BrandSafetyRequestStatus) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *BrandSafetyRequestStatus) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BrandSafetyRequestStatus) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyRequestStatus) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BrandSafetyRequestStatus) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BrandSafetyRequestStatus) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDetails returns the StatusDetails field value if set, zero value otherwise.
func (o *BrandSafetyRequestStatus) GetStatusDetails() string {
	if o == nil || IsNil(o.StatusDetails) {
		var ret string
		return ret
	}
	return *o.StatusDetails
}

// GetStatusDetailsOk returns a tuple with the StatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyRequestStatus) GetStatusDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.StatusDetails) {
		return nil, false
	}
	return o.StatusDetails, true
}

// HasStatusDetails returns a boolean if a field has been set.
func (o *BrandSafetyRequestStatus) HasStatusDetails() bool {
	if o != nil && !IsNil(o.StatusDetails) {
		return true
	}

	return false
}

// SetStatusDetails gets a reference to the given string and assigns it to the StatusDetails field.
func (o *BrandSafetyRequestStatus) SetStatusDetails(v string) {
	o.StatusDetails = &v
}

func (o BrandSafetyRequestStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusDetails) {
		toSerialize["statusDetails"] = o.StatusDetails
	}
	return toSerialize, nil
}

type NullableBrandSafetyRequestStatus struct {
	value *BrandSafetyRequestStatus
	isSet bool
}

func (v NullableBrandSafetyRequestStatus) Get() *BrandSafetyRequestStatus {
	return v.value
}

func (v *NullableBrandSafetyRequestStatus) Set(val *BrandSafetyRequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandSafetyRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandSafetyRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandSafetyRequestStatus(val *BrandSafetyRequestStatus) *NullableBrandSafetyRequestStatus {
	return &NullableBrandSafetyRequestStatus{value: val, isSet: true}
}

func (v NullableBrandSafetyRequestStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandSafetyRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
