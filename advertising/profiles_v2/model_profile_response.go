package profiles_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the ProfileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileResponse{}

// ProfileResponse struct for ProfileResponse
type ProfileResponse struct {
	ProfileId *int64  `json:"profileId,omitempty"`
	Code      *string `json:"code,omitempty"`
	Details   *string `json:"details,omitempty"`
}

// NewProfileResponse instantiates a new ProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileResponse() *ProfileResponse {
	this := ProfileResponse{}
	return &this
}

// NewProfileResponseWithDefaults instantiates a new ProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileResponseWithDefaults() *ProfileResponse {
	this := ProfileResponse{}
	return &this
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *ProfileResponse) GetProfileId() int64 {
	if o == nil || IsNil(o.ProfileId) {
		var ret int64
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileResponse) GetProfileIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *ProfileResponse) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given int64 and assigns it to the ProfileId field.
func (o *ProfileResponse) SetProfileId(v int64) {
	o.ProfileId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProfileResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProfileResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProfileResponse) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ProfileResponse) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileResponse) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ProfileResponse) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ProfileResponse) SetDetails(v string) {
	o.Details = &v
}

func (o ProfileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableProfileResponse struct {
	value *ProfileResponse
	isSet bool
}

func (v NullableProfileResponse) Get() *ProfileResponse {
	return v.value
}

func (v *NullableProfileResponse) Set(val *ProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileResponse(val *ProfileResponse) *NullableProfileResponse {
	return &NullableProfileResponse{value: val, isSet: true}
}

func (v NullableProfileResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
