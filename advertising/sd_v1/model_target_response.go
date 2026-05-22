package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the TargetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetResponse{}

// TargetResponse struct for TargetResponse
type TargetResponse struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Description *string `json:"description,omitempty"`
	TargetId    *int64  `json:"targetId,omitempty"`
}

// NewTargetResponse instantiates a new TargetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetResponse() *TargetResponse {
	this := TargetResponse{}
	return &this
}

// NewTargetResponseWithDefaults instantiates a new TargetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetResponseWithDefaults() *TargetResponse {
	this := TargetResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *TargetResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *TargetResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *TargetResponse) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TargetResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TargetResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TargetResponse) SetDescription(v string) {
	o.Description = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *TargetResponse) GetTargetId() int64 {
	if o == nil || IsNil(o.TargetId) {
		var ret int64
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetResponse) GetTargetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *TargetResponse) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given int64 and assigns it to the TargetId field.
func (o *TargetResponse) SetTargetId(v int64) {
	o.TargetId = &v
}

func (o TargetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.TargetId) {
		toSerialize["targetId"] = o.TargetId
	}
	return toSerialize, nil
}

type NullableTargetResponse struct {
	value *TargetResponse
	isSet bool
}

func (v NullableTargetResponse) Get() *TargetResponse {
	return v.value
}

func (v *NullableTargetResponse) Set(val *TargetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetResponse(val *TargetResponse) *NullableTargetResponse {
	return &NullableTargetResponse{value: val, isSet: true}
}

func (v NullableTargetResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTargetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
