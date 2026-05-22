package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the AdGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdGroupResponse{}

// AdGroupResponse struct for AdGroupResponse
type AdGroupResponse struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Description *string `json:"description,omitempty"`
	// The identifier of the ad group.
	AdGroupId *int64 `json:"adGroupId,omitempty"`
}

// NewAdGroupResponse instantiates a new AdGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdGroupResponse() *AdGroupResponse {
	this := AdGroupResponse{}
	return &this
}

// NewAdGroupResponseWithDefaults instantiates a new AdGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdGroupResponseWithDefaults() *AdGroupResponse {
	this := AdGroupResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AdGroupResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AdGroupResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AdGroupResponse) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AdGroupResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AdGroupResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AdGroupResponse) SetDescription(v string) {
	o.Description = &v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *AdGroupResponse) GetAdGroupId() int64 {
	if o == nil || IsNil(o.AdGroupId) {
		var ret int64
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupResponse) GetAdGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *AdGroupResponse) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given int64 and assigns it to the AdGroupId field.
func (o *AdGroupResponse) SetAdGroupId(v int64) {
	o.AdGroupId = &v
}

func (o AdGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	return toSerialize, nil
}

type NullableAdGroupResponse struct {
	value *AdGroupResponse
	isSet bool
}

func (v NullableAdGroupResponse) Get() *AdGroupResponse {
	return v.value
}

func (v *NullableAdGroupResponse) Set(val *AdGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAdGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAdGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdGroupResponse(val *AdGroupResponse) *NullableAdGroupResponse {
	return &NullableAdGroupResponse{value: val, isSet: true}
}

func (v NullableAdGroupResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
